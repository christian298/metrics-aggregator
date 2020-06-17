package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/christian298/metrics-aggegator/models"
)

func getRandomBrowser(browsers *[]models.Browser) models.Browser {
	rand.Seed(time.Now().UnixNano())
	return (*browsers)[rand.Intn(len(*browsers))]
}

func getRandomMetric(metrics *[]string) models.Metric {
	rand.Seed(time.Now().UnixNano())

	m := (*metrics)[rand.Intn(len(*metrics))]

	if m == "CLS" {
		return models.Metric{Id: "", Name: m, Value: (0.0 + rand.Float32()*(0.5))}
	}

	if m == "LCP" {
		return models.Metric{Id: "", Name: m, Value: (50 + rand.Float32()*(3000-50))}
	}

	if m == "FID" {
		return models.Metric{Id: "", Name: m, Value: (10 + rand.Float32()*(500-10))}
	}

	if m == "FCP" {
		return models.Metric{Id: "", Name: m, Value: (50 + rand.Float32()*(1200-50))}
	}

	return models.Metric{Id: "", Name: m, Value: (50 + rand.Float32()*(500-50))}

}

func main() {

	metrics := []string{
		"TTFB",
		"CLS",
		"LCP",
		"FCP",
		"FID",
	}

	browsers := []models.Browser{
		{Name: "Firefox", Platform: "Windows", Type: "Desktop", Version: "76.0"},
		{Name: "Firefox", Platform: "Windows", Type: "Desktop", Version: "75.0"},
		{Name: "Firefox", Platform: "Mac", Type: "Desktop", Version: "76.0"},
		{Name: "Firefox", Platform: "Android", Type: "Mobile", Version: "76.0"},
		{Name: "Chrome", Platform: "Android", Type: "Mobile", Version: "83.0"},
		{Name: "Chrome", Platform: "Android", Type: "Mobile", Version: "82.0"},
		{Name: "Chrome", Platform: "Windows", Type: "Desktop", Version: "83.0"},
		{Name: "Chrome", Platform: "Mac", Type: "Desktop", Version: "82.0"},
		{Name: "Safari", Platform: "Mac", Type: "Desktop", Version: "13.0.0"},
		{Name: "Safari", Platform: "Mac", Type: "Desktop", Version: "13.1.0"},
		{Name: "Safari", Platform: "iOS", Type: "Mobile", Version: "13.1.0"},
		{Name: "Safari", Platform: "iOS", Type: "Mobile", Version: "13.0.0"},
		{Name: "Edge", Platform: "Windows", Type: "Desktop", Version: "83.0"},
		{Name: "Edge", Platform: "Windows", Type: "Desktop", Version: "16.0"},
		{Name: "Edge", Platform: "Mac", Type: "Desktop", Version: "83.0"},
	}

	for i := 0; i < 50; i++ {
		t := time.Now()
		m := getRandomMetric(&metrics)
		b := getRandomBrowser(&browsers)

		payload := fmt.Sprintf("%v,browser=%v,platform=%v,version=%v,type=%v value=%v %v", m.Name, b.Name, b.Platform, b.Version, b.Type, m.Value, t.Add(time.Minute).UnixNano())

		resp, err := http.Post("http://localhost:8086/api/v2/write?bucket=rum_data&org=&precision=ns", "text/plain", bytes.NewBufferString(payload))

		if err == nil {
			fmt.Print(err)
		}

		defer resp.Body.Close()

	}

}
