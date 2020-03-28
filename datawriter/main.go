package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	dataProvider := newProvider()

	influxDB, err := newInfluxDB(influxConfig{
		url:    mustGetenv("INFLUXDB_URL"),
		token:  mustGetenv("INFLUXDB_TOKEN"),
		bucket: mustGetenv("INFLUXDB_BUCKET"),
		org:    mustGetenv("INFLUXDB_ORG"),
	})
	if err != nil {
		log.Fatalf("[ERROR] Cannot create InfluxDB client: %v", err)
	}

	for c := time.Tick(15 * time.Minute); ; <-c {
		// Get latest case history, for all dates for all municipalities.
		casesHistory, err := dataProvider.getCasesHistory()
		if err != nil {
			log.Fatalf("[ERROR] Getting case history failed: %v", err)
		}

		// Record metrics of daily case history in InfluxDB.
		for date, caseReports := range casesHistory {
			writeCount, err := influxDB.writeMetrics(ctx, date, caseReports)
			if err != nil {
				log.Fatalf("[ERROR] Cannot write metrics to InfluxDB: %v", err)
			}

			log.Printf("[INFO] Wrote %v case report metric(s) (%v) to InfluxDB.", writeCount, date.Format("2006-01-02"))
		}
	}
}

func mustGetenv(env string) string {
	token := os.Getenv(env)
	if token == "" {
		log.Fatalf("[FATAL] Environment variable `%v` is required.", env)
	}

	return token
}
