package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zserge/lorca"
)

func main() {
	/* Checking config file */
	var config Config

	isConfigDirExists()

	if _, err := os.Stat(ConfigFile); os.IsNotExist(err) {
		config = createDefaultConfig()
		saveDefaultConfig(config)
	} else {
		openConfig(config)
	}

	width, err := strconv.Atoi(strings.Split(config.UIConfig.ScreenSize, "x")[0])
	if err != nil {
		fmt.Println(err)
	}

	height, err := strconv.Atoi(strings.Split(config.UIConfig.ScreenSize, "x")[1])
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/update-config", updateConfigHandler)
	go func() {
		fmt.Println("Démarrage du serveur sur http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Erreur lors du démarrage du serveur: ", err)
		}
	}()

	/* Launching UI */
	ui, err := lorca.New("", "", width, height, "--remote-allow-origins=*")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ui.Close()

	cwd, _ := os.Getwd()
	ui.Load(fmt.Sprintf("file://%s", filepath.Join(cwd, "ui\\index.html")))

	<-ui.Done()
}
