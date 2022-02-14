package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Time struct {
	Start string
	End   string
	Time  float64
}

func (t Time) GetVals() []interface{} {
	return []interface{}{t.Start, t.End, t.Time}
}

func getTimesFromCSV() (timeArr []Time, err error) {
	reportFile, err := os.Open("report.csv")
	if err != nil {
		return timeArr, err
	}
	defer reportFile.Close()

	csvReader := csv.NewReader(reportFile)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return timeArr, err
		}

		var times []float64

		if len(rec) != 2 {
			return timeArr, errors.New("Wrong line length")
		}

		for idx, time := range rec {
			rec[idx] = strings.TrimSpace(time)
			var sTimes = strings.Split(rec[idx], ":")
			hours, err := strconv.ParseFloat(sTimes[0], 4)
			if err != nil {
				return timeArr, err
			}
			minutes, err := strconv.ParseFloat(sTimes[1], 4)
			if err != nil {
				return timeArr, err
			}
			times = append(times, hours+minutes/60)
		}
		var addTime = 0
		if times[0] > times[1] {
			addTime = 24
		}
		timeArr = append(timeArr, Time{
			rec[0],
			rec[1],
			times[1] + float64(addTime) - times[0],
		})
	}
	return
}

func convertToMD(arr []Time) (md string) {
	for _, time := range arr {
		md += fmt.Sprintf("| %5s | %5s | %2.2f |       |\n", time.GetVals()...)
	}
	return
}

func getHeader() string {
	return `| Start | End   | Time | Total |
| ----- | ----- | ---- | ----- |
`
}

func getFooter(timeArr []Time) string {
	var total float64
	for _, time := range timeArr {
		total += time.Time
	}
	return fmt.Sprintf("|       |       |      | %2.2f |", total)
}

func main() {
	timeArr, err := getTimesFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	md := convertToMD(timeArr)
	header := getHeader()
	footer := getFooter(timeArr)
	fmt.Println(header + md + footer)
}
