package fileserver

import (
	"net/http"
	"strconv"
)

func StartServer(port int, directory string) error {
	// Parse command line arguments
	http.Handle("/", http.FileServer(http.Dir(directory)))

	addr := ":" + strconv.Itoa(port)
	return http.ListenAndServe(addr, nil)
}