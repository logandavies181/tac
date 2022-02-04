// Code generated by MockGen. DO NOT EDIT.
// Source: mocks/mocks.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockConfigurationVersions is a mock of ConfigurationVersions interface.
type MockConfigurationVersions struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationVersionsMockRecorder
}

// MockConfigurationVersionsMockRecorder is the mock recorder for MockConfigurationVersions.
type MockConfigurationVersionsMockRecorder struct {
	mock *MockConfigurationVersions
}

// NewMockConfigurationVersions creates a new mock instance.
func NewMockConfigurationVersions(ctrl *gomock.Controller) *MockConfigurationVersions {
	mock := &MockConfigurationVersions{ctrl: ctrl}
	mock.recorder = &MockConfigurationVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationVersions) EXPECT() *MockConfigurationVersionsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockConfigurationVersions) Create(ctx context.Context, workspaceID string, options tfe.ConfigurationVersionCreateOptions) (*tfe.ConfigurationVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.ConfigurationVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockConfigurationVersionsMockRecorder) Create(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigurationVersions)(nil).Create), ctx, workspaceID, options)
}

// List mocks base method.
func (m *MockConfigurationVersions) List(ctx context.Context, workspaceID string, options tfe.ConfigurationVersionListOptions) (*tfe.ConfigurationVersionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.ConfigurationVersionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockConfigurationVersionsMockRecorder) List(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfigurationVersions)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockConfigurationVersions) Read(ctx context.Context, cvID string) (*tfe.ConfigurationVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, cvID)
	ret0, _ := ret[0].(*tfe.ConfigurationVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockConfigurationVersionsMockRecorder) Read(ctx, cvID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConfigurationVersions)(nil).Read), ctx, cvID)
}

// Upload mocks base method.
func (m *MockConfigurationVersions) Upload(ctx context.Context, url, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, url, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockConfigurationVersionsMockRecorder) Upload(ctx, url, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockConfigurationVersions)(nil).Upload), ctx, url, path)
}

// MockPlans is a mock of Plans interface.
type MockPlans struct {
	ctrl     *gomock.Controller
	recorder *MockPlansMockRecorder
}

// MockPlansMockRecorder is the mock recorder for MockPlans.
type MockPlansMockRecorder struct {
	mock *MockPlans
}

// NewMockPlans creates a new mock instance.
func NewMockPlans(ctrl *gomock.Controller) *MockPlans {
	mock := &MockPlans{ctrl: ctrl}
	mock.recorder = &MockPlansMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlans) EXPECT() *MockPlansMockRecorder {
	return m.recorder
}

// JSONOutput mocks base method.
func (m *MockPlans) JSONOutput(ctx context.Context, planID string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSONOutput", ctx, planID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JSONOutput indicates an expected call of JSONOutput.
func (mr *MockPlansMockRecorder) JSONOutput(ctx, planID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONOutput", reflect.TypeOf((*MockPlans)(nil).JSONOutput), ctx, planID)
}

// Logs mocks base method.
func (m *MockPlans) Logs(ctx context.Context, planID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, planID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockPlansMockRecorder) Logs(ctx, planID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockPlans)(nil).Logs), ctx, planID)
}

// Read mocks base method.
func (m *MockPlans) Read(ctx context.Context, planID string) (*tfe.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, planID)
	ret0, _ := ret[0].(*tfe.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockPlansMockRecorder) Read(ctx, planID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPlans)(nil).Read), ctx, planID)
}

// MockRuns is a mock of Runs interface.
type MockRuns struct {
	ctrl     *gomock.Controller
	recorder *MockRunsMockRecorder
}

// MockRunsMockRecorder is the mock recorder for MockRuns.
type MockRunsMockRecorder struct {
	mock *MockRuns
}

