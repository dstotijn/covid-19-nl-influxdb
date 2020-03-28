package main

import (
	"context"
	"strconv"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

type influxDB struct {
	client *influxdb.Client
	bucket string
	org    string
}

type influxConfig struct {
	url    string
	token  string
	bucket string
	org    string
}

func newInfluxDB(cfg influxConfig) (*influxDB, error) {
	client, err := influxdb.New(cfg.url, cfg.token)
	if err != nil {
		return nil, err
	}

	return &influxDB{
		client: client,
		bucket: cfg.bucket,
		org:    cfg.org,
	}, nil
}

func (influxDB *influxDB) writeMetrics(ctx context.Context, date time.Time, caseReports []caseReport) (int, error) {
	var metrics []influxdb.Metric
	for _, caseReport := range caseReports {
		metrics = append(metrics, influxdb.NewRowMetric(
			map[string]interface{}{
				"confirmed": caseReport.Count,
				"citizens":  caseReport.CitizenCount,
			},
			"case_reports",
			map[string]string{
				"municipality":     caseReport.Municipality,
				"municipality_num": strconv.Itoa(caseReport.MunicipalityNum),
			},
			date,
		))
	}

	return influxDB.client.Write(ctx, influxDB.bucket, influxDB.org, metrics...)
}
