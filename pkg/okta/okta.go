package okta

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/authn"
	"github.com/cmsgov/easi-app/pkg/handlers"
)

const (
	prodGRTJobCode = "EASI_P_GOVTEAM"
	testGRTJobCode = "EASI_D_GOVTEAM"
)

func (f oktaMiddlewareFactory) jwt(logger *zap.Logger, authHeader string) (*jwtverifier.Jwt, error) {
	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) < 2 {
		return nil, errors.New("invalid Bearer in auth header")
	}
	bearerToken := tokenParts[1]
	if bearerToken == "" {
		return nil, errors.New("empty bearer value")
	}

	return f.verifier.VerifyAccessToken(bearerToken)
}

func jwtGroupsContainsJobCode(jwt *jwtverifier.Jwt, jobCode string) bool {
	list, ok := jwt.Claims["groups"]
	if !ok {
		return false
	}

	// json arrays decode to `[]interface{}`
	codes, ok := list.([]interface{})
	if !ok {
		return false
	}

	for _, code := range codes {
		if c, ok := code.(string); ok {
			if strings.EqualFold(c, jobCode) {
				return true
			}
		}
	}
	return false
}

func (f oktaMiddlewareFactory) newPrincipal(jwt *jwtverifier.Jwt) (*authn.EUAPrincipal, error) {
	euaID := jwt.Claims["sub"].(string)
	if euaID == "" {
		return nil, errors.New("unable to retrieve EUA ID from JWT")
	}

	// the current assumption is that anyone with an appropriate
	// JWT provided by Okta for EASi is allowed to use EASi
	// as a viewer/submitter
	jcEASi := true

	// need to check the claims for empowerment as a reviewer
	jcGRT := jwtGroupsContainsJobCode(jwt, f.codeGRT)

	return &authn.EUAPrincipal{EUAID: euaID, JobCodeEASi: jcEASi, JobCodeGRT: jcGRT}, nil
}

func (f oktaMiddlewareFactory) newAuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := appcontext.ZLogger(r.Context())
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			f.WriteErrorResponse(
				r.Context(),
				w,
				&apperrors.UnauthorizedError{Err: errors.New("empty authorization header")},
			)
			return
		}

		jwt, err := f.jwt(logger, authHeader)
		if err != nil {
			f.WriteErrorResponse(
				r.Context(),
				w,
				&apperrors.UnauthorizedError{Err: fmt.Errorf("unable to parse jwt: %w", err)},
			)
			return
		}

		principal, err := f.newPrincipal(jwt)
		if err != nil {
			f.WriteErrorResponse(
				r.Context(),
				w,
				&apperrors.UnauthorizedError{Err: fmt.Errorf("unable to get Principal from jwt: %w", err)},
			)
			return
		}
		logger = logger.With(zap.String("user", principal.ID())).With(zap.Bool("grt", principal.AllowGRT()))

		ctx := r.Context()
		ctx = appcontext.WithPrincipal(ctx, principal)
		ctx = appcontext.WithLogger(ctx, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newJwtVerifier(clientID string, issuer string) *jwtverifier.JwtVerifier {
	toValidate := map[string]string{}
	toValidate["cid"] = clientID
	toValidate["aud"] = "EASi"

	jwtVerifierSetup := jwtverifier.JwtVerifier{
		Issuer:           issuer,
		ClaimsToValidate: toValidate,
	}

	return jwtVerifierSetup.New()
}

type oktaMiddlewareFactory struct {
	handlers.HandlerBase
	verifier *jwtverifier.JwtVerifier
	codeGRT  string
}

// NewOktaAuthorizeMiddleware returns a wrapper for HandlerFunc to authorize with Okta
func NewOktaAuthorizeMiddleware(base handlers.HandlerBase, clientID string, issuer string, testGRT bool) func(http.Handler) http.Handler {
	verifier := newJwtVerifier(clientID, issuer)

	// by default we want to use the PROD job code, and only in
	// pre-PROD environments do we want to empower the
	// alternate job code.
	jobCodeGRT := prodGRTJobCode
	if testGRT {
		jobCodeGRT = testGRTJobCode
	}

	middlewareFactory := oktaMiddlewareFactory{HandlerBase: base, verifier: verifier, codeGRT: jobCodeGRT}
	return func(next http.Handler) http.Handler {
		return middlewareFactory.newAuthorizeMiddleware(next)
	}
}
