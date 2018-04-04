package main

import (
	"fmt"
	"strings"
	"time"
)

const lightOff = "O"
const redLight = "R"
const yellowLight = "Y"

func main() {
	ticker := time.NewTicker(1 * time.Second)

	for t := range ticker.C {
		handleClock(t)
	}
}

func handleClock(t time.Time) {
	fmt.Print("\033c") // HACK: clear screem
	fmt.Println(parseClock(t.Hour(), t.Minute(), t.Second()))
}

func parseClock(hours, minutes, seconds int) string {
	var sb strings.Builder

	fmt.Print(&sb, fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds))
	fmt.Println(&sb)

	if seconds%2 == 0 {
		fmt.Print(&sb, yellowLight)
	} else {
		fmt.Print(&sb, lightOff)
	}

	fmt.Println(&sb)
	fmt.Println(&sb, buildString(redLight, redLight, hours, 4, 4))
	fmt.Print(&sb, buildString(yellowLight, redLight, minutes, 11, 4))

	return sb.String()
}

func buildString(lightColor, quarterColor string, time, numberLightsTop, numberLightsBottom int) string {
	var sb strings.Builder
	groupSize := 5

	topLightsOn := time / groupSize

	for l := 1; l <= topLightsOn; l++ {
		if l%3 == 0 {
			fmt.Print(&sb, quarterColor)
		} else {
			fmt.Print(&sb, lightColor)
		}
	}

	topLightsOff := numberLightsTop - topLightsOn
	if topLightsOff > 0 {
		fmt.Print(&sb, strings.Repeat(lightOff, topLightsOff))
	}

	fmt.Println(&sb)

	bottomLightsOn := time % groupSize
	if bottomLightsOn > 0 {
		fmt.Print(&sb, strings.Repeat(lightColor, bottomLightsOn))
	}

	bottomLightsOff := numberLightsBottom - bottomLightsOn
	if bottomLightsOff > 0 {
		fmt.Print(&sb, strings.Repeat(lightOff, bottomLightsOff))
	}

	return sb.String()
}
