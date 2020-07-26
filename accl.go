package bmx055

func (b *BMX055Driver) GetAcclData() ([]float64, error) {
	raw, err := b.GetAcclRawData()
	if err != nil {
		return nil, err
	}
	return b.convertRawDataToMs2(raw), nil
}

func (b *BMX055Driver) GetAcclRawData() ([]int, error) {
	var err error
	data := make([]uint8, 6)

	data[0], err = b.acclCon.ReadByteData(ACC_REG_ACCD_X_LSB)
	if err != nil {
		return nil, err
	}
	data[1], err = b.acclCon.ReadByteData(ACC_REG_ACCD_X_MSB)
	if err != nil {
		return nil, err
	}
	data[2], err = b.acclCon.ReadByteData(ACC_REG_ACCD_Y_LSB)
	if err != nil {
		return nil, err
	}
	data[3], err = b.acclCon.ReadByteData(ACC_REG_ACCD_Y_MSB)
	if err != nil {
		return nil, err
	}
	data[4], err = b.acclCon.ReadByteData(ACC_REG_ACCD_Z_LSB)
	if err != nil {
		return nil, err
	}
	data[5], err = b.acclCon.ReadByteData(ACC_REG_ACCD_Z_MSB)
	if err != nil {
		return nil, err
	}

	castData := make([]int, 6)

	for i := range data {
		castData[i] = int(data[i])
	}

	xAccl := ((castData[1] * 256) + (castData[0] & 0xF0)) / 16
	if xAccl > 2047 {
		xAccl -= 4096
	}
	yAccl := ((castData[3] * 256) + (castData[2] & 0xF0)) / 16
	if yAccl > 2047 {
		yAccl -= 4096
	}
	zAccl := ((castData[5] * 256) + (castData[4] & 0xF0)) / 16
	if zAccl > 2047 {
		zAccl -= 4096
	}
	return []int{xAccl, yAccl, zAccl}, nil
}

const (
	AccGravity       = 9.80665      // m/s^2
	AccResolution2g  = 0.98 * 0.001 // g/LSB
	AccResolution4g  = 1.95 * 0.001 // g/LSB
	AccResolution8g  = 3.91 * 0.001 // g/LSB
	AccResolution16g = 7.81 * 0.001 // g/LSB
)

type AcclResolution int

const (
	Resolution2G AcclResolution = iota
	Resolution4G
	Resolution8G
	Resolution16G
)

func (b *BMX055Driver) convertRawDataToMs2(val []int) []float64 {
	res := AccResolution2g
	switch b.config.resolution {
	case Resolution2G:
		res = AccResolution2g
	case Resolution4G:
		res = AccResolution4g
	case Resolution8G:
		res = AccResolution8g
	case Resolution16G:
		res = AccResolution16g
	default:
		panic("resolution")
	}

	var accl = make([]float64, 3)
	for i, g := range val {
		accl[i] = float64(g) * res * AccGravity
	}
	return accl
}
