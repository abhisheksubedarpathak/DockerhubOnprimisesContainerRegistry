// Code generated by mockery v2.46.2. DO NOT EDIT.

package v1

import (
	v1 "github.com/goharbor/harbor/src/pkg/scan/rest/v1"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetMetadata provides a mock function with given fields:
func (_m *Client) GetMetadata() (*v1.ScannerAdapterMetadata, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 *v1.ScannerAdapterMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.ScannerAdapterMetadata, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.ScannerAdapterMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ScannerAdapterMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanReport provides a mock function with given fields: scanRequestID, reportMIMEType, urlParameter
func (_m *Client) GetScanReport(scanRequestID string, reportMIMEType string, urlParameter string) (string, error) {
	ret := _m.Called(scanRequestID, reportMIMEType, urlParameter)

	if len(ret) == 0 {
		panic("no return value specified for GetScanReport")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(scanRequestID, reportMIMEType, urlParameter)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(scanRequestID, reportMIMEType, urlParameter)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(scanRequestID, reportMIMEType, urlParameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitScan provides a mock function with given fields: req
func (_m *Client) SubmitScan(req *v1.ScanRequest) (*v1.ScanResponse, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for SubmitScan")
	}

	var r0 *v1.ScanResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*v1.ScanRequest) (*v1.ScanResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*v1.ScanRequest) *v1.ScanResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ScanResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*v1.ScanRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
