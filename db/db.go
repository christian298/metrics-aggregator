package db

import (
	"fmt"
	"log"
	"time"

	"github.com/christian298/metrics-aggegator/models"
	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	dbName = "rum_data"
	dbURL  = "http://localhost:8086"
	dbUser = "user"
	dbPass = "useruser"
)

// Db client
type Db struct {
	client     client.Client
	batchPoint client.BatchPoints
}

// New DB connection
func New() (*Db, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: dbURL, Username: dbUser, Password: dbPass})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return nil, err
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "s",
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Db{client: c, batchPoint: bp}, nil
}

// Insert performance metrics into InfluxDB
func (db *Db) Insert(metrics models.Metric, browser models.Browser) {
	tags := map[string]string{"browser": browser.Name, "platform": browser.Platform, "version": browser.Version}

	//v := reflect.ValueOf(metrics)
	//typeOfS := v.Type()

	//for i := 0; i < v.NumField(); i++ {
	//	pt, err := client.NewPoint(typeOfS.Field(i).Name, tags, map[string]interface{}{"value": v.Field(i).Interface()}, time.Now())
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	db.batchPoint.AddPoint(pt)
	//}

	pt, err := client.NewPoint(metrics.Name, tags, map[string]interface{}{"value": metrics.Value}, time.Now())

	if err != nil {
		log.Fatal(err)
	}

	db.batchPoint.AddPoint(pt)

	// Write the batch
	if err := db.client.Write(db.batchPoint); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := db.client.Close(); err != nil {
		log.Fatal(err)
	}
}
