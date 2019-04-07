# Go Feature Toggle
[![Build Status](https://travis-ci.org/Piszmog/gotoggle.svg?branch=develop)](https://travis-ci.org/Piszmog/gotoggle)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/gotoggle/badge.svg?branch=develop)](https://coveralls.io/github/Piszmog/gotoggle?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/gotoggle)](https://goreportcard.com/report/github.com/Piszmog/gotoggle)
[![GitHub release](https://img.shields.io/github/release/Piszmog/gotoggle.svg)](https://github.com/Piszmog/gotoggle/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is a simple package for using feature toggles in Go.

To use, simply initialize the features with a configuration or use `gotoggle.AddFeatures(..)` / `gotoggle.AddFeature(..)` 
to setup the features. Then use `gotoggle.IsEnabled(..)` to determine if a feature is enabled.

Determining if a feature is enabled is case insensitive. There is no different between "myFeature" and "myfeature". 

## Example
```go
package main

import (
	"encoding/json"
	"github.com/Piszmog/gotoggle"
)

func main() {
	configJSON := `{"features":{"feature1":true,"feature2":false}}`
	var config gotoggle.Configuration
	err := json.Unmarshal([]byte(configJSON), &config)
	if err != nil {
		// do something
	}
	gotoggle.InitializeFeatures(config)
	// ...
	if gotoggle.IsEnabled("myFeature") {
		// do something
	}
}
```

### Usage
Initializing or adding features should be done once. Under the hood, the feature map is not safe to update in many 
different goroutines. You can access the features from different goroutines.