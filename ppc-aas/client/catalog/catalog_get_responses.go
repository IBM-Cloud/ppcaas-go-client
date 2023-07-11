// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// CatalogGetReader is a Reader for the CatalogGet structure.
type CatalogGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v2/catalog] catalog.get", response, response.Code())
	}
}

// NewCatalogGetOK creates a CatalogGetOK with default headers values
func NewCatalogGetOK() *CatalogGetOK {
	return &CatalogGetOK{}
}

/*
CatalogGetOK describes a response with status code 200, with default header values.

catalog response
*/
type CatalogGetOK struct {
	Payload *models.Catalog
}

// IsSuccess returns true when this catalog get o k response has a 2xx status code
func (o *CatalogGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog get o k response has a 3xx status code
func (o *CatalogGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog get o k response has a 4xx status code
func (o *CatalogGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog get o k response has a 5xx status code
func (o *CatalogGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog get o k response a status code equal to that given
func (o *CatalogGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog get o k response
func (o *CatalogGetOK) Code() int {
	return 200
}

func (o *CatalogGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetOK  %+v", 200, o.Payload)
}

func (o *CatalogGetOK) String() string {
	return fmt.Sprintf("[GET /v2/catalog][%d] catalogGetOK  %+v", 200, o.Payload)
}

func (o *CatalogGetOK) GetPayload() *models.Catalog {
	return o.Payload
}

func (o *CatalogGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Catalog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
