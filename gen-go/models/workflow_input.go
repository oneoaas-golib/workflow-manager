package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// WorkflowInput workflow input
// swagger:model WorkflowInput
type WorkflowInput struct {

	// data
	Data []interface{} `json:"data"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// queue
	Queue string `json:"queue,omitempty"`

	// workflow definition
	WorkflowDefinition *WorkflowDefinitionRef `json:"workflowDefinition,omitempty"`
}

// Validate validates this workflow input
func (m *WorkflowInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkflowDefinition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowInput) validateWorkflowDefinition(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowDefinition) { // not required
		return nil
	}

	if m.WorkflowDefinition != nil {

		if err := m.WorkflowDefinition.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
