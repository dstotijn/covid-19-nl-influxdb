package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://kapulara.github.io/COVID-19-NL/"

type provider struct {
	httpClient *http.Client
}

func newProvider() provider {
	return provider{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type caseReport struct {
	Count           int    `json:"Aantal"`
	CitizenCount    int    `json:"BevAant"`
	Municipality    string `json:"Gemeente"`
	MunicipalityNum int    `json:"Gemnr"`
}

type cases map[time.Time][]caseReport

func (p provider) getFilePaths() ([]string, error) {
	resp, err := p.httpClient.Get(baseURL + "Municipalities/json/files.json")
	if err != nil {
		return nil, fmt.Errorf("cannot execute HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var fileIndex []string
	err = json.NewDecoder(resp.Body).Decode(&fileIndex)
	if err != nil {
		return nil, fmt.Errorf("cannot parse HTTP response body: %v", err)
	}

	return fileIndex, nil
}

func (p provider) getCasesHistory() (cases, error) {
	filePaths, err := p.getFilePaths()
	if err != nil {
		return nil, fmt.Errorf("cannot fetch files index: %v", err)
	}

	casesHistory := make(cases, len(filePaths))

	for _, filePath := range filePaths {
		caseReports, err := p.getCaseReports(filePath)
		if err != nil {
			return nil, fmt.Errorf("cannot get cases (%v): %v", filePath, err)
		}
		date, err := time.Parse("01-02-2006", filePath[:10])
		if err != nil {
			return nil, fmt.Errorf("cannot parse file date (%v): %v", filePath, err)
		}

		casesHistory[date] = caseReports
	}

	return casesHistory, nil
}

func (p provider) getCaseReports(filePath string) ([]caseReport, error) {
	resp, err := p.httpClient.Get(baseURL + "Municipalities/json/" + filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot execute HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var caseReports []caseReport
	err = json.NewDecoder(resp.Body).Decode(&caseReports)
	if err != nil {
		return nil, fmt.Errorf("cannot parse HTTP response body: %v", err)
	}

	return caseReports, nil
}
