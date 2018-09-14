package device

import (
	"fmt"
	"strconv"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
	"github.com/jrmycanady/slurp-rtl_433/config"
	"github.com/jrmycanady/slurp-rtl_433/logger"
)

var (
	// FineOffsetWH24Name is the name that is used when storing into influxdb.
	FineOffsetWH24Name = "FineOffsetWH24"

	// FineOffsetWH24ModelName is the model name rtl_433 returns.
	FineOffsetWH24ModelName = "Fine Offset WH24"
)

// FineOffsetWH24DataPoint represents a datapoint from an FineOffsetWH24 device.
type FineOffsetWH24DataPoint struct {
	Model   string `json:"model"`
	TimeStr string `json:"time"`
	Time    time.Time
	Battery	string `json:"battery"`
	ID      int     `json:"id"`
	Rain    float64 `json:"rainfall_mm"`
	TemperatureC float64 `json:"temperature_C"`
	Humidity int `json:"humidity"`
	WindDirection	int `json:"wind_dir_deg"`
	WindSpeed	float64 `json:"wind_speed_ms"`
	WindGust	float64 `json:"gust_speed_ms"`
	UV	int `json:"uv"`
	UVI	int `json:"uvi"`
	LightLux	float64 `json:"light_lux"`
}

// GetTimeStr returns the string format of the time as provided by the device
// output.
func (a *FineOffsetWH24DataPoint) GetTimeStr() string {
	return a.TimeStr
}

// SetTime sets the time value fo the FineOffsetWH24DataPoint.
func (a *FineOffsetWH24DataPoint) SetTime(t time.Time) {
	a.Time = t
}

// InfluxData builds a new InfluxDB datapoint from the values in the DataPoint.
func (a *FineOffsetWH24DataPoint) InfluxData(sets map[string]config.MetaDataFieldSet) (*influx.Point, error) {
	tags := map[string]string{
		"model": a.Model,
		"id":    strconv.Itoa(a.ID),
	}
	// Parsing any metadata for this type if we have some.
	for _, set := range sets {
		logger.Debug.Printf("processing metadata set %s", set)
		ProcessMetaDataFieldSet(tags, &set)

	}

	fields := map[string]interface{}{
		"rainfall_mm": a.Rain,
		"humidity": a.Humidity,
		"wind_dir_deg": a.WindDirection,
		"wind_speed_ms": a.WindSpeed,
		"gust_speed_ms": a.WindGust,
		"uv": a.UV,
		"uvi": a.UVI,
		"light_lux": a.LightLux,
	}

	ParseTime(a)
	p, err := influx.NewPoint(FineOffsetWH24Name, tags, fields, a.Time)
	if err != nil {
		return nil, fmt.Errorf("failed to create point: %s", err)
	}

	return p, nil
}
