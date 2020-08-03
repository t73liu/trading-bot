package utils

import "strings"

func TrimCompanyName(company string) string {
	trimmedName := strings.TrimSpace(strings.Split(company, " Class ")[0])
	//trimmedName = strings.TrimSpace(strings.Split(company, " Inc.")[0])
	return trimmedName
}
