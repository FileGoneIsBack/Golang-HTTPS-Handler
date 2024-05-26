package models

var (
	Config *Conf = new(Conf)
)

type Conf struct {
	Name    string `json:"name"`
	Secure  bool   `json:"secure"`
	Cert    string `json:"cert"`
	Vers    string `json:"version"`
	Key     string `json:"key"`
}
