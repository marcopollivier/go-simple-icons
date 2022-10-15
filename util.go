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

func TitleToSlug(title string) string {
	var lowerTitle = strings.ToLower(title)
	replacer := strings.NewReplacer(
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
	return replacer.Replace(lowerTitle)
}

func readFile(filePath string) []byte {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return fileContent
}

func ReadDataJsonFile() string {
	return string(readFile("data/simple-icons.json"))
}

func ReadSvgFile(filename string) string {
	return string(readFile("data/icons/" + filename + ".svg"))
}
