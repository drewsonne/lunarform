// Code generated by go-swagger; DO NOT EDIT.

package stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/getlunaform/lunaform/models"
)

// ListStacksOKCode is the HTTP code returned for type ListStacksOK
const ListStacksOKCode int = 200

/*ListStacksOK OK

swagger:response listStacksOK
*/
type ListStacksOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseListTfStacks `json:"body,omitempty"`
}

// NewListStacksOK creates ListStacksOK with default headers values
func NewListStacksOK() *ListStacksOK {

	return &ListStacksOK{}
}

// WithPayload adds the payload to the list stacks o k response
func (o *ListStacksOK) WithPayload(payload *models.ResponseListTfStacks) *ListStacksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list stacks o k response
func (o *ListStacksOK) SetPayload(payload *models.ResponseListTfStacks) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStacksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStacksNotFoundCode is the HTTP code returned for type ListStacksNotFound
const ListStacksNotFoundCode int = 404

/*ListStacksNotFound Not Found

swagger:response listStacksNotFound
*/
type ListStacksNotFound struct {
}

// NewListStacksNotFound creates ListStacksNotFound with default headers values
func NewListStacksNotFound() *ListStacksNotFound {

	return &ListStacksNotFound{}
}

// WriteResponse to the client
func (o *ListStacksNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ListStacksInternalServerErrorCode is the HTTP code returned for type ListStacksInternalServerError
const ListStacksInternalServerErrorCode int = 500

/*ListStacksInternalServerError Internal Server Error

swagger:response listStacksInternalServerError
*/
type ListStacksInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewListStacksInternalServerError creates ListStacksInternalServerError with default headers values
func NewListStacksInternalServerError() *ListStacksInternalServerError {

	return &ListStacksInternalServerError{}
}

// WithPayload adds the payload to the list stacks internal server error response
func (o *ListStacksInternalServerError) WithPayload(payload *models.ServerError) *ListStacksInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list stacks internal server error response
func (o *ListStacksInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStacksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
