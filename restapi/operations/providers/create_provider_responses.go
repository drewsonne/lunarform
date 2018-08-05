// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/getlunaform/lunaform/models"
)

// CreateProviderCreatedCode is the HTTP code returned for type CreateProviderCreated
const CreateProviderCreatedCode int = 201

/*CreateProviderCreated OK

swagger:response createProviderCreated
*/
type CreateProviderCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ResourceTfProvider `json:"body,omitempty"`
}

// NewCreateProviderCreated creates CreateProviderCreated with default headers values
func NewCreateProviderCreated() *CreateProviderCreated {

	return &CreateProviderCreated{}
}

// WithPayload adds the payload to the create provider created response
func (o *CreateProviderCreated) WithPayload(payload *models.ResourceTfProvider) *CreateProviderCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create provider created response
func (o *CreateProviderCreated) SetPayload(payload *models.ResourceTfProvider) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProviderCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProviderBadRequestCode is the HTTP code returned for type CreateProviderBadRequest
const CreateProviderBadRequestCode int = 400

/*CreateProviderBadRequest Bad Request

swagger:response createProviderBadRequest
*/
type CreateProviderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewCreateProviderBadRequest creates CreateProviderBadRequest with default headers values
func NewCreateProviderBadRequest() *CreateProviderBadRequest {

	return &CreateProviderBadRequest{}
}

// WithPayload adds the payload to the create provider bad request response
func (o *CreateProviderBadRequest) WithPayload(payload *models.ServerError) *CreateProviderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create provider bad request response
func (o *CreateProviderBadRequest) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProviderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}