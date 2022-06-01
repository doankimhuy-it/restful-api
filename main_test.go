package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"restful/api"
)

func TestGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/tasks?id=1", nil)
	w := httptest.NewRecorder()
	api.Get(w, req)
	
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
        t.Fatalf("expected a %d, instead got: %d", want, got)
    }
}

func TestCreate(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "localhost:8080/tasks/create?id=10&title=XYZ&status=Done", nil)
	w := httptest.NewRecorder()
	api.Create(w, req)
	
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
        t.Fatalf("expected a %d, instead got: %d", want, got)
    }
}

func TestUpdate(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "localhost:8080/tasks/update?id=10&title=XY&status=Done", nil)
	w := httptest.NewRecorder()
	api.Update(w, req)
	
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
        t.Fatalf("expected a %d, instead got: %d", want, got)
    }
}

func TestDelete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "localhost:8080/tasks/delete?id=2", nil)
	w := httptest.NewRecorder()
	api.Del(w, req)
	
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
        t.Fatalf("expected a %d, instead got: %d", want, got)
    }
}