/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import (
	"encoding/json"
	"errors"
	"strings"
)

type License struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Icon struct {
	Title      string  `json:"title"`
	Slug       string  `json:"slug"`
	Hex        string  `json:"hex"`
	Source     string  `json:"source"`
	Svg        string  `json:"svg"`
	Path       string  `json:"path"`
	Guidelines string  `json:"guidelines"`
	License    License `json:"license"`
}

type Icons struct {
	Icons []Icon `json:"icons"`
}

func dataJsonToIcons() (Icons, error) {
	var fileContent = ReadDataJsonFile()

	data := Icons{}
	if json.Unmarshal([]byte(fileContent), &data) != nil {
		return data, errors.New("icon not found")
	}

	return data, nil
}

func findIconByTitle(title string, icons Icons) (result *Icon) {
	result = nil
	for _, icon := range icons.Icons {
		if icon.Title == title {
			result = &icon
			break
		}
	}
	return result
}

func titleToSlug(title string) string {
	var lowerTitle = strings.ToLower(title)
	return replaceSpecialCharacteres(lowerTitle)
}

// GetIcon Returns the Icon Object
func GetIcon(iconName string) (*Icon, error) {
	icons, err := dataJsonToIcons()
	if err != nil {
		return nil, err
	}

	icon := findIconByTitle(iconName, icons)
	slug := titleToSlug(icon.Title)
	icon.Svg = ReadSvgFile(slug)

	return icon, nil
}
