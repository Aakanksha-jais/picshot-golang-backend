// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package stores is a generated GoMock package.
package stores

import (
	reflect "reflect"

	models "github.com/Aakanksha-jais/picshot-golang-backend/models"
	app "github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
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
func (m *MockAccount) Create(c *app.Context, model *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, model)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountMockRecorder) Create(c, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccount)(nil).Create), c, model)
}

// Delete mocks base method.
func (m *MockAccount) Delete(c *app.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountMockRecorder) Delete(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccount)(nil).Delete), c, id)
}

// Get mocks base method.
func (m *MockAccount) Get(c *app.Context, filter *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c, filter)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountMockRecorder) Get(c, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccount)(nil).Get), c, filter)
}

// GetAll mocks base method.
func (m *MockAccount) GetAll(c *app.Context, filter *models.Account) ([]*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", c, filter)
	ret0, _ := ret[0].([]*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAccountMockRecorder) GetAll(c, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAccount)(nil).GetAll), c, filter)
}

// Update mocks base method.
func (m *MockAccount) Update(c *app.Context, model *models.Account) (*models.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, model)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAccountMockRecorder) Update(c, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccount)(nil).Update), c, model)
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
func (m *MockBlog) Create(c *app.Context, model *models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, model)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBlogMockRecorder) Create(c, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlog)(nil).Create), c, model)
}

// Delete mocks base method.
func (m *MockBlog) Delete(c *app.Context, blogID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, blogID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlogMockRecorder) Delete(c, blogID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlog)(nil).Delete), c, blogID)
}

// Get mocks base method.
func (m *MockBlog) Get(c *app.Context, filter *models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c, filter)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlogMockRecorder) Get(c, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlog)(nil).Get), c, filter)
}

// GetAll mocks base method.
func (m *MockBlog) GetAll(c *app.Context, filter models.Blog) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", c, filter)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBlogMockRecorder) GetAll(c, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBlog)(nil).GetAll), c, filter)
}

// GetByIDs mocks base method.
func (m *MockBlog) GetByIDs(c *app.Context, idList []string) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", c, idList)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockBlogMockRecorder) GetByIDs(c, idList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockBlog)(nil).GetByIDs), c, idList)
}

// Update mocks base method.
func (m *MockBlog) Update(c *app.Context, model *models.Blog) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, model)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlogMockRecorder) Update(c, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlog)(nil).Update), c, model)
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
func (m *MockTag) AddBlogID(c *app.Context, blogID string, tags []string) ([]*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBlogID", c, blogID, tags)
	ret0, _ := ret[0].([]*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBlogID indicates an expected call of AddBlogID.
func (mr *MockTagMockRecorder) AddBlogID(c, blogID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBlogID", reflect.TypeOf((*MockTag)(nil).AddBlogID), c, blogID, tags)
}

// GetByName mocks base method.
func (m *MockTag) GetByName(c *app.Context, name string) (*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", c, name)
	ret0, _ := ret[0].(*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockTagMockRecorder) GetByName(c, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockTag)(nil).GetByName), c, name)
}

// RemoveBlogID mocks base method.
func (m *MockTag) RemoveBlogID(c *app.Context, blogID string, tags []string) ([]*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBlogID", c, blogID, tags)
	ret0, _ := ret[0].([]*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBlogID indicates an expected call of RemoveBlogID.
func (mr *MockTagMockRecorder) RemoveBlogID(c, blogID, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBlogID", reflect.TypeOf((*MockTag)(nil).RemoveBlogID), c, blogID, tags)
}
