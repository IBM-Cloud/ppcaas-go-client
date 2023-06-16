// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_pod_capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudPodcapacityGetReader is a Reader for the PcloudPodcapacityGet structure.
type PcloudPodcapacityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPodcapacityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPodcapacityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudPodcapacityGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPodcapacityGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPodcapacityGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPodcapacityGetOK creates a PcloudPodcapacityGetOK with default headers values
func NewPcloudPodcapacityGetOK() *PcloudPodcapacityGetOK {
	return &PcloudPodcapacityGetOK{}
}

/*
PcloudPodcapacityGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPodcapacityGetOK struct {
	Payload *models.PodCapacity
}

// IsSuccess returns true when this pcloud podcapacity get o k response has a 2xx status code
func (o *PcloudPodcapacityGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud podcapacity get o k response has a 3xx status code
func (o *PcloudPodcapacityGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud podcapacity get o k response has a 4xx status code
func (o *PcloudPodcapacityGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud podcapacity get o k response has a 5xx status code
func (o *PcloudPodcapacityGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud podcapacity get o k response a status code equal to that given
func (o *PcloudPodcapacityGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud podcapacity get o k response
func (o *PcloudPodcapacityGetOK) Code() int {
	return 200
}

func (o *PcloudPodcapacityGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPodcapacityGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPodcapacityGetOK) GetPayload() *models.PodCapacity {
	return o.Payload
}

func (o *PcloudPodcapacityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPodcapacityGetUnauthorized creates a PcloudPodcapacityGetUnauthorized with default headers values
func NewPcloudPodcapacityGetUnauthorized() *PcloudPodcapacityGetUnauthorized {
	return &PcloudPodcapacityGetUnauthorized{}
}

/*
PcloudPodcapacityGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPodcapacityGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud podcapacity get unauthorized response has a 2xx status code
func (o *PcloudPodcapacityGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud podcapacity get unauthorized response has a 3xx status code
func (o *PcloudPodcapacityGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud podcapacity get unauthorized response has a 4xx status code
func (o *PcloudPodcapacityGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud podcapacity get unauthorized response has a 5xx status code
func (o *PcloudPodcapacityGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud podcapacity get unauthorized response a status code equal to that given
func (o *PcloudPodcapacityGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud podcapacity get unauthorized response
func (o *PcloudPodcapacityGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudPodcapacityGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPodcapacityGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPodcapacityGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPodcapacityGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPodcapacityGetForbidden creates a PcloudPodcapacityGetForbidden with default headers values
func NewPcloudPodcapacityGetForbidden() *PcloudPodcapacityGetForbidden {
	return &PcloudPodcapacityGetForbidden{}
}

/*
PcloudPodcapacityGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPodcapacityGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud podcapacity get forbidden response has a 2xx status code
func (o *PcloudPodcapacityGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud podcapacity get forbidden response has a 3xx status code
func (o *PcloudPodcapacityGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud podcapacity get forbidden response has a 4xx status code
func (o *PcloudPodcapacityGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud podcapacity get forbidden response has a 5xx status code
func (o *PcloudPodcapacityGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud podcapacity get forbidden response a status code equal to that given
func (o *PcloudPodcapacityGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud podcapacity get forbidden response
func (o *PcloudPodcapacityGetForbidden) Code() int {
	return 403
}

func (o *PcloudPodcapacityGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPodcapacityGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPodcapacityGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPodcapacityGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPodcapacityGetInternalServerError creates a PcloudPodcapacityGetInternalServerError with default headers values
func NewPcloudPodcapacityGetInternalServerError() *PcloudPodcapacityGetInternalServerError {
	return &PcloudPodcapacityGetInternalServerError{}
}

/*
PcloudPodcapacityGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPodcapacityGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud podcapacity get internal server error response has a 2xx status code
func (o *PcloudPodcapacityGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud podcapacity get internal server error response has a 3xx status code
func (o *PcloudPodcapacityGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud podcapacity get internal server error response has a 4xx status code
func (o *PcloudPodcapacityGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud podcapacity get internal server error response has a 5xx status code
func (o *PcloudPodcapacityGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud podcapacity get internal server error response a status code equal to that given
func (o *PcloudPodcapacityGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud podcapacity get internal server error response
func (o *PcloudPodcapacityGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudPodcapacityGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPodcapacityGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pod-capacity][%d] pcloudPodcapacityGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPodcapacityGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPodcapacityGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}