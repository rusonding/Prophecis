// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetCurrentUserStoragePathOKCode is the HTTP code returned for type GetCurrentUserStoragePathOK
const GetCurrentUserStoragePathOKCode int = 200

/*GetCurrentUserStoragePathOK Detailed GroupStorage and GroupStorage information.

swagger:response getCurrentUserStoragePathOK
*/
type GetCurrentUserStoragePathOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewGetCurrentUserStoragePathOK creates GetCurrentUserStoragePathOK with default headers values
func NewGetCurrentUserStoragePathOK() *GetCurrentUserStoragePathOK {

	return &GetCurrentUserStoragePathOK{}
}

// WithPayload adds the payload to the get current user storage path o k response
func (o *GetCurrentUserStoragePathOK) WithPayload(payload *models.Result) *GetCurrentUserStoragePathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user storage path o k response
func (o *GetCurrentUserStoragePathOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserStoragePathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentUserStoragePathUnauthorizedCode is the HTTP code returned for type GetCurrentUserStoragePathUnauthorized
const GetCurrentUserStoragePathUnauthorizedCode int = 401

/*GetCurrentUserStoragePathUnauthorized Unauthorized

swagger:response getCurrentUserStoragePathUnauthorized
*/
type GetCurrentUserStoragePathUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCurrentUserStoragePathUnauthorized creates GetCurrentUserStoragePathUnauthorized with default headers values
func NewGetCurrentUserStoragePathUnauthorized() *GetCurrentUserStoragePathUnauthorized {

	return &GetCurrentUserStoragePathUnauthorized{}
}

// WithPayload adds the payload to the get current user storage path unauthorized response
func (o *GetCurrentUserStoragePathUnauthorized) WithPayload(payload *models.Error) *GetCurrentUserStoragePathUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user storage path unauthorized response
func (o *GetCurrentUserStoragePathUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserStoragePathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCurrentUserStoragePathNotFoundCode is the HTTP code returned for type GetCurrentUserStoragePathNotFound
const GetCurrentUserStoragePathNotFoundCode int = 404

/*GetCurrentUserStoragePathNotFound Model with the given ID not found.

swagger:response getCurrentUserStoragePathNotFound
*/
type GetCurrentUserStoragePathNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCurrentUserStoragePathNotFound creates GetCurrentUserStoragePathNotFound with default headers values
func NewGetCurrentUserStoragePathNotFound() *GetCurrentUserStoragePathNotFound {

	return &GetCurrentUserStoragePathNotFound{}
}

// WithPayload adds the payload to the get current user storage path not found response
func (o *GetCurrentUserStoragePathNotFound) WithPayload(payload *models.Error) *GetCurrentUserStoragePathNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user storage path not found response
func (o *GetCurrentUserStoragePathNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserStoragePathNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}