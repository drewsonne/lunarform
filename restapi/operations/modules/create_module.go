// Code generated by go-swagger; DO NOT EDIT.

package modules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// CreateModuleHandlerFunc turns a function with the right signature into a create module handler
type CreateModuleHandlerFunc func(CreateModuleParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateModuleHandlerFunc) Handle(params CreateModuleParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// CreateModuleHandler interface for that can handle valid create module params
type CreateModuleHandler interface {
	Handle(CreateModuleParams, *models.ResourceAuthUser) middleware.Responder
}

// NewCreateModule creates a new http.Handler for the create module operation
func NewCreateModule(ctx *middleware.Context, handler CreateModuleHandler) *CreateModule {
	return &CreateModule{Context: ctx, Handler: handler}
}

/*CreateModule swagger:route POST /tf/modules modules createModule

Upload a Terraform module

*/
type CreateModule struct {
	Context *middleware.Context
	Handler CreateModuleHandler
}

func (o *CreateModule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateModuleParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.ResourceAuthUser
	if uprinc != nil {
		principal = uprinc.(*models.ResourceAuthUser) // this is really a models.ResourceAuthUser, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
