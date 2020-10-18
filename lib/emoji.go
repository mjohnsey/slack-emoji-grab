package lib

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

type Emoji struct {
	Name          string `json:"name"`
	IsAlias       bool   `json:"is_alias"`
	URL           string `json:"url"`
	AliasForName  string `json:"alias_for"`
	CreatedByName string `json:"user_display_name"`
	CreatedTs     int    `json:"created"`
}

func (e Emoji) SaveToFile(resultDir string) error {
	// Adapted from: https://stackoverflow.com/a/22417396/2371482
	if _, err := os.Stat(resultDir); os.IsNotExist(err) {
		return errors.Wrap(err, "The resultDir did not exist!")
	}
	response, err := http.Get(e.URL)
	if err != nil {
		newError := errors.Wrap(err, "Problem downloading the URL!")
		log.Println(newError)
		return newError
	}
	//
	defer response.Body.Close()
	//
	////open a file for writing
	var fullFileName = e.Name + e.FileExtension()
	var newFilePath = path.Join(resultDir, fullFileName)
	file, err := os.Create(newFilePath)
	if err != nil {
		return errors.Wrap(err, "Could not open the result file!")
	}
	//// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.Wrap(err, "Problem copying the file!")
	}
	file.Close()
	log.Printf("Done downloading: %s", newFilePath)
	return nil
}

func (e Emoji) FileExtension() string {
	return filepath.Ext(e.URL)
}
