package ztm_go_sdk

import (
	"encoding/json"
	"io"
	"os"
)

func LoadPermit(permfile string) (*Permit, error) {
	if _, statErr := os.Stat(permfile); statErr != nil {
		return nil, statErr
	}
	file, fileErr := os.Open(permfile)
	if fileErr != nil {
		return nil, fileErr
	}
	defer file.Close()

	bytes, readErr := io.ReadAll(file)
	if readErr != nil {
		return nil, readErr
	}
	perm := new(Permit)
	if err := json.Unmarshal(bytes, perm); err != nil {
		return nil, err
	}
	return perm, nil
}
