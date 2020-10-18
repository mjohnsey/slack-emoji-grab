package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	lib "github.com/mjohnsey/slack-emoji-grab/lib"
	"github.com/pkg/errors"
)

type EmojiFile struct {
	Emojis []lib.Emoji `json:"emoji"`
}

func (ef EmojiFile) readFromFile(fileName string) EmojiFile {
	// Adapted from https://www.chazzuka.com/2015/03/load-parse-json-file-golang/
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not open the emoji file"))
	}

	var obj EmojiFile
	json.Unmarshal(raw, &obj)
	return obj
}

func (ef EmojiFile) saveAllImages(resultDir string) error {
	// https://stackoverflow.com/questions/23635070/golang-download-multiple-files-in-parallel-using-goroutines
	if _, err := os.Stat(resultDir); os.IsNotExist(err) {
		log.Fatal(errors.Wrap(err, "This directory does not exist!"))
	}
	var w sync.WaitGroup
	for _, e := range ef.Emojis {
		w.Add(1)
		go func(emoji lib.Emoji) error {
			defer w.Done()
			err := emoji.SaveToFile(resultDir)
			if err != nil {
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
	if len(argsWithProg) <= 2 {
		log.Fatal("Must pass in a fileName and resultDir!")
	}
	fileName := argsWithProg[1]
	resultDir := argsWithProg[2]
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatal(errors.Wrap(err, "This file does not exist!"))
	}
	emojiFile := EmojiFile{}.readFromFile(fileName)
	err := emojiFile.saveAllImages(resultDir)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Done!")
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
