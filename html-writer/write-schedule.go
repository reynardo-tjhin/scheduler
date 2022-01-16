package htmlwriter

import (
	"scheduler/schedule"
)

func WriteHTML(s schedule.Schedule) string {

	var result string = "<table>\n"

	var number_of_days int = len(s.Days)
	result += "\t<tr>"
	result += "\t<th colspan=\"3\" width=\"5%\">Time</th>"

	// Default width: "Time": 5%, Other days: 15%
	i := 0
	for i < number_of_days {
		result += "\t\t<th> colspan=\"2\" id=\"" + s.Days[i] + "\" width=\"15%\">" + s.Days[i] + "</th>"
		i++
	}
	result += "\t<\tr>"

	result += "</table>"

	return result
}
