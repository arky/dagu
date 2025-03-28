// Code generated by go-swagger; DO NOT EDIT.

package dags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dagu-org/dagu/internal/frontend/gen/models"
)

// ListDAGsOKCode is the HTTP code returned for type ListDAGsOK
const ListDAGsOKCode int = 200

/*
ListDAGsOK A successful response.

swagger:response listDAGsOK
*/
type ListDAGsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListDAGsResponse `json:"body,omitempty"`
}

// NewListDAGsOK creates ListDAGsOK with default headers values
func NewListDAGsOK() *ListDAGsOK {

	return &ListDAGsOK{}
}

// WithPayload adds the payload to the list d a gs o k response
func (o *ListDAGsOK) WithPayload(payload *models.ListDAGsResponse) *ListDAGsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list d a gs o k response
func (o *ListDAGsOK) SetPayload(payload *models.ListDAGsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListDAGsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListDAGsDefault Generic error response.

swagger:response listDAGsDefault
*/
type ListDAGsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListDAGsDefault creates ListDAGsDefault with default headers values
func NewListDAGsDefault(code int) *ListDAGsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListDAGsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list d a gs default response
func (o *ListDAGsDefault) WithStatusCode(code int) *ListDAGsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list d a gs default response
func (o *ListDAGsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list d a gs default response
func (o *ListDAGsDefault) WithPayload(payload *models.Error) *ListDAGsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list d a gs default response
func (o *ListDAGsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListDAGsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
