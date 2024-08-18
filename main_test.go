package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newUploadRequest(endpoint, matrix string) *http.Request {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", "matrix.csv")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, strings.NewReader(matrix))
	if err != nil {
		panic(err)
	}

	// Close the multipart writer to set the terminating boundary
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// Create a new POST request with the multipart body
	req := httptest.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestHandleEcho(t *testing.T) {
	matrix := `1,2,3
4,5,6
7,8,9`

	req := newUploadRequest("/echo", matrix)
	resp := httptest.NewRecorder()

	handleEcho(resp, req)

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if resp.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, resp.Body.String())
	}
}


func TestHandleInvert(t *testing.T) {
	matrix := `1,2,3
4,5,6
7,8,9`

	req := newUploadRequest("/invert", matrix)
	resp := httptest.NewRecorder()

	handleInvert(resp, req)

	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if resp.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, resp.Body.String())
	}
}

func TestHandleFlatten(t *testing.T) {
	matrix := `1,2,3
4,5,6
7,8,9`

	req := newUploadRequest("/flatten", matrix)
	resp := httptest.NewRecorder()

	handleFlatten(resp, req)

	expected := "1,2,3,4,5,6,7,8,9"
	if resp.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, resp.Body.String())
	}
}

func TestHandleSum(t *testing.T) {
	matrix := `1,2,3
4,5,6
7,8,9`

	req := newUploadRequest("/sum", matrix)
	resp := httptest.NewRecorder()

	handleSum(resp, req)

	expected := "45"
	if resp.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, resp.Body.String())
	}
}

func TestHandleMultiply(t *testing.T) {
	matrix := `1,2,3
4,5,6
7,8,9`

	req := newUploadRequest("/multiply", matrix)
	resp := httptest.NewRecorder()

	handleMultiply(resp, req)

	expected := "362880"
	if resp.Body.String() != expected {
		t.Errorf("expected %q but got %q", expected, resp.Body.String())
	}
}

