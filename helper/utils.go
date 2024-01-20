package helper

import (
	"strconv"

	"github.com/a-fandy/finan/exception"
)

func ConvertStringToUint64(num string) uint64 {
	// Convert string to uint64
	uint64Value, err := strconv.ParseUint(num, 10, 64)
	if err != nil {
		exception.PanicIfError(err)
	}
	return uint64Value
}
