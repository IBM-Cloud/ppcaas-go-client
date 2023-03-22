// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// ServiceBrokerAuthDeviceTokenPostReader is a Reader for the ServiceBrokerAuthDeviceTokenPost structure.
type ServiceBrokerAuthDeviceTokenPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthDeviceTokenPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerAuthDeviceTokenPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerAuthDeviceTokenPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerAuthDeviceTokenPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewServiceBrokerAuthDeviceTokenPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewServiceBrokerAuthDeviceTokenPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServiceBrokerAuthDeviceTokenPostOK creates a ServiceBrokerAuthDeviceTokenPostOK with default headers values
func NewServiceBrokerAuthDeviceTokenPostOK() *ServiceBrokerAuthDeviceTokenPostOK {
	return &ServiceBrokerAuthDeviceTokenPostOK{}
}

/*
ServiceBrokerAuthDeviceTokenPostOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerAuthDeviceTokenPostOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this service broker auth device token post o k response has a 2xx status code
func (o *ServiceBrokerAuthDeviceTokenPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker auth device token post o k response has a 3xx status code
func (o *ServiceBrokerAuthDeviceTokenPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device token post o k response has a 4xx status code
func (o *ServiceBrokerAuthDeviceTokenPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth device token post o k response has a 5xx status code
func (o *ServiceBrokerAuthDeviceTokenPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device token post o k response a status code equal to that given
func (o *ServiceBrokerAuthDeviceTokenPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker auth device token post o k response
func (o *ServiceBrokerAuthDeviceTokenPostOK) Code() int {
	return 200
}

func (o *ServiceBrokerAuthDeviceTokenPostOK) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostOK) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceTokenPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceTokenPostBadRequest creates a ServiceBrokerAuthDeviceTokenPostBadRequest with default headers values
func NewServiceBrokerAuthDeviceTokenPostBadRequest() *ServiceBrokerAuthDeviceTokenPostBadRequest {
	return &ServiceBrokerAuthDeviceTokenPostBadRequest{}
}

/*
ServiceBrokerAuthDeviceTokenPostBadRequest describes a response with status code 400, with default header values.

Authorization pending
*/
type ServiceBrokerAuthDeviceTokenPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device token post bad request response has a 2xx status code
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device token post bad request response has a 3xx status code
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device token post bad request response has a 4xx status code
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth device token post bad request response has a 5xx status code
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device token post bad request response a status code equal to that given
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker auth device token post bad request response
func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceTokenPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceTokenPostForbidden creates a ServiceBrokerAuthDeviceTokenPostForbidden with default headers values
func NewServiceBrokerAuthDeviceTokenPostForbidden() *ServiceBrokerAuthDeviceTokenPostForbidden {
	return &ServiceBrokerAuthDeviceTokenPostForbidden{}
}

/*
ServiceBrokerAuthDeviceTokenPostForbidden describes a response with status code 403, with default header values.

User refused grant
*/
type ServiceBrokerAuthDeviceTokenPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device token post forbidden response has a 2xx status code
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device token post forbidden response has a 3xx status code
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device token post forbidden response has a 4xx status code
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth device token post forbidden response has a 5xx status code
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device token post forbidden response a status code equal to that given
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker auth device token post forbidden response
func (o *ServiceBrokerAuthDeviceTokenPostForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerAuthDeviceTokenPostForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostForbidden) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceTokenPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceTokenPostTooManyRequests creates a ServiceBrokerAuthDeviceTokenPostTooManyRequests with default headers values
func NewServiceBrokerAuthDeviceTokenPostTooManyRequests() *ServiceBrokerAuthDeviceTokenPostTooManyRequests {
	return &ServiceBrokerAuthDeviceTokenPostTooManyRequests{}
}

/*
ServiceBrokerAuthDeviceTokenPostTooManyRequests describes a response with status code 429, with default header values.

Polling too frequently
*/
type ServiceBrokerAuthDeviceTokenPostTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device token post too many requests response has a 2xx status code
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device token post too many requests response has a 3xx status code
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device token post too many requests response has a 4xx status code
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker auth device token post too many requests response has a 5xx status code
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker auth device token post too many requests response a status code equal to that given
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the service broker auth device token post too many requests response
func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) Code() int {
	return 429
}

func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceTokenPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthDeviceTokenPostInternalServerError creates a ServiceBrokerAuthDeviceTokenPostInternalServerError with default headers values
func NewServiceBrokerAuthDeviceTokenPostInternalServerError() *ServiceBrokerAuthDeviceTokenPostInternalServerError {
	return &ServiceBrokerAuthDeviceTokenPostInternalServerError{}
}

/*
ServiceBrokerAuthDeviceTokenPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ServiceBrokerAuthDeviceTokenPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker auth device token post internal server error response has a 2xx status code
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker auth device token post internal server error response has a 3xx status code
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker auth device token post internal server error response has a 4xx status code
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker auth device token post internal server error response has a 5xx status code
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this service broker auth device token post internal server error response a status code equal to that given
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the service broker auth device token post internal server error response
func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) Code() int {
	return 500
}

func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /auth/v1/device/token][%d] serviceBrokerAuthDeviceTokenPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerAuthDeviceTokenPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ServiceBrokerAuthDeviceTokenPostBody service broker auth device token post body
swagger:model ServiceBrokerAuthDeviceTokenPostBody
*/
type ServiceBrokerAuthDeviceTokenPostBody struct {

	// The deviceCode that the authorization server returned
	DeviceCode string `json:"deviceCode,omitempty"`
}

// Validate validates this service broker auth device token post body
func (o *ServiceBrokerAuthDeviceTokenPostBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service broker auth device token post body based on context it is used
func (o *ServiceBrokerAuthDeviceTokenPostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceBrokerAuthDeviceTokenPostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceBrokerAuthDeviceTokenPostBody) UnmarshalBinary(b []byte) error {
	var res ServiceBrokerAuthDeviceTokenPostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
