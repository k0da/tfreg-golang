// Code generated by MockGen. DO NOT EDIT.
// Source: internal/location/location.go

// Package location is a generated GoMock package.
package location

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocationer is a mock of Locationer interface.
type MockLocationer struct {
	ctrl     *gomock.Controller
	recorder *MockLocationerMockRecorder
}

// MockLocationerMockRecorder is the mock recorder for MockLocationer.
type MockLocationerMockRecorder struct {
	mock *MockLocationer
}

// NewMockLocationer creates a new mock instance.
func NewMockLocationer(ctrl *gomock.Controller) *MockLocationer {
	mock := &MockLocationer{ctrl: ctrl}
	mock.recorder = &MockLocationerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationer) EXPECT() *MockLocationerMockRecorder {
	return m.recorder
}

// ArtifactsPath mocks base method.
func (m *MockLocationer) ArtifactsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArtifactsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// ArtifactsPath indicates an expected call of ArtifactsPath.
func (mr *MockLocationerMockRecorder) ArtifactsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArtifactsPath", reflect.TypeOf((*MockLocationer)(nil).ArtifactsPath))
}

// BinariesPath mocks base method.
func (m *MockLocationer) BinariesPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BinariesPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// BinariesPath indicates an expected call of BinariesPath.
func (mr *MockLocationerMockRecorder) BinariesPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BinariesPath", reflect.TypeOf((*MockLocationer)(nil).BinariesPath))
}

// DownloadsPath mocks base method.
func (m *MockLocationer) DownloadsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// DownloadsPath indicates an expected call of DownloadsPath.
func (mr *MockLocationerMockRecorder) DownloadsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadsPath", reflect.TypeOf((*MockLocationer)(nil).DownloadsPath))
}

// GPGFingerprint mocks base method.
func (m *MockLocationer) GPGFingerprint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPGFingerprint")
	ret0, _ := ret[0].(string)
	return ret0
}

// GPGFingerprint indicates an expected call of GPGFingerprint.
func (mr *MockLocationerMockRecorder) GPGFingerprint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPGFingerprint", reflect.TypeOf((*MockLocationer)(nil).GPGFingerprint))
}

// GPGPubring mocks base method.
func (m *MockLocationer) GPGPubring() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPGPubring")
	ret0, _ := ret[0].(string)
	return ret0
}

// GPGPubring indicates an expected call of GPGPubring.
func (mr *MockLocationerMockRecorder) GPGPubring() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPGPubring", reflect.TypeOf((*MockLocationer)(nil).GPGPubring))
}

// GetArtifacts mocks base method.
func (m *MockLocationer) GetArtifacts() []Artifact {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArtifacts")
	ret0, _ := ret[0].([]Artifact)
	return ret0
}

// GetArtifacts indicates an expected call of GetArtifacts.
func (mr *MockLocationerMockRecorder) GetArtifacts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtifacts", reflect.TypeOf((*MockLocationer)(nil).GetArtifacts))
}

// GetShaSumFile mocks base method.
func (m *MockLocationer) GetShaSumFile() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShaSumFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetShaSumFile indicates an expected call of GetShaSumFile.
func (mr *MockLocationerMockRecorder) GetShaSumFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShaSumFile", reflect.TypeOf((*MockLocationer)(nil).GetShaSumFile))
}

// GetShaSumSignatureFile mocks base method.
func (m *MockLocationer) GetShaSumSignatureFile() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShaSumSignatureFile")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetShaSumSignatureFile indicates an expected call of GetShaSumSignatureFile.
func (mr *MockLocationerMockRecorder) GetShaSumSignatureFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShaSumSignatureFile", reflect.TypeOf((*MockLocationer)(nil).GetShaSumSignatureFile))
}

// GetVersion mocks base method.
func (m *MockLocationer) GetVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockLocationerMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockLocationer)(nil).GetVersion))
}

// TargetsPath mocks base method.
func (m *MockLocationer) TargetsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// TargetsPath indicates an expected call of TargetsPath.
func (mr *MockLocationerMockRecorder) TargetsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetsPath", reflect.TypeOf((*MockLocationer)(nil).TargetsPath))
}

// UrlBinaries mocks base method.
func (m *MockLocationer) UrlBinaries() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UrlBinaries")
	ret0, _ := ret[0].(string)
	return ret0
}

// UrlBinaries indicates an expected call of UrlBinaries.
func (mr *MockLocationerMockRecorder) UrlBinaries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UrlBinaries", reflect.TypeOf((*MockLocationer)(nil).UrlBinaries))
}

// VersionsPath mocks base method.
func (m *MockLocationer) VersionsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// VersionsPath indicates an expected call of VersionsPath.
func (mr *MockLocationerMockRecorder) VersionsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionsPath", reflect.TypeOf((*MockLocationer)(nil).VersionsPath))
}