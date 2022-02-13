package main

import (
	"encoding/csv"
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

func getTimesFromCSV() (timeArr []Time) {
	reportFile, err := os.Open("report.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer reportFile.Close()

	csvReader := csv.NewReader(reportFile)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var times []float64

		for idx, time := range rec {
			rec[idx] = strings.TrimSpace(time)
			var sTimes = strings.Split(rec[idx], ":")
			hours, _ := strconv.ParseFloat(sTimes[0], 4)
			minutes, _ := strconv.ParseFloat(sTimes[1], 4)
			times = append(times, (hours)+(minutes/60))
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
	timeArr := getTimesFromCSV()
	fmt.Println(timeArr)
	md := convertToMD(timeArr)
	header := getHeader()
	footer := getFooter(timeArr)
	fmt.Println(header + md + footer)
}
