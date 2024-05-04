package util

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func IsValidAge(numString string, operator string) int {
	numStr := strings.Replace(numString, operator, "", 1)

	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 {
		return -1
	}

	return num
}

func ParseAgeInMonth(ageInMonth string) int {
	switch {
	case strings.HasPrefix(ageInMonth, ">"):
		num := IsValidAge(ageInMonth, ">")
		return num
	case strings.HasPrefix(ageInMonth, "<"):
		num := IsValidAge(ageInMonth, "<")
		return num
	default:
		num := IsValidAge(ageInMonth, "=")
		return num
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
