package server

import (
	"net/http"

	_ "github.com/lib/pq" // pq is required to get the postgres driver into sqlx
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appses"
	"github.com/cmsgov/easi-app/pkg/cedar/cedareasi"
	"github.com/cmsgov/easi-app/pkg/cedar/cedarldap"
	"github.com/cmsgov/easi-app/pkg/email"
	"github.com/cmsgov/easi-app/pkg/handlers"
	"github.com/cmsgov/easi-app/pkg/local"
	"github.com/cmsgov/easi-app/pkg/services"
	"github.com/cmsgov/easi-app/pkg/storage"
)

func (s *Server) routes(
	authorizationMiddleware func(handler http.Handler) http.Handler,
	corsMiddleware func(handler http.Handler) http.Handler,
	traceMiddleware func(handler http.Handler) http.Handler,
	loggerMiddleware func(handler http.Handler) http.Handler) {

	// trace all requests with an ID
	s.router.Use(traceMiddleware)

	// set up handler base
	base := handlers.NewHandlerBase(s.logger)

	// health check goes directly on the main router to avoid auth
	healthCheckHandler := handlers.NewHealthCheckHandler(
		base,
		s.Config,
	)
	s.router.HandleFunc("/api/v1/healthcheck", healthCheckHandler.Handle())

	// check we have all of the configs for CEDAR clients
	if s.environment.Deployed() {
		s.NewCEDARClientCheck()
	}

	// set up CEDAR client
	cedarEasiClient := cedareasi.NewTranslatedClient(
		s.Config.GetString("CEDAR_API_URL"),
		s.Config.GetString("CEDAR_API_KEY"),
	)

	cedarLdapClient := cedarldap.NewTranslatedClient(
		s.Config.GetString("CEDAR_API_URL"),
		s.Config.GetString("CEDAR_API_KEY"),
	)

	if s.environment.Prod() {
		s.CheckCEDAREasiClientConnection(cedarEasiClient)
		s.CheckCEDARLdapClientConnection(cedarLdapClient)
	}

	// set up Email Client
	sesConfig := s.NewSESConfig()
	sesSender := appses.NewSender(sesConfig)
	emailConfig := s.NewEmailConfig()
	emailClient, err := email.NewClient(emailConfig, sesSender)
	if err != nil {
		s.logger.Fatal("Failed to create email client", zap.Error(err))
	}
	// override email client with local one
	if s.environment.Local() {
		localSender := local.NewSender(s.logger)
		emailClient, err = email.NewClient(emailConfig, localSender)
		if err != nil {
			s.logger.Fatal("Failed to create email client", zap.Error(err))
		}
	}

	if s.environment.Deployed() {
		s.CheckEmailClient(emailClient)
	}

	// API base path is versioned
	api := s.router.PathPrefix("/api/v1").Subrouter()

	// add a request based logger
	api.Use(loggerMiddleware)

	// wrap with CORs
	api.Use(corsMiddleware)

	// protect all API routes with authorization middleware
	api.Use(authorizationMiddleware)

	serviceConfig := services.NewConfig(s.logger)

	store, err := storage.NewStore(
		s.logger,
		s.NewDBConfig(),
	)
	if err != nil {
		s.logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// endpoint for system list
	systemHandler := handlers.NewSystemsListHandler(
		base,
		cedarEasiClient.FetchSystems,
	)
	api.Handle("/systems", systemHandler.Handle())

	transitions := services.StateMachineTransitions{
		CreateIntake: services.NewCreateSystemIntake(
			serviceConfig,
			store.CreateSystemIntake,
		),
		UpdateDraft: services.NewUpdateDRAFTSystemIntake(
			serviceConfig,
			services.NewAuthorizeUserIsIntakeRequester(s.logger),
			store.UpdateSystemIntake,
		),
		Submit: services.NewSubmitSystemIntake(
			serviceConfig,
			services.NewAuthorizeUserIsIntakeRequester(s.logger),
			store.UpdateSystemIntake,
			cedarEasiClient.ValidateAndSubmitSystemIntake,
			emailClient.SendSystemIntakeSubmissionEmail,
		),
		Approve: services.NewDecideSystemIntake(
			serviceConfig,
			services.NewAuthorizeUserIsGRT(),
			cedarLdapClient.FetchUserEmailAddress,
			store.UpdateSystemIntake,
			emailClient.SendSystemIntakeReviewEmail,
		),
		Accept: services.NewDecideSystemIntake(
			serviceConfig,
			services.NewAuthorizeUserIsGRT(),
			cedarLdapClient.FetchUserEmailAddress,
			store.UpdateSystemIntake,
			emailClient.SendSystemIntakeReviewEmail,
		),
		Close: services.NewDecideSystemIntake(
			serviceConfig,
			services.NewAuthorizeUserIsGRT(),
			cedarLdapClient.FetchUserEmailAddress,
			store.UpdateSystemIntake,
			emailClient.SendSystemIntakeReviewEmail,
		),
		Archive: services.NewArchiveSystemIntake(
			serviceConfig,
			store.FetchSystemIntakeByID,
			store.UpdateSystemIntake,
			services.NewArchiveBusinessCase(
				serviceConfig,
				store.FetchBusinessCaseByID,
				store.UpdateBusinessCase,
			),
			services.NewAuthorizeUserIsIntakeRequester(s.logger),
		),
	}

	systemIntakeHandler := handlers.NewSystemIntakeHandler(
		base,
		services.NewCreateSystemIntakeWrapper(transitions.CreateIntake),
		services.NewUpdateSystemIntake(
			store.FetchSystemIntakeByID,
			!s.environment.Prod(),
			transitions,
		),
		services.NewFetchSystemIntakeByID(
			serviceConfig,
			store.FetchSystemIntakeByID,
			services.NewAuthorizeFetchSystemIntakeByID(),
		),
		services.NewArchiveSystemIntakeWrapper(transitions.Archive),
	)

	api.Handle("/system_intake/{intake_id}", systemIntakeHandler.Handle())
	api.Handle("/system_intake", systemIntakeHandler.Handle())

	systemIntakesHandler := handlers.NewSystemIntakesHandler(
		base,
		services.NewFetchSystemIntakesByEuaID(
			serviceConfig,
			store.FetchSystemIntakesByEuaID,
			services.NewAuthorizeFetchSystemIntakesByEuaID(),
		),
	)
	api.Handle("/system_intakes", systemIntakesHandler.Handle())

	businessCaseHandler := handlers.NewBusinessCaseHandler(
		base,
		services.NewFetchBusinessCaseByID(
			serviceConfig,
			store.FetchBusinessCaseByID,
			services.NewAuthorizeFetchBusinessCaseByID(),
		),
		services.NewCreateBusinessCase(
			serviceConfig,
			store.FetchSystemIntakeByID,
			services.NewAuthorizeCreateBusinessCase(s.logger),
			store.CreateBusinessCase,
		),
		services.NewUpdateBusinessCase(
			serviceConfig,
			store.FetchBusinessCaseByID,
			services.NewAuthorizeUpdateBusinessCase(s.logger),
			store.UpdateBusinessCase,
			emailClient.SendBusinessCaseSubmissionEmail,
		),
	)
	api.Handle("/business_case/{business_case_id}", businessCaseHandler.Handle())
	api.Handle("/business_case", businessCaseHandler.Handle())

	businessCasesHandler := handlers.NewBusinessCasesHandler(
		base,
		services.NewFetchBusinessCasesByEuaID(
			serviceConfig,
			store.FetchBusinessCasesByEuaID,
			services.NewAuthorizeFetchBusinessCasesByEuaID(),
		),
	)
	api.Handle("/business_cases", businessCasesHandler.Handle())

	metricsHandler := handlers.NewMetricsHandler(
		base,
		services.NewFetchMetrics(serviceConfig, store.FetchSystemIntakeMetrics),
	)
	api.Handle("/metrics", metricsHandler.Handle())

	s.router.PathPrefix("/").Handler(handlers.NewCatchAllHandler(
		base,
	).Handle())
}
