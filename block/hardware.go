package block

type HardwareDevice struct {
	*FileDevice
}

// no allocation on hardware
func (h *HardwareDevice) Prepare() error {
	return nil
}
