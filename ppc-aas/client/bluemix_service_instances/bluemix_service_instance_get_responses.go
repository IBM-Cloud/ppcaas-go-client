// Code generated by go-swagger; DO NOT EDIT.

package bluemix_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// BluemixServiceInstanceGetReader is a Reader for the BluemixServiceInstanceGet structure.
type BluemixServiceInstanceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BluemixServiceInstanceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBluemixServiceInstanceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBluemixServiceInstanceGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBluemixServiceInstanceGetOK creates a BluemixServiceInstanceGetOK with default headers values
func NewBluemixServiceInstanceGetOK() *BluemixServiceInstanceGetOK {
	return &BluemixServiceInstanceGetOK{}
}

/*
BluemixServiceInstanceGetOK describes a response with status code 200, with default header values.

OK
*/
type BluemixServiceInstanceGetOK struct {
	Payload *models.ServiceInstance
}

// IsSuccess returns true when this bluemix service instance get o k response has a 2xx status code
func (o *BluemixServiceInstanceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bluemix service instance get o k response has a 3xx status code
func (o *BluemixServiceInstanceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bluemix service instance get o k response has a 4xx status code
func (o *BluemixServiceInstanceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bluemix service instance get o k response has a 5xx status code
func (o *BluemixServiceInstanceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bluemix service instance get o k response a status code equal to that given
func (o *BluemixServiceInstanceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *BluemixServiceInstanceGetOK) Error() string {
	return fmt.Sprintf("[GET /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstanceGetOK  %+v", 200, o.Payload)
}

func (o *BluemixServiceInstanceGetOK) String() string {
	return fmt.Sprintf("[GET /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstanceGetOK  %+v", 200, o.Payload)
}

func (o *BluemixServiceInstanceGetOK) GetPayload() *models.ServiceInstance {
	return o.Payload
}

func (o *BluemixServiceInstanceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBluemixServiceInstanceGetBadRequest creates a BluemixServiceInstanceGetBadRequest with default headers values
func NewBluemixServiceInstanceGetBadRequest() *BluemixServiceInstanceGetBadRequest {
	return &BluemixServiceInstanceGetBadRequest{}
}

/*
BluemixServiceInstanceGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BluemixServiceInstanceGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this bluemix service instance get bad request response has a 2xx status code
func (o *BluemixServiceInstanceGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this bluemix service instance get bad request response has a 3xx status code
func (o *BluemixServiceInstanceGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bluemix service instance get bad request response has a 4xx status code
func (o *BluemixServiceInstanceGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this bluemix service instance get bad request response has a 5xx status code
func (o *BluemixServiceInstanceGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this bluemix service instance get bad request response a status code equal to that given
func (o *BluemixServiceInstanceGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BluemixServiceInstanceGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstanceGetBadRequest  %+v", 400, o.Payload)
}

func (o *BluemixServiceInstanceGetBadRequest) String() string {
	return fmt.Sprintf("[GET /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstanceGetBadRequest  %+v", 400, o.Payload)
}

func (o *BluemixServiceInstanceGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *BluemixServiceInstanceGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
