// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudTenantsSshkeysPutReader is a Reader for the PcloudTenantsSshkeysPut structure.
type PcloudTenantsSshkeysPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudTenantsSshkeysPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysPutOK creates a PcloudTenantsSshkeysPutOK with default headers values
func NewPcloudTenantsSshkeysPutOK() *PcloudTenantsSshkeysPutOK {
	return &PcloudTenantsSshkeysPutOK{}
}

/*
PcloudTenantsSshkeysPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysPutOK struct {
	Payload *models.SSHKey
}

// IsSuccess returns true when this pcloud tenants sshkeys put o k response has a 2xx status code
func (o *PcloudTenantsSshkeysPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tenants sshkeys put o k response has a 3xx status code
func (o *PcloudTenantsSshkeysPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys put o k response has a 4xx status code
func (o *PcloudTenantsSshkeysPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys put o k response has a 5xx status code
func (o *PcloudTenantsSshkeysPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys put o k response a status code equal to that given
func (o *PcloudTenantsSshkeysPutOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudTenantsSshkeysPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysPutOK) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysPutOK) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPutBadRequest creates a PcloudTenantsSshkeysPutBadRequest with default headers values
func NewPcloudTenantsSshkeysPutBadRequest() *PcloudTenantsSshkeysPutBadRequest {
	return &PcloudTenantsSshkeysPutBadRequest{}
}

/*
PcloudTenantsSshkeysPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys put bad request response has a 2xx status code
func (o *PcloudTenantsSshkeysPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys put bad request response has a 3xx status code
func (o *PcloudTenantsSshkeysPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys put bad request response has a 4xx status code
func (o *PcloudTenantsSshkeysPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys put bad request response has a 5xx status code
func (o *PcloudTenantsSshkeysPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys put bad request response a status code equal to that given
func (o *PcloudTenantsSshkeysPutBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudTenantsSshkeysPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysPutBadRequest) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPutUnauthorized creates a PcloudTenantsSshkeysPutUnauthorized with default headers values
func NewPcloudTenantsSshkeysPutUnauthorized() *PcloudTenantsSshkeysPutUnauthorized {
	return &PcloudTenantsSshkeysPutUnauthorized{}
}

/*
PcloudTenantsSshkeysPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys put unauthorized response has a 2xx status code
func (o *PcloudTenantsSshkeysPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys put unauthorized response has a 3xx status code
func (o *PcloudTenantsSshkeysPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys put unauthorized response has a 4xx status code
func (o *PcloudTenantsSshkeysPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys put unauthorized response has a 5xx status code
func (o *PcloudTenantsSshkeysPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys put unauthorized response a status code equal to that given
func (o *PcloudTenantsSshkeysPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudTenantsSshkeysPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysPutUnauthorized) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPutUnprocessableEntity creates a PcloudTenantsSshkeysPutUnprocessableEntity with default headers values
func NewPcloudTenantsSshkeysPutUnprocessableEntity() *PcloudTenantsSshkeysPutUnprocessableEntity {
	return &PcloudTenantsSshkeysPutUnprocessableEntity{}
}

/*
PcloudTenantsSshkeysPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudTenantsSshkeysPutUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys put unprocessable entity response has a 2xx status code
func (o *PcloudTenantsSshkeysPutUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys put unprocessable entity response has a 3xx status code
func (o *PcloudTenantsSshkeysPutUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys put unprocessable entity response has a 4xx status code
func (o *PcloudTenantsSshkeysPutUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys put unprocessable entity response has a 5xx status code
func (o *PcloudTenantsSshkeysPutUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys put unprocessable entity response a status code equal to that given
func (o *PcloudTenantsSshkeysPutUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PcloudTenantsSshkeysPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudTenantsSshkeysPutUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudTenantsSshkeysPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysPutInternalServerError creates a PcloudTenantsSshkeysPutInternalServerError with default headers values
func NewPcloudTenantsSshkeysPutInternalServerError() *PcloudTenantsSshkeysPutInternalServerError {
	return &PcloudTenantsSshkeysPutInternalServerError{}
}

/*
PcloudTenantsSshkeysPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys put internal server error response has a 2xx status code
func (o *PcloudTenantsSshkeysPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys put internal server error response has a 3xx status code
func (o *PcloudTenantsSshkeysPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys put internal server error response has a 4xx status code
func (o *PcloudTenantsSshkeysPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys put internal server error response has a 5xx status code
func (o *PcloudTenantsSshkeysPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tenants sshkeys put internal server error response a status code equal to that given
func (o *PcloudTenantsSshkeysPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudTenantsSshkeysPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysPutInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
