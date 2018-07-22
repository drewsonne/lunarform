// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/getlunaform/lunaform/models"
)

// ListResourcesOKCode is the HTTP code returned for type ListResourcesOK
const ListResourcesOKCode int = 200

/*ListResourcesOK OK

swagger:response listResourcesOK
*/
type ListResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseListResources `json:"body,omitempty"`
}

// NewListResourcesOK creates ListResourcesOK with default headers values
func NewListResourcesOK() *ListResourcesOK {

	return &ListResourcesOK{}
}

// WithPayload adds the payload to the list resources o k response
func (o *ListResourcesOK) WithPayload(payload *models.ResponseListResources) *ListResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list resources o k response
func (o *ListResourcesOK) SetPayload(payload *models.ResponseListResources) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}