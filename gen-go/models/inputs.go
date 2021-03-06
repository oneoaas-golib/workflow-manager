package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// These imports may not be used depending on the input parameters
var _ = json.Marshal
var _ = fmt.Sprintf
var _ = url.QueryEscape
var _ = strconv.FormatInt
var _ = strings.Replace
var _ = validate.Maximum
var _ = strfmt.NewFormats

// HealthCheckInput holds the input parameters for a healthCheck operation.
type HealthCheckInput struct {
}

// Validate returns an error if any of the HealthCheckInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i HealthCheckInput) Validate() error {
	return nil
}

// Path returns the URI path for the input.
func (i HealthCheckInput) Path() (string, error) {
	path := "/_health"
	urlVals := url.Values{}

	return path + "?" + urlVals.Encode(), nil
}

// DeleteStateResourceInput holds the input parameters for a deleteStateResource operation.
type DeleteStateResourceInput struct {
	Namespace string
	Name      string
}

// Validate returns an error if any of the DeleteStateResourceInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i DeleteStateResourceInput) Validate() error {

	return nil
}

// Path returns the URI path for the input.
func (i DeleteStateResourceInput) Path() (string, error) {
	path := "/state-resources/{namespace}/{name}"
	urlVals := url.Values{}

	pathnamespace := i.Namespace
	if pathnamespace == "" {
		err := fmt.Errorf("namespace cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{namespace}", pathnamespace, -1)

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	return path + "?" + urlVals.Encode(), nil
}

// GetStateResourceInput holds the input parameters for a getStateResource operation.
type GetStateResourceInput struct {
	Namespace string
	Name      string
}

// Validate returns an error if any of the GetStateResourceInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i GetStateResourceInput) Validate() error {

	return nil
}

// Path returns the URI path for the input.
func (i GetStateResourceInput) Path() (string, error) {
	path := "/state-resources/{namespace}/{name}"
	urlVals := url.Values{}

	pathnamespace := i.Namespace
	if pathnamespace == "" {
		err := fmt.Errorf("namespace cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{namespace}", pathnamespace, -1)

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	return path + "?" + urlVals.Encode(), nil
}

// PutStateResourceInput holds the input parameters for a putStateResource operation.
type PutStateResourceInput struct {
	Namespace        string
	Name             string
	NewStateResource *NewStateResource
}

// Validate returns an error if any of the PutStateResourceInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i PutStateResourceInput) Validate() error {

	if err := i.NewStateResource.Validate(nil); err != nil {
		return err
	}
	return nil
}

// Path returns the URI path for the input.
func (i PutStateResourceInput) Path() (string, error) {
	path := "/state-resources/{namespace}/{name}"
	urlVals := url.Values{}

	pathnamespace := i.Namespace
	if pathnamespace == "" {
		err := fmt.Errorf("namespace cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{namespace}", pathnamespace, -1)

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	return path + "?" + urlVals.Encode(), nil
}

// GetWorkflowDefinitionsInput holds the input parameters for a getWorkflowDefinitions operation.
type GetWorkflowDefinitionsInput struct {
}

// Validate returns an error if any of the GetWorkflowDefinitionsInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i GetWorkflowDefinitionsInput) Validate() error {
	return nil
}

// Path returns the URI path for the input.
func (i GetWorkflowDefinitionsInput) Path() (string, error) {
	path := "/workflow-definitions"
	urlVals := url.Values{}

	return path + "?" + urlVals.Encode(), nil
}

// GetWorkflowDefinitionVersionsByNameInput holds the input parameters for a getWorkflowDefinitionVersionsByName operation.
type GetWorkflowDefinitionVersionsByNameInput struct {
	Name   string
	Latest *bool
}

// Validate returns an error if any of the GetWorkflowDefinitionVersionsByNameInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i GetWorkflowDefinitionVersionsByNameInput) Validate() error {

	return nil
}

// Path returns the URI path for the input.
func (i GetWorkflowDefinitionVersionsByNameInput) Path() (string, error) {
	path := "/workflow-definitions/{name}"
	urlVals := url.Values{}

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	if i.Latest != nil {
		urlVals.Add("latest", strconv.FormatBool(*i.Latest))
	}

	return path + "?" + urlVals.Encode(), nil
}

