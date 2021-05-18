// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package stores is a generated GoMock package.
package stores

import (
	context "context"
	reflect "reflect"

	models "github.com/Aakanksha-jais/picshot-golang-backend/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccount) Create(ctx context.Context, model *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, model)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountMockRecorder) Create(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccount)(nil).Create), ctx, model)
}

// Delete mocks base method.
func (m *MockAccount) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccount)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockAccount) Get(ctx context.Context, filter *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, filter)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAccount)(nil).Get), ctx, filter)
}

// GetAll mocks base method.
func (m *MockAccount) GetAll(ctx context.Context, filter *models.Account) ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, filter)
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAccountMockRecorder) GetAll(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAccount)(nil).GetAll), ctx, filter)
}

// Update mocks base method.
func (m *MockAccount) Update(ctx context.Context, model *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, model)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountMockRecorder) Update(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccount)(nil).Update), ctx, model)
}

// MockBlog is a mock of Blog interface.
type MockBlog struct {
	ctrl     *gomock.Controller
	recorder *MockBlogMockRecorder
}

// MockBlogMockRecorder is the mock recorder for MockBlog.
type MockBlogMockRecorder struct {
	mock *MockBlog
}

// NewMockBlog creates a new mock instance.
func NewMockBlog(ctrl *gomock.Controller) *MockBlog {
	mock := &MockBlog{ctrl: ctrl}
	mock.recorder = &MockBlogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlog) EXPECT() *MockBlogMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlog) Create(ctx context.Context, model models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, model)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlogMockRecorder) Create(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlog)(nil).Create), ctx, model)
}

// Delete mocks base method.
func (m *MockBlog) Delete(ctx context.Context, blogID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, blogID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlogMockRecorder) Delete(ctx, blogID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlog)(nil).Delete), ctx, blogID)
}

// Get mocks base method.
func (m *MockBlog) Get(ctx context.Context, filter models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, filter)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlogMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockBlog)(nil).Get), ctx, filter)
}

// GetAll mocks base method.
func (m *MockBlog) GetAll(ctx context.Context, filter models.Blog) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, filter)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBlogMockRecorder) GetAll(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBlog)(nil).GetAll), ctx, filter)
}

// GetByIDs mocks base method.
func (m *MockBlog) GetByIDs(ctx context.Context, idList []string) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, idList)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockBlogMockRecorder) GetByIDs(ctx, idList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockBlog)(nil).GetByIDs), ctx, idList)
}

// Update mocks base method.
func (m *MockBlog) Update(ctx context.Context, model models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, model)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlogMockRecorder) Update(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlog)(nil).Update), ctx, model)
}

// MockTag is a mock of Tag interface.
type MockTag struct {
	ctrl     *gomock.Controller
	recorder *MockTagMockRecorder
}

// MockTagMockRecorder is the mock recorder for MockTag.
type MockTagMockRecorder struct {
	mock *MockTag
}

// NewMockTag creates a new mock instance.
func NewMockTag(ctrl *gomock.Controller) *MockTag {
	mock := &MockTag{ctrl: ctrl}
	mock.recorder = &MockTagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTag) EXPECT() *MockTagMockRecorder {
	return m.recorder
}

// AddBlogID mocks base method.
func (m *MockTag) AddBlogID(ctx context.Context, blogID string, tags []string) ([]*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlogID", ctx, blogID, tags)
	ret0, _ := ret[0].([]*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBlogID indicates an expected call of AddBlogID.
func (mr *MockTagMockRecorder) AddBlogID(ctx, blogID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlogID", reflect.TypeOf((*MockTag)(nil).AddBlogID), ctx, blogID, tags)
}

// GetByName mocks base method.
func (m *MockTag) GetByName(ctx context.Context, name string) (*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockTagMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockTag)(nil).GetByName), ctx, name)
}

// RemoveBlogID mocks base method.
func (m *MockTag) RemoveBlogID(ctx context.Context, blogID string, tags []string) ([]*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBlogID", ctx, blogID, tags)
	ret0, _ := ret[0].([]*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBlogID indicates an expected call of RemoveBlogID.
func (mr *MockTagMockRecorder) RemoveBlogID(ctx, blogID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBlogID", reflect.TypeOf((*MockTag)(nil).RemoveBlogID), ctx, blogID, tags)
}
