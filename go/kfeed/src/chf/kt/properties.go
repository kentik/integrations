package kt

import (
	"github.com/kentik/eggs/pkg/properties"
)

const ()

// Default properties for the alerting backend (can be overridden by env vars or filesystem entries).
var DefaultKFeedProperties properties.PropertyBacking

// Initialize DefaultAlertingProperties
func init() {
	defaultProps := make(map[string]string, 0)

	// features that are globally enabled by default
	for _, featureName := range FEATURES_ENABLED_BY_DEFAULT {
		defaultProps["features."+featureName+"-global"] = "true"
	}

	DefaultKFeedProperties = properties.NewStaticMapPropertyBacking(defaultProps)
}
