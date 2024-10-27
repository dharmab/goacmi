package events

type Event string

const (
	// Mesage is a generic event.
	Message Event = "Message"
	// Bookmark highlights a point in the timeline and event log.
	Bookmark Event = "Bookmark"
	// Debug highlights a point in the timeline and event log. Debug events are only visible when running Tacview with debugging enabled.
	Debug Event = "Debug"
	// LeftArea occurs when an aircraft (or any object) is cleanly removed from the battlefield (i.e. not destroyed).
	LeftArea Event = "LeftArea"
	// Destroyed occurs when an object has been destroyed.
	Destroyed Event = "Destroyed"
	// TakenOff may occur when an aircraft takes off to explicitly indicate a takeoff.
	TakenOff Event = "TakenOff"
	// Landed may occur when an aircraft lands to explicitly indicate a landing.
	LandedEvent Event = "Landed"
	// Timeout occurs when a simulated weapon reaches or misses its target. This is mainly used for shot logs for real-world training.
	TimeoutEvent = "Timeout"
)
