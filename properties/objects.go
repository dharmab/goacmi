package properties

const (
	// Name is the common notation for each object.
	Name = "Name"
	// Type is the object's type tags.
	Type = "Type"
	// Transform contains the object's coordinates, relative to the global reference point.
	Transform = "T"
	// AdditionalType adds additional tags to the object's type tags. This is only present in the Tacview database, not in telemetry files.
	AdditionalType = "AdditionalType"
	// Parent object ID in hexadecimal format. Ex: A missile's parent is the aircraft that launched it.
	Parent = "Parent"
	// Next object ID in hexadecimal format. Ex: The next waypoint in a flight plan.
	Next = "Next"
	// ShortName is an abbreviated name of the object, e.g. "A-10A".
	ShortName = "ShortName"
	// LongName is a detailed name of the object, e.g. "A-10A Thunderbolt II".
	LongName = "LongName"
	// FullName is the full name of the object, e.g. "Fairchild Republic A-10A Thunderbolt II".
	FullName = "FullName"
	// CallSign is the object's callsign, e.g. "Dodge 1-1".
	CallSign = "CallSign"
	// Registration is the aircraft registration number (tail number).
	Registration = "Registration"
	// Squawk is the current transponder code.
	Squawk = "Squawk"
	// ICAO24 is the aircraft's ICAO 24-bit address in hexadecimal format.
	ICAO24 = "ICAO24"
	// Pilot is the name of the pilot in command. In DCS World, this is the player's name.
	Pilot = "Pilot"
	// Group the object belongs to.
	Group = "Group"
	// Country is an ISO 3166-1 alpha-2 country code.
	Country   = "Country"
	Coalition = "Coalition"
	// Color is the display color of the object.
	Color = "Color"
	// Filename of the displayed 3D model.
	Shape = "Shape"
	// Debug text visibile if debugging is enabled in Tacview.
	Debug = "Debug"
	// Label is free text visible in Tacview.
	Label = "Label"
	// FocusedTarget is the ID of an object that this object is focused on. Ex: Object designated by laster.
	FocusedTarget = "FocusedTarget"
	// LockedTarget is the ID of an object that this object is locked on with any sensor.
	LockedTarget = "LockedTarget"
	// LockedTarget2 is the ID of an object that this object is locked on with any sensor.
	LockedTarget2 = "LockedTarget2"
	// LockedTarget3 is the ID of an object that this object is locked on with any sensor.
	LockedTarget3 = "LockedTarget3"
	// LockedTarget4 is the ID of an object that this object is locked on with any sensor.
	LockedTarget4 = "LockedTarget4"
	// LockedTarget5 is the ID of an object that this object is locked on with any sensor.
	LockedTarget5 = "LockedTarget5"
	// LockedTarget6 is the ID of an object that this object is locked on with any sensor.
	LockedTarget6 = "LockedTarget6"
	// LockedTarget7 is the ID of an object that this object is locked on with any sensor.
	LockedTarget7 = "LockedTarget7"
	// LockedTarget8 is the ID of an object that this object is locked on with any sensor.
	LockedTarget8 = "LockedTarget8"
	// LockedTarget9 is the ID of an object that this object is locked on with any sensor.
	LockedTarget9 = "LockedTarget9"
	// Important is a ratio indicating the object's relative importance.
	Importance = "Importance"
	// Slot indexes the objection's position within a group. Ex: 0 is the first object in a group.
	Slot = "Slot"
	// Disabled indicates if the object is disabled (but not destroyed). This is useful for real-world training.
	// 0 indicates the object is enabled, 1 indicates the object is disabled.
	Disabled = "Disabled"
	// Visible is a ratio indicating the object's visibility in the client. 0 is invisible, 1 is fully visible.
	Visible = "Visible"
	// Health is a ratio indicating the object's health. 0 is destroyed, 1 is fully healthy.
	Health = "Health"
	// Length of the object in meters.
	Length = "Length"
	// Width of the object in meters.
	Width = "Width"
	// Height of the object in meters.
	Height = "Height"
	// Radius of the object's boudning sphere in meters.
	Radius = "Radius"
	// IAS is the indicated airspeed in meters per second.
	IAS = "IAS"
	// CAS is the calibrated airspeed in meters per second.
	CAS = "CAS"
	// TAS is the true airspeed in meters per second.
	TAS = "TAS"
	// Mach is the Mach number.
	Mach = "Mach"
	// AOA is the angle of attack in degrees.
	AOA = "AOA"
	// AOS is the angle of sideslip in degrees.
	AOS = "AOS"
	// AGL is the altitude above ground level in meters.
	AGL = "AGL"
	// HDG is the aircraft's true heading in degrees.
	HDG = "HDG"
	// HDM is the aircraft's magnetic heading in degrees.
	HDM = "HDM"
	// Throttle is the throtle handle position for engine 1. Values below 0 indicate reverse thrust, and values above 1 indicate afterburner.
	Throttle = "Throttle"
	// Throttle2 is the throtle handle position for engine 2. Values below 0 indicate reverse thrust, and values above 1 indicate afterburner.
	Throttle2 = "Throttle2"
	// EngineRPM is the revolutions per minute of engine 1.
	EngineRPM = "EngineRPM"
	// EngineRPM2 is the revolutions per minute of engine 2.
	EngineRPM2 = "EngineRPM2"
	// NR is the normalized rotation speed of engine 1, expressed as a ratio compared to 100% RPM.
	NR = "NR"
	// NR2 is the normalized rotation speed of engine 2, expressed as a ratio compared to 100% RPM.
	NR2 = "NR2"
	// RotorRPM is the revolutions per minute of the main rotor.
	RotorRPM = "RotorRPM"
	// RotorRPM2 is the revolutions per minute of the secondary rotor.
	RotorRPM2 = "RotorRPM2"
	// Afterburner is the afterburner status of Engine 1, expressed as a ratio. 0 is off, 1 is full afterburner.
	Afterburner = "Afterburner"
	// AirBrakes is the airbrake extension expressed as a ratio.
	AirBrakes = "AirBrakes"
	// Flaps position expressed as a ratio.
	Flaps = "Flaps"
	// LandingGear is the position of the landing gear expressed as a ratio.
	LandingGear = "LandingGear"
	// LandingGearHandle is the position of the landing gear handle expressed as a ratio.
	LandingGearHandle = "LandingGearHandle"
	// Tailhook position expressed as a ratio.
	Tailhook = "Tailhook"
	// Parachute status expressed as a ratio. Note this is not the dragging chute.
	Parachute = "Parachute"
	// DragChute is the dragging chute status expressed as a ratio.
	DragChute = "DragChute"
	// FuelWeight is the quantiy of fuel in the first tank in kilograms.
	FuelWeight = "FuelWeight"
	// FuelWeight2 is the quantity of fuel in the second tank in kilograms.
	FuelWeight2 = "FuelWeight2"
	// FuelWeight3 is the quantity of fuel in the third tank in kilograms.
	FuelWeight3 = "FuelWeight3"
	// FuelWeight4 is the quantity of fuel in the fourth tank in kilograms.
	FuelWeight4 = "FuelWeight4"
	// FuelWeight5 is the quantity of fuel in the fifth tank in kilograms.
	FuelWeight5 = "FuelWeight5"
	// FuelWeight6 is the quantity of fuel in the sixth tank in kilograms.
	FuelWeight6 = "FuelWeight6"
	// FuelWeight7 is the quantity of fuel in the seventh tank in kilograms.
	FuelWeight7 = "FuelWeight7"
	// FuelWeight8 is the quantity of fuel in the eighth tank in kilograms.
	FuelWeight8 = "FuelWeight8"
	// FuelWeight9 is the quantity of fuel in the ninth tank in kilograms.
	FuelWeight9 = "FuelWeight9"
	// FuelVolume is the volume of fuel in the first tank in liters.
	FuelVolume = "FuelVolume"
	// FuelFlowWeight is the fueld flow for the first engine in kilograms per hour.
	FuelFlowWeight = "FuelFlowWeight"
	// FuelFlowWeight2 is the fuel flow for the second engine in kilograms per hour.
	FuelFlowWeight2 = "FuelFlowWeight2"
	// FuelFlowWeight3 is the fuel flow for the third engine in kilograms per hour.
	FuelFlowWeight3 = "FuelFlowWeight3"
	// FuelFlowWeight4 is the fuel flow for the fourth engine in kilograms per hour.
	FuelFlowWeight4 = "FuelFlowWeight4"
	// FuelFlowWeight5 is the fuel flow for the fifth engine in kilograms per hour.
	FuelFlowWeight5 = "FuelFlowWeight5"
	// FuelFlowWeight6 is the fuel flow for the sixth engine in kilograms per hour.
	FuelFlowWeight6 = "FuelFlowWeight6"
	// FuelFlowWeight7 is the fuel flow for the seventh engine in kilograms per hour.
	FuelFlowWeight7 = "FuelFlowWeight7"
	// FuelFlowWeight8 is the fuel flow for the eighth engine in kilograms per hour.
	FuelFlowWeight8 = "FuelFlowWeight8"
	// FuelFlowVolume is the fuel flow for the first engine in liters per hour.
	FuelFlowVolume = "FuelFlowVolume"
	// FuelFlowVolume2 is the fuel flow for the second engine in liters per hour.
	FuelFlowVolume2 = "FuelFlowVolume2"
	// FuelFlowVolume3 is the fuel flow for the third engine in liters per hour.
	FuelFlowVolume3 = "FuelFlowVolume3"
	// FuelFlowVolume4 is the fuel flow for the fourth engine in liters per hour.
	FuelFlowVolume4 = "FuelFlowVolume4"
	// FuelFlowVolume5 is the fuel flow for the fifth engine in liters per hour.
	FuelFlowVolume5 = "FuelFlowVolume5"
	// FuelFlowVolume6 is the fuel flow for the sixth engine in liters per hour.
	FuelFlowVolume6 = "FuelFlowVolume6"
	// FuelFlowVolume7 is the fuel flow for the seventh engine in liters per hour.
	FuelFlowVolume7 = "FuelFlowVolume7"
	// FuelFlowVolume8 is the fuel flow for the eighth engine in liters per hour.
	FuelFlowVolume8 = "FuelFlowVolume8"
	// RadarMode indicates the radar mode. 0 indicates off.
	RadarMode = "RadarMode"
	// RadarAzimuth is the radar azimuth in degrees.
	RadarAzimuth = "RadarAzimuth"
	// RadarElevation is the radar elevation in degrees.
	RadarElevation = "RadarElevation"
	// RadarRoll is the radar roll angle in degrees.
	RadarRoll = "RadarRoll"
	// RadaraRange is the radar scan range in meters.
	RadarRange = "RadarRange"
	// RadarHorizontalBeamwidth is the radar azimuth beam width in degrees.
	RadarHorizontalBeamwidth = "RadarHorizontalBeamwidth"
	// RadarVerticalBeamwidth is the radar elevation beam width in degrees.
	RadarVerticalBeamwidth = "RadarVerticalBeamwidth"
	// RadarRangeGateAzimuth is the radar azimuth range gate in degrees.
	RadarRangeGateAzimuth = "RadarRangeGateAzimuth"
	// RadarRangeGateElevation is the radar elevation range gate in degrees.
	RadarRangeGateElevation = "RadarRangeGateElevation"
	// RadarRangeGateRoll is the radar roll angle range gate in degrees.
	RadarRangeGateRoll = "RadarRangeGateRoll"
	// RadarRangeGateMin is the closer limit of the radar range in meters.
	RadarRangeGateMin = "RadarRangeGateMin"
	// RadarRangeGateMax is the further limit of the radar range in meters.
	RadarRangeGateMax = "RadarRangeGateMax"
	// RadarRangeGateHorizontalBeamwidth is the radar azimuth beam width  in degrees.
	RadarRangeGateHorizontalBeamwidth = "RadarRangeGateHorizontalBeamwidth"
	// RadarRangeGateVerticalBeamwidth is the radar elevation beam width in degrees.
	RadarRangeGateVerticalBeamwidth = "RadarRangeGateVerticalBeamwidth"
	// LockedTargetMode indicates the targer lock mode for the primary target, with 0 indicating no lock.
	LockedTargetMode = "LockedTargetMode"
	// LockedTargetAzimuth indicates the azimuth of the primary locked target relaitve to this aircraft in degrees.
	LockedTargetAzimuth = "LockedTargetAzimuth"
	// LockedTargetElevation indicates the elevation of the primary locked target relaitve to this aircraft in degrees.
	LockedTargetElevation = "LockedTargetElevation"
	// LockedTargetRange indicates the range to the primary locked target in meters.
	LockedTargetRange = "LockedTargetRange"
	// EnagementMode indicates the engagement mode, such as for SAM sites toggling their radars. 0 indicates off.
	EngagementMode = "EngagementMode"
	// EngagementMode2 indicates the engagement mode, such as for SAM sites toggling their radars. 0 indicates off.
	EngagementMode2 = "EngagementMode2"
	// EngagementRange is the radius of the engagement range displayed in Tacview.
	EngagementRange = "EngagementRange"
	// EngagementRange2 is the radius of the engagement range displayed in Tacview.
	EngagementRange2 = "EngagementRange2"
	// VerticalEngagementRange is the vertical engagement range displayed in Tacview. If this is not present, the vertical engagement range is the same as the horizontal engagement range.
	VerticalEngagementRange = "VerticalEngagementRange"
	// VerticalEngagementRange2 is the vertical engagement range displayed in Tacview. If this is not present, the vertical engagement range is the same as the horizontal engagement range.
	VerticalEngagementRange2 = "VerticalEngagementRange2"
	// RollControlInput is the position of the input device's roll axis.
	RollControlInput = "RollControlInput"
	// PitchControlInput is the position of the input device's pitch axis.
	PitchControlInput = "PitchControlInput"
	// YawControlInput is the position of the input device's yaw axis.
	YawControlInput = "YawControlInput"
	// RollControlPosition is the position of aircraft's roll control device (after input curves are applied).
	RollControlPosition = "RollControlPosition"
	// PitchControlPosition is the position of aircraft's pitch control device (after input curves are applied).
	PitchControlPosition = "PitchControlPosition"
	// YawControlPosition is the position of aircraft's yaw control device (after input curves are applied).
	YawControlPosition = "YawControlPosition"
	// RollTrimTab is the position of the roll trimmer expressed as a ratio.
	RollTrimTab = "RollTrimTab"
	// PitchTrimTab is the position of the pitch trimmer expressed as a ratio.
	PitchTrimTab = "PitchTrimTab"
	// YawTrimTab is the position of the yaw trimmer expressed as a ratio.
	YawTrimTab = "YawTrimTab"
	// AileronLeft is the position of the left aileron expressed as a ratio.
	AileronLeft = "AileronLeft"
	// AileronRight is the position of the right aileron expressed as a ratio.
	AileronRight = "AileronRight"
	// Elevator is the position of the elevator expressed as a ratio.
	Elevator = "Elevator"
	// Rudder is the position of the rudder expressed as a ratio.
	Rudder = "Rudder"
	// LocalizerLateralDeviation is the lateral deviation from the runway centerline in meters. Negative values are to the left and positive values are to the right.
	LocalizerLateralDeviation = "LocalizerLateralDeviation"
	// GlideslopeVerticalDeviation is the vertical deviation from the glideslope in meters. Negative values are below the glideslope and positive values are above the glideslope.
	GlideslopeVerticalDeviation = "GlideslopeVerticalDeviation"
	// LocalizerAngularDeviation is the angular deviation from the runway centerline in degrees. Negative values are to the left and positive values are to the right.
	LocalizerAngularDeviation = "LocalizerAngularDeviation"
	// GlideslopeAngularDeviation is the angular deviation from the glideslope in degrees. Negative values are below the glideslope and positive values are above the glideslope.
	GlideslopeAngularDeviation = "GlideslopeAngularDeviation"
	// PilotHeadRoll is the roll angle of the pilot's head in degrees.
	PilotHeadRoll = "PilotHeadRoll"
	// PilotHeadPitch is the pitch angle of the pilot's head in degrees.
	PilotHeadPitch = "PilotHeadPitch"
	// PilotHeadYaw is the yaw angle of the pilot's head in degrees.
	PilotHeadYaw = "PilotHeadYaw"
	// PilotEyeGazePitch is the vertical orientation of the pilot's eyes in degrees, relative to the pilot's head.
	PilotEyeGazePitch = "PilotEyeGazePitch"
	// PilotEyeGazeYaw is the horizontal orientation of the pilot's eyes in degrees, relative to the pilot's head.
	PilotEyeGazeYaw = "PilotEyeGazeYaw"
	// VerticalGForce is the force on the aircraft's vertical axis in g.
	VerticalGForce = "VerticalGForce"
	// LongitudinalGForce is the force on the aircraft's longitudinal axis in g.
	LongitudinalGForce = "LongitudinalGForce"
	// LateralGForce is the force on the aircraft's lateral axis in g.
	LateralGForce = "LateralGForce"
	// QNH is the regiontal barometric altimitmer setting that calibrates the alimter to read altitude above sea level, in hectopascals.
	QNH = "QNH"
	// WindDirection is the wind direction relative to true north in degrees.
	WindDirection = "WindDirection"
	// WindPitch is the wind pitch in degrees.
	WindPitch = "WindPitch"
	// WindSpeed is the wind 3D speed in meters per second.
	WindSpeed = "WindSpeed"
	// TriggerPressed is the position of the trigger expressed as a ratio, with 1.0 indicating full depression.
	TriggerPressed = "TriggerPressed"
	// HeartRate is the pilot's heart rate in beats per minute.
	HeartRate = "HeartRate"
	// SpO2 is the pilot's blood oxygen saturation level in percent.
	SpO2 = "SpO2"
)
