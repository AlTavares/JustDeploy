package dto

type ServiceDto struct {
	Name    string   `json:"name"`
	Icon    string   `json:"icon"`
	Image   string   `json:"image"`
	Port    string   `json:"port"`
	Envs    []string `json:"envs"`
	Secrets []string `json:"secrets"`
	// TODO: move to array
	Volume     string `json:"volume"`
	ConnectUrl string `json:"connect_url"`
}
