package account

//
//import (
//	"context"
//	"database/sql"
//	"fmt"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//
//	"github.com/Aakanksha-jais/picshot-golang-backend/models"
//
//	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
//	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
//)
//
//func initializeTest() (*app.Context, stores.Account) {
//	app.InitializeTestAccountsTable(a.SQL.GetDB(), a.Logger, "../../db")
//	return &app.Context{Context: context.TODO(), App: a}, New()
//}
//
//func TestAccount_GetAll(t *testing.T) {
//	ctx, account := initializeTest()
//
//	tests := []struct {
//		description string
//		input       *models.Account
//		output      []*models.Account
//		err         error
//	}{
//		{
//			input:  &models.Account{Status: "ACTIVE"},
//			output: getAllOutput(),
//		},
//	}
//
//	for i := range tests {
//		output, err := account.GetAll(ctx, tests[i].input)
//
//		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
//			for j := range output {
//				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
//				fmt.Println(tests[i].output[j].CreatedAt.String())
//				fmt.Println(output[j].CreatedAt.String())
//			}
//		}
//
//		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
//	}
//}
//
//func getAllOutput() []*models.Account {
//	loc, _ := time.LoadLocation("Asia/Kolkata")
//
//	return []*models.Account{
//		{User: models.User{ID: 1, UserName: "aakanksha_jais", FName: "Aakanksha", LName: "Jaiswal", Email: sql.NullString{String: "jaiswal14aakanksha@gmail.com", Valid: true}, PhoneNo: sql.NullString{String: "7807052049", Valid: true}}, CreatedAt: time.Now().In(loc)., Status: "ACTIVE"},
//		{User: models.User{ID: 2, UserName: "mainak_pandit", FName: "Mainak", LName: "Pandit", Email: sql.NullString{String: "mainakpandit@gmail.com", Valid: true}, PhoneNo: sql.NullString{String: "9149137433", Valid: true}}, CreatedAt: time.Now().In(loc), Status: "ACTIVE"},
//		{User: models.User{ID: 3, UserName: "divij_gupta", FName: "Divij", LName: "Gupta", Email: sql.NullString{String: "divijgupta@gmail.com", Valid: true}, PhoneNo: sql.NullString{String: "9682622125", Valid: true}}, CreatedAt: time.Now().In(loc), Status: "ACTIVE"},
//	}
//}
