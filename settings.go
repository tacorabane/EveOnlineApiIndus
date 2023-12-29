package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type UIConfig struct {
	ScreenSize string `json:"screenSize"`
	Theme      string `json:"theme"`
	FontSize   int    `json:"font_size"`
}

type Config struct {
	UIConfig        UIConfig `json:"ui_config"`
	SavedCharacters []int    `json:"saved_characters"`
}

const ConfigDir = "config"
const ConfigFile = ConfigDir + "\\config.json"

func isConfigDirExists() {
	if _, err := os.Stat(ConfigDir); os.IsNotExist(err) {
		err := os.Mkdir(ConfigDir, 0755)
		if err != nil {
			fmt.Println("Erreur lors de la création du dossier de configuration: ", err)
			os.Exit(1)
		}
	}
}

func createDefaultConfig() Config {
	return Config{
		UIConfig: UIConfig{
			ScreenSize: "1028x768",
			Theme:      "light",
			FontSize:   12,
		},
		SavedCharacters: []int{},
	}
}

func saveDefaultConfig(config Config) error {
	configFile, err := os.Create(ConfigFile)
	if err != nil {
		return err
	}
	defer configFile.Close()

	configBytes, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	_, err = configFile.Write(configBytes)
	if err != nil {
		return err
	}

	return nil
}

func openConfig(config Config) {
	configFile, err := os.Open(ConfigFile)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier de configuration: ", err)
		return
	}
	defer configFile.Close()

	bytes, _ := io.ReadAll(configFile)
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Println("Erreur lors du parsing du fichier de configuration: ", err)
		return
	}
}

func updateConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Methode non authorisée", http.StatusMethodNotAllowed)
		return
	}

	var config Config
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Nouvelle configuration reçue: ", config)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
