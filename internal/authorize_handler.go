// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: AuthorizeEndpointHandler)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	context "golang.org/x/net/context"
	http "net/http"
)

// Mock of AuthorizeEndpointHandler interface
type MockAuthorizeEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthorizeEndpointHandlerRecorder
}

// Recorder for MockAuthorizeEndpointHandler (not exported)
type _MockAuthorizeEndpointHandlerRecorder struct {
	mock *MockAuthorizeEndpointHandler
}

func NewMockAuthorizeEndpointHandler(ctrl *gomock.Controller) *MockAuthorizeEndpointHandler {
	mock := &MockAuthorizeEndpointHandler{ctrl: ctrl}
	mock.recorder = &_MockAuthorizeEndpointHandlerRecorder{mock}
	return mock
}

func (_m *MockAuthorizeEndpointHandler) EXPECT() *_MockAuthorizeEndpointHandlerRecorder {
	return _m.recorder
}

func (_m *MockAuthorizeEndpointHandler) HandleAuthorizeEndpointRequest(_param0 context.Context, _param1 *http.Request, _param2 fosite.AuthorizeRequester, _param3 fosite.AuthorizeResponder, _param4 interface{}) error {
	ret := _m.ctrl.Call(_m, "HandleAuthorizeEndpointRequest", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeEndpointHandlerRecorder) HandleAuthorizeEndpointRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleAuthorizeEndpointRequest", arg0, arg1, arg2, arg3, arg4)
}
