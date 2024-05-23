package internal

type Drop struct {
	Indices bool `default:"true" json:"indices,omitempty" yaml:"indices" xml:"indices" toml:"indices"`
}
