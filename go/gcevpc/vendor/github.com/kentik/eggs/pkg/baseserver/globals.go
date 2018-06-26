package baseserver

import (
	"github.com/kentik/eggs/pkg/features"
	"github.com/kentik/eggs/pkg/preconditions"
	"github.com/kentik/eggs/pkg/properties"
)

var globalBaseServer *BaseServer

func GetGlobalBaseServer() *BaseServer {
	preconditions.AssertNonNil(globalBaseServer, "globalBaseServer has not been set")
	return globalBaseServer
}

func setGlobalBaseServer(bs *BaseServer) {
	preconditions.AssertNil(globalBaseServer, "globalBaseServer has already been set")
	globalBaseServer = bs
}

// for testing
func resetGlobalBaseServer() {
	globalBaseServer = nil
}

// GetGlobalPropertyService looks up the property service on globalBaseServer.
// If it doesn't exist (i.e. during tests), that's ok, give us an empty one.
func GetGlobalPropertyService() properties.PropertyService {
	if globalBaseServer == nil {
		return properties.NewPropertyService()
	}

	return globalBaseServer.GetPropertyService()
}

// GetGlobalFeatureService looks up the feature service on globalBaseServer.
// If it doesn't exist (i.e. during tests), that's ok, give us an empty one.
func GetGlobalFeatureService() features.FeatureService {
	if globalBaseServer == nil {
		return features.NewFeatureService(properties.NewPropertyService())
	}

	return globalBaseServer.GetFeatureService()
}

// InitGlobalBaseServerTestingOnly is for tests to create a BaseServer and
// play with its PropertyService and FeatureService.
func InitGlobalBaseServerTestingOnly(propertyMap map[string]string, defaultPropertyBacking properties.PropertyBacking) {
	props := properties.NewPropertyService(
		properties.NewStaticMapPropertyBacking(propertyMap),
		defaultPropertyBacking,
	)

	globalBaseServer = &BaseServer{
		propertyService: props,
		featureService:  features.NewFeatureService(props),
	}
}

// ResetGlobalBaseServerTestingOnly should be defered after a call to
// InitGlobalBaseServerTestingOnly.
func ResetGlobalBaseServerTestingOnly() {
	resetGlobalBaseServer()
}
