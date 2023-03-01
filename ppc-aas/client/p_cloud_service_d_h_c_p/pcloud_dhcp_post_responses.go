// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudDhcpPostReader is a Reader for the PcloudDhcpPost structure.
type PcloudDhcpPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudDhcpPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudDhcpPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudDhcpPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudDhcpPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudDhcpPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudDhcpPostAccepted creates a PcloudDhcpPostAccepted with default headers values
func NewPcloudDhcpPostAccepted() *PcloudDhcpPostAccepted {
	return &PcloudDhcpPostAccepted{}
}

/*
PcloudDhcpPostAccepted describes a response with status code 202, with default header values.

OK
*/
type PcloudDhcpPostAccepted struct {
	Payload *models.DHCPServer
}

// IsSuccess returns true when this pcloud dhcp post accepted response has a 2xx status code
func (o *PcloudDhcpPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud dhcp post accepted response has a 3xx status code
func (o *PcloudDhcpPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp post accepted response has a 4xx status code
func (o *PcloudDhcpPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp post accepted response has a 5xx status code
func (o *PcloudDhcpPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp post accepted response a status code equal to that given
func (o *PcloudDhcpPostAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PcloudDhcpPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudDhcpPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudDhcpPostAccepted) GetPayload() *models.DHCPServer {
	return o.Payload
}

func (o *PcloudDhcpPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DHCPServer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpPostBadRequest creates a PcloudDhcpPostBadRequest with default headers values
func NewPcloudDhcpPostBadRequest() *PcloudDhcpPostBadRequest {
	return &PcloudDhcpPostBadRequest{}
}

/*
PcloudDhcpPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudDhcpPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp post bad request response has a 2xx status code
func (o *PcloudDhcpPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp post bad request response has a 3xx status code
func (o *PcloudDhcpPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp post bad request response has a 4xx status code
func (o *PcloudDhcpPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp post bad request response has a 5xx status code
func (o *PcloudDhcpPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp post bad request response a status code equal to that given
func (o *PcloudDhcpPostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudDhcpPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudDhcpPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudDhcpPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpPostForbidden creates a PcloudDhcpPostForbidden with default headers values
func NewPcloudDhcpPostForbidden() *PcloudDhcpPostForbidden {
	return &PcloudDhcpPostForbidden{}
}

/*
PcloudDhcpPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudDhcpPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp post forbidden response has a 2xx status code
func (o *PcloudDhcpPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp post forbidden response has a 3xx status code
func (o *PcloudDhcpPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp post forbidden response has a 4xx status code
func (o *PcloudDhcpPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud dhcp post forbidden response has a 5xx status code
func (o *PcloudDhcpPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud dhcp post forbidden response a status code equal to that given
func (o *PcloudDhcpPostForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudDhcpPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudDhcpPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudDhcpPostInternalServerError creates a PcloudDhcpPostInternalServerError with default headers values
func NewPcloudDhcpPostInternalServerError() *PcloudDhcpPostInternalServerError {
	return &PcloudDhcpPostInternalServerError{}
}

/*
PcloudDhcpPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudDhcpPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud dhcp post internal server error response has a 2xx status code
func (o *PcloudDhcpPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud dhcp post internal server error response has a 3xx status code
func (o *PcloudDhcpPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud dhcp post internal server error response has a 4xx status code
func (o *PcloudDhcpPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud dhcp post internal server error response has a 5xx status code
func (o *PcloudDhcpPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud dhcp post internal server error response a status code equal to that given
func (o *PcloudDhcpPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudDhcpPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/services/dhcp][%d] pcloudDhcpPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudDhcpPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudDhcpPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
