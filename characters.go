package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const API_URL = "https://esi.evetech.net/latest/universe/ids/?datasource=tranquility&language=fr"

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ApiResponse struct {
	Characters []Character `json:"characters"`
}

func fetchCharacterID(characterName string) {
	postData, err := json.Marshal([]string{characterName})
	if err != nil {
		fmt.Println("Erreur lors de la création du corps de la requête: ", err)
		return
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("Erreur lors de la requête: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "fr")
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse: ", err)
		return
	}

	var result ApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON: ", err)
		return
	}

	for _, character := range result.Characters {
		fmt.Println("ID du personnage: ", character.ID)
	}
}
