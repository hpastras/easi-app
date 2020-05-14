// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/gen/models"
)

// IntakegovernanceidGET5Reader is a Reader for the IntakegovernanceidGET5 structure.
type IntakegovernanceidGET5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntakegovernanceidGET5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntakegovernanceidGET5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIntakegovernanceidGET5BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIntakegovernanceidGET5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIntakegovernanceidGET5InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntakegovernanceidGET5OK creates a IntakegovernanceidGET5OK with default headers values
func NewIntakegovernanceidGET5OK() *IntakegovernanceidGET5OK {
	return &IntakegovernanceidGET5OK{}
}

/*IntakegovernanceidGET5OK handles this case with default header values.

OK
*/
type IntakegovernanceidGET5OK struct {
	Payload *models.IntakegovernanceidGETResponse
}

func (o *IntakegovernanceidGET5OK) Error() string {
	return fmt.Sprintf("[GET /intake/governance/{id}][%d] intakegovernanceidGET5OK  %+v", 200, o.Payload)
}

func (o *IntakegovernanceidGET5OK) GetPayload() *models.IntakegovernanceidGETResponse {
	return o.Payload
}

func (o *IntakegovernanceidGET5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntakegovernanceidGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntakegovernanceidGET5BadRequest creates a IntakegovernanceidGET5BadRequest with default headers values
func NewIntakegovernanceidGET5BadRequest() *IntakegovernanceidGET5BadRequest {
	return &IntakegovernanceidGET5BadRequest{}
}

/*IntakegovernanceidGET5BadRequest handles this case with default header values.

Bad Request
*/
type IntakegovernanceidGET5BadRequest struct {
}

func (o *IntakegovernanceidGET5BadRequest) Error() string {
	return fmt.Sprintf("[GET /intake/governance/{id}][%d] intakegovernanceidGET5BadRequest ", 400)
}

func (o *IntakegovernanceidGET5BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIntakegovernanceidGET5Unauthorized creates a IntakegovernanceidGET5Unauthorized with default headers values
func NewIntakegovernanceidGET5Unauthorized() *IntakegovernanceidGET5Unauthorized {
	return &IntakegovernanceidGET5Unauthorized{}
}

/*IntakegovernanceidGET5Unauthorized handles this case with default header values.

Access Denied
*/
type IntakegovernanceidGET5Unauthorized struct {
}

func (o *IntakegovernanceidGET5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /intake/governance/{id}][%d] intakegovernanceidGET5Unauthorized ", 401)
}

func (o *IntakegovernanceidGET5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIntakegovernanceidGET5InternalServerError creates a IntakegovernanceidGET5InternalServerError with default headers values
func NewIntakegovernanceidGET5InternalServerError() *IntakegovernanceidGET5InternalServerError {
	return &IntakegovernanceidGET5InternalServerError{}
}

/*IntakegovernanceidGET5InternalServerError handles this case with default header values.

Internal Server Error
*/
type IntakegovernanceidGET5InternalServerError struct {
}

func (o *IntakegovernanceidGET5InternalServerError) Error() string {
	return fmt.Sprintf("[GET /intake/governance/{id}][%d] intakegovernanceidGET5InternalServerError ", 500)
}

func (o *IntakegovernanceidGET5InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}