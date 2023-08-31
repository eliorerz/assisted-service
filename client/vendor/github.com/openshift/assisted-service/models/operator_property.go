// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OperatorProperty operator property
//
// swagger:model operator-property
type OperatorProperty struct {

	// Type of the property
	// Enum: [boolean string integer float]
	DataType string `json:"data_type,omitempty"`

	// Default value for the property
	DefaultValue string `json:"default_value,omitempty"`

	// Description of a property
	Description string `json:"description,omitempty"`

	// Indicates whether the property is reqired
	Mandatory bool `json:"mandatory,omitempty"`

	// Name of the property
	Name string `json:"name,omitempty"`

	// Values to select from
	Options []string `json:"options"`
}

// Validate validates this operator property
func (m *OperatorProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var operatorPropertyTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["boolean","string","integer","float"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operatorPropertyTypeDataTypePropEnum = append(operatorPropertyTypeDataTypePropEnum, v)
	}
}

const (

	// OperatorPropertyDataTypeBoolean captures enum value "boolean"
	OperatorPropertyDataTypeBoolean string = "boolean"

	// OperatorPropertyDataTypeString captures enum value "string"
	OperatorPropertyDataTypeString string = "string"

	// OperatorPropertyDataTypeInteger captures enum value "integer"
	OperatorPropertyDataTypeInteger string = "integer"

	// OperatorPropertyDataTypeFloat captures enum value "float"
	OperatorPropertyDataTypeFloat string = "float"
)

// prop value enum
func (m *OperatorProperty) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, operatorPropertyTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OperatorProperty) validateDataType(formats strfmt.Registry) error {
	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("data_type", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this operator property based on context it is used
func (m *OperatorProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperatorProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatorProperty) UnmarshalBinary(b []byte) error {
	var res OperatorProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}