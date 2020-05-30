package db

import (
	"fmt"
	"log"
	"time"

	"github.com/christian298/metrics-aggegator/config"
	"github.com/christian298/metrics-aggegator/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

// Db client
type Db struct {
	client     client.Client
	batchPoint client.BatchPoints
}

// New DB connection
func New(config *config.Config) (*Db, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: config.Db.URL, Username: config.Db.User, Password: config.Db.Password})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return nil, err
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Db.Name,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Db{client: c, batchPoint: bp}, nil
}

// Insert performance metrics into InfluxDB
func (db *Db) Insert(metrics models.Metric, browser models.Browser) {
	tags := map[string]string{"browser": browser.Name, "platform": browser.Platform, "version": browser.Version, "type": browser.Type}

	pt, err := client.NewPoint(metrics.Name, tags, map[string]interface{}{"value": metrics.Value}, time.Now())

	if err != nil {
		fmt.Println(err)
	}

	db.batchPoint.AddPoint(pt)

	// Write the batch
	if err := db.client.Write(db.batchPoint); err != nil {
		fmt.Println(err)
	}

	// Close client resources
	if err := db.client.Close(); err != nil {
		fmt.Println(err)
	}
}
