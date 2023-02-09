package models

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

var UserRole = newUserRoleRegistry()

func newUserRoleRegistry() *userRoleRegistry {
	return &userRoleRegistry{
		Administrator: 1,
		Manufacturer:  2,
		EndUser:       3,
	}
}

type userRoleRegistry struct {
	Administrator int16
	Manufacturer  int16
	EndUser       int16
}

var TrackerDeviceStatus = newTrackerDeviceStatusRegistry()

func newTrackerDeviceStatusRegistry() *trackerDeviceStatusRegistry {
	return &trackerDeviceStatusRegistry{
		Unsold:    1,
		Sold:      2,
		Certified: 3,
	}
}

type trackerDeviceStatusRegistry struct {
	Unsold    int16
	Sold      int16
	Certified int16
}
