package domain

type Deploy struct {
	Name     string
	ServerIP string
	Path     string
	ZipPath  string
	Email    string
	Domain   string
}

type DeployConfigDto struct {
	PathToProject    string           `json:"pathToProject"`
	DockerFileValid  bool             `json:"dockerFileValid"`
	ServerConfig     ConnectServerDto `json:"serverConfig"`
	AppConfig        AppConfigDto     `json:"appConfig"`
	DeployFromStatus string           `json:"deployFormStatus"`
	Url              string           `json:"url"`
}

type ConnectServerDto struct {
	Domain   string  `json:"domain"`
	SshKey   *string `json:"sshKey"`
	Password *string `json:"password"`
	User     string  `json:"user"`
}

type Env struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type AppConfigDto struct {
	Name         string `json:"name"`
	EnableTls    bool   `json:"enableTls"`
	Email        string `json:"email"`
	PathToSource string `json:"pathToSource"`
	Envs         []Env  `json:"envs"`
}

type ApplicationConfigResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponseApi struct {
	Message string `json:"message"`
}
