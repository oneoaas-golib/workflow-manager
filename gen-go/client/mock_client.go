// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

package client

import (
	context "context"
	models "github.com/Clever/workflow-manager/gen-go/models"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockClient) EXPECT() *MockClientMockRecorder {
	return _m.recorder
}

// HealthCheck mocks base method
func (_m *MockClient) HealthCheck(ctx context.Context) error {
	ret := _m.ctrl.Call(_m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (_mr *MockClientMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HealthCheck", arg0)
}

// GetJobsForWorkflow mocks base method
func (_m *MockClient) GetJobsForWorkflow(ctx context.Context, i *models.GetJobsForWorkflowInput) ([]models.Job, error) {
	ret := _m.ctrl.Call(_m, "GetJobsForWorkflow", ctx, i)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobsForWorkflow indicates an expected call of GetJobsForWorkflow
func (_mr *MockClientMockRecorder) GetJobsForWorkflow(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobsForWorkflow", arg0, arg1)
}

// StartJobForWorkflow mocks base method
func (_m *MockClient) StartJobForWorkflow(ctx context.Context, i *models.JobInput) (*models.Job, error) {
	ret := _m.ctrl.Call(_m, "StartJobForWorkflow", ctx, i)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartJobForWorkflow indicates an expected call of StartJobForWorkflow
func (_mr *MockClientMockRecorder) StartJobForWorkflow(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartJobForWorkflow", arg0, arg1)
}

// CancelJob mocks base method
func (_m *MockClient) CancelJob(ctx context.Context, i *models.CancelJobInput) error {
	ret := _m.ctrl.Call(_m, "CancelJob", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelJob indicates an expected call of CancelJob
func (_mr *MockClientMockRecorder) CancelJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelJob", arg0, arg1)
}

// GetJob mocks base method
func (_m *MockClient) GetJob(ctx context.Context, jobId string) (*models.Job, error) {
	ret := _m.ctrl.Call(_m, "GetJob", ctx, jobId)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (_mr *MockClientMockRecorder) GetJob(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJob", arg0, arg1)
}

// PostStateResource mocks base method
func (_m *MockClient) PostStateResource(ctx context.Context, i *models.NewStateResource) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "PostStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStateResource indicates an expected call of PostStateResource
func (_mr *MockClientMockRecorder) PostStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostStateResource", arg0, arg1)
}

// DeleteStateResource mocks base method
func (_m *MockClient) DeleteStateResource(ctx context.Context, i *models.DeleteStateResourceInput) error {
	ret := _m.ctrl.Call(_m, "DeleteStateResource", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStateResource indicates an expected call of DeleteStateResource
func (_mr *MockClientMockRecorder) DeleteStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStateResource", arg0, arg1)
}

// GetStateResource mocks base method
func (_m *MockClient) GetStateResource(ctx context.Context, i *models.GetStateResourceInput) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "GetStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateResource indicates an expected call of GetStateResource
func (_mr *MockClientMockRecorder) GetStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStateResource", arg0, arg1)
}

// PutStateResource mocks base method
func (_m *MockClient) PutStateResource(ctx context.Context, i *models.PutStateResourceInput) (*models.StateResource, error) {
	ret := _m.ctrl.Call(_m, "PutStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutStateResource indicates an expected call of PutStateResource
func (_mr *MockClientMockRecorder) PutStateResource(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutStateResource", arg0, arg1)
}

// GetWorkflows mocks base method
func (_m *MockClient) GetWorkflows(ctx context.Context) ([]models.Workflow, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflows", ctx)
	ret0, _ := ret[0].([]models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflows indicates an expected call of GetWorkflows
func (_mr *MockClientMockRecorder) GetWorkflows(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflows", arg0)
}

// NewWorkflow mocks base method
func (_m *MockClient) NewWorkflow(ctx context.Context, i *models.NewWorkflowRequest) (*models.Workflow, error) {
	ret := _m.ctrl.Call(_m, "NewWorkflow", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWorkflow indicates an expected call of NewWorkflow
func (_mr *MockClientMockRecorder) NewWorkflow(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewWorkflow", arg0, arg1)
}

// GetWorkflowVersionsByName mocks base method
func (_m *MockClient) GetWorkflowVersionsByName(ctx context.Context, i *models.GetWorkflowVersionsByNameInput) ([]models.Workflow, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflowVersionsByName", ctx, i)
	ret0, _ := ret[0].([]models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowVersionsByName indicates an expected call of GetWorkflowVersionsByName
func (_mr *MockClientMockRecorder) GetWorkflowVersionsByName(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflowVersionsByName", arg0, arg1)
}

// UpdateWorkflow mocks base method
func (_m *MockClient) UpdateWorkflow(ctx context.Context, i *models.UpdateWorkflowInput) (*models.Workflow, error) {
	ret := _m.ctrl.Call(_m, "UpdateWorkflow", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflow indicates an expected call of UpdateWorkflow
func (_mr *MockClientMockRecorder) UpdateWorkflow(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateWorkflow", arg0, arg1)
}

// GetWorkflowByNameAndVersion mocks base method
func (_m *MockClient) GetWorkflowByNameAndVersion(ctx context.Context, i *models.GetWorkflowByNameAndVersionInput) (*models.Workflow, error) {
	ret := _m.ctrl.Call(_m, "GetWorkflowByNameAndVersion", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByNameAndVersion indicates an expected call of GetWorkflowByNameAndVersion
func (_mr *MockClientMockRecorder) GetWorkflowByNameAndVersion(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWorkflowByNameAndVersion", arg0, arg1)
}
