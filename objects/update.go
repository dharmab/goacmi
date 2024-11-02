package objects

// / Update models the changes to an object within a line of ACMI data.
type Update struct {
	// ID of the object.
	ID uint64
	// IsRemove indicates if the update removes the object.
	IsRemoval bool
	// Properties provided in the update.
	Properties map[string]string
}

func (u *Update) String() string {
	return asString(u.IsRemoval, u.ID, u.Properties)
}
