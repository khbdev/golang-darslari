package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler_GetMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil) // Fake GET so‘rov
	w := httptest.NewRecorder()                          // Javob yozish uchun recorder

	HelloHandler(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Kutilgan status 200, lekin %d chiqdi", resp.StatusCode)
	}

	if !strings.Contains(body, "Hello, World!") {
		t.Errorf("Body noto‘g‘ri: %s", body)
	}
}

func TestHelloHandler_NotGetMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil) // Fake POST so‘rov
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Kutilgan status 405, lekin %d chiqdi", resp.StatusCode)
	}

	if !strings.Contains(body, "Faqat GET method ishlaydi!") {
		t.Errorf("Body noto‘g‘ri: %s", body)
	}
}
