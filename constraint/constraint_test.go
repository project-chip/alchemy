package constraint

import (
	"strconv"
	"testing"

	"github.com/hasty/alchemy/matter"
)

var constraints []string = []string{
	"0 to 80000",
	"max (NumberOfEventsPerProgram * (1 + NumberOfLoadControlPrograms))",
	"max (MaxTemperature - 1)",
	"InstalledOpenLimitLift to InstalledClosedLimitLift",
	"0x00 to 0x3C",
	"-32767 to MaxScaledValue-1",
	"MaxScaledValue-1",
	"-10000 to +10000",
	"-127 to 127",
	"-2.5°C to 2.5°C",
	"-27315 to MaxMeasuredValue-1",
	"-32767 to MaxMeasuredValue-1",
	"-32767 to MaxScaledValue-1",
	"0 to 0x001F",
	"0 to 0xFEFF",
	"0 to 1",
	"0 to 100",
	"0 to 10000",
	"0 to 1000000",
	"0 to 1439",
	"0 to 1440",
	"0 to 15",
	"0 to 2",
	"0 to 2048",
	"0 to 254",
	"0 to 3",
	"0 to 31",
	"0 to 4",
	"0 to 5",
	"0 to 6",
	"0 to 65534",
	"0 to 7",
	"0 to MaxFrequency",
	"0 to MaxLevel",
	"0 to MaxMeasuredValue-1",
	"0 to NumberOfCredentialsSupportedPerUser",
	"0 to NumberOfPositions-1",
	"0 to OccupiedSetbackMax",
	"0 to SetTime",
	"0 to SpeedMax",
	"0 to UnoccupiedSetbackMax",
	"0% to 100%",
	"0, MinMeasuredValue to MaxMeasuredValue",
	"00000xxx",
	"0b0000 xxxx",
	"0b00xx xxxx",
	"0x00 to 0x0FE",
	"0x01 to 0xFF",
	"0x954D to 0x7FFF",
	"0°C to 2.5°C",
	"1 to 100",
	"1 to 2047",
	"1 to 254",
	"1 to 255",
	"1 to 2880",
	"1 to 65535",
	"1 to 8",
	"1 to MaxLevel",
	"1 to MaxMeasuredValue-1",
	"1 to MaxPower",
	"1 to MultiPressMax",
	"1 to NumberOfCredentialsSupportedPerUser",
	"1 to NumberOfTotalUsersSupported",
	"1 to NumberOfTotalUsersSupported, 0xFFFE",
	"1 to SupportedFabrics",
	"100 to MS",
	"16",
	"16 to 100",
	"16[2]",
	"2",
	"2 to 10",
	"2 to 255",
	"2 to MultiPressMax",
	"20",
	"32",
	"48",
	"500 to 2047",
	"65",
	"7,9",
	"Add, Modify",
	"ColorTempPhysicalMinMireds",
	"InstalledOpenLimitLift to InstalledClosedLimitLift",
	"InstalledOpenLimitTilt to InstalledClosedLimitTilt",
	"MS",
	"MinFrequency to MaxFrequency",
	"MinLevel to 254",
	"MinLevel to MaxLevel",
	"MinLevel to PhysicalMaxLevel",
	"MinMeasuredValue to MaxMeasuredValue",
	"MinMeasuredValue+1 to 10000",
	"MinMeasuredValue+1 to 32767",
	"MinMeasuredValue+1 to 65534",
	"MinPower to 100",
	"MinPower to MaxPower",
	"MinScaledValue to MaxScaledValue",
	"MinScaledValue+1 to 32767",
	"MinTemperature to MaxTemperature",
	"OccupiedEnabled, OccupiedDisabled",
	"OccupiedSetbackMin to 25.4°C",
	"OccupiedSetbackMin to OccupiedSetbackMax",
	"PhysicalMinLevel to MaxLevel",
	"TODO",
	"UnoccupiedSetbackMin to 25.4°C",
	"UnoccupiedSetbackMin to UnoccupiedSetbackMax",
	"UnrestrictedUser,",
	"Unspecified,",
	"all",
	"all[min 1]",
	"any",
	"desc",
	"max (MaxTemperature - 1)",
	"max (MaxTemperature - MinTemperature)",
	"max 0xFFFE",
	"max 10",
	"max 10 [max 50]",
	"max 100[max 1024]",
	"max 1024",
	"max 12",
	"max 120",
	"max 140",
	"max 1440",
	"max 16",
	"max 16[max 64]",
	"max 20",
	"max 253",
	"max 254",
	"max 255",
	"max 256",
	"max 259200",
	"max 3",
	"max 30",
	"max 32",
	"max 32 chars",
	"max 32[max 16]",
	"max 32[max 64]",
	"max 4",
	"max 5",
	"max 6",
	"max 6000",
	"max 60000",
	"max 604800",
	"max 64",
	"max 8",
	"max 8192",
	"max 900",
	"max 999",
	"max ClientTableSize",
	"max MaxMeasuredValue",
	"max NumberOfPositions-1",
	"min -27315",
	"min 0",
	"min 1",
	"min 10",
	"min 11",
	"min 2",
	"min 3",
	"min 5",
	"min 8",
	"min MinFrequency",
	"min MinMeasuredValue",
	"percent",
}

func TestConstraints(t *testing.T) {
	for _, s := range constraints {
		c := ParseConstraint(s)

		t.Logf("conformance: \"%s\" => \"%v\"", s, c.AsciiDocString())

		break
	}
}

func TestComplex(t *testing.T) {

	var fs matter.FieldSet
	fs = append(fs, &matter.Field{})
	for _, s := range constraints {
		c := ParseConstraint(s)

		t.Logf("conformance: \"%s\" => \"%v\" %T", s, c.AsciiDocString(), c)
		min, max := c.MinMax(&matter.ConstraintContext{Fields: fs})
		var from, to string
		switch min.Type {
		case matter.ConstraintExtremeTypeInt64:
			from = strconv.FormatInt(min.Int64, 10)
		case matter.ConstraintExtremeTypeUInt64:
			from = strconv.FormatUint(min.UInt64, 10)
		case matter.ConstraintExtremeTypeUndefined:
			from = "undefined"
		}
		switch max.Type {
		case matter.ConstraintExtremeTypeInt64:
			to = strconv.FormatInt(max.Int64, 10)
		case matter.ConstraintExtremeTypeUInt64:
			to = strconv.FormatUint(max.UInt64, 10)
		case matter.ConstraintExtremeTypeUndefined:
			to = "undefined"
		}
		t.Logf("\t: %s => %s", from, to)

		//break
	}
}
