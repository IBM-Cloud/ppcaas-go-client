// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudV2VolumescloneGetallReader is a Reader for the PcloudV2VolumescloneGetall structure.
type PcloudV2VolumescloneGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudV2VolumescloneGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumescloneGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumescloneGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumescloneGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumescloneGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudV2VolumescloneGetallOK creates a PcloudV2VolumescloneGetallOK with default headers values
func NewPcloudV2VolumescloneGetallOK() *PcloudV2VolumescloneGetallOK {
	return &PcloudV2VolumescloneGetallOK{}
}

/*
PcloudV2VolumescloneGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudV2VolumescloneGetallOK struct {
	Payload *models.VolumesClones
}

// IsSuccess returns true when this pcloud v2 volumesclone getall o k response has a 2xx status code
func (o *PcloudV2VolumescloneGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumesclone getall o k response has a 3xx status code
func (o *PcloudV2VolumescloneGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone getall o k response has a 4xx status code
func (o *PcloudV2VolumescloneGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone getall o k response has a 5xx status code
func (o *PcloudV2VolumescloneGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone getall o k response a status code equal to that given
func (o *PcloudV2VolumescloneGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudV2VolumescloneGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneGetallOK) GetPayload() *models.VolumesClones {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClones)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallBadRequest creates a PcloudV2VolumescloneGetallBadRequest with default headers values
func NewPcloudV2VolumescloneGetallBadRequest() *PcloudV2VolumescloneGetallBadRequest {
	return &PcloudV2VolumescloneGetallBadRequest{}
}

/*
PcloudV2VolumescloneGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumescloneGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone getall bad request response has a 2xx status code
func (o *PcloudV2VolumescloneGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone getall bad request response has a 3xx status code
func (o *PcloudV2VolumescloneGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone getall bad request response has a 4xx status code
func (o *PcloudV2VolumescloneGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone getall bad request response has a 5xx status code
func (o *PcloudV2VolumescloneGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone getall bad request response a status code equal to that given
func (o *PcloudV2VolumescloneGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudV2VolumescloneGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumescloneGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallUnauthorized creates a PcloudV2VolumescloneGetallUnauthorized with default headers values
func NewPcloudV2VolumescloneGetallUnauthorized() *PcloudV2VolumescloneGetallUnauthorized {
	return &PcloudV2VolumescloneGetallUnauthorized{}
}

/*
PcloudV2VolumescloneGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumescloneGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone getall unauthorized response has a 2xx status code
func (o *PcloudV2VolumescloneGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone getall unauthorized response has a 3xx status code
func (o *PcloudV2VolumescloneGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone getall unauthorized response has a 4xx status code
func (o *PcloudV2VolumescloneGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone getall unauthorized response has a 5xx status code
func (o *PcloudV2VolumescloneGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone getall unauthorized response a status code equal to that given
func (o *PcloudV2VolumescloneGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudV2VolumescloneGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallForbidden creates a PcloudV2VolumescloneGetallForbidden with default headers values
func NewPcloudV2VolumescloneGetallForbidden() *PcloudV2VolumescloneGetallForbidden {
	return &PcloudV2VolumescloneGetallForbidden{}
}

/*
PcloudV2VolumescloneGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumescloneGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone getall forbidden response has a 2xx status code
func (o *PcloudV2VolumescloneGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone getall forbidden response has a 3xx status code
func (o *PcloudV2VolumescloneGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone getall forbidden response has a 4xx status code
func (o *PcloudV2VolumescloneGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone getall forbidden response has a 5xx status code
func (o *PcloudV2VolumescloneGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone getall forbidden response a status code equal to that given
func (o *PcloudV2VolumescloneGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudV2VolumescloneGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneGetallInternalServerError creates a PcloudV2VolumescloneGetallInternalServerError with default headers values
func NewPcloudV2VolumescloneGetallInternalServerError() *PcloudV2VolumescloneGetallInternalServerError {
	return &PcloudV2VolumescloneGetallInternalServerError{}
}

/*
PcloudV2VolumescloneGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone getall internal server error response has a 2xx status code
func (o *PcloudV2VolumescloneGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone getall internal server error response has a 3xx status code
func (o *PcloudV2VolumescloneGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone getall internal server error response has a 4xx status code
func (o *PcloudV2VolumescloneGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone getall internal server error response has a 5xx status code
func (o *PcloudV2VolumescloneGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumesclone getall internal server error response a status code equal to that given
func (o *PcloudV2VolumescloneGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudV2VolumescloneGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumescloneGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
