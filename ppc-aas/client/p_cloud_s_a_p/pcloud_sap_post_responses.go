// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// PcloudSapPostReader is a Reader for the PcloudSapPost structure.
type PcloudSapPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSapPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudSapPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPcloudSapPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudSapPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudSapPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudSapPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudSapPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudSapPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudSapPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudSapPostOK creates a PcloudSapPostOK with default headers values
func NewPcloudSapPostOK() *PcloudSapPostOK {
	return &PcloudSapPostOK{}
}

/*
PcloudSapPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudSapPostOK struct {
	Payload models.PVMInstanceList
}

// IsSuccess returns true when this pcloud sap post o k response has a 2xx status code
func (o *PcloudSapPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sap post o k response has a 3xx status code
func (o *PcloudSapPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post o k response has a 4xx status code
func (o *PcloudSapPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap post o k response has a 5xx status code
func (o *PcloudSapPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post o k response a status code equal to that given
func (o *PcloudSapPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud sap post o k response
func (o *PcloudSapPostOK) Code() int {
	return 200
}

func (o *PcloudSapPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSapPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostOK  %+v", 200, o.Payload)
}

func (o *PcloudSapPostOK) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudSapPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostCreated creates a PcloudSapPostCreated with default headers values
func NewPcloudSapPostCreated() *PcloudSapPostCreated {
	return &PcloudSapPostCreated{}
}

/*
PcloudSapPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudSapPostCreated struct {
	Payload models.PVMInstanceList
}

// IsSuccess returns true when this pcloud sap post created response has a 2xx status code
func (o *PcloudSapPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sap post created response has a 3xx status code
func (o *PcloudSapPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post created response has a 4xx status code
func (o *PcloudSapPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap post created response has a 5xx status code
func (o *PcloudSapPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post created response a status code equal to that given
func (o *PcloudSapPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud sap post created response
func (o *PcloudSapPostCreated) Code() int {
	return 201
}

func (o *PcloudSapPostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudSapPostCreated) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostCreated  %+v", 201, o.Payload)
}

func (o *PcloudSapPostCreated) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudSapPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostAccepted creates a PcloudSapPostAccepted with default headers values
func NewPcloudSapPostAccepted() *PcloudSapPostAccepted {
	return &PcloudSapPostAccepted{}
}

/*
PcloudSapPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudSapPostAccepted struct {
	Payload models.PVMInstanceList
}

// IsSuccess returns true when this pcloud sap post accepted response has a 2xx status code
func (o *PcloudSapPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud sap post accepted response has a 3xx status code
func (o *PcloudSapPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post accepted response has a 4xx status code
func (o *PcloudSapPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap post accepted response has a 5xx status code
func (o *PcloudSapPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post accepted response a status code equal to that given
func (o *PcloudSapPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud sap post accepted response
func (o *PcloudSapPostAccepted) Code() int {
	return 202
}

func (o *PcloudSapPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudSapPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudSapPostAccepted) GetPayload() models.PVMInstanceList {
	return o.Payload
}

func (o *PcloudSapPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostBadRequest creates a PcloudSapPostBadRequest with default headers values
func NewPcloudSapPostBadRequest() *PcloudSapPostBadRequest {
	return &PcloudSapPostBadRequest{}
}

/*
PcloudSapPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudSapPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap post bad request response has a 2xx status code
func (o *PcloudSapPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap post bad request response has a 3xx status code
func (o *PcloudSapPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post bad request response has a 4xx status code
func (o *PcloudSapPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap post bad request response has a 5xx status code
func (o *PcloudSapPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post bad request response a status code equal to that given
func (o *PcloudSapPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud sap post bad request response
func (o *PcloudSapPostBadRequest) Code() int {
	return 400
}

func (o *PcloudSapPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostUnauthorized creates a PcloudSapPostUnauthorized with default headers values
func NewPcloudSapPostUnauthorized() *PcloudSapPostUnauthorized {
	return &PcloudSapPostUnauthorized{}
}

/*
PcloudSapPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudSapPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap post unauthorized response has a 2xx status code
func (o *PcloudSapPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap post unauthorized response has a 3xx status code
func (o *PcloudSapPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post unauthorized response has a 4xx status code
func (o *PcloudSapPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap post unauthorized response has a 5xx status code
func (o *PcloudSapPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post unauthorized response a status code equal to that given
func (o *PcloudSapPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud sap post unauthorized response
func (o *PcloudSapPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudSapPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudSapPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostConflict creates a PcloudSapPostConflict with default headers values
func NewPcloudSapPostConflict() *PcloudSapPostConflict {
	return &PcloudSapPostConflict{}
}

/*
PcloudSapPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudSapPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap post conflict response has a 2xx status code
func (o *PcloudSapPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap post conflict response has a 3xx status code
func (o *PcloudSapPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post conflict response has a 4xx status code
func (o *PcloudSapPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap post conflict response has a 5xx status code
func (o *PcloudSapPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post conflict response a status code equal to that given
func (o *PcloudSapPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud sap post conflict response
func (o *PcloudSapPostConflict) Code() int {
	return 409
}

func (o *PcloudSapPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSapPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudSapPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostUnprocessableEntity creates a PcloudSapPostUnprocessableEntity with default headers values
func NewPcloudSapPostUnprocessableEntity() *PcloudSapPostUnprocessableEntity {
	return &PcloudSapPostUnprocessableEntity{}
}

/*
PcloudSapPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudSapPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap post unprocessable entity response has a 2xx status code
func (o *PcloudSapPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap post unprocessable entity response has a 3xx status code
func (o *PcloudSapPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post unprocessable entity response has a 4xx status code
func (o *PcloudSapPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud sap post unprocessable entity response has a 5xx status code
func (o *PcloudSapPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud sap post unprocessable entity response a status code equal to that given
func (o *PcloudSapPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud sap post unprocessable entity response
func (o *PcloudSapPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudSapPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSapPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudSapPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapPostInternalServerError creates a PcloudSapPostInternalServerError with default headers values
func NewPcloudSapPostInternalServerError() *PcloudSapPostInternalServerError {
	return &PcloudSapPostInternalServerError{}
}

/*
PcloudSapPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudSapPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud sap post internal server error response has a 2xx status code
func (o *PcloudSapPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud sap post internal server error response has a 3xx status code
func (o *PcloudSapPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud sap post internal server error response has a 4xx status code
func (o *PcloudSapPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud sap post internal server error response has a 5xx status code
func (o *PcloudSapPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud sap post internal server error response a status code equal to that given
func (o *PcloudSapPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud sap post internal server error response
func (o *PcloudSapPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudSapPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudSapPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
