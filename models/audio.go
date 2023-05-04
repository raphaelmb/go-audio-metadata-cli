package models

import (
	"bytes"
	"encoding/json"
)

type Audio struct {
	Id       string   `json:"id"`
	Path     string   `json:"path"`
	Metadata Metadata `json:"metadata"`
	Status   string   `json:"status"`
	Error    []string `json:"error"`
}

func (a *Audio) JSON() (string, error) {
	audioJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(audioJSON), "", "   "); err != nil {
		return "", nil
	}
	return prettyJSON.String(), nil
}