// NewMockRuns creates a new mock instance.
func NewMockRuns(ctrl *gomock.Controller) *MockRuns {
	mock := &MockRuns{ctrl: ctrl}
	mock.recorder = &MockRunsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuns) EXPECT() *MockRunsMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockRuns) Apply(ctx context.Context, runID string, options tfe.RunApplyOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockRunsMockRecorder) Apply(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockRuns)(nil).Apply), ctx, runID, options)
}

// Cancel mocks base method.
func (m *MockRuns) Cancel(ctx context.Context, runID string, options tfe.RunCancelOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockRunsMockRecorder) Cancel(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockRuns)(nil).Cancel), ctx, runID, options)
}

// Create mocks base method.
func (m *MockRuns) Create(ctx context.Context, options tfe.RunCreateOptions) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRunsMockRecorder) Create(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRuns)(nil).Create), ctx, options)
}

// Discard mocks base method.
func (m *MockRuns) Discard(ctx context.Context, runID string, options tfe.RunDiscardOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockRunsMockRecorder) Discard(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockRuns)(nil).Discard), ctx, runID, options)
}

// ForceCancel mocks base method.
func (m *MockRuns) ForceCancel(ctx context.Context, runID string, options tfe.RunForceCancelOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceCancel", ctx, runID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceCancel indicates an expected call of ForceCancel.
func (mr *MockRunsMockRecorder) ForceCancel(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceCancel", reflect.TypeOf((*MockRuns)(nil).ForceCancel), ctx, runID, options)
}

// List mocks base method.
func (m *MockRuns) List(ctx context.Context, workspaceID string, options tfe.RunListOptions) (*tfe.RunList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.RunList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRunsMockRecorder) List(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRuns)(nil).List), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockRuns) Read(ctx context.Context, runID string) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, runID)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRunsMockRecorder) Read(ctx, runID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRuns)(nil).Read), ctx, runID)
}

// ReadWithOptions mocks base method.
func (m *MockRuns) ReadWithOptions(ctx context.Context, runID string, options *tfe.RunReadOptions) (*tfe.Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, runID, options)
	ret0, _ := ret[0].(*tfe.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockRunsMockRecorder) ReadWithOptions(ctx, runID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockRuns)(nil).ReadWithOptions), ctx, runID, options)
}

// MockWorkspaces is a mock of Workspaces interface.
type MockWorkspaces struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspacesMockRecorder
}

// MockWorkspacesMockRecorder is the mock recorder for MockWorkspaces.
type MockWorkspacesMockRecorder struct {
	mock *MockWorkspaces
}

// NewMockWorkspaces creates a new mock instance.
func NewMockWorkspaces(ctrl *gomock.Controller) *MockWorkspaces {
	mock := &MockWorkspaces{ctrl: ctrl}
	mock.recorder = &MockWorkspacesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaces) EXPECT() *MockWorkspacesMockRecorder {
	return m.recorder
}

// AddRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) AddRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceAddRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRemoteStateConsumers indicates an expected call of AddRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) AddRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).AddRemoteStateConsumers), ctx, workspaceID, options)
}

// AssignSSHKey mocks base method.
func (m *MockWorkspaces) AssignSSHKey(ctx context.Context, workspaceID string, options tfe.WorkspaceAssignSSHKeyOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignSSHKey", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignSSHKey indicates an expected call of AssignSSHKey.
func (mr *MockWorkspacesMockRecorder) AssignSSHKey(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).AssignSSHKey), ctx, workspaceID, options)
}

// Create mocks base method.
func (m *MockWorkspaces) Create(ctx context.Context, organization string, options tfe.WorkspaceCreateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWorkspacesMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWorkspaces)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockWorkspaces) Delete(ctx context.Context, organization, workspace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization, workspace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWorkspacesMockRecorder) Delete(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWorkspaces)(nil).Delete), ctx, organization, workspace)
}

// DeleteByID mocks base method.
func (m *MockWorkspaces) DeleteByID(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockWorkspacesMockRecorder) DeleteByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockWorkspaces)(nil).DeleteByID), ctx, workspaceID)
}

