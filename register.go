package bmx055

const (
	GYR_VAL_LPM1_MODE_NOMAL           = 0b00
	GYR_VAL_LPM1_MODE_DEEP_SUSPEND    = 0b01
	GYR_VAL_LPM1_MODE_SUSPEND         = 0b10
	ACC_VAL_RANGE_2                   = 0b0011
	ACC_VAL_RANGE_4                   = 0b0101
	ACC_VAL_RANGE_8                   = 0b1000
	ACC_VAL_RANGE_16                  = 0b1100
	ACC_VAL_PMU_BW_7_81               = 0b01000
	ACC_VAL_PMU_BW_15_63              = 0b01001
	ACC_VAL_PMU_BW_31_25              = 0b01010
	ACC_VAL_PMU_BW_62_5               = 0b01011
	ACC_VAL_PMU_BW_125                = 0b01100
	ACC_VAL_PMU_BW_250                = 0b01101
	ACC_VAL_PMU_BW_500                = 0b01110
	ACC_VAL_PMU_BW_1000               = 0b01111
	ACC_VAL_PMU_LPW_MODE_NOMAL        = 0b00000000
	ACC_VAL_PMU_LPW_MODE_DEEP_SUSPEND = 0b00100000
	ACC_VAL_PMU_LPW_MODE_LOW_POWER    = 0b01000000
	ACC_VAL_PMU_LPW_MODE_SUSPEND      = 0b10000000
	ACC_VAL_PMU_LPW_SLEEP_DUR_0_5MS   = 0b00000
	ACC_VAL_PMU_LPW_SLEEP_DUR_1MS     = 0b01100
	ACC_VAL_PMU_LPW_SLEEP_DUR_2MS     = 0b01110
	ACC_VAL_PMU_LPW_SLEEP_DUR_4MS     = 0b10000
	ACC_VAL_PMU_LPW_SLEEP_DUR_6MS     = 0b10010
	ACC_VAL_PMU_LPW_SLEEP_DUR_10MS    = 0b10100
	ACC_VAL_PMU_LPW_SLEEP_DUR_25MS    = 0b10110
	ACC_VAL_PMU_LPW_SLEEP_DUR_50MS    = 0b11000
	ACC_VAL_PMU_LPW_SLEEP_DUR_100MS   = 0b11010
	ACC_VAL_PMU_LPW_SLEEP_DUR_500MS   = 0b11100
	ACC_VAL_PMU_LPW_SLEEP_DUR_1S      = 0b11110
	GYR_VAL_RANGE_16_4                = 0b000
	GYR_VAL_RANGE_32_8                = 0b001
	GYR_VAL_RANGE_65_6                = 0b010
	GYR_VAL_RANGE_131_2               = 0b011
	GYR_VAL_RANGE_262_4               = 0b100
	GYR_VAL_BW_32                     = 0b0111
	GYR_VAL_BW_64                     = 0b0110
	GYR_VAL_BW_12                     = 0b0101
	GYR_VAL_BW_23                     = 0b0100
	GYR_VAL_BW_47                     = 0b0011
	GYR_VAL_BW_116                    = 0b0010
	GYR_VAL_BW_230                    = 0b0001
	GYR_VAL_BW_523                    = 0b0000
	GYR_VAL_LPM1_SLEEP_DUR_2MS        = 0b000
	GYR_VAL_LPM1_SLEEP_DUR_4MS        = 0b001
	GYR_VAL_LPM1_SLEEP_DUR_5MS        = 0b010
	GYR_VAL_LPM1_SLEEP_DUR_8MS        = 0b011
	GYR_VAL_LPM1_SLEEP_DUR_10MS       = 0b100
	GYR_VAL_LPM1_SLEEP_DUR_15MS       = 0b101
	GYR_VAL_LPM1_SLEEP_DUR_18MS       = 0b110
	GYR_VAL_LPM1_SLEEP_DUR_20MS       = 0b111
	MAG_AVAL_DV_SELF_TEST_NORMAL      = 0b00000000
	MAG_VAL_ADV_SELF_TEST_NEGATIVE    = 0b10000000
	MAG_VAL_ADV_SELF_TEST_POSITIVE    = 0b11000000
	MAG_VAL_DATA_RATE_10              = 0b000000
	MAG_VAL_DATA_RATE_2               = 0b001000
	MAG_VAL_DATA_RATE_6               = 0b010000
	MAG_VAL_DATA_RATE_8               = 0b011000
	MAG_VAL_DATA_RATE_15              = 0b100000
	MAG_VAL_DATA_RATE_20              = 0b101000
	MAG_VAL_DATA_RATE_25              = 0b110000
	MAG_VAL_DATA_RATE_30              = 0b111000
	MAG_VAL_OP_MODE_NORMAL            = 0b000
	MAG_VAL_OP_MODE_FORCED            = 0b010
	MAG_VAL_OP_MODE_SLEEP             = 0b110
	MAG_VAL_TEST_NORMAL               = 0b0
	MAG_VAL_TEST_SELF_TEST            = 0b1
	MAG_VAL_POW_CTL_SOFT_RESET        = 0b10000010
	MAG_VAL_POW_CTL_SLEEP_MODE        = 0b00000001
	MAG_VAL_POW_CTL_SUSPEND_MODE      = 0b00000000
)

