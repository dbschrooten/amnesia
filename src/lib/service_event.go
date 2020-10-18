package lib

type ServiceEvent struct {
	Id       string                 `toml:"id"`
	Required map[string]interface{} `toml:"required"`
}
