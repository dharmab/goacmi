package objects

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dharmab/goacmi/properties"
	"github.com/martinlindhe/unit"
)

// Object models the ID and properties of an object within ACMI data.
type Object struct {
	ID         uint64
	Properties map[string]string
}

// / Update models the changes to an object within a line of ACMI data.
type Update struct {
	// ID of the object.
	ID uint64
	// IsRemove indicates if the update removes the object.
	IsRemoval bool
	// Properties provided in the update.
	Properties map[string]string
}

// ErrParseFailure indicates that a property is present but failed to parse.
var ErrParseFailure = errors.New("error parsing object")

// ErrNotFound indicates that a property is missing.
var ErrNotFound = errors.New("property not found")

// New instantiates a new object with the given ID.
func New(id uint64) *Object {
	return &Object{
		ID:         id,
		Properties: make(map[string]string),
	}
}

// Update applies the given update to the object. If the update contains a coordinate transformation, the object's
// coordinates are updated. Other properties in the update are merged into the object's properties.
// The referenceLongitude and referenceLatitude are the reference point from the global properties.
func (o *Object) Update(update *Update, referenceLongitude, referenceLatitude float64) error {
	var newTransform string
	for k, v := range update.Properties {
		if k == properties.Transform {
			newTransform = v
		} else {
			o.SetProperty(k, v)
		}
	}

	if newTransform == "" {
		return nil
	}
	currentCoordinates, err := o.GetCoordinates(referenceLongitude, referenceLatitude)
	if err != nil {
		o.SetProperty(properties.Transform, newTransform)
		//lint:ignore nilerr intentional overwrite of invalid coordinates
		return nil // ignore
	}
	err = currentCoordinates.Parse(newTransform, referenceLongitude, referenceLatitude)
	if err != nil {
		return fmt.Errorf("error parsing coordinates: %w", err)
	}
	o.SetProperty(properties.Transform, currentCoordinates.Transform(referenceLongitude, referenceLatitude))
	return nil
}

func (o *Object) SetProperty(p, v string) {
	o.Properties[p] = v
}

func (o *Object) GetProperty(p string) (string, bool) {
	val, ok := o.Properties[p]
	if !ok {
		return "", false
	}
	return val, true
}

// GetTypes returns all object type tags.
func (o *Object) GetTypes() ([]string, error) {
	val, ok := o.GetProperty(properties.Type)
	if !ok {
		return nil, errors.Join(ErrNotFound, errors.New("object does not contain Type property"))
	}
	return strings.Split(val, "+"), nil
}

// GetCoordinates returns the coordinates of the object, if possible.
// Objects may have insufficient information to determine their coordinates.
// In such a case, the function returns nil and no error.
// The referenceLongitude and referenceLatitude are the reference point from the global properties.
func (o *Object) GetCoordinates(referenceLongitude, referenceLatitude float64) (*Coordinates, error) {
	val, ok := o.GetProperty(properties.Transform)
	if !ok {
		return nil, errors.Join(ErrNotFound, errors.New("object does not contain Transform property"))
	}

	c := &Coordinates{}
	err := c.Parse(val, referenceLongitude, referenceLatitude)
	if err != nil {
		return nil, fmt.Errorf("error parsing coordinates: %w", err)
	}
	return c, nil
}

// getNumericProperty is a helper to read the given property as a number.
func (o *Object) getNumericProperty(property string) (float64, error) {
	val, ok := o.GetProperty(property)
	if !ok {
		return 0, fmt.Errorf("object does not contain %s", property)
	}
	n, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing %s: %w", property, err)
	}
	return n, nil
}

// GetSpeed is a helper to read the given property as a speed.
func (o *Object) GetSpeed(property string) (unit.Speed, error) {
	n, err := o.getNumericProperty(property)
	if err != nil {
		return 0, err
	}
	return unit.Speed(n) * unit.MetersPerSecond, nil
}

// GetAngle is a helper to read the given property as an angle.
func (o *Object) GetAngle(property string) (unit.Angle, error) {
	n, err := o.getNumericProperty(property)
	if err != nil {
		return 0, err
	}
	return unit.Angle(n) * unit.Degree, nil
}

// GetLength is a helper to read the given property as a length or distance.
func (o *Object) GetLength(property string) (unit.Length, error) {
	n, err := o.getNumericProperty(property)
	if err != nil {
		return 0, err
	}
	return unit.Length(n) * unit.Meter, nil
}
