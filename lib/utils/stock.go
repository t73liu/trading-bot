package utils

import "strings"

func TrimCompanyName(company string) string {
	trimmedName := strings.TrimSpace(strings.Split(company, " Class ")[0])
	trimmedName = strings.TrimSpace(strings.Split(trimmedName, " Series ")[0])
	trimmedName = strings.TrimSpace(strings.Split(trimmedName, " Inc. ")[0])
	trimmedName = strings.TrimSpace(strings.Split(trimmedName, " Inc ")[0])
	trimmedName = strings.TrimSpace(strings.Split(trimmedName, " Ltd. ")[0])
	trimmedName = strings.TrimSpace(strings.Split(trimmedName, " Ltd ")[0])
	return trimmedName
}
