/**
 * Author: Marco Ollivier
 * File: init.go
 */

package lib

import "testing"

func TestGetIcon(testing *testing.T) {
	icon, err := GetIcon("Simple Icons")

	if err != nil {
		testing.Error("Unexpected error")
	}

	if icon.Title != "Simple Icons" {
		testing.Error("The Icon name is wrong")
	}

}
