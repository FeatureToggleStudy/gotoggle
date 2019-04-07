package gotoggle

import (
	"encoding/json"
	"testing"
)

func resetFeatures() {
	features = make(map[string]bool)
}

func TestInitializeFeatures(t *testing.T) {
	resetFeatures()
	configJSON := `{"features":{"feature1":true,"feature2":false}}`
	var config Configuration
	err := json.Unmarshal([]byte(configJSON), &config)
	if err != nil {
		t.Error("failed to unmarshal config JSON")
	}
	InitializeFeatures(config)
	if !IsEnabled("feature1") {
		t.Error("feature1 is disabled. It should be enabled")
	}
	if IsEnabled("feature2") {
		t.Error("feature2 is enabled. It should be disabled")
	}
}

func TestAddFeatures(t *testing.T) {
	resetFeatures()
	newFeatures := make(map[string]bool)
	newFeatures["newFeature1"] = true
	newFeatures["newFeature2"] = false
	AddFeatures(newFeatures)
	if !IsEnabled("newFeature1") {
		t.Error("newFeature1 is disabled. It should be enabled")
	}
	if IsEnabled("newFeature2") {
		t.Error("newFeature2 is enabled. It should be disabled")
	}
}

func TestAddFeature(t *testing.T) {
	resetFeatures()
	AddFeature("newFeature3", true)
	if !IsEnabled("newFeature3") {
		t.Error("newFeature3 is disabled. It should be enabled")
	}
}

func TestIsEnabled(t *testing.T) {
	resetFeatures()
	AddFeature("newFeature4", true)
	if !IsEnabled("newfeature4") {
		t.Error("newFeature4 is disabled. It should be enabled")
	}
}

func TestIsEnabled_notAdded(t *testing.T) {
	resetFeatures()
	if IsEnabled("newfeature5") {
		t.Error("newFeature5 is enabled. It should be disabled")
	}
}
