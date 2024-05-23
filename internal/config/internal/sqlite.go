package internal

type Sqlite struct {
	Name     string `default:"_auth" json:"name,omitempty" yaml:"name" xml:"name" toml:"name"`
	User     string `default:"_auth_user" json:"user,omitempty" yaml:"user" xml:"user" toml:"user"`
	Password string `default:"_auth_pass" json:"password,omitempty" yaml:"password" xml:"password" toml:"password"`
	Crypt    string `default:"_auth_crypt" json:"crypt,omitempty" yaml:"crypt" xml:"crypt" toml:"crypt"`
}
