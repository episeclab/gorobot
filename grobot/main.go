package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var cfg *string = flag.String("config", "grobot.json", "path to a json configuration file")
var newcfg *bool = flag.Bool("generate-config", false, "generate a default configuration file")

// Handles command line arguments and runs the bot
func main() {
	flag.Parse()
	if *newcfg == true {
		file, err := os.Create(*cfg)
		if err != nil {
			log.Fatalf("Can't create configuration file: %v", err)
		}
		config := DefaultConfig()
		data, err := json.MarshalIndent(*config, "", " ")
		if err != nil {
			log.Fatalf("Can't create configuration file: %v", err)
		}
		file.Write(data)
		file.Close()
	} else {
		bot := NewGrobot(*cfg)
		bot.Run()
	}
}
