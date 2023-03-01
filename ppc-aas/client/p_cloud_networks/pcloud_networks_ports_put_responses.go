// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudNetworksPortsPutReader is a Reader for the PcloudNetworksPortsPut structure.
type PcloudNetworksPortsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksPortsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksPortsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksPortsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksPortsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksPortsPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudNetworksPortsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksPortsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudNetworksPortsPutOK creates a PcloudNetworksPortsPutOK with default headers values
func NewPcloudNetworksPortsPutOK() *PcloudNetworksPortsPutOK {
	return &PcloudNetworksPortsPutOK{}
}

/*
PcloudNetworksPortsPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksPortsPutOK struct {
	Payload *models.NetworkPort
}

// IsSuccess returns true when this pcloud networks ports put o k response has a 2xx status code
func (o *PcloudNetworksPortsPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks ports put o k response has a 3xx status code
func (o *PcloudNetworksPortsPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put o k response has a 4xx status code
func (o *PcloudNetworksPortsPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports put o k response has a 5xx status code
func (o *PcloudNetworksPortsPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put o k response a status code equal to that given
func (o *PcloudNetworksPortsPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudNetworksPortsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsPutOK) GetPayload() *models.NetworkPort {
	return o.Payload
}

func (o *PcloudNetworksPortsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutBadRequest creates a PcloudNetworksPortsPutBadRequest with default headers values
func NewPcloudNetworksPortsPutBadRequest() *PcloudNetworksPortsPutBadRequest {
	return &PcloudNetworksPortsPutBadRequest{}
}

/*
PcloudNetworksPortsPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksPortsPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put bad request response has a 2xx status code
func (o *PcloudNetworksPortsPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put bad request response has a 3xx status code
func (o *PcloudNetworksPortsPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put bad request response has a 4xx status code
func (o *PcloudNetworksPortsPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports put bad request response has a 5xx status code
func (o *PcloudNetworksPortsPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put bad request response a status code equal to that given
func (o *PcloudNetworksPortsPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudNetworksPortsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPortsPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPortsPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutUnauthorized creates a PcloudNetworksPortsPutUnauthorized with default headers values
func NewPcloudNetworksPortsPutUnauthorized() *PcloudNetworksPortsPutUnauthorized {
	return &PcloudNetworksPortsPutUnauthorized{}
}

/*
PcloudNetworksPortsPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksPortsPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put unauthorized response has a 2xx status code
func (o *PcloudNetworksPortsPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put unauthorized response has a 3xx status code
func (o *PcloudNetworksPortsPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put unauthorized response has a 4xx status code
func (o *PcloudNetworksPortsPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports put unauthorized response has a 5xx status code
func (o *PcloudNetworksPortsPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put unauthorized response a status code equal to that given
func (o *PcloudNetworksPortsPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudNetworksPortsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutForbidden creates a PcloudNetworksPortsPutForbidden with default headers values
func NewPcloudNetworksPortsPutForbidden() *PcloudNetworksPortsPutForbidden {
	return &PcloudNetworksPortsPutForbidden{}
}

/*
PcloudNetworksPortsPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksPortsPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put forbidden response has a 2xx status code
func (o *PcloudNetworksPortsPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put forbidden response has a 3xx status code
func (o *PcloudNetworksPortsPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put forbidden response has a 4xx status code
func (o *PcloudNetworksPortsPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports put forbidden response has a 5xx status code
func (o *PcloudNetworksPortsPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put forbidden response a status code equal to that given
func (o *PcloudNetworksPortsPutForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudNetworksPortsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsPutForbidden) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutNotFound creates a PcloudNetworksPortsPutNotFound with default headers values
func NewPcloudNetworksPortsPutNotFound() *PcloudNetworksPortsPutNotFound {
	return &PcloudNetworksPortsPutNotFound{}
}

/*
PcloudNetworksPortsPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksPortsPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put not found response has a 2xx status code
func (o *PcloudNetworksPortsPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put not found response has a 3xx status code
func (o *PcloudNetworksPortsPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put not found response has a 4xx status code
func (o *PcloudNetworksPortsPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports put not found response has a 5xx status code
func (o *PcloudNetworksPortsPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put not found response a status code equal to that given
func (o *PcloudNetworksPortsPutNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudNetworksPortsPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsPutNotFound) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutUnprocessableEntity creates a PcloudNetworksPortsPutUnprocessableEntity with default headers values
func NewPcloudNetworksPortsPutUnprocessableEntity() *PcloudNetworksPortsPutUnprocessableEntity {
	return &PcloudNetworksPortsPutUnprocessableEntity{}
}

/*
PcloudNetworksPortsPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudNetworksPortsPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put unprocessable entity response has a 2xx status code
func (o *PcloudNetworksPortsPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put unprocessable entity response has a 3xx status code
func (o *PcloudNetworksPortsPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put unprocessable entity response has a 4xx status code
func (o *PcloudNetworksPortsPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports put unprocessable entity response has a 5xx status code
func (o *PcloudNetworksPortsPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports put unprocessable entity response a status code equal to that given
func (o *PcloudNetworksPortsPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudNetworksPortsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudNetworksPortsPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudNetworksPortsPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsPutInternalServerError creates a PcloudNetworksPortsPutInternalServerError with default headers values
func NewPcloudNetworksPortsPutInternalServerError() *PcloudNetworksPortsPutInternalServerError {
	return &PcloudNetworksPortsPutInternalServerError{}
}

/*
PcloudNetworksPortsPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksPortsPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports put internal server error response has a 2xx status code
func (o *PcloudNetworksPortsPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports put internal server error response has a 3xx status code
func (o *PcloudNetworksPortsPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports put internal server error response has a 4xx status code
func (o *PcloudNetworksPortsPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports put internal server error response has a 5xx status code
func (o *PcloudNetworksPortsPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks ports put internal server error response a status code equal to that given
func (o *PcloudNetworksPortsPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudNetworksPortsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
