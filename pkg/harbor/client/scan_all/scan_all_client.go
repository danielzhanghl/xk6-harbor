// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the scan all client
type API interface {
	/*
	   CreateScanAllSchedule creates a schedule or a manual trigger for the scan all job

	   This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor.*/
	CreateScanAllSchedule(ctx context.Context, params *CreateScanAllScheduleParams) (*CreateScanAllScheduleCreated, error)
	/*
	   GetLatestScanAllMetrics gets the metrics of the latest scan all process

	   Get the metrics of the latest scan all process*/
	GetLatestScanAllMetrics(ctx context.Context, params *GetLatestScanAllMetricsParams) (*GetLatestScanAllMetricsOK, error)
	/*
	   GetLatestScheduledScanAllMetrics gets the metrics of the latest scheduled scan all process

	   Get the metrics of the latest scheduled scan all process*/
	GetLatestScheduledScanAllMetrics(ctx context.Context, params *GetLatestScheduledScanAllMetricsParams) (*GetLatestScheduledScanAllMetricsOK, error)
	/*
	   GetScanAllSchedule gets scan all s schedule

	   This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.*/
	GetScanAllSchedule(ctx context.Context, params *GetScanAllScheduleParams) (*GetScanAllScheduleOK, error)
	/*
	   UpdateScanAllSchedule updates scan all s schedule

	   This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor.*/
	UpdateScanAllSchedule(ctx context.Context, params *UpdateScanAllScheduleParams) (*UpdateScanAllScheduleOK, error)
}

// New creates a new scan all API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for scan all API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateScanAllSchedule creates a schedule or a manual trigger for the scan all job

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor.
*/
func (a *Client) CreateScanAllSchedule(ctx context.Context, params *CreateScanAllScheduleParams) (*CreateScanAllScheduleCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createScanAllSchedule",
		Method:             "POST",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateScanAllScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateScanAllScheduleCreated), nil

}

/*
GetLatestScanAllMetrics gets the metrics of the latest scan all process

Get the metrics of the latest scan all process
*/
func (a *Client) GetLatestScanAllMetrics(ctx context.Context, params *GetLatestScanAllMetricsParams) (*GetLatestScanAllMetricsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLatestScanAllMetrics",
		Method:             "GET",
		PathPattern:        "/scans/all/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLatestScanAllMetricsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLatestScanAllMetricsOK), nil

}

/*
GetLatestScheduledScanAllMetrics gets the metrics of the latest scheduled scan all process

Get the metrics of the latest scheduled scan all process
*/
func (a *Client) GetLatestScheduledScanAllMetrics(ctx context.Context, params *GetLatestScheduledScanAllMetricsParams) (*GetLatestScheduledScanAllMetricsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLatestScheduledScanAllMetrics",
		Method:             "GET",
		PathPattern:        "/scans/schedule/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLatestScheduledScanAllMetricsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLatestScheduledScanAllMetricsOK), nil

}

/*
GetScanAllSchedule gets scan all s schedule

This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.
*/
func (a *Client) GetScanAllSchedule(ctx context.Context, params *GetScanAllScheduleParams) (*GetScanAllScheduleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScanAllSchedule",
		Method:             "GET",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetScanAllScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScanAllScheduleOK), nil

}

/*
UpdateScanAllSchedule updates scan all s schedule

This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor.
*/
func (a *Client) UpdateScanAllSchedule(ctx context.Context, params *UpdateScanAllScheduleParams) (*UpdateScanAllScheduleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateScanAllSchedule",
		Method:             "PUT",
		PathPattern:        "/system/scanAll/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateScanAllScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateScanAllScheduleOK), nil

}
