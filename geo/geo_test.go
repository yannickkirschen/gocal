package geo

import "testing"

func TestString(t *testing.T) {
	coordinates := Coordinates{
		Latitude:  49.474193,
		Longitude: 8.534854,
	}

	expected := "49.474193,8.534854"
	if coordinates.String() != expected {
		t.Errorf("Coordinates string is not as expected: %s", coordinates.String())
	}
}

func TestStringSep(t *testing.T) {
	coordinates := Coordinates{
		Latitude:  49.474193,
		Longitude: 8.534854,
	}

	expected := "49.474193;8.534854"
	if coordinates.StringSep(";") != expected {
		t.Errorf("Coordinates string is not as expected: %s", coordinates.StringSep(";"))
	}
}
