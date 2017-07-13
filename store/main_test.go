package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a App

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func ensureTableExists() {
	if _, err := a.DB.Exec("CREATE TABLE IF NOT EXISTS users ( id serial PRIMARY KEY, name varchar(32), email varchar(32), password varchar(8) )"); err != nil {
		log.Fatal(err)
	}
}

func assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(
		thost, tport, tuser, tpassword, tdbname,
	)
	ensureTableExists()
	listServe := m.Run()
	clearTable()
	os.Exit(listServe)
}

func TestUserCreateHandler(t *testing.T) {
	clearTable()
	goodparams := User{Name: "testName", Email: "test@mail.com", Password: "testpass"}

	jsreq, err := json.Marshal(&goodparams)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	req, err := http.NewRequest("POST", "/api/users/create", bytes.NewReader(jsreq))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating request", err)
	}
	rr := executeRequest(req)

	if rr.Code != 201 {
		t.Fatalf("Status code '%d' was not expected", rr.Code)
	}
}

func TestUserCreateHandlerInvalidInput(t *testing.T) {
	clearTable()
	badparams := `{}`

	jsreq, err := json.Marshal(&badparams)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	req, err := http.NewRequest("POST", "/api/users/create", bytes.NewReader(jsreq))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating request", err)
	}
	rr := executeRequest(req)

	if rr.Code != 400 {
		t.Fatalf("Status code '%d' was not expected", rr.Code)
	}
}

func TestUserCreateHandlerInvalidUserInput(t *testing.T) {
	clearTable()
	badparams := User{Name: "", Password: ""}

	jsreq, err := json.Marshal(&badparams)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	req, err := http.NewRequest("POST", "/api/users/create", bytes.NewReader(jsreq))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating request", err)
	}
	rr := executeRequest(req)

	if rr.Code != 400 {
		t.Fatalf("Status code '%d' was not expected", rr.Code)
	}
	assert.Equal(t, "Bad arguments\n", rr.Body.String(), "Response body is not correct")
}

func TestShowUsersHandler(t *testing.T) {
	clearTable()
	var userid int
	a.DB.QueryRow(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`, "testname", "test@mail.com", "testpass").Scan(&userid)
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating request", err)
	}

	rr := executeRequest(req)

	if rr.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", rr.Code)
	}

	data := []User{User{ID: userid, Name: "testname", Email: "test@mail.com", Password: "testpass"}}
	assertJSON(rr.Body.Bytes(), data, t)
}
