// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// PostStepReplyReader is a Reader for the PostStepReply structure.
type PostStepReplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStepReplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostStepReplyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStepReplyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostStepReplyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostStepReplyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostStepReplyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostStepReplyMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStepReplyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostStepReplyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostStepReplyNoContent creates a PostStepReplyNoContent with default headers values
func NewPostStepReplyNoContent() *PostStepReplyNoContent {
	return &PostStepReplyNoContent{}
}

/*PostStepReplyNoContent handles this case with default header values.

Success.
*/
type PostStepReplyNoContent struct {
}

func (o *PostStepReplyNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyNoContent ", 204)
}

func (o *PostStepReplyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostStepReplyBadRequest creates a PostStepReplyBadRequest with default headers values
func NewPostStepReplyBadRequest() *PostStepReplyBadRequest {
	return &PostStepReplyBadRequest{}
}

/*PostStepReplyBadRequest handles this case with default header values.

Error.
*/
type PostStepReplyBadRequest struct {
	Payload *models.Error
}

func (o *PostStepReplyBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyBadRequest  %+v", 400, o.Payload)
}

func (o *PostStepReplyBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStepReplyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyUnauthorized creates a PostStepReplyUnauthorized with default headers values
func NewPostStepReplyUnauthorized() *PostStepReplyUnauthorized {
	return &PostStepReplyUnauthorized{}
}

/*PostStepReplyUnauthorized handles this case with default header values.

Unauthorized.
*/
type PostStepReplyUnauthorized struct {
	Payload *models.InfraError
}

func (o *PostStepReplyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostStepReplyUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *PostStepReplyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyForbidden creates a PostStepReplyForbidden with default headers values
func NewPostStepReplyForbidden() *PostStepReplyForbidden {
	return &PostStepReplyForbidden{}
}

/*PostStepReplyForbidden handles this case with default header values.

Forbidden.
*/
type PostStepReplyForbidden struct {
	Payload *models.InfraError
}

func (o *PostStepReplyForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyForbidden  %+v", 403, o.Payload)
}

func (o *PostStepReplyForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *PostStepReplyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyNotFound creates a PostStepReplyNotFound with default headers values
func NewPostStepReplyNotFound() *PostStepReplyNotFound {
	return &PostStepReplyNotFound{}
}

/*PostStepReplyNotFound handles this case with default header values.

Error.
*/
type PostStepReplyNotFound struct {
	Payload *models.Error
}

func (o *PostStepReplyNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyNotFound  %+v", 404, o.Payload)
}

func (o *PostStepReplyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStepReplyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyMethodNotAllowed creates a PostStepReplyMethodNotAllowed with default headers values
func NewPostStepReplyMethodNotAllowed() *PostStepReplyMethodNotAllowed {
	return &PostStepReplyMethodNotAllowed{}
}

/*PostStepReplyMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type PostStepReplyMethodNotAllowed struct {
	Payload *models.Error
}

func (o *PostStepReplyMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostStepReplyMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStepReplyMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyInternalServerError creates a PostStepReplyInternalServerError with default headers values
func NewPostStepReplyInternalServerError() *PostStepReplyInternalServerError {
	return &PostStepReplyInternalServerError{}
}

/*PostStepReplyInternalServerError handles this case with default header values.

Error.
*/
type PostStepReplyInternalServerError struct {
	Payload *models.Error
}

func (o *PostStepReplyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStepReplyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStepReplyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStepReplyServiceUnavailable creates a PostStepReplyServiceUnavailable with default headers values
func NewPostStepReplyServiceUnavailable() *PostStepReplyServiceUnavailable {
	return &PostStepReplyServiceUnavailable{}
}

/*PostStepReplyServiceUnavailable handles this case with default header values.

Unavailable.
*/
type PostStepReplyServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostStepReplyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/instructions][%d] postStepReplyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostStepReplyServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStepReplyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
