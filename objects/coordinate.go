package objects

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/martinlindhe/unit"
)

// Coordinates models the position and orientation of an object within ACMI data.
type Coordinates struct {
	// Longitude. This is the object's actual longitude, after applying any reference longitude.
	Longitude *float64
	// Latitude. This is the object's actual latitude, after applying any reference latitude.
	Latitude *float64
	// Altitude above sea level
	Altitude *unit.Length
	// X is the object's native coordinate within the sim
	X *float64
	// Y is the object's native coordiante within the sim
	Y *float64
	// Roll angle
	Roll *unit.Angle
	// Pitch angle
	Pitch *unit.Angle
	// Yaw angle (this may be different from heading)
	Yaw *unit.Angle
	// Heading is the object's flat earth heading
	Heading *unit.Angle
}

// NewCoordinates instantiates Coordinates with the given values.
func NewCoordinates(
	longitude, latitude *float64,
	altitude *unit.Length,
	x, y *float64,
	roll, pitch, yaw, heading *unit.Angle,
) *Coordinates {
	return &Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
		Altitude:  altitude,
		X:         x,
		Y:         y,
		Roll:      roll,
		Pitch:     pitch,
		Yaw:       yaw,
		Heading:   heading,
	}
}

// Update updates the coordinates with the values from the next coordinates.
func (c *Coordinates) Update(next *Coordinates) {
	if next.Longitude != nil {
		c.Longitude = next.Longitude
	}
	if next.Latitude != nil {
		c.Latitude = next.Latitude
	}
	if next.Altitude != nil {
		c.Altitude = next.Altitude
	}
	if next.X != nil {
		c.X = next.X
	}
	if next.Y != nil {
		c.Y = next.Y
	}
	if next.Roll != nil {
		c.Roll = next.Roll
	}
	if next.Pitch != nil {
		c.Pitch = next.Pitch
	}
	if next.Yaw != nil {
		c.Yaw = next.Yaw
	}
	if next.Heading != nil {
		c.Heading = next.Heading
	}
}

// Parse reads the value of a Transform property and updates the coordinates accordingly. The referenceLongitude and referenceLatitude parameters must be the reference point from the global properties.
func (c *Coordinates) Parse(transform string, referenceLongitude float64, referenceLatitude float64) error {
	fields := strings.Split(transform, "|")

	if len(fields) < 3 {
		return errors.Join(ErrParseFailure, errors.New("unexpected number of fields in coordinate transformation"))
	}

	var longitude, latitude float64
	if fields[0] != "" {
		offset, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			return errors.Join(ErrParseFailure, fmt.Errorf("error parsing longitude: %w", err))
		}
		longitude = referenceLongitude + offset
		c.Longitude = &longitude
	}
	if fields[1] != "" {
		offset, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return errors.Join(ErrParseFailure, fmt.Errorf("error parsing latitude: %w", err))
		}
		latitude = referenceLatitude + offset
		c.Latitude = &latitude
	}

	var alt, x, y, roll, pitch, yaw, heading string
	switch len(fields) {
	case 3:
		alt = fields[2]
	case 5:
		alt = fields[2]
		x = fields[3]
		y = fields[4]
	case 6:
		alt = fields[2]
		roll = fields[3]
		pitch = fields[4]
		yaw = fields[5]
	case 9:
		alt = fields[2]
		roll = fields[3]
		pitch = fields[4]
		yaw = fields[5]
		x = fields[6]
		y = fields[7]
		heading = fields[8]
	default:
		return errors.Join(ErrParseFailure, errors.New("unexpected number of fields in coordinate transformation"))
	}

	for s, fn := range map[string]func(float64){
		alt: func(v float64) {
			a := unit.Length(v) * unit.Meter
			c.Altitude = &a
		},
		x: func(v float64) {
			c.X = &v
		},
		y: func(v float64) {
			c.Y = &v
		},
		roll: func(v float64) {
			θ := unit.Angle(v) * unit.Degree
			c.Roll = &θ
		},
		pitch: func(v float64) {
			θ := unit.Angle(v) * unit.Degree
			c.Pitch = &θ
		},
		yaw: func(v float64) {
			θ := unit.Angle(v) * unit.Degree
			c.Yaw = &θ
		},
		heading: func(v float64) {
			θ := unit.Angle(v) * unit.Degree
			c.Heading = &θ
		},
	} {
		if s != "" {
			n, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return errors.Join(ErrParseFailure, fmt.Errorf("error parsing coordinate field: %w", err))
			}
			fn(n)
		}
	}

	return nil
}

// Transform serializes the coordinates into a coordinate transformation suitable for use in the Transform property in an ACMI object update.
// The referenceLongitude and referenceLatitude parameters must be the reference point in the global properties.
func (c *Coordinates) Transform(referenceLongitude float64, referenceLatitude float64) string {
	fields := make([]string, 9)
	if c.Longitude != nil {
		fields[0] = fmt.Sprintf("%f", *c.Longitude-referenceLongitude)
	}
	if c.Latitude != nil {
		fields[1] = fmt.Sprintf("%f", *c.Latitude-referenceLatitude)
	}
	if c.Altitude != nil {
		fields[2] = fmt.Sprintf("%f", c.Altitude.Meters())
	}
	if c.Roll != nil {
		fields[3] = fmt.Sprintf("%f", c.Roll.Degrees())
	}
	if c.Pitch != nil {
		fields[4] = fmt.Sprintf("%f", c.Pitch.Degrees())
	}
	if c.Yaw != nil {
		fields[5] = fmt.Sprintf("%f", c.Yaw.Degrees())
	}
	if c.X != nil {
		fields[6] = fmt.Sprintf("%f", *c.X)
	}
	if c.Y != nil {
		fields[7] = fmt.Sprintf("%f", *c.Y)
	}
	if c.Heading != nil {
		fields[8] = fmt.Sprintf("%f", c.Heading.Degrees())
	}
	return strings.Join(fields, "|")
}
