package api

import (
	"github.com/Yggdrasil80/home-plugin-api/event"
	"github.com/Yggdrasil80/home-plugin-api/storage"
)

type Api struct {
	EventBus event.Bus
	Storage  storage.Storage
}
