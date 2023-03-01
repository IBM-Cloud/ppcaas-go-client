// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// PcloudCloudinstancesImagesGetReader is a Reader for the PcloudCloudinstancesImagesGet structure.
type PcloudCloudinstancesImagesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesImagesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesImagesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesImagesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesImagesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesImagesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesImagesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesImagesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesImagesGetOK creates a PcloudCloudinstancesImagesGetOK with default headers values
func NewPcloudCloudinstancesImagesGetOK() *PcloudCloudinstancesImagesGetOK {
	return &PcloudCloudinstancesImagesGetOK{}
}

/*
PcloudCloudinstancesImagesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesImagesGetOK struct {
	Payload *models.Image
}

// IsSuccess returns true when this pcloud cloudinstances images get o k response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances images get o k response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get o k response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images get o k response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images get o k response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesImagesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetOK) GetPayload() *models.Image {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesGetBadRequest creates a PcloudCloudinstancesImagesGetBadRequest with default headers values
func NewPcloudCloudinstancesImagesGetBadRequest() *PcloudCloudinstancesImagesGetBadRequest {
	return &PcloudCloudinstancesImagesGetBadRequest{}
}

/*
PcloudCloudinstancesImagesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesImagesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images get bad request response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images get bad request response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get bad request response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images get bad request response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images get bad request response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesImagesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesGetUnauthorized creates a PcloudCloudinstancesImagesGetUnauthorized with default headers values
func NewPcloudCloudinstancesImagesGetUnauthorized() *PcloudCloudinstancesImagesGetUnauthorized {
	return &PcloudCloudinstancesImagesGetUnauthorized{}
}

/*
PcloudCloudinstancesImagesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesImagesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesImagesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesGetForbidden creates a PcloudCloudinstancesImagesGetForbidden with default headers values
func NewPcloudCloudinstancesImagesGetForbidden() *PcloudCloudinstancesImagesGetForbidden {
	return &PcloudCloudinstancesImagesGetForbidden{}
}

/*
PcloudCloudinstancesImagesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesImagesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images get forbidden response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images get forbidden response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get forbidden response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images get forbidden response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images get forbidden response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudCloudinstancesImagesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesGetNotFound creates a PcloudCloudinstancesImagesGetNotFound with default headers values
func NewPcloudCloudinstancesImagesGetNotFound() *PcloudCloudinstancesImagesGetNotFound {
	return &PcloudCloudinstancesImagesGetNotFound{}
}

/*
PcloudCloudinstancesImagesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesImagesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images get not found response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images get not found response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get not found response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances images get not found response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances images get not found response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudinstancesImagesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesImagesGetInternalServerError creates a PcloudCloudinstancesImagesGetInternalServerError with default headers values
func NewPcloudCloudinstancesImagesGetInternalServerError() *PcloudCloudinstancesImagesGetInternalServerError {
	return &PcloudCloudinstancesImagesGetInternalServerError{}
}

/*
PcloudCloudinstancesImagesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesImagesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances images get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesImagesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances images get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesImagesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances images get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesImagesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances images get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesImagesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances images get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesImagesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesImagesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/images/{image_id}][%d] pcloudCloudinstancesImagesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesImagesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesImagesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
