package gotoggle

import "strings"

// Configuration for initializing the feature toggles
type Configuration struct {
	Features map[string]bool `json:"features"`
}

var features = make(map[string]bool)

// InitializeFeatures adds all the features from the configurations
func InitializeFeatures(config Configuration) {
	AddFeatures(config.Features)
}

// AddFeatures adds the features from the provided map
func AddFeatures(features map[string]bool) {
	for key, value := range features {
		AddFeature(key, value)
	}
}

// AddFeature adds a single feature
func AddFeature(key string, isEnabled bool) {
	features[strings.ToLower(key)] = isEnabled
}

// IsEnabled determines if the specified feature is enabled. Determining if a feature is enabled is
// case insensitive.
// If a feature is not present, it defaults to false.
func IsEnabled(key string) bool {
	return features[strings.ToLower(key)]
}
