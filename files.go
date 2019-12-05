package helferlein

import (
	"errors"
	"io"
	"net/http"
	"os"
)

// DownloadFiles downloads a file to a given path
func DownloadFile(url string, filepath string) error {
	// Create the file
	if _, err := os.Stat(filepath); err == nil {
		return errors.New("File already exists")
	} else if os.IsNotExist(err) {
		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}