// ForceUnlock mocks base method.
func (m *MockWorkspaces) ForceUnlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUnlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForceUnlock indicates an expected call of ForceUnlock.
func (mr *MockWorkspacesMockRecorder) ForceUnlock(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUnlock", reflect.TypeOf((*MockWorkspaces)(nil).ForceUnlock), ctx, workspaceID)
}

// List mocks base method.
func (m *MockWorkspaces) List(ctx context.Context, organization string, options tfe.WorkspaceListOptions) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkspacesMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkspaces)(nil).List), ctx, organization, options)
}

// Lock mocks base method.
func (m *MockWorkspaces) Lock(ctx context.Context, workspaceID string, options tfe.WorkspaceLockOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockWorkspacesMockRecorder) Lock(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockWorkspaces)(nil).Lock), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockWorkspaces) Read(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockWorkspacesMockRecorder) Read(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockWorkspaces)(nil).Read), ctx, organization, workspace)
}

// ReadByID mocks base method.
func (m *MockWorkspaces) ReadByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByID indicates an expected call of ReadByID.
func (mr *MockWorkspacesMockRecorder) ReadByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByID", reflect.TypeOf((*MockWorkspaces)(nil).ReadByID), ctx, workspaceID)
}

// Readme mocks base method.
func (m *MockWorkspaces) Readme(ctx context.Context, workspaceID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readme", ctx, workspaceID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readme indicates an expected call of Readme.
func (mr *MockWorkspacesMockRecorder) Readme(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readme", reflect.TypeOf((*MockWorkspaces)(nil).Readme), ctx, workspaceID)
}

// RemoteStateConsumers mocks base method.
func (m *MockWorkspaces) RemoteStateConsumers(ctx context.Context, workspaceID string) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteStateConsumers", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteStateConsumers indicates an expected call of RemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) RemoteStateConsumers(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).RemoteStateConsumers), ctx, workspaceID)
}

// RemoveRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) RemoveRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceRemoveRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRemoteStateConsumers indicates an expected call of RemoveRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) RemoveRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).RemoveRemoteStateConsumers), ctx, workspaceID, options)
}

// RemoveVCSConnection mocks base method.
func (m *MockWorkspaces) RemoveVCSConnection(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnection", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnection indicates an expected call of RemoveVCSConnection.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnection(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnection", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnection), ctx, organization, workspace)
}

// RemoveVCSConnectionByID mocks base method.
func (m *MockWorkspaces) RemoveVCSConnectionByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnectionByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnectionByID indicates an expected call of RemoveVCSConnectionByID.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnectionByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnectionByID", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnectionByID), ctx, workspaceID)
}

// UnassignSSHKey mocks base method.
func (m *MockWorkspaces) UnassignSSHKey(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignSSHKey", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnassignSSHKey indicates an expected call of UnassignSSHKey.
func (mr *MockWorkspacesMockRecorder) UnassignSSHKey(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).UnassignSSHKey), ctx, workspaceID)
}

// Unlock mocks base method.
func (m *MockWorkspaces) Unlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockWorkspacesMockRecorder) Unlock(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockWorkspaces)(nil).Unlock), ctx, workspaceID)
}

// Update mocks base method.
func (m *MockWorkspaces) Update(ctx context.Context, organization, workspace string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organization, workspace, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWorkspacesMockRecorder) Update(ctx, organization, workspace, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWorkspaces)(nil).Update), ctx, organization, workspace, options)
}

// UpdateByID mocks base method.
func (m *MockWorkspaces) UpdateByID(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockWorkspacesMockRecorder) UpdateByID(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockWorkspaces)(nil).UpdateByID), ctx, workspaceID, options)
}

// UpdateRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) UpdateRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteStateConsumers indicates an expected call of UpdateRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) UpdateRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).UpdateRemoteStateConsumers), ctx, workspaceID, options)
}
