// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// ServiceBindingLastOperationGetReader is a Reader for the ServiceBindingLastOperationGet structure.
type ServiceBindingLastOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingLastOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingLastOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBindingLastOperationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewServiceBindingLastOperationGetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBindingLastOperationGetOK creates a ServiceBindingLastOperationGetOK with default headers values
func NewServiceBindingLastOperationGetOK() *ServiceBindingLastOperationGetOK {
	return &ServiceBindingLastOperationGetOK{}
}

/*
ServiceBindingLastOperationGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBindingLastOperationGetOK struct {
	Payload *models.LastOperationResource
}

// IsSuccess returns true when this service binding last operation get o k response has a 2xx status code
func (o *ServiceBindingLastOperationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service binding last operation get o k response has a 3xx status code
func (o *ServiceBindingLastOperationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get o k response has a 4xx status code
func (o *ServiceBindingLastOperationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service binding last operation get o k response has a 5xx status code
func (o *ServiceBindingLastOperationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get o k response a status code equal to that given
func (o *ServiceBindingLastOperationGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceBindingLastOperationGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingLastOperationGetOK) String() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingLastOperationGetOK) GetPayload() *models.LastOperationResource {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LastOperationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetBadRequest creates a ServiceBindingLastOperationGetBadRequest with default headers values
func NewServiceBindingLastOperationGetBadRequest() *ServiceBindingLastOperationGetBadRequest {
	return &ServiceBindingLastOperationGetBadRequest{}
}

/*
ServiceBindingLastOperationGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBindingLastOperationGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get bad request response has a 2xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get bad request response has a 3xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get bad request response has a 4xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get bad request response has a 5xx status code
func (o *ServiceBindingLastOperationGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get bad request response a status code equal to that given
func (o *ServiceBindingLastOperationGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ServiceBindingLastOperationGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingLastOperationGetBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingLastOperationGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetGone creates a ServiceBindingLastOperationGetGone with default headers values
func NewServiceBindingLastOperationGetGone() *ServiceBindingLastOperationGetGone {
	return &ServiceBindingLastOperationGetGone{}
}

/*
ServiceBindingLastOperationGetGone describes a response with status code 410, with default header values.

Gone
*/
type ServiceBindingLastOperationGetGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this service binding last operation get gone response has a 2xx status code
func (o *ServiceBindingLastOperationGetGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service binding last operation get gone response has a 3xx status code
func (o *ServiceBindingLastOperationGetGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service binding last operation get gone response has a 4xx status code
func (o *ServiceBindingLastOperationGetGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this service binding last operation get gone response has a 5xx status code
func (o *ServiceBindingLastOperationGetGone) IsServerError() bool {
	return false
}

// IsCode returns true when this service binding last operation get gone response a status code equal to that given
func (o *ServiceBindingLastOperationGetGone) IsCode(code int) bool {
	return code == 410
}

func (o *ServiceBindingLastOperationGetGone) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingLastOperationGetGone) String() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingLastOperationGetGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
