// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/aws/aws-sdk-go-v2/service/eks/types"
	mock "github.com/stretchr/testify/mock"

	v1alpha5 "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

// PodIdentityIAMUpdater is an autogenerated mock type for the PodIdentityIAMUpdater type
type PodIdentityIAMUpdater struct {
	mock.Mock
}

// DeleteRole provides a mock function with given fields: ctx, addonName, serviceAccountName
func (_m *PodIdentityIAMUpdater) DeleteRole(ctx context.Context, addonName string, serviceAccountName string) (bool, error) {
	ret := _m.Called(ctx, addonName, serviceAccountName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRole")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, addonName, serviceAccountName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, addonName, serviceAccountName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, addonName, serviceAccountName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: ctx, podIdentityAssociations, addonName
func (_m *PodIdentityIAMUpdater) UpdateRole(ctx context.Context, podIdentityAssociations []v1alpha5.PodIdentityAssociation, addonName string) ([]types.AddonPodIdentityAssociations, error) {
	ret := _m.Called(ctx, podIdentityAssociations, addonName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 []types.AddonPodIdentityAssociations
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []v1alpha5.PodIdentityAssociation, string) ([]types.AddonPodIdentityAssociations, error)); ok {
		return rf(ctx, podIdentityAssociations, addonName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []v1alpha5.PodIdentityAssociation, string) []types.AddonPodIdentityAssociations); ok {
		r0 = rf(ctx, podIdentityAssociations, addonName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.AddonPodIdentityAssociations)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []v1alpha5.PodIdentityAssociation, string) error); ok {
		r1 = rf(ctx, podIdentityAssociations, addonName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPodIdentityIAMUpdater creates a new instance of PodIdentityIAMUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPodIdentityIAMUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *PodIdentityIAMUpdater {
	mock := &PodIdentityIAMUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
