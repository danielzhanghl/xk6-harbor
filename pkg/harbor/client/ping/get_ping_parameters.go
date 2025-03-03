// Code generated by go-swagger; DO NOT EDIT.

package ping

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

// NewGetPingParams creates a new GetPingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPingParams() *GetPingParams {
	return &GetPingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPingParamsWithTimeout creates a new GetPingParams object
// with the ability to set a timeout on a request.
func NewGetPingParamsWithTimeout(timeout time.Duration) *GetPingParams {
	return &GetPingParams{
		timeout: timeout,
	}
}

// NewGetPingParamsWithContext creates a new GetPingParams object
// with the ability to set a context for a request.
func NewGetPingParamsWithContext(ctx context.Context) *GetPingParams {
	return &GetPingParams{
		Context: ctx,
	}
}

// NewGetPingParamsWithHTTPClient creates a new GetPingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPingParamsWithHTTPClient(client *http.Client) *GetPingParams {
	return &GetPingParams{
		HTTPClient: client,
	}
}

/* GetPingParams contains all the parameters to send to the API endpoint
   for the get ping operation.

   Typically these are written to a http.Request.
*/
type GetPingParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get ping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPingParams) WithDefaults() *GetPingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ping params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ping params
func (o *GetPingParams) WithTimeout(timeout time.Duration) *GetPingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ping params
func (o *GetPingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ping params
func (o *GetPingParams) WithContext(ctx context.Context) *GetPingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ping params
func (o *GetPingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ping params
func (o *GetPingParams) WithHTTPClient(client *http.Client) *GetPingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ping params
func (o *GetPingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get ping params
func (o *GetPingParams) WithXRequestID(xRequestID *string) *GetPingParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get ping params
func (o *GetPingParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
