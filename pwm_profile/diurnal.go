package pwm_profile

import (
	"encoding/json"
	"math"
	"time"
)

type diurnal struct {
	temporal
}

func Diurnal(conf json.RawMessage, min, max float64) (*diurnal, error) {
	t, err := Temporal(conf, min, max)
	if err != nil {
		return nil, err
	}
	d := diurnal{t}
	return &d, nil
}

func (d *diurnal) Get(t time.Time) float64 {
	if d.IsOutside(t) {
		return 0
	}
	totalMinutes := d.TotalMinutes()
	pastMinutes := d.PastMinutes(t)
	percent := float64(pastMinutes) * 2 * math.Pi / float64(totalMinutes)
	k := math.Pow(math.Cos(percent), 3)
	v := (1 - k) * d.ValueRange()
	v = v + float64(d.min)
	if v > float64(d.max) {
		v = float64(d.max)
	}
	return v
}
