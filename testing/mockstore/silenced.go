package mockstore

import (
	"context"

	"github.com/sensu/sensu-go/types"
)

// DeleteSilencedEntryByID ...
func (s *MockStore) DeleteSilencedEntryByID(ctx context.Context, silencedID string) error {
	args := s.Called(ctx, silencedID)
	return args.Error(0)
}

// DeleteSilencedEntriesBySubscription ...
func (s *MockStore) DeleteSilencedEntriesBySubscription(ctx context.Context, subscription string) error {
	args := s.Called(ctx, subscription)
	return args.Error(0)
}

// DeleteSilencedEntriesByCheckName ...
func (s *MockStore) DeleteSilencedEntriesByCheckName(ctx context.Context, checkName string) error {
	args := s.Called(ctx, checkName)
	return args.Error(0)
}

// GetSilencedEntries ...
func (s *MockStore) GetSilencedEntries(ctx context.Context) ([]*types.Silenced, error) {
	args := s.Called(ctx)
	return args.Get(0).([]*types.Silenced), args.Error(1)
}

// GetSilencedEntryByID ...
func (s *MockStore) GetSilencedEntryByID(ctx context.Context, silencedID string) (*types.Silenced, error) {
	args := s.Called(ctx, silencedID)
	return args.Get(0).(*types.Silenced), args.Error(1)
}

// GetSilencedEntriesBySubscription ...
func (s *MockStore) GetSilencedEntriesBySubscription(ctx context.Context, subscription string) ([]*types.Silenced, error) {
	args := s.Called(ctx)
	return args.Get(0).([]*types.Silenced), args.Error(1)
}

// GetSilencedEntriesByCheckName ...
func (s *MockStore) GetSilencedEntriesByCheckName(ctx context.Context, checkName string) ([]*types.Silenced, error) {
	args := s.Called(ctx)
	return args.Get(0).([]*types.Silenced), args.Error(1)
}

// UpdateSilencedEntry ...
func (s *MockStore) UpdateSilencedEntry(ctx context.Context, silenced *types.Silenced) error {
	args := s.Called(ctx, silenced)
	return args.Error(0)
}
