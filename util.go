/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import (
	"io/ioutil"
	"log"
	"strings"
)

func replaceSpecialCharacteres(original string) string {
	var replacer = strings.NewReplacer(
		" ", "",
		".", "dot",
		"+", "plus",
		"&", "and",
		"đ", "d",
		"ħ", "h",
		"ı", "i",
		"ĸ", "k",
		"ŀ", "l",
		"ł", "l",
		"ß", "ss",
		"ŧ", "t")
	return replacer.Replace(original)
}

func readFile(filePath string) []byte {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return fileContent
}

// ReadDataJsonFile return as string the data JSON file content
func ReadDataJsonFile() string {
	return string(readFile("data/simple-icons.json"))
}

// ReadSvgFile return as string the data of the SVG file content
func ReadSvgFile(filename string) string {
	return string(readFile("data/icons/" + filename + ".svg"))
}
