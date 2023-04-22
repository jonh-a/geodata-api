package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		Name      string `json:"NAME"`
		NameLong  string `json:"NAME_LONG"`
		Continent string `json:"CONTINENT"`
		Subregion string `json:"SUBREGION"`
		IsoA3     string `json:"ISO_A3"`
		Postal    string `json:"POSTAL"`
	} `json:"properties"`
	Geometry struct {
		Type        string          `json:"type"`
		Coordinates json.RawMessage `json:"coordinates"`
	} `json:"geometry"`
	Bbox json.RawMessage `json:"bbox"`
}

type GeoJson struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type FeatureBrief struct {
	Name      string `json:"name"`
	NameLong  string `json:"name_long"`
	Continent string `json:"continent"`
	IsoA3     string `json:"iso_a3"`
}

func readGeojson(detail string) (GeoJson, error) {
	file, err := os.Open(fmt.Sprintf("%s.geojson", detail))
	if err != nil {
		fmt.Println(err)
		return GeoJson{}, err
	}
	defer file.Close()

	b, _ := ioutil.ReadAll(file)
	var data GeoJson

	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err)
		return GeoJson{}, err
	}

	return data, nil
}

func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
}

func getCountries(c *gin.Context) {
	geojson, err := readGeojson("10m")

	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to read geoJSON"})
		return
	}

	countries := make([]FeatureBrief, len(geojson.Features))

	for k, v := range geojson.Features {
		country := FeatureBrief{v.Properties.Name, v.Properties.NameLong, v.Properties.Continent, v.Properties.IsoA3}
		countries[k] = country
	}

	c.IndentedJSON(http.StatusOK, countries)
}

func getCountry(c *gin.Context) {
	name := strings.ToLower(c.Param("name"))
	detail := strings.ToLower(c.DefaultQuery("detail", "110m"))

	if detail != "10m" && detail != "50m" && detail != "110m" {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"error": "Detail must be 10m, 50m, or 110m."})
		return
	}

	geojson, err := readGeojson(detail)

	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to read geoJSON"})
		return
	}

	for _, v := range geojson.Features {
		if strings.ToLower(v.Properties.Name) == name ||
			strings.ToLower(v.Properties.NameLong) == name ||
			strings.ToLower(v.Properties.IsoA3) == name {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func main() {
	router := gin.Default()

	router.GET("/health", health)
	router.GET("/countries", getCountries)
	router.GET("/countries/:name", getCountry)

	router.Run(":3000")
}
