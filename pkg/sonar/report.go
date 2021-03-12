package sonar

import (
	"encoding/json"
	"os"
	"path/filepath"
)

//ReportData is representing the data of the step report JSON
type ReportData struct {
	ServerURL      string `json:"serverUrl"`
	ProjectKey     string `json:"projectKey"`
	NumberOfIssues Issues `json:"numberOfIssues"`
}

// Issues ...
type Issues struct {
	Blocker  int `json:"blocker"`
	Critical int `json:"critical"`
	Major    int `json:"major"`
	Minor    int `json:"minor"`
	Info     int `json:"info"`
}

// WriteReport ...
func WriteReport(data ReportData, reportPath string, reportFileName string, writeToFile func(f string, d []byte, p os.FileMode) error) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return writeToFile(filepath.Join(reportPath, reportFileName), jsonData, 0644)
}
