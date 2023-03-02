// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerAuthDeviceCodePostReader is a Reader for the ServiceBrokerAuthDeviceCodePost structure.
type ServiceBrokerAuthDeviceCodePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthDeviceCodePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthDeviceCodePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewServiceBrokerAuthDeviceCodePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthDeviceCodePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerAuthDeviceCodePostOK creates a ServiceBrokerAuthDeviceCodePostOK with default headers values
func NewServiceBrokerAuthDeviceCodePostOK() *ServiceBrokerAuthDeviceCodePostOK {
	return &ServiceBrokerAuthDeviceCodePostOK{}
}

/*
ServiceBrokerAuthDeviceCodePostOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthDeviceCodePostOK struct {
	Payload *models.DeviceCode
}

// IsSuccess returns true when this service broker auth device code post o k response has a 2xx status code
func (o *ServiceBrokerAuthDeviceCodePostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker auth device code post o k response has a 3xx status code
func (o *ServiceBrokerAuthDeviceCodePostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device code post o k response has a 4xx status code
func (o *ServiceBrokerAuthDeviceCodePostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth device code post o k response has a 5xx status code
func (o *ServiceBrokerAuthDeviceCodePostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device code post o k response a status code equal to that given
func (o *ServiceBrokerAuthDeviceCodePostOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServiceBrokerAuthDeviceCodePostOK) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostOK) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostOK) GetPayload() *models.DeviceCode {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceCodePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceCodePostForbidden creates a ServiceBrokerAuthDeviceCodePostForbidden with default headers values
func NewServiceBrokerAuthDeviceCodePostForbidden() *ServiceBrokerAuthDeviceCodePostForbidden {
	return &ServiceBrokerAuthDeviceCodePostForbidden{}
}

/*
ServiceBrokerAuthDeviceCodePostForbidden describes a response with status code 403, with default header values.

Quota exceeded
*/
type ServiceBrokerAuthDeviceCodePostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device code post forbidden response has a 2xx status code
func (o *ServiceBrokerAuthDeviceCodePostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device code post forbidden response has a 3xx status code
func (o *ServiceBrokerAuthDeviceCodePostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device code post forbidden response has a 4xx status code
func (o *ServiceBrokerAuthDeviceCodePostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth device code post forbidden response has a 5xx status code
func (o *ServiceBrokerAuthDeviceCodePostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device code post forbidden response a status code equal to that given
func (o *ServiceBrokerAuthDeviceCodePostForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ServiceBrokerAuthDeviceCodePostForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostForbidden) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceCodePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceCodePostInternalServerError creates a ServiceBrokerAuthDeviceCodePostInternalServerError with default headers values
func NewServiceBrokerAuthDeviceCodePostInternalServerError() *ServiceBrokerAuthDeviceCodePostInternalServerError {
	return &ServiceBrokerAuthDeviceCodePostInternalServerError{}
}

/*
ServiceBrokerAuthDeviceCodePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthDeviceCodePostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device code post internal server error response has a 2xx status code
func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device code post internal server error response has a 3xx status code
func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device code post internal server error response has a 4xx status code
func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth device code post internal server error response has a 5xx status code
func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker auth device code post internal server error response a status code equal to that given
func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/code][%d] serviceBrokerAuthDeviceCodePostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceCodePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}