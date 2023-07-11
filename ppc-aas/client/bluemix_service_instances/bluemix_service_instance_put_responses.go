// Code generated by go-swagger; DO NOT EDIT.

package bluemix_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// BluemixServiceInstancePutReader is a Reader for the BluemixServiceInstancePut structure.
type BluemixServiceInstancePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BluemixServiceInstancePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBluemixServiceInstancePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBluemixServiceInstancePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /bluemix_v1/service_instances/{instance_id}] bluemix.serviceInstance.put", response, response.Code())
	}
}

// NewBluemixServiceInstancePutOK creates a BluemixServiceInstancePutOK with default headers values
func NewBluemixServiceInstancePutOK() *BluemixServiceInstancePutOK {
	return &BluemixServiceInstancePutOK{}
}

/*
BluemixServiceInstancePutOK describes a response with status code 200, with default header values.

OK
*/
type BluemixServiceInstancePutOK struct {
	Payload *models.ServiceInstance
}

// IsSuccess returns true when this bluemix service instance put o k response has a 2xx status code
func (o *BluemixServiceInstancePutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bluemix service instance put o k response has a 3xx status code
func (o *BluemixServiceInstancePutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bluemix service instance put o k response has a 4xx status code
func (o *BluemixServiceInstancePutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bluemix service instance put o k response has a 5xx status code
func (o *BluemixServiceInstancePutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bluemix service instance put o k response a status code equal to that given
func (o *BluemixServiceInstancePutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bluemix service instance put o k response
func (o *BluemixServiceInstancePutOK) Code() int {
	return 200
}

func (o *BluemixServiceInstancePutOK) Error() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutOK  %+v", 200, o.Payload)
}

func (o *BluemixServiceInstancePutOK) String() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutOK  %+v", 200, o.Payload)
}

func (o *BluemixServiceInstancePutOK) GetPayload() *models.ServiceInstance {
	return o.Payload
}

func (o *BluemixServiceInstancePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBluemixServiceInstancePutBadRequest creates a BluemixServiceInstancePutBadRequest with default headers values
func NewBluemixServiceInstancePutBadRequest() *BluemixServiceInstancePutBadRequest {
	return &BluemixServiceInstancePutBadRequest{}
}

/*
BluemixServiceInstancePutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BluemixServiceInstancePutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this bluemix service instance put bad request response has a 2xx status code
func (o *BluemixServiceInstancePutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bluemix service instance put bad request response has a 3xx status code
func (o *BluemixServiceInstancePutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bluemix service instance put bad request response has a 4xx status code
func (o *BluemixServiceInstancePutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this bluemix service instance put bad request response has a 5xx status code
func (o *BluemixServiceInstancePutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this bluemix service instance put bad request response a status code equal to that given
func (o *BluemixServiceInstancePutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the bluemix service instance put bad request response
func (o *BluemixServiceInstancePutBadRequest) Code() int {
	return 400
}

func (o *BluemixServiceInstancePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutBadRequest  %+v", 400, o.Payload)
}

func (o *BluemixServiceInstancePutBadRequest) String() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutBadRequest  %+v", 400, o.Payload)
}

func (o *BluemixServiceInstancePutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *BluemixServiceInstancePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
