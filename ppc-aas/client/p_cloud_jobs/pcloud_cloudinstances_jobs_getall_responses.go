// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudCloudinstancesJobsGetallReader is a Reader for the PcloudCloudinstancesJobsGetall structure.
type PcloudCloudinstancesJobsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesJobsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesJobsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesJobsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesJobsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesJobsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesJobsGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesJobsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesJobsGetallOK creates a PcloudCloudinstancesJobsGetallOK with default headers values
func NewPcloudCloudinstancesJobsGetallOK() *PcloudCloudinstancesJobsGetallOK {
	return &PcloudCloudinstancesJobsGetallOK{}
}

/*
PcloudCloudinstancesJobsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesJobsGetallOK struct {
	Payload *models.Jobs
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall o k response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall o k response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall o k response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs getall o k response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs getall o k response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesJobsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallOK) GetPayload() *models.Jobs {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetallBadRequest creates a PcloudCloudinstancesJobsGetallBadRequest with default headers values
func NewPcloudCloudinstancesJobsGetallBadRequest() *PcloudCloudinstancesJobsGetallBadRequest {
	return &PcloudCloudinstancesJobsGetallBadRequest{}
}

/*
PcloudCloudinstancesJobsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesJobsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall bad request response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall bad request response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall bad request response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs getall bad request response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs getall bad request response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesJobsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetallUnauthorized creates a PcloudCloudinstancesJobsGetallUnauthorized with default headers values
func NewPcloudCloudinstancesJobsGetallUnauthorized() *PcloudCloudinstancesJobsGetallUnauthorized {
	return &PcloudCloudinstancesJobsGetallUnauthorized{}
}

/*
PcloudCloudinstancesJobsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesJobsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs getall unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs getall unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesJobsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetallForbidden creates a PcloudCloudinstancesJobsGetallForbidden with default headers values
func NewPcloudCloudinstancesJobsGetallForbidden() *PcloudCloudinstancesJobsGetallForbidden {
	return &PcloudCloudinstancesJobsGetallForbidden{}
}

/*
PcloudCloudinstancesJobsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesJobsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall forbidden response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall forbidden response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall forbidden response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs getall forbidden response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs getall forbidden response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudCloudinstancesJobsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetallNotFound creates a PcloudCloudinstancesJobsGetallNotFound with default headers values
func NewPcloudCloudinstancesJobsGetallNotFound() *PcloudCloudinstancesJobsGetallNotFound {
	return &PcloudCloudinstancesJobsGetallNotFound{}
}

/*
PcloudCloudinstancesJobsGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesJobsGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall not found response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall not found response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall not found response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances jobs getall not found response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances jobs getall not found response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudinstancesJobsGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetallInternalServerError creates a PcloudCloudinstancesJobsGetallInternalServerError with default headers values
func NewPcloudCloudinstancesJobsGetallInternalServerError() *PcloudCloudinstancesJobsGetallInternalServerError {
	return &PcloudCloudinstancesJobsGetallInternalServerError{}
}

/*
PcloudCloudinstancesJobsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesJobsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances jobs getall internal server error response has a 2xx status code
func (o *PcloudCloudinstancesJobsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances jobs getall internal server error response has a 3xx status code
func (o *PcloudCloudinstancesJobsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances jobs getall internal server error response has a 4xx status code
func (o *PcloudCloudinstancesJobsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances jobs getall internal server error response has a 5xx status code
func (o *PcloudCloudinstancesJobsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances jobs getall internal server error response a status code equal to that given
func (o *PcloudCloudinstancesJobsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesJobsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs][%d] pcloudCloudinstancesJobsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesJobsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
