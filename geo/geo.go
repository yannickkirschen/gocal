package geo

import "fmt"

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinates(latitude, longitude float64) *Coordinates {
	return &Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (c Coordinates) String() string {
	return fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
}

func (c Coordinates) StringSep(sep string) string {
	return fmt.Sprintf("%f%s%f", c.Latitude, sep, c.Longitude)
}
