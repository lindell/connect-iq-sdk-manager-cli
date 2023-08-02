package connectiq

// DeviceGroup express the type of device
type DeviceGroup string

const (
	DeviceGroupWatchesWearables DeviceGroup = "Watches/Wearables"
	DeviceGroupEdge             DeviceGroup = "Edge"
	DeviceGroupOutdoorHandhelds DeviceGroup = "Outdoor Handhelds"
)