const (
	ACC_REG_BGW_CHIPID    = 0x00
	ACC_REG_ACCD_X_LSB    = 0x02
	ACC_REG_ACCD_X_MSB    = 0x03
	ACC_REG_ACCD_Y_LSB    = 0x04
	ACC_REG_ACCD_Y_MSB    = 0x05
	ACC_REG_ACCD_Z_LSB    = 0x06
	ACC_REG_ACCD_Z_MSB    = 0x07
	ACC_REG_ACCD_TEMP     = 0x08
	ACC_REG_INT_STATUS_0  = 0x09
	ACC_REG_INT_STATUS_1  = 0x0A
	ACC_REG_INT_STATUS_2  = 0x0B
	ACC_REG_INT_STATUS_3  = 0x0C
	ACC_REG_FIFO_STATUS   = 0x0E
	ACC_REG_PMU_RANGE     = 0x0F
	ACC_REG_PMU_BW        = 0x10
	ACC_REG_PMU_LPW       = 0x11
	ACC_REG_PMU_LOW_POWER = 0x12
	ACC_REG_ACCD_HBW      = 0x13
	ACC_REG_BGW_SOFTRESET = 0x14
	ACC_REG_INT_EN_0      = 0x16
	ACC_REG_INT_EN_1      = 0x17
	ACC_REG_INT_EN_2      = 0x18
	ACC_REG_INT_MAP_0     = 0x19
	ACC_REG_INT_MAP_1     = 0x1A
	ACC_REG_INT_MAP_2     = 0x1B
	ACC_REG_INT_SRC       = 0x1E
	ACC_REG_INT_OUT_CTRL  = 0x20
	ACC_REG_INT_RST_LATCH = 0x21
	ACC_REG_INT_0         = 0x22
	ACC_REG_INT_1         = 0x23
	ACC_REG_INT_2         = 0x24
	ACC_REG_INT_3         = 0x25
	ACC_REG_INT_4         = 0x26
	ACC_REG_INT_5         = 0x27
	ACC_REG_INT_6         = 0x28
	ACC_REG_INT_7         = 0x29
	ACC_REG_INT_8         = 0x2A
	ACC_REG_INT_9         = 0x2B
	ACC_REG_INT_A         = 0x2C
	ACC_REG_INT_B         = 0x2D
	ACC_REG_INT_C         = 0x2E
	ACC_REG_INT_D         = 0x2F
	ACC_REG_FIFO_CONFIG_0 = 0x30
	ACC_REG_PMU_SELF_TEST = 0x32
	ACC_REG_TRIM_NVM_CTRL = 0x33
	ACC_REG_BGW_SPI3_WDT  = 0x34
	ACC_REG_OFC_CTRL      = 0x36
	ACC_REG_OFC_SETTING   = 0x37
	ACC_REG_OFC_OFFSET_X  = 0x38
	ACC_REG_OFC_OFFSET_Y  = 0x39
	ACC_REG_OFC_OFFSET_Z  = 0x3A
	ACC_REG_TRIM_GP0      = 0x3B
	ACC_REG_TRIM_GP1      = 0x3C
	ACC_REG_FIFO_CONFIG_1 = 0x3E
	ACC_REG_FIFO_DATA     = 0x3F
)

