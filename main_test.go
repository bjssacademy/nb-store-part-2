package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeartbeatActive(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if http.StatusOK != w.Result().StatusCode {
		t.Fail()
		t.Log("Expected status code 200 but got", w.Result().StatusCode)
	}

	if w.Body.String() != `{"message":"I'm alive!"}` {
		t.Fail()
		t.Log("Unexpected message from heartbeat:", w.Body.String())
	}

}
