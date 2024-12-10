package model

type (
	ServiceStatus string
	Env           string
)

const (
	ServiceStatusOnline      ServiceStatus = "online"
	ServiceStatusMaintenance ServiceStatus = "maintenance"

	EnvProduction Env = "production"
	EnvStaging    Env = "staging"
)

type GeneralStatusResponse struct {
	Status  ServiceStatus `json:"status"`
	Version struct {
		Version    string `json:"version"`
		CommitHash string `json:"commitHash"`
		BuildTime  string `json:"buildTime"`
	} `json:"version"`
	Env    Env    `json:"env"`
	Uptime string `json:"uptime"`
}
