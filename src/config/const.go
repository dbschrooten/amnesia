package config

import "time"

const (
	DefaultInterval = "5m"
	DefaultTimeout  = "30s"

	SrvWriteTimeout = 15 * time.Second
	SrvReadTimeout  = 15 * time.Second
)
