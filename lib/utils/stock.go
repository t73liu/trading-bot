package utils

import "strings"

func TrimCompanyName(company string) string {
	trimmedName := strings.TrimSpace(strings.Split(company, " Class ")[0])
	trimmedName = strings.TrimSpace(strings.Split(company, " Series ")[0])
	trimmedName = strings.TrimSpace(strings.Split(company, " Inc. ")[0])
	trimmedName = strings.TrimSpace(strings.Split(company, " Inc ")[0])
	trimmedName = strings.TrimSpace(strings.Split(company, " Ltd. ")[0])
	trimmedName = strings.TrimSpace(strings.Split(company, " Ltd ")[0])
	return trimmedName
}
