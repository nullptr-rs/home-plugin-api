package plugin

import (
	"github.com/Yggdrasil80/home-plugin-api/api"
)

type Plugin struct {
	Info Info `json:"info"`

	OnEnable  func(api *api.Api) `json:"on_enable"`
	OnDisable func(api *api.Api) `json:"on_disable"`
}

type Info struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