const (
	GYR_REG_CHIP_ID       = 0x00
	GYR_REG_RATE_X_LSB    = 0x02
	GYR_REG_RATE_X_MSB    = 0x03
	GYR_REG_RATE_Y_LSB    = 0x04
	GYR_REG_RATE_Y_MSB    = 0x05
	GYR_REG_RATE_Z_LSB    = 0x06
	GYR_REG_RATE_Z_MSB    = 0x07
	GYR_REG_INT_STATUS_0  = 0x09
	GYR_REG_INT_STATUS_1  = 0x0A
	GYR_REG_INT_STATUS_2  = 0x0B
	GYR_REG_INT_STATUS_3  = 0x0C
	GYR_REG_FIFO_STATUS   = 0x0E
	GYR_REG_RANGE         = 0x0F
	GYR_REG_BW            = 0x10
	GYR_REG_LPM1          = 0x11
	GYR_REG_LPM2          = 0x12
	GYR_REG_RATE_HBW      = 0x13
	GYR_REG_BGW_SOFTRESET = 0x14
	GYR_REG_INT_EN_0      = 0x15
	GYR_REG_INT_EN_1      = 0x16
	GYR_REG_INT_MAP_0     = 0x17
	GYR_REG_INT_MAP_1     = 0x18
	GYR_REG_INT_MAP_2     = 0x19
	GYR_REG_1A            = 0x1A
	GYR_REG_1B            = 0x1B
	GYR_REG_1C            = 0x1C
	GYR_REG_1E            = 0x1E
	GYR_REG_INT_RST_LATCH = 0x21
	GYR_REG_HIGH_TH_X     = 0x22
	GYR_REG_HIGH_DUR_X    = 0x23
	GYR_REG_HIGH_TH_Y     = 0x24
	GYR_REG_HIGH_DUR_Y    = 0x25
	GYR_REG_HIGH_TH_Z     = 0x26
	GYR_REG_HIGH_DUR_Z    = 0x27
	GYR_REG_SOC           = 0x31
	GYR_REG_A_FOC         = 0x32
	GYR_REG_TRIM_NVM_CTRL = 0x33
	GYR_REG_BGW_SPI3_WDT  = 0x34
	GYR_REG_OFC1          = 0x36
	GYR_REG_OFC2          = 0x37
	GYR_REG_OFC3          = 0x38
	GYR_REG_OFC4          = 0x39
	GYR_REG_TRIM_GP0      = 0x3A
	GYR_REG_TRIM_GP1      = 0x3B
	GYR_REG_BIST          = 0x3C
	GYR_REG_FIFO_CONFIG_0 = 0x3D
	GYR_REG_FIFO_CONFIG_1 = 0x3E
)

const (
	MAG_REG_CHIP_ID                  = 0x40
	MAG_REG_MAGNETIC_FIELD_DATA_42   = 0x42
	MAG_REG_MAGNETIC_FIELD_DATA_43   = 0x43
	MAG_REG_MAGNETIC_FIELD_DATA_44   = 0x44
	MAG_REG_MAGNETIC_FIELD_DATA_45   = 0x45
	MAG_REG_MAGNETIC_FIELD_DATA_46   = 0x46
	MAG_REG_MAGNETIC_FIELD_DATA_47   = 0x47
	MAG_REG_MAGNETIC_FIELD_DATA_48   = 0x48
	MAG_REG_MAGNETIC_FIELD_DATA_49   = 0x49
	MAG_REG_INTERRUPT_STATUS         = 0x4A
	MAG_REG_POWER_4B                 = 0x4B
	MAG_REG_POWER_4C                 = 0x4C
	MAG_REG_INTERRUPT_4D             = 0x4D
	MAG_REG_INTERRUPT_4E             = 0x4E
	MAG_REG_INTERRUPT_4F             = 0x4F
	MAG_REG_INTERRUPT_50             = 0x50
	MAG_REG_NUMBER_OF_REPETITIONS_51 = 0x51
	MAG_REG_NUMBER_OF_REPETITIONS_52 = 0x52
)
