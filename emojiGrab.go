package main

import (
	"os"
	"github.com/pkg/errors"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"io"
	"path"
	"sync"
)

type Emoji struct {
	Extension string `json:"extension"`
	Name      string `json:"name"`
	URL       string `json:"url"`
}

func (e Emoji) saveToFile (resultDir string) error{
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
	var fullFileName = e.Name + e.Extension
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

type EmojiFile struct {
	Emojis []Emoji `json:"emojis"`
}

func (ef EmojiFile) readFromFile ( fileName string) (EmojiFile){
	// Adapted from https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(errors.Wrap(err,"Could not open the emoji file"))
	}

	var obj EmojiFile
	json.Unmarshal(raw, &obj)
	return obj
}

func (ef EmojiFile) saveAllImages(resultDir string) (error){
	// https://stackoverflow.com/questions/23635070/golang-download-multiple-files-in-parallel-using-goroutines
	if _, err := os.Stat(resultDir); os.IsNotExist(err) {
		log.Fatal(errors.Wrap(err, "This directory does not exist!"))
	}
	var w sync.WaitGroup
	for _, e := range ef.Emojis{
		w.Add(1)
		go func (emoji Emoji) error{
			defer w.Done()
			err := emoji.saveToFile(resultDir)
			if err != nil{
				newError := errors.Wrap(err, "Problem saving one of the emojis!")
				//log.Println(newError)
				return newError
			}
			return nil
		}(e)
	}
	w.Wait()
	return nil

}

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) <= 2{
		log.Fatal("Must pass in a fileName!")
	}
	fileName := argsWithProg[1]
	resultDir := argsWithProg[2]
	log.Printf("Loading %s", fileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal(errors.Wrap(err, "This file does not exist!"))
	}
	var emojiFile = EmojiFile{}.readFromFile(fileName)
	err := emojiFile.saveAllImages(resultDir)
	if err != nil{
		log.Fatal(errors.Wrap(err, "WTF happened?"))
	} else{
		log.Println("DONE!")
	}

}

func removeDuplicates(elements []string) ([]string, []string) {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	uniq := []string{}
	dupe := []string{}

	for i, v := range elements {
		if encountered[elements[i]] == true {
			// Do not add duplicate.
			dupe = append(dupe, v)
		} else {
			// Record this element as an encountered element.
			encountered[elements[i]] = true
			// Append to result slice.
			uniq = append(uniq, v)
		}
	}
	// Return the new slice.
	return uniq, dupe
}