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

// DeleteSiteReader is a Reader for the DeleteSite structure.
type DeleteSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteOK creates a DeleteSiteOK with default headers values
func NewDeleteSiteOK() *DeleteSiteOK {
	return &DeleteSiteOK{}
}

/*DeleteSiteOK handles this case with default header values.

OK
*/
type DeleteSiteOK struct {
}

func (o *DeleteSiteOK) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}][%d] deleteSiteOK ", 200)
}

func (o *DeleteSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteDefault creates a DeleteSiteDefault with default headers values
func NewDeleteSiteDefault(code int) *DeleteSiteDefault {
	return &DeleteSiteDefault{
		_statusCode: code,
	}
}

/*DeleteSiteDefault handles this case with default header values.

error
*/
type DeleteSiteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete site default response
func (o *DeleteSiteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteDefault) Error() string {
	return fmt.Sprintf("[DELETE /sites/{site_id}][%d] deleteSite default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
