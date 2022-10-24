/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import "testing"

func TestTitleToSlug(testing *testing.T) {
	if replaceSpecialCharacteres("Default") != "Default" {
		testing.Error("Error to convert a regular string")
	}

	if replaceSpecialCharacteres("Default+") != "Defaultplus" {
		testing.Error("Error to convert a regular string with +")
	}

	if replaceSpecialCharacteres("Default.") != "Defaultdot" {
		testing.Error("Error to convert a regular string with .")
	}

	if replaceSpecialCharacteres("Default&") != "Defaultand" {
		testing.Error("Error to convert a regular string with &")
	}

	if replaceSpecialCharacteres("Defaultđ") != "Defaultd" {
		testing.Error("Error to convert a regular string with đ")
	}

	if replaceSpecialCharacteres("Defaultħ") != "Defaulth" {
		testing.Error("Error to convert a regular string with ħ")
	}

	if replaceSpecialCharacteres("Defaultı") != "Defaulti" {
		testing.Error("Error to convert a regular string with ı")
	}

	if replaceSpecialCharacteres("Defaultĸ") != "Defaultk" {
		testing.Error("Error to convert a regular string with ĸ")
	}

	if replaceSpecialCharacteres("Defaultŀ") != "Defaultl" {
		testing.Error("Error to convert a regular string with ŀ")
	}

	if replaceSpecialCharacteres("Defaultł") != "Defaultl" {
		testing.Error("Error to convert a regular string with ł")
	}

	if replaceSpecialCharacteres("Defaultß") != "Defaultss" {
		testing.Error("Error to convert a regular string with ß")
	}

	if replaceSpecialCharacteres("Defaultŧ") != "Defaultt" {
		testing.Error("Error to convert a regular string with ŧ")
	}

	if replaceSpecialCharacteres("Simple Icons") != "SimpleIcons" {
		testing.Error("Error to convert a regular string with white spaces")
	}

	if replaceSpecialCharacteres("Simple Icons ßŧ") != "SimpleIconssst" {
		testing.Error("Error to convert a regular string with multiple replaces")
	}

}
