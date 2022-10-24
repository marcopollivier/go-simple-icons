/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import "testing"

func TestTitleToSlug(testing *testing.T) {
	if titleToSlug("Default") != "default" {
		testing.Error("Error to convert a regular string")
	}

	if titleToSlug("Default+") != "defaultplus" {
		testing.Error("Error to convert a regular string with +")
	}

	if titleToSlug("Default.") != "defaultdot" {
		testing.Error("Error to convert a regular string with .")
	}

	if titleToSlug("Default&") != "defaultand" {
		testing.Error("Error to convert a regular string with &")
	}

	if titleToSlug("Defaultđ") != "defaultd" {
		testing.Error("Error to convert a regular string with đ")
	}

	if titleToSlug("Defaultħ") != "defaulth" {
		testing.Error("Error to convert a regular string with ħ")
	}

	if titleToSlug("Defaultı") != "defaulti" {
		testing.Error("Error to convert a regular string with ı")
	}

	if titleToSlug("Defaultĸ") != "defaultk" {
		testing.Error("Error to convert a regular string with ĸ")
	}

	if titleToSlug("Defaultŀ") != "defaultl" {
		testing.Error("Error to convert a regular string with ŀ")
	}

	if titleToSlug("Defaultł") != "defaultl" {
		testing.Error("Error to convert a regular string with ł")
	}

	if titleToSlug("Defaultß") != "defaultss" {
		testing.Error("Error to convert a regular string with ß")
	}

	if titleToSlug("Defaultŧ") != "defaultt" {
		testing.Error("Error to convert a regular string with ŧ")
	}

	if titleToSlug("Simple Icons") != "simpleicons" {
		testing.Error("Error to convert a regular string with white spaces")
	}

	if titleToSlug("Simple Icons ßŧ") != "simpleiconssst" {
		testing.Error("Error to convert a regular string with multiple replaces")
	}

}
