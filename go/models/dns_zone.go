// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DNSZone dns zone
// swagger:model dnsZone
type DNSZone struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// records
	Records []*DNSRecord `json:"records"`
}

// Validate validates this dns zone
func (m *DNSZone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNSZone) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {

		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {

			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DNSZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSZone) UnmarshalBinary(b []byte) error {
	var res DNSZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