// UpdateWorkflowDefinitionInput holds the input parameters for a updateWorkflowDefinition operation.
type UpdateWorkflowDefinitionInput struct {
	NewWorkflowDefinitionRequest *NewWorkflowDefinitionRequest
	Name                         string
}

// Validate returns an error if any of the UpdateWorkflowDefinitionInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i UpdateWorkflowDefinitionInput) Validate() error {

	if err := i.NewWorkflowDefinitionRequest.Validate(nil); err != nil {
		return err
	}

	return nil
}

// Path returns the URI path for the input.
func (i UpdateWorkflowDefinitionInput) Path() (string, error) {
	path := "/workflow-definitions/{name}"
	urlVals := url.Values{}

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	return path + "?" + urlVals.Encode(), nil
}

// GetWorkflowDefinitionByNameAndVersionInput holds the input parameters for a getWorkflowDefinitionByNameAndVersion operation.
type GetWorkflowDefinitionByNameAndVersionInput struct {
	Name    string
	Version int64
}

// Validate returns an error if any of the GetWorkflowDefinitionByNameAndVersionInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i GetWorkflowDefinitionByNameAndVersionInput) Validate() error {

	return nil
}

// Path returns the URI path for the input.
func (i GetWorkflowDefinitionByNameAndVersionInput) Path() (string, error) {
	path := "/workflow-definitions/{name}/{version}"
	urlVals := url.Values{}

	pathname := i.Name
	if pathname == "" {
		err := fmt.Errorf("name cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{name}", pathname, -1)

	pathversion := strconv.FormatInt(i.Version, 10)
	if pathversion == "" {
		err := fmt.Errorf("version cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{version}", pathversion, -1)

	return path + "?" + urlVals.Encode(), nil
}

// GetWorkflowsInput holds the input parameters for a getWorkflows operation.
type GetWorkflowsInput struct {
	WorkflowDefinitionName string
}

// Validate returns an error if any of the GetWorkflowsInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i GetWorkflowsInput) Validate() error {

	return nil
}

// Path returns the URI path for the input.
func (i GetWorkflowsInput) Path() (string, error) {
	path := "/workflows"
	urlVals := url.Values{}

	urlVals.Add("workflowDefinitionName", i.WorkflowDefinitionName)

	return path + "?" + urlVals.Encode(), nil
}

// CancelWorkflowInput holds the input parameters for a CancelWorkflow operation.
type CancelWorkflowInput struct {
	WorkflowId string
	Reason     *CancelReason
}

// Validate returns an error if any of the CancelWorkflowInput parameters don't satisfy the
// requirements from the swagger yml file.
func (i CancelWorkflowInput) Validate() error {

	if err := i.Reason.Validate(nil); err != nil {
		return err
	}
	return nil
}

// Path returns the URI path for the input.
func (i CancelWorkflowInput) Path() (string, error) {
	path := "/workflows/{workflowId}"
	urlVals := url.Values{}

	pathworkflowId := i.WorkflowId
	if pathworkflowId == "" {
		err := fmt.Errorf("workflowId cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{workflowId}", pathworkflowId, -1)

	return path + "?" + urlVals.Encode(), nil
}

// GetWorkflowByIDInput holds the input parameters for a getWorkflowByID operation.
type GetWorkflowByIDInput struct {
	WorkflowId string
}

// ValidateGetWorkflowByIDInput returns an error if the input parameter doesn't
// satisfy the requirements in the swagger yml file.
func ValidateGetWorkflowByIDInput(workflowId string) error {

	return nil
}

// GetWorkflowByIDInputPath returns the URI path for the input.
func GetWorkflowByIDInputPath(workflowId string) (string, error) {
	path := "/workflows/{workflowId}"
	urlVals := url.Values{}

	pathworkflowId := workflowId
	if pathworkflowId == "" {
		err := fmt.Errorf("workflowId cannot be empty because it's a path parameter")
		if err != nil {
			return "", err
		}
	}
	path = strings.Replace(path, "{workflowId}", pathworkflowId, -1)

	return path + "?" + urlVals.Encode(), nil
}
