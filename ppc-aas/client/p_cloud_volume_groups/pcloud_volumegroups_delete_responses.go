// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudVolumegroupsDeleteReader is a Reader for the PcloudVolumegroupsDelete structure.
type PcloudVolumegroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumegroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVolumegroupsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumegroupsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumegroupsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumegroupsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudVolumegroupsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumegroupsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumegroupsDeleteAccepted creates a PcloudVolumegroupsDeleteAccepted with default headers values
func NewPcloudVolumegroupsDeleteAccepted() *PcloudVolumegroupsDeleteAccepted {
	return &PcloudVolumegroupsDeleteAccepted{}
}

/*
PcloudVolumegroupsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVolumegroupsDeleteAccepted struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud volumegroups delete accepted response has a 2xx status code
func (o *PcloudVolumegroupsDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volumegroups delete accepted response has a 3xx status code
func (o *PcloudVolumegroupsDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete accepted response has a 4xx status code
func (o *PcloudVolumegroupsDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups delete accepted response has a 5xx status code
func (o *PcloudVolumegroupsDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups delete accepted response a status code equal to that given
func (o *PcloudVolumegroupsDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PcloudVolumegroupsDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVolumegroupsDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteAccepted  %+v", 202, o.Payload)
}

func (o *PcloudVolumegroupsDeleteAccepted) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsDeleteBadRequest creates a PcloudVolumegroupsDeleteBadRequest with default headers values
func NewPcloudVolumegroupsDeleteBadRequest() *PcloudVolumegroupsDeleteBadRequest {
	return &PcloudVolumegroupsDeleteBadRequest{}
}

/*
PcloudVolumegroupsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumegroupsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups delete bad request response has a 2xx status code
func (o *PcloudVolumegroupsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups delete bad request response has a 3xx status code
func (o *PcloudVolumegroupsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete bad request response has a 4xx status code
func (o *PcloudVolumegroupsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups delete bad request response has a 5xx status code
func (o *PcloudVolumegroupsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups delete bad request response a status code equal to that given
func (o *PcloudVolumegroupsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudVolumegroupsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumegroupsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsDeleteUnauthorized creates a PcloudVolumegroupsDeleteUnauthorized with default headers values
func NewPcloudVolumegroupsDeleteUnauthorized() *PcloudVolumegroupsDeleteUnauthorized {
	return &PcloudVolumegroupsDeleteUnauthorized{}
}

/*
PcloudVolumegroupsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumegroupsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups delete unauthorized response has a 2xx status code
func (o *PcloudVolumegroupsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups delete unauthorized response has a 3xx status code
func (o *PcloudVolumegroupsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete unauthorized response has a 4xx status code
func (o *PcloudVolumegroupsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups delete unauthorized response has a 5xx status code
func (o *PcloudVolumegroupsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups delete unauthorized response a status code equal to that given
func (o *PcloudVolumegroupsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudVolumegroupsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumegroupsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsDeleteForbidden creates a PcloudVolumegroupsDeleteForbidden with default headers values
func NewPcloudVolumegroupsDeleteForbidden() *PcloudVolumegroupsDeleteForbidden {
	return &PcloudVolumegroupsDeleteForbidden{}
}

/*
PcloudVolumegroupsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumegroupsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups delete forbidden response has a 2xx status code
func (o *PcloudVolumegroupsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups delete forbidden response has a 3xx status code
func (o *PcloudVolumegroupsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete forbidden response has a 4xx status code
func (o *PcloudVolumegroupsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups delete forbidden response has a 5xx status code
func (o *PcloudVolumegroupsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups delete forbidden response a status code equal to that given
func (o *PcloudVolumegroupsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudVolumegroupsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumegroupsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsDeleteNotFound creates a PcloudVolumegroupsDeleteNotFound with default headers values
func NewPcloudVolumegroupsDeleteNotFound() *PcloudVolumegroupsDeleteNotFound {
	return &PcloudVolumegroupsDeleteNotFound{}
}

/*
PcloudVolumegroupsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudVolumegroupsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups delete not found response has a 2xx status code
func (o *PcloudVolumegroupsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups delete not found response has a 3xx status code
func (o *PcloudVolumegroupsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete not found response has a 4xx status code
func (o *PcloudVolumegroupsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volumegroups delete not found response has a 5xx status code
func (o *PcloudVolumegroupsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volumegroups delete not found response a status code equal to that given
func (o *PcloudVolumegroupsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudVolumegroupsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVolumegroupsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudVolumegroupsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumegroupsDeleteInternalServerError creates a PcloudVolumegroupsDeleteInternalServerError with default headers values
func NewPcloudVolumegroupsDeleteInternalServerError() *PcloudVolumegroupsDeleteInternalServerError {
	return &PcloudVolumegroupsDeleteInternalServerError{}
}

/*
PcloudVolumegroupsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumegroupsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volumegroups delete internal server error response has a 2xx status code
func (o *PcloudVolumegroupsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volumegroups delete internal server error response has a 3xx status code
func (o *PcloudVolumegroupsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volumegroups delete internal server error response has a 4xx status code
func (o *PcloudVolumegroupsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volumegroups delete internal server error response has a 5xx status code
func (o *PcloudVolumegroupsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volumegroups delete internal server error response a status code equal to that given
func (o *PcloudVolumegroupsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudVolumegroupsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/volume-groups/{volume_group_id}][%d] pcloudVolumegroupsDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumegroupsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumegroupsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
