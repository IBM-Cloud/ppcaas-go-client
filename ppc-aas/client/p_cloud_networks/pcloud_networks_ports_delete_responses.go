// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudNetworksPortsDeleteReader is a Reader for the PcloudNetworksPortsDelete structure.
type PcloudNetworksPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudNetworksPortsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudNetworksPortsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudNetworksPortsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudNetworksPortsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudNetworksPortsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudNetworksPortsDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudNetworksPortsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudNetworksPortsDeleteOK creates a PcloudNetworksPortsDeleteOK with default headers values
func NewPcloudNetworksPortsDeleteOK() *PcloudNetworksPortsDeleteOK {
	return &PcloudNetworksPortsDeleteOK{}
}

/*
PcloudNetworksPortsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudNetworksPortsDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud networks ports delete o k response has a 2xx status code
func (o *PcloudNetworksPortsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud networks ports delete o k response has a 3xx status code
func (o *PcloudNetworksPortsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete o k response has a 4xx status code
func (o *PcloudNetworksPortsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports delete o k response has a 5xx status code
func (o *PcloudNetworksPortsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete o k response a status code equal to that given
func (o *PcloudNetworksPortsDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudNetworksPortsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteBadRequest creates a PcloudNetworksPortsDeleteBadRequest with default headers values
func NewPcloudNetworksPortsDeleteBadRequest() *PcloudNetworksPortsDeleteBadRequest {
	return &PcloudNetworksPortsDeleteBadRequest{}
}

/*
PcloudNetworksPortsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudNetworksPortsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete bad request response has a 2xx status code
func (o *PcloudNetworksPortsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete bad request response has a 3xx status code
func (o *PcloudNetworksPortsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete bad request response has a 4xx status code
func (o *PcloudNetworksPortsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports delete bad request response has a 5xx status code
func (o *PcloudNetworksPortsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete bad request response a status code equal to that given
func (o *PcloudNetworksPortsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudNetworksPortsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPortsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudNetworksPortsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteUnauthorized creates a PcloudNetworksPortsDeleteUnauthorized with default headers values
func NewPcloudNetworksPortsDeleteUnauthorized() *PcloudNetworksPortsDeleteUnauthorized {
	return &PcloudNetworksPortsDeleteUnauthorized{}
}

/*
PcloudNetworksPortsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudNetworksPortsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete unauthorized response has a 2xx status code
func (o *PcloudNetworksPortsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete unauthorized response has a 3xx status code
func (o *PcloudNetworksPortsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete unauthorized response has a 4xx status code
func (o *PcloudNetworksPortsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports delete unauthorized response has a 5xx status code
func (o *PcloudNetworksPortsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete unauthorized response a status code equal to that given
func (o *PcloudNetworksPortsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudNetworksPortsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudNetworksPortsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteForbidden creates a PcloudNetworksPortsDeleteForbidden with default headers values
func NewPcloudNetworksPortsDeleteForbidden() *PcloudNetworksPortsDeleteForbidden {
	return &PcloudNetworksPortsDeleteForbidden{}
}

/*
PcloudNetworksPortsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudNetworksPortsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete forbidden response has a 2xx status code
func (o *PcloudNetworksPortsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete forbidden response has a 3xx status code
func (o *PcloudNetworksPortsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete forbidden response has a 4xx status code
func (o *PcloudNetworksPortsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports delete forbidden response has a 5xx status code
func (o *PcloudNetworksPortsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete forbidden response a status code equal to that given
func (o *PcloudNetworksPortsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudNetworksPortsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudNetworksPortsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteNotFound creates a PcloudNetworksPortsDeleteNotFound with default headers values
func NewPcloudNetworksPortsDeleteNotFound() *PcloudNetworksPortsDeleteNotFound {
	return &PcloudNetworksPortsDeleteNotFound{}
}

/*
PcloudNetworksPortsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudNetworksPortsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete not found response has a 2xx status code
func (o *PcloudNetworksPortsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete not found response has a 3xx status code
func (o *PcloudNetworksPortsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete not found response has a 4xx status code
func (o *PcloudNetworksPortsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports delete not found response has a 5xx status code
func (o *PcloudNetworksPortsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete not found response a status code equal to that given
func (o *PcloudNetworksPortsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudNetworksPortsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteGone creates a PcloudNetworksPortsDeleteGone with default headers values
func NewPcloudNetworksPortsDeleteGone() *PcloudNetworksPortsDeleteGone {
	return &PcloudNetworksPortsDeleteGone{}
}

/*
PcloudNetworksPortsDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudNetworksPortsDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete gone response has a 2xx status code
func (o *PcloudNetworksPortsDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete gone response has a 3xx status code
func (o *PcloudNetworksPortsDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete gone response has a 4xx status code
func (o *PcloudNetworksPortsDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud networks ports delete gone response has a 5xx status code
func (o *PcloudNetworksPortsDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud networks ports delete gone response a status code equal to that given
func (o *PcloudNetworksPortsDeleteGone) IsCode(code int) bool {
	return code == 410
}

func (o *PcloudNetworksPortsDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudNetworksPortsDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudNetworksPortsDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsDeleteInternalServerError creates a PcloudNetworksPortsDeleteInternalServerError with default headers values
func NewPcloudNetworksPortsDeleteInternalServerError() *PcloudNetworksPortsDeleteInternalServerError {
	return &PcloudNetworksPortsDeleteInternalServerError{}
}

/*
PcloudNetworksPortsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudNetworksPortsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud networks ports delete internal server error response has a 2xx status code
func (o *PcloudNetworksPortsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud networks ports delete internal server error response has a 3xx status code
func (o *PcloudNetworksPortsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud networks ports delete internal server error response has a 4xx status code
func (o *PcloudNetworksPortsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud networks ports delete internal server error response has a 5xx status code
func (o *PcloudNetworksPortsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud networks ports delete internal server error response a status code equal to that given
func (o *PcloudNetworksPortsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudNetworksPortsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudNetworksPortsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
