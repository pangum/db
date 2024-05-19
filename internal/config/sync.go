package config

import (
	"github.com/pangum/db/internal/config/internal"
)

type Sync struct {
	Enabled bool            `json:"enabled,omitempty" yaml:"enabled" xml:"enabled" toml:"enabled"`
	Ignore  internal.Ignore `json:"ignore,omitempty" yaml:"ignore" xml:"ignore" toml:"ignore"`
}
