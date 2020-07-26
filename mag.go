package bmx055

func (b *BMX055Driver) GetMagData() ([]float64, error) {
	raw, err := b.GetMagRawData()
	if err !=nil {
		return nil, err
	}

	var val = make([]float64, 3)
	for i, r := range raw {
		val[i] = float64(r)
	}
	return val, nil
}

func (b *BMX055Driver) GetMagRawData() ([]int, error) {
	var err error
	data := make([]uint8, 6)

	data[0], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_42)
	if err != nil {
		return nil, err
	}
	data[1], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_43)
	if err != nil {
		return nil, err
	}
	data[2], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_44)
	if err != nil {
		return nil, err
	}
	data[3], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_45)
	if err != nil {
		return nil, err
	}
	data[4], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_46)
	if err != nil {
		return nil, err
	}
	data[5], err = b.magCon.ReadByteData(MAG_REG_MAGNETIC_FIELD_DATA_47)
	if err != nil {
		return nil, err
	}

	castData := make([]int, 6)

	for i := range data {
		castData[i] = int(data[i])
	}

	xMag := (castData[1] << 8) | (castData[0] >> 3)
	if xMag > 4095 {
		xMag -= 8192
	}
	yMag := (castData[3] << 8) | (castData[2] >> 3)
	if yMag > 4095 {
		yMag -= 8192
	}
	zMag := (castData[5] << 8) | (castData[4] >> 3)
	if zMag > 16383 {
		zMag -= 32768
	}

	return []int{xMag, yMag, zMag}, nil
}
