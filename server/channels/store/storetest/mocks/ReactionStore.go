// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ReactionStore is an autogenerated mock type for the ReactionStore type
type ReactionStore struct {
	mock.Mock
}

// BulkGetForPosts provides a mock function with given fields: postIds
func (_m *ReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, error) {
	ret := _m.Called(postIds)

	var r0 []*model.Reaction
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*model.Reaction, error)); ok {
		return rf(postIds)
	}
	if rf, ok := ret.Get(0).(func([]string) []*model.Reaction); ok {
		r0 = rf(postIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(postIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: reaction
func (_m *ReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, error) {
	ret := _m.Called(reaction)

	var r0 *model.Reaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Reaction) (*model.Reaction, error)); ok {
		return rf(reaction)
	}
	if rf, ok := ret.Get(0).(func(*model.Reaction) *model.Reaction); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Reaction) error); ok {
		r1 = rf(reaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllWithEmojiName provides a mock function with given fields: emojiName
func (_m *ReactionStore) DeleteAllWithEmojiName(emojiName string) error {
	ret := _m.Called(emojiName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(emojiName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedRows provides a mock function with given fields: limit
func (_m *ReactionStore) DeleteOrphanedRows(limit int) (int64, error) {
	ret := _m.Called(limit)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForPost provides a mock function with given fields: postID, allowFromCache
func (_m *ReactionStore) GetForPost(postID string, allowFromCache bool) ([]*model.Reaction, error) {
	ret := _m.Called(postID, allowFromCache)

	var r0 []*model.Reaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) ([]*model.Reaction, error)); ok {
		return rf(postID, allowFromCache)
	}
	if rf, ok := ret.Get(0).(func(string, bool) []*model.Reaction); ok {
		r0 = rf(postID, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(postID, allowFromCache)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForPostSince provides a mock function with given fields: postId, since, excludeRemoteId, inclDeleted
func (_m *ReactionStore) GetForPostSince(postId string, since int64, excludeRemoteId string, inclDeleted bool) ([]*model.Reaction, error) {
	ret := _m.Called(postId, since, excludeRemoteId, inclDeleted)

	var r0 []*model.Reaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, string, bool) ([]*model.Reaction, error)); ok {
		return rf(postId, since, excludeRemoteId, inclDeleted)
	}
	if rf, ok := ret.Get(0).(func(string, int64, string, bool) []*model.Reaction); ok {
		r0 = rf(postId, since, excludeRemoteId, inclDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, string, bool) error); ok {
		r1 = rf(postId, since, excludeRemoteId, inclDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopForTeamSince provides a mock function with given fields: teamID, userID, since, offset, limit
func (_m *ReactionStore) GetTopForTeamSince(teamID string, userID string, since int64, offset int, limit int) (*model.TopReactionList, error) {
	ret := _m.Called(teamID, userID, since, offset, limit)

	var r0 *model.TopReactionList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) (*model.TopReactionList, error)); ok {
		return rf(teamID, userID, since, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopReactionList); ok {
		r0 = rf(teamID, userID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopReactionList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(teamID, userID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTopForUserSince provides a mock function with given fields: userID, teamID, since, offset, limit
func (_m *ReactionStore) GetTopForUserSince(userID string, teamID string, since int64, offset int, limit int) (*model.TopReactionList, error) {
	ret := _m.Called(userID, teamID, since, offset, limit)

	var r0 *model.TopReactionList
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) (*model.TopReactionList, error)); ok {
		return rf(userID, teamID, since, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int, int) *model.TopReactionList); ok {
		r0 = rf(userID, teamID, since, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TopReactionList)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int, int) error); ok {
		r1 = rf(userID, teamID, since, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *ReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, error) {
	ret := _m.Called(endTime, limit)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64) (int64, error)); ok {
		return rf(endTime, limit)
	}
	if rf, ok := ret.Get(0).(func(int64, int64) int64); ok {
		r0 = rf(endTime, limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: reaction
func (_m *ReactionStore) Save(reaction *model.Reaction) (*model.Reaction, error) {
	ret := _m.Called(reaction)

	var r0 *model.Reaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Reaction) (*model.Reaction, error)); ok {
		return rf(reaction)
	}
	if rf, ok := ret.Get(0).(func(*model.Reaction) *model.Reaction); ok {
		r0 = rf(reaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Reaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Reaction) error); ok {
		r1 = rf(reaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReactionStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewReactionStore creates a new instance of ReactionStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReactionStore(t mockConstructorTestingTNewReactionStore) *ReactionStore {
	mock := &ReactionStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}