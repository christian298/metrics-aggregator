package db

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/api"
	"time"

	"github.com/christian298/metrics-aggegator/config"
	"github.com/christian298/metrics-aggegator/models"
	"github.com/influxdata/influxdb-client-go"
)

// Db client
type Db struct {
	api    api.WriteApi
	Client influxdb2.Client
}

// New DB connection
func New(config *config.Config) (*Db, error) {
	c := influxdb2.NewClient(config.Db.URL, fmt.Sprintf("%s:%s", config.Db.User, config.Db.Password))
	writeApi := c.WriteApi("", config.Db.Name)

	return &Db{api: writeApi, Client: c}, nil
}

// Insert performance metrics into InfluxDB
func (db *Db) Insert(metrics models.Metric, browser models.Browser) {
	tags := map[string]string{"browser": browser.Name, "platform": browser.Platform, "version": browser.Version, "type": browser.Type}

	p := influxdb2.NewPoint(metrics.Name, tags, map[string]interface{}{"value": metrics.Value}, time.Now())
	db.api.WritePoint(p)

	db.api.Flush()
}
