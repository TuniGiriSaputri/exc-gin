package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1.memebuat endpoint menggunakan gin yang mengembalikan random quote anime https://animechan.vercel.app/docs
// 1a. membuat HTTP server

type Quote struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func main() {
	r := gin.Default()
	r.GET("/quote", func(c *gin.Context) {
		quote, err := getAnimeQuote()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"data":    nil,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Succes",
			"data":    quote,
		})
	})
	r.Run(":8080")
}

// 1a1. mmebuat design endpoint
// (API constrant :
// - HTTP method GET
// - path / quote)

// 1b. membuat HTTP client
func getAnimeQuote() (*Quote, error) {
	resp, err := http.Get("https://animechan.vercel.app/api/random/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch anime quote")
	}

	body, err := ioutil.ReadAll(resp.Body)

	var quote Quote

	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}
