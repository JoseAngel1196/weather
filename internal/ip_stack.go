package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/JoseAngel1196/weather/config"
	"github.com/JoseAngel1196/weather/print"
)

type Location struct {
	GeonameID int    `json:"geoname_id"`
	Capital   string `json:"capital"`
	Languages []struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Native string `json:"native"`
	} `json:"languages"`
	CountryFlag             string `json:"country_flag"`
	CountryFlagEmoji        string `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string `json:"country_flag_emoji_unicode"`
	CallingCode             string `json:"calling_code"`
	IsEu                    bool   `json:"is_eu"`
}

type TimeZone struct {
	ID               string `json:"id"`
	CurrentTime      string `json:"current_time"`
	GmtOffset        int    `json:"gmt_offset"`
	Code             string `json:"code"`
	IsDaylightSaving bool   `json:"is_daylight_saving"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Plural       string `json:"plural"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

type Connection struct {
	ASN int    `json:"asn"`
	ISP string `json:"isp"`
}

type Security struct {
	IsProxy     bool        `json:"is_proxy"`
	ProxyType   interface{} `json:"proxy_type"`
	IsCrawler   bool        `json:"is_crawler"`
	CrawlerName interface{} `json:"crawler_name"`
	CrawlerType interface{} `json:"crawler_type"`
	IsTor       bool        `json:"is_tor"`
	ThreatLevel string      `json:"threat_level"`
	ThreatTypes interface{} `json:"threat_types"`
}

type UserLocation struct {
	IP            string     `json:"ip"`
	Hostname      string     `json:"hostname"`
	Type          string     `json:"type"`
	ContinentCode string     `json:"continent_code"`
	ContinentName string     `json:"continent_name"`
	CountryCode   string     `json:"country_code"`
	CountryName   string     `json:"country_name"`
	RegionCode    string     `json:"region_code"`
	RegionName    string     `json:"region_name"`
	City          string     `json:"city"`
	Zip           string     `json:"zip"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	Location      Location   `json:"location"`
	TimeZone      TimeZone   `json:"time_zone"`
	Currency      Currency   `json:"currency"`
	Connection    Connection `json:"connection"`
	Security      Security   `json:"security"`
}

func GetUserLocation() string {
	// Get user location using ipstack
	ipStackApiKey := config.GetEnv("IP_STACK_API_KEY")
	var publicIpAddress string = getUserPublicIp()
	httpUrl := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", publicIpAddress, url.QueryEscape(ipStackApiKey))

	resp, err := http.Get(httpUrl)
	if err != nil {
		print.ExitOnError("", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		print.ExitOnError("", err)
	}

	var bodyAPI UserLocation
	err = json.Unmarshal(body, &bodyAPI)
	if err != nil {
		print.ExitOnError("Failed to parse the user location", err)
	}
	fmt.Println(bodyAPI)
	return bodyAPI.IP
}

func getUserPublicIp() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		print.ExitOnError("Failed to get public IP address:", err)
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		print.ExitOnError("Failed to read response body:", err)
	}

	return string(ip)
}
