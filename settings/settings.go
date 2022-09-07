/*
 * Borsch Playground API
 *
 * Copyright (C) 2022 Yuriy Lisovskiy - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the MIT license.
 */

package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Settings struct {
	Debug               bool          `json:"debug"`
	ShutdownTimeoutSec  time.Duration `json:"shutdown_timeout_sec"`
	BorschVersions      []string      `json:"borsch_versions"`
	ApiDocumentationUrl string        `json:"api_documentation_url"`
	Runner              *Runner       `json:"runner"`
	Queue               *Queue        `json:"queue"`
	Database            *Database     `json:"database"`
}

func Load() (*Settings, error) {
	localSettings, err := loadFromJson("settings.local.json")
	if err == nil {
		return localSettings, nil
	}

	return loadFromJson("settings.json")
}

func loadFromJson(filename string) (*Settings, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var s Settings
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
