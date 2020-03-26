// Package location contains kafka message payloads for the location topic
// Related to locations in the world
package location

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/polygens/models/shared"
)

// Type of the location
type Type uint

const (
	// Hotel is a place to sleep
	Hotel Type = iota
	// Parc is place of green
	Parc
	// Restaurant is a place to eat
	Restaurant
	// Bank is place to steal money
	Bank
)

// Location is any type of point of interest
type Location struct {
	// Name of the location
	Name string
	// LocationType of the location
	LocationType Type
	// Location in the world
	Location shared.Coordinate
	// Geometry of the point of interest
	Geometry *geojson.Feature `json:",omitempty"`
}
