// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/cedareasi/gen/models"
)

// NewIntakebusinessCasePOST7Params creates a new IntakebusinessCasePOST7Params object
// with the default values initialized.
func NewIntakebusinessCasePOST7Params() *IntakebusinessCasePOST7Params {
	var ()
	return &IntakebusinessCasePOST7Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntakebusinessCasePOST7ParamsWithTimeout creates a new IntakebusinessCasePOST7Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntakebusinessCasePOST7ParamsWithTimeout(timeout time.Duration) *IntakebusinessCasePOST7Params {
	var ()
	return &IntakebusinessCasePOST7Params{

		timeout: timeout,
	}
}

// NewIntakebusinessCasePOST7ParamsWithContext creates a new IntakebusinessCasePOST7Params object
// with the default values initialized, and the ability to set a context for a request
func NewIntakebusinessCasePOST7ParamsWithContext(ctx context.Context) *IntakebusinessCasePOST7Params {
	var ()
	return &IntakebusinessCasePOST7Params{

		Context: ctx,
	}
}

// NewIntakebusinessCasePOST7ParamsWithHTTPClient creates a new IntakebusinessCasePOST7Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntakebusinessCasePOST7ParamsWithHTTPClient(client *http.Client) *IntakebusinessCasePOST7Params {
	var ()
	return &IntakebusinessCasePOST7Params{
		HTTPClient: client,
	}
}

/*IntakebusinessCasePOST7Params contains all the parameters to send to the API endpoint
for the intakebusiness case p o s t 7 operation typically these are written to a http.Request
*/
type IntakebusinessCasePOST7Params struct {

	/*Body*/
	Body *models.Intake2

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) WithTimeout(timeout time.Duration) *IntakebusinessCasePOST7Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) WithContext(ctx context.Context) *IntakebusinessCasePOST7Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) WithHTTPClient(client *http.Client) *IntakebusinessCasePOST7Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) WithBody(body *models.Intake2) *IntakebusinessCasePOST7Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the intakebusiness case p o s t 7 params
func (o *IntakebusinessCasePOST7Params) SetBody(body *models.Intake2) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IntakebusinessCasePOST7Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
