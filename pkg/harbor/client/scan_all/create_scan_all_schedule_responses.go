// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// CreateScanAllScheduleReader is a Reader for the CreateScanAllSchedule structure.
type CreateScanAllScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScanAllScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateScanAllScheduleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateScanAllScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateScanAllScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateScanAllScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateScanAllScheduleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewCreateScanAllSchedulePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateScanAllScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateScanAllScheduleCreated creates a CreateScanAllScheduleCreated with default headers values
func NewCreateScanAllScheduleCreated() *CreateScanAllScheduleCreated {
	return &CreateScanAllScheduleCreated{}
}

/* CreateScanAllScheduleCreated describes a response with status code 201, with default header values.

Created
*/
type CreateScanAllScheduleCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateScanAllScheduleCreated) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleCreated ", 201)
}

func (o *CreateScanAllScheduleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateScanAllScheduleBadRequest creates a CreateScanAllScheduleBadRequest with default headers values
func NewCreateScanAllScheduleBadRequest() *CreateScanAllScheduleBadRequest {
	return &CreateScanAllScheduleBadRequest{}
}

/* CreateScanAllScheduleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateScanAllScheduleBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateScanAllScheduleBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScanAllScheduleUnauthorized creates a CreateScanAllScheduleUnauthorized with default headers values
func NewCreateScanAllScheduleUnauthorized() *CreateScanAllScheduleUnauthorized {
	return &CreateScanAllScheduleUnauthorized{}
}

/* CreateScanAllScheduleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateScanAllScheduleUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateScanAllScheduleUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScanAllScheduleForbidden creates a CreateScanAllScheduleForbidden with default headers values
func NewCreateScanAllScheduleForbidden() *CreateScanAllScheduleForbidden {
	return &CreateScanAllScheduleForbidden{}
}

/* CreateScanAllScheduleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateScanAllScheduleForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleForbidden  %+v", 403, o.Payload)
}
func (o *CreateScanAllScheduleForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScanAllScheduleConflict creates a CreateScanAllScheduleConflict with default headers values
func NewCreateScanAllScheduleConflict() *CreateScanAllScheduleConflict {
	return &CreateScanAllScheduleConflict{}
}

/* CreateScanAllScheduleConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateScanAllScheduleConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllScheduleConflict) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleConflict  %+v", 409, o.Payload)
}
func (o *CreateScanAllScheduleConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllScheduleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScanAllSchedulePreconditionFailed creates a CreateScanAllSchedulePreconditionFailed with default headers values
func NewCreateScanAllSchedulePreconditionFailed() *CreateScanAllSchedulePreconditionFailed {
	return &CreateScanAllSchedulePreconditionFailed{}
}

/* CreateScanAllSchedulePreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type CreateScanAllSchedulePreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllSchedulePreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllSchedulePreconditionFailed  %+v", 412, o.Payload)
}
func (o *CreateScanAllSchedulePreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllSchedulePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScanAllScheduleInternalServerError creates a CreateScanAllScheduleInternalServerError with default headers values
func NewCreateScanAllScheduleInternalServerError() *CreateScanAllScheduleInternalServerError {
	return &CreateScanAllScheduleInternalServerError{}
}

/* CreateScanAllScheduleInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateScanAllScheduleInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateScanAllScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/scanAll/schedule][%d] createScanAllScheduleInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateScanAllScheduleInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateScanAllScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
