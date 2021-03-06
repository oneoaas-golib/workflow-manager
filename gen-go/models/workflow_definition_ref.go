package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// WorkflowDefinitionRef workflow definition ref
// swagger:model WorkflowDefinitionRef
type WorkflowDefinitionRef struct {

	// name
	Name string `json:"name,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this workflow definition ref
func (m *WorkflowDefinitionRef) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
