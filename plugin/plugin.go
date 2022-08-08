package plugin

type Plugin struct {
	Info Info `json:"info"`

	OnEnable  func(api *home_plugin_api.Api) `json:"on_enable"`
	OnDisable func(api *home_plugin_api.Api) `json:"on_disable"`
}

type Info struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
