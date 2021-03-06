// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: replicationTaskProcessor.go

// Package history is a generated GoMock package.
package history

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReplicationTaskProcessor is a mock of ReplicationTaskProcessor interface
type MockReplicationTaskProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockReplicationTaskProcessorMockRecorder
}

// MockReplicationTaskProcessorMockRecorder is the mock recorder for MockReplicationTaskProcessor
type MockReplicationTaskProcessorMockRecorder struct {
	mock *MockReplicationTaskProcessor
}

// NewMockReplicationTaskProcessor creates a new mock instance
func NewMockReplicationTaskProcessor(ctrl *gomock.Controller) *MockReplicationTaskProcessor {
	mock := &MockReplicationTaskProcessor{ctrl: ctrl}
	mock.recorder = &MockReplicationTaskProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReplicationTaskProcessor) EXPECT() *MockReplicationTaskProcessorMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockReplicationTaskProcessor) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockReplicationTaskProcessorMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockReplicationTaskProcessor)(nil).Start))
}

// Stop mocks base method
func (m *MockReplicationTaskProcessor) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockReplicationTaskProcessorMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockReplicationTaskProcessor)(nil).Stop))
}
