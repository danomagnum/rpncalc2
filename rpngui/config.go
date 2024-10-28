package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"log"
	"os"
)

//go:embed icon.png
var iconpng []byte

type configfile struct {
	Favorites  map[string]string `json:"favorites"`
	History    int               `json:"history"`
	UserMenu   map[string]string `json:"usermenu"`
	ShowNumpad bool              `json:"show_numpad"`
}

func LoadConfig() configfile {

	f, err := os.Open("config.json")
	if err != nil {
		log.Printf("couldn't open config file.: %v", err)
		return configfile{}
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	var cfg configfile
	dec.Decode(&cfg)
	return cfg
}

func LoadHistory() []string {

	f, err := os.Open("history.json")
	if err != nil {
		log.Printf("couldn't open history file.: %v", err)
		return []string{}
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	history := make([]string, 0)
	err = dec.Decode(&history)
	if err != nil {
		log.Printf("problem opening history file: %v", err)
	}
	return history
}

func SaveHistory(history []string) {

	var f *os.File
	if _, err := os.Stat("history.json"); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create("history.json")
		if err != nil {
			log.Printf("couldn't create history file.: %v", err)
			return
		}
	} else {
		f, err = os.OpenFile("history.json", os.O_RDWR, 0755)
		if err != nil {
			log.Printf("couldn't open history file.: %v", err)
			return
		}
	}
	defer f.Close()
	dec := json.NewEncoder(f)
	dec.SetIndent("", "  ")
	err := dec.Encode(history)
	if err != nil {
		log.Printf("problem saving history file: %v", err)
	}
}
