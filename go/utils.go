package main

import (
	"strings"
	"math/rand"
)

func StrExtract(total, start, end string) string {

	initArray := strings.Split(total, start)

	if len(initArray) <= 1 {
		return ""
	}

	member := initArray[1]
	endArray := strings.Split(member, end)

	if len(endArray) == 1 {
		return ""
	}

	return endArray[0]
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

