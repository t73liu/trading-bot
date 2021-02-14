package utils

import (
	"fmt"
	"strings"
)

func CreatePositionalParams(startingPosition, numberOfParams int) string {
	var positionParams string
	for i := 0; i < numberOfParams; i++ {
		positionParams += fmt.Sprintf("$%d,", i+startingPosition)
	}
	return fmt.Sprintf("(%s),", strings.TrimSuffix(positionParams, ","))
}
