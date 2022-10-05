// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// GetClusterValuesReader is a Reader for the GetClusterValues structure.
type GetClusterValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterValuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterValuesOK creates a GetClusterValuesOK with default headers values
func NewGetClusterValuesOK() *GetClusterValuesOK {
	return &GetClusterValuesOK{}
}

/*
GetClusterValuesOK describes a response with status code 200, with default header values.

OK
*/
type GetClusterValuesOK struct {
	Payload *models.GetClusterValuesOKBody
}

// IsSuccess returns true when this get cluster values o k response has a 2xx status code
func (o *GetClusterValuesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster values o k response has a 3xx status code
func (o *GetClusterValuesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster values o k response has a 4xx status code
func (o *GetClusterValuesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster values o k response has a 5xx status code
func (o *GetClusterValuesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster values o k response a status code equal to that given
func (o *GetClusterValuesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterValuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{account}/clusters/{name}/values][%d] getClusterValuesOK  %+v", 200, o.Payload)
}

func (o *GetClusterValuesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/{account}/clusters/{name}/values][%d] getClusterValuesOK  %+v", 200, o.Payload)
}

func (o *GetClusterValuesOK) GetPayload() *models.GetClusterValuesOKBody {
	return o.Payload
}

func (o *GetClusterValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetClusterValuesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterValuesDefault creates a GetClusterValuesDefault with default headers values
func NewGetClusterValuesDefault(code int) *GetClusterValuesDefault {
	return &GetClusterValuesDefault{
		_statusCode: code,
	}
}

/*
GetClusterValuesDefault describes a response with status code -1, with default header values.

Error
*/
type GetClusterValuesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cluster values default response
func (o *GetClusterValuesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get cluster values default response has a 2xx status code
func (o *GetClusterValuesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster values default response has a 3xx status code
func (o *GetClusterValuesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster values default response has a 4xx status code
func (o *GetClusterValuesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster values default response has a 5xx status code
func (o *GetClusterValuesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster values default response a status code equal to that given
func (o *GetClusterValuesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetClusterValuesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{account}/clusters/{name}/values][%d] GetClusterValues default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterValuesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/{account}/clusters/{name}/values][%d] GetClusterValues default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterValuesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterValuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
