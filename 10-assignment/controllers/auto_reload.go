package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/retry19/challenge-hacktiv8/10-assignment/models"
)

func AutoReloadController(c *gin.Context) {
	c.String(http.StatusOK, "Auto reload is working")
}

func TriggerAutoReload() {
	for {
		min := 1
		max := 100

		numberWater := rand.Intn(max - min)
		numberWind := rand.Intn(max - min)

		data := models.Data{
			Water: numberWater,
			Wind:  numberWind,
		}

		updateData(&data)
		logging(&data)
		makeApiRequest(&data)

		time.Sleep(time.Second * 15)
	}
}

func updateData(data *models.Data) {
	dataJson, _ := json.MarshalIndent(data, "", "  ")

	err := os.WriteFile("data.json", dataJson, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataJson))
}

func logging(data *models.Data) {
	var waterStatus string

	if data.Water <= 5 {
		waterStatus = "aman"
	} else if data.Water > 5 && data.Water <= 8 {
		waterStatus = "siaga"
	} else if data.Water > 8 {
		waterStatus = "bahaya"
	}

	fmt.Println("status water: ", waterStatus)

	var windStatus string

	if data.Water <= 5 {
		windStatus = "aman"
	} else if data.Water > 5 && data.Water <= 8 {
		windStatus = "siaga"
	} else if data.Water > 8 {
		windStatus = "bahaya"
	}

	fmt.Println("status wind: ", windStatus)
}

func makeApiRequest(data *models.Data) {
	_, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/auto-reload", "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
