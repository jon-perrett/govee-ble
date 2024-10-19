package bluetooth

import bluetooth "tinygo.org/x/bluetooth"

var adaptor = bluetooth.DefaultAdapter

func GetAdvertisementManufacturerData() map[int][]byte {
	err := adaptor.Enable()

	if err != nil {
		panic("Failed to enable adaptor.")
	}
	data := make(map[int][]byte)
	err = adaptor.Scan(func(adaptor *bluetooth.Adapter, device bluetooth.ScanResult) {
		if device.LocalName() == "GVH5075_A1B3" {
			adaptor.StopScan()
			for _, el := range device.ManufacturerData() {
				data[int(el.CompanyID)] = el.Data
			}
		}

	})
	if err != nil {
		panic("Failed to scan device")
	}
	return data
}
