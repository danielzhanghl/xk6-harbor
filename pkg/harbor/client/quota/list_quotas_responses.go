// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// ListQuotasReader is a Reader for the ListQuotas structure.
type ListQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListQuotasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListQuotasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListQuotasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListQuotasOK creates a ListQuotasOK with default headers values
func NewListQuotasOK() *ListQuotasOK {
	return &ListQuotasOK{}
}

/* ListQuotasOK describes a response with status code 200, with default header values.

Successfully retrieved the quotas.
*/
type ListQuotasOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of access logs
	 */
	XTotalCount int64

	Payload []*models.Quota
}

func (o *ListQuotasOK) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasOK  %+v", 200, o.Payload)
}
func (o *ListQuotasOK) GetPayload() []*models.Quota {
	return o.Payload
}

func (o *ListQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListQuotasUnauthorized creates a ListQuotasUnauthorized with default headers values
func NewListQuotasUnauthorized() *ListQuotasUnauthorized {
	return &ListQuotasUnauthorized{}
}

/* ListQuotasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListQuotasUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListQuotasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasUnauthorized  %+v", 401, o.Payload)
}
func (o *ListQuotasUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListQuotasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListQuotasForbidden creates a ListQuotasForbidden with default headers values
func NewListQuotasForbidden() *ListQuotasForbidden {
	return &ListQuotasForbidden{}
}

/* ListQuotasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListQuotasForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListQuotasForbidden) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasForbidden  %+v", 403, o.Payload)
}
func (o *ListQuotasForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListQuotasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListQuotasInternalServerError creates a ListQuotasInternalServerError with default headers values
func NewListQuotasInternalServerError() *ListQuotasInternalServerError {
	return &ListQuotasInternalServerError{}
}

/* ListQuotasInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListQuotasInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListQuotasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quotas][%d] listQuotasInternalServerError  %+v", 500, o.Payload)
}
func (o *ListQuotasInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListQuotasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
