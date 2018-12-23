// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/nbio/st"
	"os"
	"testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.yml")

	err := config.Load(file.NewSource(
		file.WithPath(configFile),
	))

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", config.Get("app", "port").String("8080"))
}

// TestHealthStatus test cases
func TestHealthStatus(t *testing.T) {
	ok, err := HealthStatus()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
}
