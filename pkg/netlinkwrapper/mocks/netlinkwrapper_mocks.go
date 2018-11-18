// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/netlinkwrapper (interfaces: NetLink)

// Package mock_netlinkwrapper is a generated GoMock package.
package mock_netlinkwrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
)

// MockNetLink is a mock of NetLink interface
type MockNetLink struct {
	ctrl     *gomock.Controller
	recorder *MockNetLinkMockRecorder
}

// MockNetLinkMockRecorder is the mock recorder for MockNetLink
type MockNetLinkMockRecorder struct {
	mock *MockNetLink
}

// NewMockNetLink creates a new mock instance
func NewMockNetLink(ctrl *gomock.Controller) *MockNetLink {
	mock := &MockNetLink{ctrl: ctrl}
	mock.recorder = &MockNetLinkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetLink) EXPECT() *MockNetLinkMockRecorder {
	return m.recorder
}

// AddrAdd mocks base method
func (m *MockNetLink) AddrAdd(arg0 netlink.Link, arg1 *netlink.Addr) error {
	ret := m.ctrl.Call(m, "AddrAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddrAdd indicates an expected call of AddrAdd
func (mr *MockNetLinkMockRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrAdd", reflect.TypeOf((*MockNetLink)(nil).AddrAdd), arg0, arg1)
}

// AddrList mocks base method
func (m *MockNetLink) AddrList(arg0 netlink.Link, arg1 int) ([]netlink.Addr, error) {
	ret := m.ctrl.Call(m, "AddrList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddrList indicates an expected call of AddrList
func (mr *MockNetLinkMockRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddrList", reflect.TypeOf((*MockNetLink)(nil).AddrList), arg0, arg1)
}

// LinkAdd mocks base method
func (m *MockNetLink) LinkAdd(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkAdd indicates an expected call of LinkAdd
func (mr *MockNetLinkMockRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkAdd", reflect.TypeOf((*MockNetLink)(nil).LinkAdd), arg0)
}

// LinkByName mocks base method
func (m *MockNetLink) LinkByName(arg0 string) (netlink.Link, error) {
	ret := m.ctrl.Call(m, "LinkByName", arg0)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkByName indicates an expected call of LinkByName
func (mr *MockNetLinkMockRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkByName", reflect.TypeOf((*MockNetLink)(nil).LinkByName), arg0)
}

// LinkDel mocks base method
func (m *MockNetLink) LinkDel(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDel indicates an expected call of LinkDel
func (mr *MockNetLinkMockRecorder) LinkDel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDel", reflect.TypeOf((*MockNetLink)(nil).LinkDel), arg0)
}

// LinkList mocks base method
func (m *MockNetLink) LinkList() ([]netlink.Link, error) {
	ret := m.ctrl.Call(m, "LinkList")
	ret0, _ := ret[0].([]netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LinkList indicates an expected call of LinkList
func (mr *MockNetLinkMockRecorder) LinkList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkList", reflect.TypeOf((*MockNetLink)(nil).LinkList))
}

// LinkSetDown mocks base method
func (m *MockNetLink) LinkSetDown(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkSetDown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetDown indicates an expected call of LinkSetDown
func (mr *MockNetLinkMockRecorder) LinkSetDown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetDown", reflect.TypeOf((*MockNetLink)(nil).LinkSetDown), arg0)
}

// LinkSetMTU mocks base method
func (m *MockNetLink) LinkSetMTU(arg0 netlink.Link, arg1 int) error {
	ret := m.ctrl.Call(m, "LinkSetMTU", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetMTU indicates an expected call of LinkSetMTU
func (mr *MockNetLinkMockRecorder) LinkSetMTU(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetMTU", reflect.TypeOf((*MockNetLink)(nil).LinkSetMTU), arg0, arg1)
}

// LinkSetNsFd mocks base method
func (m *MockNetLink) LinkSetNsFd(arg0 netlink.Link, arg1 int) error {
	ret := m.ctrl.Call(m, "LinkSetNsFd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetNsFd indicates an expected call of LinkSetNsFd
func (mr *MockNetLinkMockRecorder) LinkSetNsFd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetNsFd", reflect.TypeOf((*MockNetLink)(nil).LinkSetNsFd), arg0, arg1)
}

// LinkSetUp mocks base method
func (m *MockNetLink) LinkSetUp(arg0 netlink.Link) error {
	ret := m.ctrl.Call(m, "LinkSetUp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkSetUp indicates an expected call of LinkSetUp
func (mr *MockNetLinkMockRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkSetUp", reflect.TypeOf((*MockNetLink)(nil).LinkSetUp), arg0)
}

// NeighAdd mocks base method
func (m *MockNetLink) NeighAdd(arg0 *netlink.Neigh) error {
	ret := m.ctrl.Call(m, "NeighAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NeighAdd indicates an expected call of NeighAdd
func (mr *MockNetLinkMockRecorder) NeighAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeighAdd", reflect.TypeOf((*MockNetLink)(nil).NeighAdd), arg0)
}

// NewRule mocks base method
func (m *MockNetLink) NewRule() *netlink.Rule {
	ret := m.ctrl.Call(m, "NewRule")
	ret0, _ := ret[0].(*netlink.Rule)
	return ret0
}

// NewRule indicates an expected call of NewRule
func (mr *MockNetLinkMockRecorder) NewRule() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRule", reflect.TypeOf((*MockNetLink)(nil).NewRule))
}

// ParseAddr mocks base method
func (m *MockNetLink) ParseAddr(arg0 string) (*netlink.Addr, error) {
	ret := m.ctrl.Call(m, "ParseAddr", arg0)
	ret0, _ := ret[0].(*netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseAddr indicates an expected call of ParseAddr
func (mr *MockNetLinkMockRecorder) ParseAddr(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseAddr", reflect.TypeOf((*MockNetLink)(nil).ParseAddr), arg0)
}

// RouteAdd mocks base method
func (m *MockNetLink) RouteAdd(arg0 *netlink.Route) error {
	ret := m.ctrl.Call(m, "RouteAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteAdd indicates an expected call of RouteAdd
func (mr *MockNetLinkMockRecorder) RouteAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteAdd", reflect.TypeOf((*MockNetLink)(nil).RouteAdd), arg0)
}

// RouteDel mocks base method
func (m *MockNetLink) RouteDel(arg0 *netlink.Route) error {
	ret := m.ctrl.Call(m, "RouteDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RouteDel indicates an expected call of RouteDel
func (mr *MockNetLinkMockRecorder) RouteDel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDel", reflect.TypeOf((*MockNetLink)(nil).RouteDel), arg0)
}

// RouteList mocks base method
func (m *MockNetLink) RouteList(arg0 netlink.Link, arg1 int) ([]netlink.Route, error) {
	ret := m.ctrl.Call(m, "RouteList", arg0, arg1)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteList indicates an expected call of RouteList
func (mr *MockNetLinkMockRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteList", reflect.TypeOf((*MockNetLink)(nil).RouteList), arg0, arg1)
}

// RuleAdd mocks base method
func (m *MockNetLink) RuleAdd(arg0 *netlink.Rule) error {
	ret := m.ctrl.Call(m, "RuleAdd", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleAdd indicates an expected call of RuleAdd
func (mr *MockNetLinkMockRecorder) RuleAdd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleAdd", reflect.TypeOf((*MockNetLink)(nil).RuleAdd), arg0)
}

// RuleDel mocks base method
func (m *MockNetLink) RuleDel(arg0 *netlink.Rule) error {
	ret := m.ctrl.Call(m, "RuleDel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RuleDel indicates an expected call of RuleDel
func (mr *MockNetLinkMockRecorder) RuleDel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleDel", reflect.TypeOf((*MockNetLink)(nil).RuleDel), arg0)
}

// RuleList mocks base method
func (m *MockNetLink) RuleList(arg0 int) ([]netlink.Rule, error) {
	ret := m.ctrl.Call(m, "RuleList", arg0)
	ret0, _ := ret[0].([]netlink.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RuleList indicates an expected call of RuleList
func (mr *MockNetLinkMockRecorder) RuleList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuleList", reflect.TypeOf((*MockNetLink)(nil).RuleList), arg0)
}
