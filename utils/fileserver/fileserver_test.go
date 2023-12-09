package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServer(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Test!"))
	}))
	defer ts.Close()

	// Test the StartServer function
	port := 8080
	go func() {	
		err := StartServer(port, "./Cuphead.zip")
		if err != nil {
			t.Errorf("Error starting the server: %v", err)
		}
	}()
}
