// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudPvminstancesSnapshotsGetallReader is a Reader for the PcloudPvminstancesSnapshotsGetall structure.
type PcloudPvminstancesSnapshotsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesSnapshotsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesSnapshotsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesSnapshotsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesSnapshotsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesSnapshotsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesSnapshotsGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesSnapshotsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPvminstancesSnapshotsGetallOK creates a PcloudPvminstancesSnapshotsGetallOK with default headers values
func NewPcloudPvminstancesSnapshotsGetallOK() *PcloudPvminstancesSnapshotsGetallOK {
	return &PcloudPvminstancesSnapshotsGetallOK{}
}

/*
PcloudPvminstancesSnapshotsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesSnapshotsGetallOK struct {
	Payload *models.Snapshots
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall o k response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall o k response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall o k response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances snapshots getall o k response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots getall o k response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances snapshots getall o k response
func (o *PcloudPvminstancesSnapshotsGetallOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesSnapshotsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallOK) GetPayload() *models.Snapshots {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshots)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsGetallBadRequest creates a PcloudPvminstancesSnapshotsGetallBadRequest with default headers values
func NewPcloudPvminstancesSnapshotsGetallBadRequest() *PcloudPvminstancesSnapshotsGetallBadRequest {
	return &PcloudPvminstancesSnapshotsGetallBadRequest{}
}

/*
PcloudPvminstancesSnapshotsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesSnapshotsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall bad request response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall bad request response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall bad request response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots getall bad request response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots getall bad request response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances snapshots getall bad request response
func (o *PcloudPvminstancesSnapshotsGetallBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesSnapshotsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsGetallUnauthorized creates a PcloudPvminstancesSnapshotsGetallUnauthorized with default headers values
func NewPcloudPvminstancesSnapshotsGetallUnauthorized() *PcloudPvminstancesSnapshotsGetallUnauthorized {
	return &PcloudPvminstancesSnapshotsGetallUnauthorized{}
}

/*
PcloudPvminstancesSnapshotsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesSnapshotsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall unauthorized response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall unauthorized response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall unauthorized response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots getall unauthorized response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots getall unauthorized response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances snapshots getall unauthorized response
func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsGetallForbidden creates a PcloudPvminstancesSnapshotsGetallForbidden with default headers values
func NewPcloudPvminstancesSnapshotsGetallForbidden() *PcloudPvminstancesSnapshotsGetallForbidden {
	return &PcloudPvminstancesSnapshotsGetallForbidden{}
}

/*
PcloudPvminstancesSnapshotsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesSnapshotsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall forbidden response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall forbidden response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall forbidden response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots getall forbidden response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots getall forbidden response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances snapshots getall forbidden response
func (o *PcloudPvminstancesSnapshotsGetallForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesSnapshotsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsGetallNotFound creates a PcloudPvminstancesSnapshotsGetallNotFound with default headers values
func NewPcloudPvminstancesSnapshotsGetallNotFound() *PcloudPvminstancesSnapshotsGetallNotFound {
	return &PcloudPvminstancesSnapshotsGetallNotFound{}
}

/*
PcloudPvminstancesSnapshotsGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesSnapshotsGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall not found response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall not found response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall not found response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances snapshots getall not found response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances snapshots getall not found response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances snapshots getall not found response
func (o *PcloudPvminstancesSnapshotsGetallNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesSnapshotsGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesSnapshotsGetallInternalServerError creates a PcloudPvminstancesSnapshotsGetallInternalServerError with default headers values
func NewPcloudPvminstancesSnapshotsGetallInternalServerError() *PcloudPvminstancesSnapshotsGetallInternalServerError {
	return &PcloudPvminstancesSnapshotsGetallInternalServerError{}
}

/*
PcloudPvminstancesSnapshotsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesSnapshotsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances snapshots getall internal server error response has a 2xx status code
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances snapshots getall internal server error response has a 3xx status code
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances snapshots getall internal server error response has a 4xx status code
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances snapshots getall internal server error response has a 5xx status code
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances snapshots getall internal server error response a status code equal to that given
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances snapshots getall internal server error response
func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/snapshots][%d] pcloudPvminstancesSnapshotsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesSnapshotsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
