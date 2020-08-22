package earthquake

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
)

type Earthquake struct {
	Event         string
	Date          time.Time
	Lat           float64
	Long          float64
	Deep          int8
	Magnitude     float64
	MagnitudeType string
	Location      string
}

func GetEarthquakes(c *colly.Collector, from string) ([]Earthquake, error) {
	var earthquakes []Earthquake

	c.OnHTML("div#content tbody tr:not(:first-child)", func(e *colly.HTMLElement) {
		var earthquakeInfo []string
		e.ForEach("td", func(i int, element *colly.HTMLElement) {
			earthquakeInfo = append(earthquakeInfo, element.Text)
		})

		date, err := time.Parse("02/01/2006 15:04:05", earthquakeInfo[1]+" "+earthquakeInfo[2])
		if err != nil {
			panic(err)
		}

		lat, err := strconv.ParseFloat(strings.TrimSpace(earthquakeInfo[4]), 64)
		if err != nil {
			panic(err)
		}

		long, err := strconv.ParseFloat(strings.TrimSpace(earthquakeInfo[5]), 64)
		if err != nil {
			panic(err)
		}

		deep := 0
		deepString := strings.TrimSpace(earthquakeInfo[6])
		if len(deepString) > 0 {
			deep, err = strconv.Atoi(deepString)
			if err != nil {
				panic(err)
			}
		}

		magnitude, err := strconv.ParseFloat(strings.TrimSpace(earthquakeInfo[7]), 64)
		if err != nil {
			panic(err)
		}

		earthquakes = append(earthquakes, Earthquake{
			Event:         earthquakeInfo[0],
			Date:          date,
			Lat:           lat,
			Long:          long,
			Deep:          int8(deep),
			Magnitude:     magnitude,
			MagnitudeType: strings.TrimSpace(earthquakeInfo[8]),
			Location:      strings.TrimSpace(earthquakeInfo[10]),
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit(from)
	if err != nil {
		return nil, err
	}

	return earthquakes, nil
}
