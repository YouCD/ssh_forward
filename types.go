package main

//nolint:tagliatelle
type PublicKeys struct {
	PrivateKeyPath string `json:"PrivateKeyPath"`
}

//nolint:tagliatelle
type Forward struct {
	ServerAddr       string                 `json:"ServerAddr" yaml:"ServerAddr"`
	ServerUser       string                 `json:"ServerUser" yaml:"ServerUser"`
	ServerAuthMethod map[string]interface{} `json:"ServerAuthMethod" yaml:"ServerAuthMethod"`
	Project          []*forward             `json:"Project" yaml:"Project"`
}

//nolint:tagliatelle
type forward struct {
	Project          string                 `json:"Project" yaml:"Project"`
	RemoteAddr       string                 `json:"RemoteAddr" yaml:"RemoteAddr"`
	LocalAddr        string                 `json:"LocalAddr" yaml:"LocalAddr"`
	ServerAddr       string                 `json:"ServerAddr" yaml:"ServerAddr"`
	ServerUser       string                 `json:"ServerUser" yaml:"ServerUser"`
	ServerAuthMethod map[string]interface{} `json:"ServerAuthMethod" yaml:"ServerAuthMethod"`
}
