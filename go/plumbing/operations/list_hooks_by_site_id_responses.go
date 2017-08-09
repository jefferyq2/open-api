// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ListHooksBySiteIDReader is a Reader for the ListHooksBySiteID structure.
type ListHooksBySiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListHooksBySiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListHooksBySiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListHooksBySiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListHooksBySiteIDOK creates a ListHooksBySiteIDOK with default headers values
func NewListHooksBySiteIDOK() *ListHooksBySiteIDOK {
	return &ListHooksBySiteIDOK{}
}

/*ListHooksBySiteIDOK handles this case with default header values.

OK
*/
type ListHooksBySiteIDOK struct {
	Payload []*models.Hook
}

func (o *ListHooksBySiteIDOK) Error() string {
	return fmt.Sprintf("[GET /hooks][%d] listHooksBySiteIdOK  %+v", 200, o.Payload)
}

func (o *ListHooksBySiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListHooksBySiteIDDefault creates a ListHooksBySiteIDDefault with default headers values
func NewListHooksBySiteIDDefault(code int) *ListHooksBySiteIDDefault {
	return &ListHooksBySiteIDDefault{
		_statusCode: code,
	}
}

/*ListHooksBySiteIDDefault handles this case with default header values.

error
*/
type ListHooksBySiteIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list hooks by site Id default response
func (o *ListHooksBySiteIDDefault) Code() int {
	return o._statusCode
}

func (o *ListHooksBySiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /hooks][%d] listHooksBySiteId default  %+v", o._statusCode, o.Payload)
}

func (o *ListHooksBySiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
