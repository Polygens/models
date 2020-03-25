// Package holiday contains api endpoint payloads related to holiday locations
package holiday

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/polygens/models/shared"
)

// Hotel is a hotel point of interest
type Hotel struct {
	// Name of the location
	Name string

	// Location in the world
	Location shared.Coordinate

	// Shape of the hotel
	Shape geojson.Feature
}
