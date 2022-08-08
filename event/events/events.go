package events

import (
	"github.com/Yggdrasil80/home-plugin-api/event"
	"github.com/labstack/echo/v4"
)

func PluginLoadedEvent(pluginName string) *event.Event {
	return &event.Event{
		Name: "plugin_loaded",
		Data: map[string]interface{}{
			"plugin_name": pluginName,
		},
	}
}

func RouteAddedEvent(method string, path string, handler func(ctx *echo.Context) error) *event.Event {
	return &event.Event{
		Name: "route_added",
		Data: map[string]interface{}{
			"method":  method,
			"path":    path,
			"handler": handler,
		},
	}
}

func RequestReceivedEvent(ctx *echo.Context) *event.Event {
	return &event.Event{
		Name: "request_received",
		Data: map[string]interface{}{
			"ctx": ctx,
		},
	}
}
