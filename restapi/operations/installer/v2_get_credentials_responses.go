// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2GetCredentialsOKCode is the HTTP code returned for type V2GetCredentialsOK
const V2GetCredentialsOKCode int = 200

/*
V2GetCredentialsOK Success.

swagger:response v2GetCredentialsOK
*/
type V2GetCredentialsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Credentials `json:"body,omitempty"`
}

// NewV2GetCredentialsOK creates V2GetCredentialsOK with default headers values
func NewV2GetCredentialsOK() *V2GetCredentialsOK {

	return &V2GetCredentialsOK{}
}

// WithPayload adds the payload to the v2 get credentials o k response
func (o *V2GetCredentialsOK) WithPayload(payload *models.Credentials) *V2GetCredentialsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials o k response
func (o *V2GetCredentialsOK) SetPayload(payload *models.Credentials) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsUnauthorizedCode is the HTTP code returned for type V2GetCredentialsUnauthorized
const V2GetCredentialsUnauthorizedCode int = 401

/*
V2GetCredentialsUnauthorized Unauthorized.

swagger:response v2GetCredentialsUnauthorized
*/
type V2GetCredentialsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetCredentialsUnauthorized creates V2GetCredentialsUnauthorized with default headers values
func NewV2GetCredentialsUnauthorized() *V2GetCredentialsUnauthorized {

	return &V2GetCredentialsUnauthorized{}
}

// WithPayload adds the payload to the v2 get credentials unauthorized response
func (o *V2GetCredentialsUnauthorized) WithPayload(payload *models.InfraError) *V2GetCredentialsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials unauthorized response
func (o *V2GetCredentialsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsForbiddenCode is the HTTP code returned for type V2GetCredentialsForbidden
const V2GetCredentialsForbiddenCode int = 403

/*
V2GetCredentialsForbidden Forbidden.

swagger:response v2GetCredentialsForbidden
*/
type V2GetCredentialsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2GetCredentialsForbidden creates V2GetCredentialsForbidden with default headers values
func NewV2GetCredentialsForbidden() *V2GetCredentialsForbidden {

	return &V2GetCredentialsForbidden{}
}

// WithPayload adds the payload to the v2 get credentials forbidden response
func (o *V2GetCredentialsForbidden) WithPayload(payload *models.InfraError) *V2GetCredentialsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials forbidden response
func (o *V2GetCredentialsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsNotFoundCode is the HTTP code returned for type V2GetCredentialsNotFound
const V2GetCredentialsNotFoundCode int = 404

/*
V2GetCredentialsNotFound Error.

swagger:response v2GetCredentialsNotFound
*/
type V2GetCredentialsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetCredentialsNotFound creates V2GetCredentialsNotFound with default headers values
func NewV2GetCredentialsNotFound() *V2GetCredentialsNotFound {

	return &V2GetCredentialsNotFound{}
}

// WithPayload adds the payload to the v2 get credentials not found response
func (o *V2GetCredentialsNotFound) WithPayload(payload *models.Error) *V2GetCredentialsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials not found response
func (o *V2GetCredentialsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsMethodNotAllowedCode is the HTTP code returned for type V2GetCredentialsMethodNotAllowed
const V2GetCredentialsMethodNotAllowedCode int = 405

/*
V2GetCredentialsMethodNotAllowed Method Not Allowed.

swagger:response v2GetCredentialsMethodNotAllowed
*/
type V2GetCredentialsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetCredentialsMethodNotAllowed creates V2GetCredentialsMethodNotAllowed with default headers values
func NewV2GetCredentialsMethodNotAllowed() *V2GetCredentialsMethodNotAllowed {

	return &V2GetCredentialsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 get credentials method not allowed response
func (o *V2GetCredentialsMethodNotAllowed) WithPayload(payload *models.Error) *V2GetCredentialsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials method not allowed response
func (o *V2GetCredentialsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsConflictCode is the HTTP code returned for type V2GetCredentialsConflict
const V2GetCredentialsConflictCode int = 409

/*
V2GetCredentialsConflict Error.

swagger:response v2GetCredentialsConflict
*/
type V2GetCredentialsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetCredentialsConflict creates V2GetCredentialsConflict with default headers values
func NewV2GetCredentialsConflict() *V2GetCredentialsConflict {

	return &V2GetCredentialsConflict{}
}

// WithPayload adds the payload to the v2 get credentials conflict response
func (o *V2GetCredentialsConflict) WithPayload(payload *models.Error) *V2GetCredentialsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials conflict response
func (o *V2GetCredentialsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2GetCredentialsInternalServerErrorCode is the HTTP code returned for type V2GetCredentialsInternalServerError
const V2GetCredentialsInternalServerErrorCode int = 500

/*
V2GetCredentialsInternalServerError Error.

swagger:response v2GetCredentialsInternalServerError
*/
type V2GetCredentialsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2GetCredentialsInternalServerError creates V2GetCredentialsInternalServerError with default headers values
func NewV2GetCredentialsInternalServerError() *V2GetCredentialsInternalServerError {

	return &V2GetCredentialsInternalServerError{}
}

// WithPayload adds the payload to the v2 get credentials internal server error response
func (o *V2GetCredentialsInternalServerError) WithPayload(payload *models.Error) *V2GetCredentialsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 get credentials internal server error response
func (o *V2GetCredentialsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2GetCredentialsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
