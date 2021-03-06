// Code generated by go-swagger; DO NOT EDIT.

package modules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteModuleParams creates a new DeleteModuleParams object
// with the default values initialized.
func NewDeleteModuleParams() *DeleteModuleParams {
	var ()
	return &DeleteModuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModuleParamsWithTimeout creates a new DeleteModuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteModuleParamsWithTimeout(timeout time.Duration) *DeleteModuleParams {
	var ()
	return &DeleteModuleParams{

		timeout: timeout,
	}
}

// NewDeleteModuleParamsWithContext creates a new DeleteModuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteModuleParamsWithContext(ctx context.Context) *DeleteModuleParams {
	var ()
	return &DeleteModuleParams{

		Context: ctx,
	}
}

// NewDeleteModuleParamsWithHTTPClient creates a new DeleteModuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteModuleParamsWithHTTPClient(client *http.Client) *DeleteModuleParams {
	var ()
	return &DeleteModuleParams{
		HTTPClient: client,
	}
}

/*DeleteModuleParams contains all the parameters to send to the API endpoint
for the delete module operation typically these are written to a http.Request
*/
type DeleteModuleParams struct {

	/*ID
	  Unique identifier for this module

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete module params
func (o *DeleteModuleParams) WithTimeout(timeout time.Duration) *DeleteModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete module params
func (o *DeleteModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete module params
func (o *DeleteModuleParams) WithContext(ctx context.Context) *DeleteModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete module params
func (o *DeleteModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete module params
func (o *DeleteModuleParams) WithHTTPClient(client *http.Client) *DeleteModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete module params
func (o *DeleteModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete module params
func (o *DeleteModuleParams) WithID(id string) *DeleteModuleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete module params
func (o *DeleteModuleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
