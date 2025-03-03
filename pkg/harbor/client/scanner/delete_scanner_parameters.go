// Code generated by go-swagger; DO NOT EDIT.

package scanner

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
)

// NewDeleteScannerParams creates a new DeleteScannerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteScannerParams() *DeleteScannerParams {
	return &DeleteScannerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScannerParamsWithTimeout creates a new DeleteScannerParams object
// with the ability to set a timeout on a request.
func NewDeleteScannerParamsWithTimeout(timeout time.Duration) *DeleteScannerParams {
	return &DeleteScannerParams{
		timeout: timeout,
	}
}

// NewDeleteScannerParamsWithContext creates a new DeleteScannerParams object
// with the ability to set a context for a request.
func NewDeleteScannerParamsWithContext(ctx context.Context) *DeleteScannerParams {
	return &DeleteScannerParams{
		Context: ctx,
	}
}

// NewDeleteScannerParamsWithHTTPClient creates a new DeleteScannerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteScannerParamsWithHTTPClient(client *http.Client) *DeleteScannerParams {
	return &DeleteScannerParams{
		HTTPClient: client,
	}
}

/* DeleteScannerParams contains all the parameters to send to the API endpoint
   for the delete scanner operation.

   Typically these are written to a http.Request.
*/
type DeleteScannerParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* RegistrationID.

	   The scanner registration identifier.
	*/
	RegistrationID string `js:"registrationID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the delete scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScannerParams) WithDefaults() *DeleteScannerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete scanner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScannerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete scanner params
func (o *DeleteScannerParams) WithTimeout(timeout time.Duration) *DeleteScannerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scanner params
func (o *DeleteScannerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scanner params
func (o *DeleteScannerParams) WithContext(ctx context.Context) *DeleteScannerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scanner params
func (o *DeleteScannerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scanner params
func (o *DeleteScannerParams) WithHTTPClient(client *http.Client) *DeleteScannerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scanner params
func (o *DeleteScannerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the delete scanner params
func (o *DeleteScannerParams) WithXRequestID(xRequestID *string) *DeleteScannerParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete scanner params
func (o *DeleteScannerParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRegistrationID adds the registrationID to the delete scanner params
func (o *DeleteScannerParams) WithRegistrationID(registrationID string) *DeleteScannerParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the delete scanner params
func (o *DeleteScannerParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScannerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
