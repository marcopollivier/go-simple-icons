package lib

import "testing"

func TestTitleToSlug(testing *testing.T) {
	if TitleToSlug("Default") != "default" {
		testing.Error("Error to convert a regular string")
	}

	if TitleToSlug("Default+") != "defaultplus" {
		testing.Error("Error to convert a regular string with +")
	}

	if TitleToSlug("Default.") != "defaultdot" {
		testing.Error("Error to convert a regular string with .")
	}

	if TitleToSlug("Default&") != "defaultand" {
		testing.Error("Error to convert a regular string with &")
	}

	if TitleToSlug("Defaultđ") != "defaultd" {
		testing.Error("Error to convert a regular string with đ")
	}

	if TitleToSlug("Defaultħ") != "defaulth" {
		testing.Error("Error to convert a regular string with ħ")
	}

	if TitleToSlug("Defaultı") != "defaulti" {
		testing.Error("Error to convert a regular string with ı")
	}

	if TitleToSlug("Defaultĸ") != "defaultk" {
		testing.Error("Error to convert a regular string with ĸ")
	}

	if TitleToSlug("Defaultŀ") != "defaultl" {
		testing.Error("Error to convert a regular string with ŀ")
	}

	if TitleToSlug("Defaultł") != "defaultl" {
		testing.Error("Error to convert a regular string with ł")
	}

	if TitleToSlug("Defaultß") != "defaultss" {
		testing.Error("Error to convert a regular string with ß")
	}

	if TitleToSlug("Defaultŧ") != "defaultt" {
		testing.Error("Error to convert a regular string with ŧ")
	}

}
