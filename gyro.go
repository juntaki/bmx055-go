package bmx055

import "math"

func (b *BMX055Driver) GetGyroData() ([]float64, error) {
	raw, err := b.GetGyroRawData()
	if err !=nil {
		return nil, err
	}

	var val = make([]float64, 3)
	for i, r := range raw {
		val[i] = float64(r) * 125.0 / 32767.0 * (math.Pi /180.0)
	}
	return val, nil
}

func (b *BMX055Driver) GetGyroRawData() ([]int, error) {
	var err error
	data := make([]uint8, 6)

	data[0], err = b.gyroCon.ReadByteData(GYR_REG_RATE_X_LSB)
	if err != nil {
		return nil, err
	}
	data[1], err = b.gyroCon.ReadByteData(GYR_REG_RATE_X_MSB)
	if err != nil {
		return nil, err
	}
	data[2], err = b.gyroCon.ReadByteData(GYR_REG_RATE_Y_LSB)
	if err != nil {
		return nil, err
	}
	data[3], err = b.gyroCon.ReadByteData(GYR_REG_RATE_Y_MSB)
	if err != nil {
		return nil, err
	}
	data[4], err = b.gyroCon.ReadByteData(GYR_REG_RATE_Z_LSB)
	if err != nil {
		return nil, err
	}
	data[5], err = b.gyroCon.ReadByteData(GYR_REG_RATE_Z_MSB)
	if err != nil {
		return nil, err
	}

	castData := make([]int, 6)

	for i := range data {
		castData[i] = int(data[i])
	}

	xGyro := (castData[1] * 256) + castData[0]
	if xGyro > 32767 {
		xGyro -= 65536
	}
	yGyro := (castData[3] * 256) + castData[2]
	if yGyro > 32767 {
		yGyro -= 65536
	}
	zGyro := (castData[5] * 256) + castData[4]
	if zGyro > 32767 {
		zGyro -= 65536
	}

	return []int{xGyro, yGyro, zGyro}, nil
}
