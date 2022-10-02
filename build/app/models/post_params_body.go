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

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// api
	// Required: true
	API *PostParamsBodyAPI `json:"api"`

	// auth
	// Required: true
	Auth *PostParamsBodyAuth `json:"auth"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) validateAPI(formats strfmt.Registry) error {

	if err := validate.Required("api", "body", m.API); err != nil {
		return err
	}

	if m.API != nil {
		if err := m.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *PostParamsBody) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post params body based on the context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if m.API != nil {
		if err := m.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("api")
			}
			return err
		}
	}

	return nil
}

func (m *PostParamsBody) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.Auth != nil {
		if err := m.Auth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}