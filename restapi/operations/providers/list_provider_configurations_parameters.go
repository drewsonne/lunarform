// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListProviderConfigurationsParams creates a new ListProviderConfigurationsParams object
// no default values defined in spec.
func NewListProviderConfigurationsParams() ListProviderConfigurationsParams {

	return ListProviderConfigurationsParams{}
}

// ListProviderConfigurationsParams contains all the bound params for the list provider configurations operation
// typically these are obtained from a http.Request
//
// swagger:parameters list-provider-configurations
type ListProviderConfigurationsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Terraform Provider Name
	  Required: true
	  In: path
	*/
	ProviderName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListProviderConfigurationsParams() beforehand.
func (o *ListProviderConfigurationsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProviderName, rhkProviderName, _ := route.Params.GetOK("provider-name")
	if err := o.bindProviderName(rProviderName, rhkProviderName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProviderName binds and validates parameter ProviderName from path.
func (o *ListProviderConfigurationsParams) bindProviderName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProviderName = raw

	return nil
}
