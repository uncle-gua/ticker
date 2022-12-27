package config

import (
	"encoding/json"
	"os"

	"github.com/uncle-gua/log"
)

const (
	FILE = "database.json"
)

type database struct {
	Host string `json:"host"`
	Name string `json:"name"`
}

var Database = &database{}

func init() {
	if err := Database.Load(); err != nil {
		log.Fatal(err)
	}
}

func (c *database) Load() error {
	body, err := os.ReadFile(FILE)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, c)
	return err
}

func (c *database) Save() error {
	body, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(FILE, body, 0644)
}
