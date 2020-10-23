package lib

type ServiceEvent struct {
	ID       string                 `toml:"id" json:"id"`
	Required map[string]interface{} `toml:"required" json:"required"`
}
