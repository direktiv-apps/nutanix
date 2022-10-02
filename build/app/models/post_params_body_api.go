// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostParamsBodyAPI post params body Api
//
// swagger:model postParamsBodyApi
type PostParamsBodyAPI struct {

	// body
	Body interface{} `json:"body,omitempty"`

	// HTTP method to use
	// Example: POST
	Method *string `json:"method,omitempty"`

	// API path to access
	// Example: /vms/list
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this post params body Api
func (m *PostParamsBodyAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBodyAPI) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post params body Api based on context it is used
func (m *PostParamsBodyAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyAPI) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}