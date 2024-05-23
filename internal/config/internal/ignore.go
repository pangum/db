package internal

type Ignore struct {
	Constrains *bool `default:"true" json:"constrains,omitempty" yaml:"constrains" xml:"constrains" toml:"constrains"`
	Indices    *bool `default:"true" json:"indices,omitempty" yaml:"indices" xml:"indices" toml:"indices"`
}
