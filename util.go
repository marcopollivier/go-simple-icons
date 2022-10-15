/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import "strings"

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
