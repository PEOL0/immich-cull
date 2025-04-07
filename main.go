package main

import (
	"fmt"
	"log"
	"os"

	"io"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getAlbumList()
	//fmt.Printf("Api is: %s ", os.Getenv("ImmichAPI"))

}

func getAlbumList() {
	url := os.Getenv("ImmichURL") + "/api/albums"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", os.Getenv("ImmichAPI"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
