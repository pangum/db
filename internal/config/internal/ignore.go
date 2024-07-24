package internal

type Ignore struct {
	Constrains bool `json:"constrains,omitempty" yaml:"constrains" xml:"constrains" toml:"constrains"`
	Indices    bool `json:"indices,omitempty" yaml:"indices" xml:"indices" toml:"indices"`
}
