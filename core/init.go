package core

import (
	"api/core/models"
	"api/modules/goconfig"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

var (
	Options = &goconfig.Options{
		Config: goconfig.NewConfig(),
	}
	Vers string = models.Config.Vers
)

func Initialize() {
	err := Options.Config.Parse("assets")
	if err != nil {
		log.Println(err)
		return
	}

	Options.Config.NewInclusion(".json", func(content []byte, file string, m map[string]interface{}) error {
		switch file {
		case filepath.Join("assets/config", "config.json"):
			models.Config = new(models.Conf)
			return json.Unmarshal(content, &models.Config)
		default:
			fmt.Println(file, m, string(content))
			return json.Unmarshal(content, &m)
		}
	})

	Options, err = Options.Config.Options()
	if err != nil {
		log.Println(err)
		return
	}
}
