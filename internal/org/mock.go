package org

import (
	"context"

	"github.com/artifacthub/hub/internal/hub"
	"github.com/stretchr/testify/mock"
)

// ManagerMock is a mock implementation of the OrganizationManager interface.
type ManagerMock struct {
	mock.Mock
}

// Add implements the OrganizationManager interface.
func (m *ManagerMock) Add(ctx context.Context, org *hub.Organization) error {
	args := m.Called(ctx, org)
	return args.Error(0)
}

// AddMember implements the OrganizationManager interface.
func (m *ManagerMock) AddMember(ctx context.Context, orgName, userAlias, baseURL string) error {
	args := m.Called(ctx, orgName, userAlias, baseURL)
	return args.Error(0)
}

// ConfirmMembership implements the OrganizationManager interface.
func (m *ManagerMock) ConfirmMembership(ctx context.Context, orgName string) error {
	args := m.Called(ctx, orgName)
	return args.Error(0)
}

// DeleteMember implements the OrganizationManager interface.
func (m *ManagerMock) DeleteMember(ctx context.Context, orgName, userAlias string) error {
	args := m.Called(ctx, orgName, userAlias)
	return args.Error(0)
}

// GetJSON implements the OrganizationManager interface.
func (m *ManagerMock) GetJSON(ctx context.Context, orgName string) ([]byte, error) {
	args := m.Called(ctx, orgName)
	data, _ := args.Get(0).([]byte)
	return data, args.Error(1)
}

// GetByUserJSON implements the OrganizationManager interface.
func (m *ManagerMock) GetByUserJSON(ctx context.Context) ([]byte, error) {
	args := m.Called(ctx)
	data, _ := args.Get(0).([]byte)
	return data, args.Error(1)
}

// GetMembersJSON implements the OrganizationManager interface.
func (m *ManagerMock) GetMembersJSON(ctx context.Context, orgName string) ([]byte, error) {
	args := m.Called(ctx, orgName)
	data, _ := args.Get(0).([]byte)
	return data, args.Error(1)
}

// Update implements the OrganizationManager interface.
func (m *ManagerMock) Update(ctx context.Context, org *hub.Organization) error {
	args := m.Called(ctx, org)
	return args.Error(0)
}
