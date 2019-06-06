package utils

import (
	"regexp"
)

func CleanupString(data string) string {

	reg, err := regexp.Compile(`[^a-zA-Z0-9\.]+`)
	if err != nil {
		return ""
	}

	return reg.ReplaceAllString(data, "")

}
