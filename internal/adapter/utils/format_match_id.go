package utils

import (
	"fmt"
)

func FormatMatchID(source interface{}) string {
	return fmt.Sprintf("game_%v", source)
}
