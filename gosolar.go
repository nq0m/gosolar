package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Solar struct {
	XMLName   xml.Name `xml:"solar"`
	Text      string   `xml:",chardata"`
	Solardata struct {
		Text   string `xml:",chardata"`
		Source struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url,attr"`
		} `xml:"source"`
		Updated              string `xml:"updated"`
		Solarflux            string `xml:"solarflux"`
		Aindex               string `xml:"aindex"`
		Kindex               string `xml:"kindex"`
		Kindexnt             string `xml:"kindexnt"`
		Xray                 string `xml:"xray"`
		Sunspots             string `xml:"sunspots"`
		Heliumline           string `xml:"heliumline"`
		Protonflux           string `xml:"protonflux"`
		Electonflux          string `xml:"electonflux"`
		Aurora               string `xml:"aurora"`
		Normalization        string `xml:"normalization"`
		Latdegree            string `xml:"latdegree"`
		Solarwind            string `xml:"solarwind"`
		Magneticfield        string `xml:"magneticfield"`
		Calculatedconditions struct {
			Text string `xml:",chardata"`
			Band []struct {
				Text string `xml:",chardata"` // Poor, Fair, Good, Good, F...
				Name string `xml:"name,attr"`
				Time string `xml:"time,attr"`
			} `xml:"band"`
		} `xml:"calculatedconditions"`
		Calculatedvhfconditions struct {
			Text       string `xml:",chardata"`
			Phenomenon []struct {
				Text     string `xml:",chardata"` // Band Closed, Band Closed,...
				Name     string `xml:"name,attr"`
				Location string `xml:"location,attr"`
			} `xml:"phenomenon"`
		} `xml:"calculatedvhfconditions"`
		Geomagfield string `xml:"geomagfield"` // UNSETTLD
		Signalnoise string `xml:"signalnoise"` // S2-S3
		Fof2        string `xml:"fof2"`        // 9.05
		Muffactor   string `xml:"muffactor"`   // 2.84
		Muf         string `xml:"muf"`         // 25.70
	} `xml:"solardata"`
}

func getSolarConditions() []uint8 {
	// Grab the solar conditions from HamQSL
	resp, err := http.Get("https://www.hamqsl.com/solarxml.php")
	// Check for error
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func parseXML(body []uint8) Solar {
	var solar Solar
	xml.Unmarshal(body, &solar)
	return solar
}

func main() {
	solarXML := getSolarConditions()
	solar := parseXML(solarXML)
	fmt.Println(solar.Solardata.Aindex)
	var sfi, aindex, kindex uint64
	sfi, _ = strconv.ParseUint(strings.TrimSpace(solar.Solardata.Solarflux), 10, 8)
	aindex, _ = strconv.ParseUint(strings.TrimSpace(solar.Solardata.Aindex), 10, 8)
	kindex, _ = strconv.ParseUint(strings.TrimSpace(solar.Solardata.Kindex), 10, 8)
	fmt.Println("SFI: " + strconv.FormatUint(sfi, 10) + " A: " + strconv.FormatUint(aindex, 10) + " K: " + strconv.FormatUint(kindex, 10))

}
