/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsPropertiesUpdateReader is a Reader for the WeaviateThingsPropertiesUpdate structure.
type WeaviateThingsPropertiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsPropertiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsPropertiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsPropertiesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsPropertiesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsPropertiesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsPropertiesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsPropertiesUpdateOK creates a WeaviateThingsPropertiesUpdateOK with default headers values
func NewWeaviateThingsPropertiesUpdateOK() *WeaviateThingsPropertiesUpdateOK {
	return &WeaviateThingsPropertiesUpdateOK{}
}

/*WeaviateThingsPropertiesUpdateOK handles this case with default header values.

Successfully replaced all the references (success is based on the behavior of the datastore).
*/
type WeaviateThingsPropertiesUpdateOK struct {
}

func (o *WeaviateThingsPropertiesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}/references/{propertyName}][%d] weaviateThingsPropertiesUpdateOK ", 200)
}

func (o *WeaviateThingsPropertiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesUpdateUnauthorized creates a WeaviateThingsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateThingsPropertiesUpdateUnauthorized() *WeaviateThingsPropertiesUpdateUnauthorized {
	return &WeaviateThingsPropertiesUpdateUnauthorized{}
}

/*WeaviateThingsPropertiesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsPropertiesUpdateUnauthorized struct {
}

func (o *WeaviateThingsPropertiesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}/references/{propertyName}][%d] weaviateThingsPropertiesUpdateUnauthorized ", 401)
}

func (o *WeaviateThingsPropertiesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesUpdateForbidden creates a WeaviateThingsPropertiesUpdateForbidden with default headers values
func NewWeaviateThingsPropertiesUpdateForbidden() *WeaviateThingsPropertiesUpdateForbidden {
	return &WeaviateThingsPropertiesUpdateForbidden{}
}

/*WeaviateThingsPropertiesUpdateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateThingsPropertiesUpdateForbidden struct {
}

func (o *WeaviateThingsPropertiesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}/references/{propertyName}][%d] weaviateThingsPropertiesUpdateForbidden ", 403)
}

func (o *WeaviateThingsPropertiesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesUpdateUnprocessableEntity creates a WeaviateThingsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateThingsPropertiesUpdateUnprocessableEntity() *WeaviateThingsPropertiesUpdateUnprocessableEntity {
	return &WeaviateThingsPropertiesUpdateUnprocessableEntity{}
}

/*WeaviateThingsPropertiesUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type WeaviateThingsPropertiesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsPropertiesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}/references/{propertyName}][%d] weaviateThingsPropertiesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsPropertiesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsPropertiesUpdateInternalServerError creates a WeaviateThingsPropertiesUpdateInternalServerError with default headers values
func NewWeaviateThingsPropertiesUpdateInternalServerError() *WeaviateThingsPropertiesUpdateInternalServerError {
	return &WeaviateThingsPropertiesUpdateInternalServerError{}
}

/*WeaviateThingsPropertiesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsPropertiesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsPropertiesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /things/{thingId}/references/{propertyName}][%d] weaviateThingsPropertiesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsPropertiesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
