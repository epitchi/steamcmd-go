package zipfolder

import (
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	ZipFile("./README.md", "test")
	
	_, err := os.Stat("test" + ".zip")

	if err != nil {
		t.Errorf("Expected zip file to exist, but got an error: %v", err)
	}

	// Delete the generated zip file after the test
	if err := os.Remove("test" + ".zip"); err != nil {
		t.Errorf("Error deleting the zip file: %v", err)
	}

}
