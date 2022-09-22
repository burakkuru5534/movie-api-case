package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/api"
	"github.com/burakkuru5534/src/helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/Movies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MovieList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Movie with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Movie with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `[{"id":4,"name":"The Lord Of The Rings","description":"desc","typ":"fantasy"},{"id":5,"name":"Harry Potter","description":"desc","typ":"fantasy"}]
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestCreate(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	var jsonStr = []byte(`{"name":"Warrior","description":"desc","typ":"action"}`)

	req, err := http.NewRequest("POST", "/api/movie", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MovieCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Movie with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Movie with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		var id int64

		err = db.Get(&id, "SELECT id from usr order by id desc limit 1")
		if err != nil {
			errors.New("get id error.")
		}

		expected := fmt.Sprintf(`{"id":%d,"name":"Warrior","description":"desc","typ":"action"}
`, id)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

func TestGet(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/movie", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MovieGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Movie with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Movie with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `{"id":4,"name":"The Lord Of The Rings","description":"desc","typ":"fantasy"}
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestDelete(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("DELETE", "/api/movie", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MovieDelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Movie with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Movie with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	} else {
		// Check the response body is what we expect.
		expected := `"Movie deleted"
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

func TestUpdate(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	var jsonStr = []byte(`{"name":"TestMovieUpdatedName","description":"desc","typ":"fantasy"}`)

	req, err := http.NewRequest("PATCH", "/api/movie", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MovieUpdate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Movie with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Movie with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := fmt.Sprintf(`{"name":"TestMovieUpdatedName","description":"desc","typ":"fantasy"}
`)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}
