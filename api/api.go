package api

import (
	"github.com/Yggdrasil80/home-plugin-api/event"
	"github.com/Yggdrasil80/home-plugin-api/storage"
	"github.com/labstack/echo/v4"
)

type Api struct {
	EventBus event.Bus
	Storage  storage.Storage

	Echo *echo.Echo
}
