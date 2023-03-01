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

// PcloudImagesGetallReader is a Reader for the PcloudImagesGetall structure.
type PcloudImagesGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudImagesGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudImagesGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudImagesGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudImagesGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudImagesGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudImagesGetallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudImagesGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudImagesGetallOK creates a PcloudImagesGetallOK with default headers values
func NewPcloudImagesGetallOK() *PcloudImagesGetallOK {
	return &PcloudImagesGetallOK{}
}

/*
PcloudImagesGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudImagesGetallOK struct {
	Payload *models.Images
}

// IsSuccess returns true when this pcloud images getall o k response has a 2xx status code
func (o *PcloudImagesGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud images getall o k response has a 3xx status code
func (o *PcloudImagesGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall o k response has a 4xx status code
func (o *PcloudImagesGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud images getall o k response has a 5xx status code
func (o *PcloudImagesGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images getall o k response a status code equal to that given
func (o *PcloudImagesGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudImagesGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudImagesGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudImagesGetallOK) GetPayload() *models.Images {
	return o.Payload
}

func (o *PcloudImagesGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetallBadRequest creates a PcloudImagesGetallBadRequest with default headers values
func NewPcloudImagesGetallBadRequest() *PcloudImagesGetallBadRequest {
	return &PcloudImagesGetallBadRequest{}
}

/*
PcloudImagesGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudImagesGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images getall bad request response has a 2xx status code
func (o *PcloudImagesGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images getall bad request response has a 3xx status code
func (o *PcloudImagesGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall bad request response has a 4xx status code
func (o *PcloudImagesGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images getall bad request response has a 5xx status code
func (o *PcloudImagesGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images getall bad request response a status code equal to that given
func (o *PcloudImagesGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudImagesGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudImagesGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudImagesGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetallUnauthorized creates a PcloudImagesGetallUnauthorized with default headers values
func NewPcloudImagesGetallUnauthorized() *PcloudImagesGetallUnauthorized {
	return &PcloudImagesGetallUnauthorized{}
}

/*
PcloudImagesGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudImagesGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images getall unauthorized response has a 2xx status code
func (o *PcloudImagesGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images getall unauthorized response has a 3xx status code
func (o *PcloudImagesGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall unauthorized response has a 4xx status code
func (o *PcloudImagesGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images getall unauthorized response has a 5xx status code
func (o *PcloudImagesGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images getall unauthorized response a status code equal to that given
func (o *PcloudImagesGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudImagesGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudImagesGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudImagesGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetallForbidden creates a PcloudImagesGetallForbidden with default headers values
func NewPcloudImagesGetallForbidden() *PcloudImagesGetallForbidden {
	return &PcloudImagesGetallForbidden{}
}

/*
PcloudImagesGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudImagesGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images getall forbidden response has a 2xx status code
func (o *PcloudImagesGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images getall forbidden response has a 3xx status code
func (o *PcloudImagesGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall forbidden response has a 4xx status code
func (o *PcloudImagesGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images getall forbidden response has a 5xx status code
func (o *PcloudImagesGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images getall forbidden response a status code equal to that given
func (o *PcloudImagesGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudImagesGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudImagesGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudImagesGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetallNotFound creates a PcloudImagesGetallNotFound with default headers values
func NewPcloudImagesGetallNotFound() *PcloudImagesGetallNotFound {
	return &PcloudImagesGetallNotFound{}
}

/*
PcloudImagesGetallNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudImagesGetallNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images getall not found response has a 2xx status code
func (o *PcloudImagesGetallNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images getall not found response has a 3xx status code
func (o *PcloudImagesGetallNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall not found response has a 4xx status code
func (o *PcloudImagesGetallNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud images getall not found response has a 5xx status code
func (o *PcloudImagesGetallNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud images getall not found response a status code equal to that given
func (o *PcloudImagesGetallNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudImagesGetallNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudImagesGetallNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallNotFound  %+v", 404, o.Payload)
}

func (o *PcloudImagesGetallNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudImagesGetallInternalServerError creates a PcloudImagesGetallInternalServerError with default headers values
func NewPcloudImagesGetallInternalServerError() *PcloudImagesGetallInternalServerError {
	return &PcloudImagesGetallInternalServerError{}
}

/*
PcloudImagesGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudImagesGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud images getall internal server error response has a 2xx status code
func (o *PcloudImagesGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud images getall internal server error response has a 3xx status code
func (o *PcloudImagesGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud images getall internal server error response has a 4xx status code
func (o *PcloudImagesGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud images getall internal server error response has a 5xx status code
func (o *PcloudImagesGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud images getall internal server error response a status code equal to that given
func (o *PcloudImagesGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudImagesGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudImagesGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/images][%d] pcloudImagesGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudImagesGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudImagesGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
