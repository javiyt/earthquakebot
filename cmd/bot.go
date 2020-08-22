package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/javiyt/earthquakebot/pkg/earthquake"
)

func main() {
	c := colly.NewCollector()

	_, err := earthquake.GetEarthquakes(c,"https://www.ign.es/web/ign/portal/ultimos-terremotos")
	if err != nil {
		panic(err)
	}
}
