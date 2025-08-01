package mapperplugins

import (
	"fmt"
	"sync"
)

type DESTINATION = interface{}
type OBJECT = interface{}

type MapperPlugin interface {
	GetDestinationObject() DESTINATION
	// should be returned through reflection
	GetDestinationObjectString() string

	GetObject() OBJECT
	// should be returned through reflection
	GetObjectString() string

	Map(obj OBJECT) (DESTINATION, MapperError)
	ReverseMap(obj DESTINATION) (OBJECT, MapperError)

	MapAll(objects []OBJECT) ([]DESTINATION, MapperError)
	ReverseMapAll(objects []DESTINATION) ([]OBJECT, MapperError)
}

//	{
//		 "time": {
//		   string: MapperPlugin
//		 },
//	}
type PluginMap map[string]MapperPlugin
type PluginInOutMap map[string]PluginMap

// GetPlugin retrieves a MapperPlugin from the PluginMap based on the destination object type.
func (pm PluginMap) GetPlugin(destinationObjectType string) MapperPlugin {
	// Retrieve the plugin for the given destination object type
	plugin, exists := pm[destinationObjectType]
	if !exists {
		return nil
	}
	return plugin
}

func (pio PluginInOutMap) GetPlugin(objectType, destinationObjectType string) MapperPlugin {
	// Retrieve the plugin map for the given object type
	pluginMap, exists := pio[objectType]
	if !exists {
		return nil
	}
	// Retrieve the plugin for the given destination object type
	return pluginMap.GetPlugin(destinationObjectType)
}

var (
	MapperPlugins     = make(PluginInOutMap)
	mapperPluginsLock sync.RWMutex
)

func reverseRegisterMapperPlugin(plugin MapperPlugin) {
	if plugin == nil {
		return
	}
	if plugin.GetObjectString() == "" || plugin.GetDestinationObjectString() == "" {
		return
	}

	if _, exists := MapperPlugins[plugin.GetDestinationObjectString()]; !exists {
		MapperPlugins[plugin.GetDestinationObjectString()] = make(PluginMap)
	}
	MapperPlugins[plugin.GetDestinationObjectString()][plugin.GetObjectString()] = plugin
}

// registerMapperPlugin registers a MapperPlugin by its origin object type.
func registerMapperPlugin(plugin MapperPlugin) {
	if plugin == nil {
		return
	}
	if plugin.GetObjectString() == "" || plugin.GetDestinationObjectString() == "" {
		return
	}
	for locked := mapperPluginsLock.TryLock(); !locked; locked = mapperPluginsLock.TryLock() {
	}
	defer mapperPluginsLock.Unlock()
	if _, exists := MapperPlugins[plugin.GetObjectString()]; !exists {
		MapperPlugins[plugin.GetObjectString()] = make(PluginMap)
	}
	MapperPlugins[plugin.GetObjectString()][plugin.GetDestinationObjectString()] = plugin
	reverseRegisterMapperPlugin(plugin)
}

func initializeMapperPlugins() {
	// This function can be used to initialize or register all mapper plugins.
	// For example, you can call registerMapperPlugin for each plugin here.
	registerMapperPlugin(TimeMapper{})
	registerMapperPlugin(NullStringMapper{})
}

func GetMapperPlugin(object string, destination string) MapperPlugin {

	initializeMapperPlugins()

	for locked := mapperPluginsLock.TryLock(); !locked; locked = mapperPluginsLock.TryLock() {
	}
	defer mapperPluginsLock.Unlock()

	plugin := MapperPlugins.GetPlugin(
		object,
		destination,
	)

	if plugin == nil {
		fmt.Println("No plugin found for object:", object, "and destination:", destination)
		return nil
	}

	return plugin
}
