// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package stores is a generated GoMock package.
package stores

import (
	multipart "mime/multipart"
	reflect "reflect"

	models "github.com/Aakanksha-jais/picshot-golang-backend/models"
	app "github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	constants "github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"
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
func (m *MockBlog) GetAll(c *app.Context, filter *models.Blog, page *models.Page) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", c, filter, page)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBlogMockRecorder) GetAll(c, filter, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBlog)(nil).GetAll), c, filter, page)
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

// Create mocks base method.
func (m *MockTag) Create(c *app.Context, tag *models.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTagMockRecorder) Create(c, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTag)(nil).Create), c, tag)
}

// Delete mocks base method.
func (m *MockTag) Delete(c *app.Context, tag string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, tag)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTagMockRecorder) Delete(c, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTag)(nil).Delete), c, tag)
}

// Get mocks base method.
func (m *MockTag) Get(c *app.Context, name string) (*models.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c, name)
	ret0, _ := ret[0].(*models.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTagMockRecorder) Get(c, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTag)(nil).Get), c, name)
}

// Update mocks base method.
func (m *MockTag) Update(c *app.Context, blogID, tag string, operation constants.Operation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, blogID, tag, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTagMockRecorder) Update(c, blogID, tag, operation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTag)(nil).Update), c, blogID, tag, operation)
}

// MockImage is a mock of Image interface.
type MockImage struct {
	ctrl     *gomock.Controller
	recorder *MockImageMockRecorder
}

// MockImageMockRecorder is the mock recorder for MockImage.
type MockImageMockRecorder struct {
	mock *MockImage
}

// NewMockImage creates a new mock instance.
func NewMockImage(ctrl *gomock.Controller) *MockImage {
	mock := &MockImage{ctrl: ctrl}
	mock.recorder = &MockImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImage) EXPECT() *MockImageMockRecorder {
	return m.recorder
}

// DeleteBulk mocks base method.
func (m *MockImage) DeleteBulk(ctx *app.Context, names []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBulk", ctx, names)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBulk indicates an expected call of DeleteBulk.
func (mr *MockImageMockRecorder) DeleteBulk(ctx, names interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBulk", reflect.TypeOf((*MockImage)(nil).DeleteBulk), ctx, names)
}

// Upload mocks base method.
func (m *MockImage) Upload(c *app.Context, fileHeader *multipart.FileHeader, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", c, fileHeader, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockImageMockRecorder) Upload(c, fileHeader, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockImage)(nil).Upload), c, fileHeader, name)
}
