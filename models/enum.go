package models

// tracker device type
var TrackerDeviceType = newTrackerDeviceTypeRegistry()

func newTrackerDeviceTypeRegistry() *trackerDeviceTypeRegistry {
	return &trackerDeviceTypeRegistry{
		Gemeni: 1,
		Tedi:   2,
		Pian:   3,
	}
}

type trackerDeviceTypeRegistry struct {
	Gemeni int16
	Tedi   int16
	Pian   int16
}

// vehicle type
var VehicleType = newVehicleTypeRegistry()

func newVehicleTypeRegistry() *vehicleTypeRegistry {
	return &vehicleTypeRegistry{
		LMV:  1,
		HMV:  2,
		Bike: 3,
	}
}

type vehicleTypeRegistry struct {
	LMV  int16
	HMV  int16
	Bike int16
}
