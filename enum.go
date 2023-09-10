package models

var TrackerDeviceType = newTrackerDeviceTypeRegistry()

func newTrackerDeviceTypeRegistry() *trackerDeviceTypeRegistry {
	return &trackerDeviceTypeRegistry{
		Quantifys: 1,
		Tedi:      2,
		Pian:      3,
		Gemeni:    4,
	}
}

type trackerDeviceTypeRegistry struct {
	Quantifys int16
	Tedi      int16
	Pian      int16
	Gemeni    int16
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
		Distributor:   4,
		Dealer:        5,
		SubDealer:     6,
	}
}

type userRoleRegistry struct {
	Administrator int16
	Manufacturer  int16
	EndUser       int16
	Distributor   int16
	Dealer        int16
	SubDealer     int16
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

var VehicleStatus = newVehicleStatusRegistry()

func newVehicleStatusRegistry() *vehicleStatusRegistry {
	return &vehicleStatusRegistry{
		New:    1,
		Issued: 2,
	}
}

type vehicleStatusRegistry struct {
	New    int16
	Issued int16
}

// vehicle events
var VehicleEvents = newVehicleEventRegistry()

func newVehicleEventRegistry() *vehicleEventRegistry {
	return &vehicleEventRegistry{
		OverSpeed:         1,
		Ignition:          2,
		TamperAlert:       3,
		BatteryDisconnect: 4,
		BatteryInternal:   5,
		HarshBraking:      6,
		HarshAcceleration: 7,
		RashTurning:       8,
	}
}

type vehicleEventRegistry struct {
	OverSpeed         int16
	Ignition          int16
	TamperAlert       int16
	BatteryDisconnect int16
	BatteryInternal   int16
	HarshBraking      int16
	HarshAcceleration int16
	RashTurning       int16
}

// packet status
var PacketStatus = newPacketStatusRegistry()

func newPacketStatusRegistry() *packetStatusRegistry {
	return &packetStatusRegistry{
		Live:    "L",
		History: "H",
	}
}

type packetStatusRegistry struct {
	Live    string
	History string
}

// location provider
var LocationProvider = newLocationProviderRegistry()

func newLocationProviderRegistry() *locationProviderRegistry {
	return &locationProviderRegistry{
		Fine:   1,
		Coarse: 2,
	}
}

type locationProviderRegistry struct {
	Fine   int16
	Coarse int16
}

// packet types
var PacketTypes = newPacketTypeRegistry()

func newPacketTypeRegistry() *packetTypeRegistry {
	return &packetTypeRegistry{
		Normal:            "NR",
		TamperAlert:       "TA",
		BatteryDisconnect: "BD",
		BatteryReconnect:  "BR",
		InternalBattery:   "BL",
		IgnitionOn:        "IN",
		IgnitionOff:       "IF",
		HarshBraking:      "HB",
		HarshAcceleration: "HA",
		EmergencyAlert:    "EA",
		RashTurning:       "RT",
	}
}

type packetTypeRegistry struct {
	Normal            string
	TamperAlert       string
	BatteryDisconnect string
	BatteryReconnect  string
	InternalBattery   string
	IgnitionOn        string
	IgnitionOff       string
	HarshBraking      string
	HarshAcceleration string
	EmergencyAlert    string
	RashTurning       string
}
