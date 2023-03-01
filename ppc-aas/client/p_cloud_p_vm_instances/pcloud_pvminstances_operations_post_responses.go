// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudPvminstancesOperationsPostReader is a Reader for the PcloudPvminstancesOperationsPost structure.
type PcloudPvminstancesOperationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesOperationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesOperationsPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesOperationsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesOperationsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesOperationsPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPvminstancesOperationsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesOperationsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesOperationsPostOK creates a PcloudPvminstancesOperationsPostOK with default headers values
func NewPcloudPvminstancesOperationsPostOK() *PcloudPvminstancesOperationsPostOK {
	return &PcloudPvminstancesOperationsPostOK{}
}

/*
PcloudPvminstancesOperationsPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesOperationsPostOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances operations post o k response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances operations post o k response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post o k response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances operations post o k response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post o k response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudPvminstancesOperationsPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostBadRequest creates a PcloudPvminstancesOperationsPostBadRequest with default headers values
func NewPcloudPvminstancesOperationsPostBadRequest() *PcloudPvminstancesOperationsPostBadRequest {
	return &PcloudPvminstancesOperationsPostBadRequest{}
}

/*
PcloudPvminstancesOperationsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesOperationsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post bad request response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post bad request response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post bad request response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post bad request response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post bad request response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudPvminstancesOperationsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostUnauthorized creates a PcloudPvminstancesOperationsPostUnauthorized with default headers values
func NewPcloudPvminstancesOperationsPostUnauthorized() *PcloudPvminstancesOperationsPostUnauthorized {
	return &PcloudPvminstancesOperationsPostUnauthorized{}
}

/*
PcloudPvminstancesOperationsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesOperationsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post unauthorized response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post unauthorized response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post unauthorized response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post unauthorized response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post unauthorized response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostNotFound creates a PcloudPvminstancesOperationsPostNotFound with default headers values
func NewPcloudPvminstancesOperationsPostNotFound() *PcloudPvminstancesOperationsPostNotFound {
	return &PcloudPvminstancesOperationsPostNotFound{}
}

/*
PcloudPvminstancesOperationsPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesOperationsPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post not found response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post not found response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post not found response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post not found response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post not found response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudPvminstancesOperationsPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostUnprocessableEntity creates a PcloudPvminstancesOperationsPostUnprocessableEntity with default headers values
func NewPcloudPvminstancesOperationsPostUnprocessableEntity() *PcloudPvminstancesOperationsPostUnprocessableEntity {
	return &PcloudPvminstancesOperationsPostUnprocessableEntity{}
}

/*
PcloudPvminstancesOperationsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesOperationsPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post unprocessable entity response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post unprocessable entity response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post unprocessable entity response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances operations post unprocessable entity response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances operations post unprocessable entity response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesOperationsPostInternalServerError creates a PcloudPvminstancesOperationsPostInternalServerError with default headers values
func NewPcloudPvminstancesOperationsPostInternalServerError() *PcloudPvminstancesOperationsPostInternalServerError {
	return &PcloudPvminstancesOperationsPostInternalServerError{}
}

/*
PcloudPvminstancesOperationsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesOperationsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances operations post internal server error response has a 2xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances operations post internal server error response has a 3xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances operations post internal server error response has a 4xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances operations post internal server error response has a 5xx status code
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances operations post internal server error response a status code equal to that given
func (o *PcloudPvminstancesOperationsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/operations][%d] pcloudPvminstancesOperationsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesOperationsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
