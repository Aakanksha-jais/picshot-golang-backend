package account

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/handlers"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initializeTest(t *testing.T) (handlers.Account, *services.MockAccount) {
	os.Setenv("ENV", "test")
	mockCtrl := gomock.NewController(t)
	mockService := services.NewMockAccount(mockCtrl)

	config, _ := configs.NewConfigLoader("../../configs")
	s := New(mockService, log.NewLogger(), config)

	return s, mockService
}

func TestHandler_Account_Signup(t *testing.T) {
	s, mockService := initializeTest(t)

	user := models.User{
		UserName: "asmita_prajapati",
		FName:    "Asmita",
		LName:    "Prajapati",
		Email:    sql.NullString{String: "asmita.prajapati@gmail.com", Valid: true},
		PhoneNo:  sql.NullString{String: "7500823463", Valid: true},
		Password: "hello123",
	}

	createdAt := time.Now()
	mockService.EXPECT().Create(gomock.Any(), &user).
		Return(&models.Account{User: user, PwdUpdate: &sql.NullTime{Valid: false}, CreatedAt: createdAt, DelRequest: &sql.NullTime{Valid: false}, Status: "ACTIVE"}, nil)

	tests := []struct {
		description string
		body        []byte
		status      int
		resp        []byte
	}{
		{
			description: "valid signup details",
			body:        []byte(`{"user_name":"asmita_prajapati","f_name":"Asmita","l_name":"Prajapati","email":{"string":"asmita.prajapati@gmail.com","valid":true},"phone_no":{"string":"7500823463","valid":true},"password":"hello123"}`),
			status:      http.StatusCreated,
			resp:        []byte(`{"status":"success"}`),
		},
	}

	for _, tc := range tests {
		r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(tc.body))
		w := httptest.NewRecorder()
		s.Signup(w, r)

		resp, _ := ioutil.ReadAll(w.Body)

		assert.Equal(t, tc.resp, resp)

		assert.Equal(t, tc.status, w.Result().StatusCode)
	}
}
