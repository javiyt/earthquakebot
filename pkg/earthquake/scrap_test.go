package earthquake_test

import (
	"github.com/gocolly/colly/v2"
	"github.com/javiyt/earthquakebot/pkg/earthquake"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestGetEarthquakes(t *testing.T) {
	transport := &http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir("testdata/")))

	c := colly.NewCollector()
	c.WithTransport(transport)

	earthquakes, err := earthquake.GetEarthquakes(c,"file://testdata/earth.html")

	require.NoError(t, err)
	require.Len(t, earthquakes, 59)
	require.Contains(t, earthquakes, earthquake.Earthquake{
		Event: "es2020lwbcz",
		Date: time.Date(2020, 6, 17, 13, 34, 55, 0, time.UTC),
		Lat: 43.1085,
		Long: -2.225,
		Deep: 15,
		Magnitude: 3.2,
		MagnitudeType: "mbLg",
		Location: "SW BEIZAMA_SS",
	})
}
