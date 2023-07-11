// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerTestTimeoutReader is a Reader for the ServiceBrokerTestTimeout structure.
type ServiceBrokerTestTimeoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerTestTimeoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerTestTimeoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/test/timeout] serviceBroker.test.timeout", response, response.Code())
	}
}

// NewServiceBrokerTestTimeoutOK creates a ServiceBrokerTestTimeoutOK with default headers values
func NewServiceBrokerTestTimeoutOK() *ServiceBrokerTestTimeoutOK {
	return &ServiceBrokerTestTimeoutOK{}
}

/*
ServiceBrokerTestTimeoutOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerTestTimeoutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service broker test timeout o k response has a 2xx status code
func (o *ServiceBrokerTestTimeoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker test timeout o k response has a 3xx status code
func (o *ServiceBrokerTestTimeoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout o k response has a 4xx status code
func (o *ServiceBrokerTestTimeoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker test timeout o k response has a 5xx status code
func (o *ServiceBrokerTestTimeoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout o k response a status code equal to that given
func (o *ServiceBrokerTestTimeoutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker test timeout o k response
func (o *ServiceBrokerTestTimeoutOK) Code() int {
	return 200
}

func (o *ServiceBrokerTestTimeoutOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerTestTimeoutOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerTestTimeoutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
