package config

import (
	"github.com/pangum/db/internal/config/internal"
)

type Sync struct {
	Enabled bool            `default:"true" json:"enabled,omitempty" yaml:"enabled" xml:"enabled" toml:"enabled"`
	Ignore  internal.Ignore `json:"ignore,omitempty" yaml:"ignore" xml:"ignore" toml:"ignore"`
	Drop    internal.Drop   `json:"drop,omitempty" yaml:"drop" xml:"drop" toml:"drop"`
	Warn    internal.Warn   `json:"warn,omitempty" yaml:"warn" xml:"warn" toml:"warn"`
}
