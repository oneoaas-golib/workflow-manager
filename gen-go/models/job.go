package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// Job job
// swagger:model Job
type Job struct {

	// container
	Container string `json:"Container,omitempty"`

	// attempts
	Attempts []*JobAttempt `json:"attempts"`

	// created at
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// started at
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"statusReason,omitempty"`

	// stopped at
	StoppedAt strfmt.DateTime `json:"stoppedAt,omitempty"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttempts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateAttempts(formats strfmt.Registry) error {

	if swag.IsZero(m.Attempts) { // not required
		return nil
	}

	for i := 0; i < len(m.Attempts); i++ {

		if swag.IsZero(m.Attempts[i]) { // not required
			continue
		}

		if m.Attempts[i] != nil {

			if err := m.Attempts[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}