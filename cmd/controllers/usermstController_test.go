package controllers

import (
	"GOTEST/cmd/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAllusermst(t *testing.T) {

	idealdata := []models.AllUsemst{
		{Usercd: "a0006802", User_f_name: "Noriufmi", User_l_name: "Sato", Belonged_company_cd: "0001", Belonged_dept_cd: "32P2", Authcd: "01", Activeflg: true, Createdby: "a0006802", Created_date: "2022-04-01", Updatedby: "a0006802", Updated_date: "2022-04-01", UpdateCounter: 1},
		{Usercd: "a0006801", User_f_name: "Noriufmi", User_l_name: "Sato", Belonged_company_cd: "0001", Belonged_dept_cd: "32P2", Authcd: "01", Activeflg: true, Createdby: "a0006802", Created_date: "2022-04-01", Updatedby: "a0006802", Updated_date: "2022-04-01", UpdateCounter: 1},
	}
	byteusermst, _ := json.Marshal(idealdata)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/mastermaintenance/usermst/allusermst",
		bytes.NewBuffer(byteusermst),
	)

	// Action
	Allusermst(c)

	// Assert ---
	var product []models.AllUsemst
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	for getdatacount := 0; getdatacount < len(product)-1; getdatacount++ {
		assert.EqualValues(t, idealdata[getdatacount].Usercd, product[getdatacount].Usercd)
		assert.EqualValues(t, idealdata[getdatacount].User_f_name, product[getdatacount].User_f_name)
		assert.EqualValues(t, idealdata[getdatacount].User_l_name, product[getdatacount].User_l_name)
		assert.EqualValues(t, idealdata[getdatacount].Belonged_company_cd, product[getdatacount].Belonged_company_cd)
		assert.EqualValues(t, idealdata[getdatacount].Belonged_dept_cd, product[getdatacount].Belonged_dept_cd)
		assert.EqualValues(t, idealdata[getdatacount].Authcd, product[getdatacount].Authcd)
		assert.EqualValues(t, idealdata[getdatacount].Activeflg, product[getdatacount].Activeflg)
		assert.EqualValues(t, idealdata[getdatacount].Createdby, product[getdatacount].Createdby)
		assert.EqualValues(t, idealdata[getdatacount].Created_date, product[getdatacount].Created_date)
		assert.EqualValues(t, idealdata[getdatacount].Updatedby, product[getdatacount].Updatedby)
		assert.EqualValues(t, idealdata[getdatacount].Updated_date, product[getdatacount].Updated_date)
		assert.EqualValues(t, idealdata[getdatacount].UpdateCounter, product[getdatacount].UpdateCounter)
	}
}
