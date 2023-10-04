package helpers

func ConvertByteToGB(byt int64) float64 {
	return float64(byt) / 1073741824
}

func ConvertByteToTB(byt int64) float64 {
	return float64(byt) / (1024 * 1024 * 1024 * 1024)
}

func ConvertGBtoTB(gb float64) float64 {
	return gb / 1024
}
