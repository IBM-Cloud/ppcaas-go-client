// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudCloudinstancesGetReader is a Reader for the PcloudCloudinstancesGet structure.
type PcloudCloudinstancesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesGetOK creates a PcloudCloudinstancesGetOK with default headers values
func NewPcloudCloudinstancesGetOK() *PcloudCloudinstancesGetOK {
	return &PcloudCloudinstancesGetOK{}
}

/*
PcloudCloudinstancesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesGetOK struct {
	Payload *models.CloudInstance
}

// IsSuccess returns true when this pcloud cloudinstances get o k response has a 2xx status code
func (o *PcloudCloudinstancesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances get o k response has a 3xx status code
func (o *PcloudCloudinstancesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances get o k response has a 4xx status code
func (o *PcloudCloudinstancesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances get o k response has a 5xx status code
func (o *PcloudCloudinstancesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances get o k response a status code equal to that given
func (o *PcloudCloudinstancesGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesGetOK) GetPayload() *models.CloudInstance {
	return o.Payload
}

func (o *PcloudCloudinstancesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetBadRequest creates a PcloudCloudinstancesGetBadRequest with default headers values
func NewPcloudCloudinstancesGetBadRequest() *PcloudCloudinstancesGetBadRequest {
	return &PcloudCloudinstancesGetBadRequest{}
}

/*
PcloudCloudinstancesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances get bad request response has a 2xx status code
func (o *PcloudCloudinstancesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances get bad request response has a 3xx status code
func (o *PcloudCloudinstancesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances get bad request response has a 4xx status code
func (o *PcloudCloudinstancesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances get bad request response has a 5xx status code
func (o *PcloudCloudinstancesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances get bad request response a status code equal to that given
func (o *PcloudCloudinstancesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetUnauthorized creates a PcloudCloudinstancesGetUnauthorized with default headers values
func NewPcloudCloudinstancesGetUnauthorized() *PcloudCloudinstancesGetUnauthorized {
	return &PcloudCloudinstancesGetUnauthorized{}
}

/*
PcloudCloudinstancesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetNotFound creates a PcloudCloudinstancesGetNotFound with default headers values
func NewPcloudCloudinstancesGetNotFound() *PcloudCloudinstancesGetNotFound {
	return &PcloudCloudinstancesGetNotFound{}
}

/*
PcloudCloudinstancesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances get not found response has a 2xx status code
func (o *PcloudCloudinstancesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances get not found response has a 3xx status code
func (o *PcloudCloudinstancesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances get not found response has a 4xx status code
func (o *PcloudCloudinstancesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances get not found response has a 5xx status code
func (o *PcloudCloudinstancesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances get not found response a status code equal to that given
func (o *PcloudCloudinstancesGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudinstancesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetInternalServerError creates a PcloudCloudinstancesGetInternalServerError with default headers values
func NewPcloudCloudinstancesGetInternalServerError() *PcloudCloudinstancesGetInternalServerError {
	return &PcloudCloudinstancesGetInternalServerError{}
}

/*
PcloudCloudinstancesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
