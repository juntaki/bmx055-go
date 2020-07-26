package bmx055

import (
	"gobot.io/x/gobot/drivers/i2c"
	"time"
)

type Config struct {
	resolution AcclResolution
	AddrAccl   int
	AddrGyro   int
	AddrMag    int
}

type BMX055Driver struct {
	acclCon i2c.Connection
	gyroCon i2c.Connection
	magCon  i2c.Connection
	config  *Config
}

func NewBMX055Driver(conn i2c.Connector, options ...func(*Config)) (*BMX055Driver, error) {
	config := &Config{
		resolution: Resolution2G,
		AddrAccl:   0x19,
		AddrGyro:   0x69,
		AddrMag:    0x13,
	}
	for _, option := range options {
		option(config)
	}

	acclCon, err := conn.GetConnection(config.AddrAccl, 1)
	if err != nil {
		return nil, err
	}
	gyroCon, err := conn.GetConnection(config.AddrGyro, 1)
	if err != nil {
		return nil, err
	}
	magCon, err := conn.GetConnection(config.AddrMag, 1)
	if err != nil {
		return nil, err
	}

	return &BMX055Driver{
		acclCon: acclCon,
		gyroCon: gyroCon,
		magCon:  magCon,
		config:  config,
	}, nil
}

func (b *BMX055Driver) Start() {
	var rangeReg uint8 = ACC_VAL_RANGE_2
	switch b.config.resolution {
	case Resolution2G:
		rangeReg = ACC_VAL_RANGE_2
	case Resolution4G:
		rangeReg = ACC_VAL_RANGE_4
	case Resolution8G:
		rangeReg = ACC_VAL_RANGE_8
	case Resolution16G:
		rangeReg = ACC_VAL_RANGE_16
	default:
		panic("resolution")
	}

	// SOFTRESET
	b.acclCon.WriteByteData(ACC_REG_BGW_SOFTRESET, 0xB6)
	time.Sleep(2 * time.Millisecond)
	b.gyroCon.WriteByteData(GYR_REG_BGW_SOFTRESET, 0xB6)
	time.Sleep(2 * time.Millisecond)
	b.magCon.WriteByteData(MAG_REG_POWER_4B, 0b10000010)
	time.Sleep(2 * time.Millisecond)

	// accel
	b.acclCon.WriteByteData(ACC_REG_PMU_RANGE, rangeReg)
	//TODO: fixed config
	b.acclCon.WriteByteData(ACC_REG_PMU_BW, ACC_VAL_PMU_BW_7_81)
	b.acclCon.WriteByteData(ACC_REG_PMU_LPW, ACC_VAL_PMU_LPW_MODE_NOMAL|ACC_VAL_PMU_LPW_SLEEP_DUR_0_5MS)

	// gyro
	b.gyroCon.WriteByteData(GYR_REG_RANGE, GYR_VAL_RANGE_262_4)
	b.gyroCon.WriteByteData(GYR_REG_BW, GYR_VAL_BW_32)
	b.gyroCon.WriteByteData(GYR_REG_LPM1, GYR_VAL_LPM1_MODE_NOMAL|GYR_VAL_LPM1_SLEEP_DUR_2MS)

	// mag
	b.magCon.WriteByteData(MAG_REG_POWER_4B, MAG_VAL_POW_CTL_SLEEP_MODE)
	time.Sleep(100 * time.Millisecond)

	b.magCon.WriteByteData(MAG_REG_POWER_4C, MAG_VAL_DATA_RATE_10)
	b.magCon.WriteByteData(MAG_REG_INTERRUPT_4E, 0x84)
	b.magCon.WriteByteData(MAG_REG_NUMBER_OF_REPETITIONS_51, 0x04)
	b.magCon.WriteByteData(MAG_REG_NUMBER_OF_REPETITIONS_52, 0x15)
}
