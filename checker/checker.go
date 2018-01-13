package checker

import (
	"math"
	"regexp"
	"strconv"

	"github.com/binatify/whoami/types"
)

var (
	complexRegexp  = regexp.MustCompile("complex")
	boolRegexp     = regexp.MustCompile("^true|false$")
	numberRegexp   = regexp.MustCompile("^[0-9|\\-|\\.]{0,}$")
	floatRegexp    = regexp.MustCompile("\\.")
	negativeRegexp = regexp.MustCompile("^\\-")
)

func Assert(input string) {
	switch {
	case complexRegexp.MatchString(input):
		types.Print("complex64", "complex128")
	case boolRegexp.MatchString(input):
		types.Print("bool")
	case numberRegexp.MatchString(input):
		if floatRegexp.MatchString(input) {
			inputValue, _ := strconv.ParseFloat(input, 64)
			switch {
			case (inputValue > math.MaxFloat32) ||
				(inputValue < -math.MaxFloat32):
				types.Print("float64")
			default:
				types.Print("float32", "float64")
			}
		} else {
			switch {
			case negativeRegexp.MatchString(input):
				inputValue, _ := strconv.ParseInt(input, 10, 64)
				switch {
				case inputValue < -math.MaxInt32:
					types.Print("int64")
				case inputValue < -math.MaxInt16:
					types.Print("int32", "int64")
				case inputValue < -math.MaxInt8:
					types.Print("int16", "int32", "int64")
				default:
					types.Print("int8", "int16", "int16", "int32", "int64")
				}
			default:
				inputValue, _ := strconv.ParseUint(input, 10, 64)
				switch {
				case inputValue > math.MaxInt64:
					types.Print("uint64")
				case inputValue > math.MaxUint32:
					types.Print("int64", "uint64")
				case inputValue > math.MaxInt32:
					types.Print("int64", "", "uint32", "uint64")
				case inputValue > math.MaxUint16:
					types.Print("int32", "int64", "", "uint32", "uint64", "", "rune")
				case inputValue > math.MaxInt16:
					types.Print("int32", "int64", "", "uint16", "uint32", "uint64", "", "rune")
				case inputValue > math.MaxUint8:
					types.Print("int16", "int32", "int64", "", "uint16", "uint32", "uint64", "", "rune")
				case inputValue > math.MaxInt8:
					types.Print("int16", "int32", "int64", "", "uint8", "uint16", "uint32", "uint64", "", "byte", "rune")

				default:
					types.Print("int8", "int16", "int32", "int64", "", "uint8", "uint16", "uint32", "uint64", "", "byte", "rune")

				}
			}
		}
	default:
		types.Print("string")
	}
}
