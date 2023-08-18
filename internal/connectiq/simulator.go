package connectiq

// SimulatorFile represent useful information in the simulator.json file located in the root/Devices/[Device]/simulator.json
type SimulatorFile struct {
	Fonts []struct {
		FontSet string `json:"fontSet"`
		Fonts   []struct {
			Filename string `json:"filename"`
			Name     string `json:"name"`
		} `json:"fonts"`
	} `json:"fonts"`
}
