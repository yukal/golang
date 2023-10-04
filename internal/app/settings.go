package app

import (
	"encoding/json"
	"io"
	"os"
)

type AppSettings struct {
	Paths  AppSettingsPaths  `json:"paths"`
	Server AppSettingsServer `json:"server"`
	Task   AppSettingsTask   `json:"task"`
}

type AppSettingsPaths struct {
	Database string `json:"database"`
	Images   string `json:"images"`
}

type AppSettingsServer struct {
	Protocol string                `json:"protocol"`
	Host     string                `json:"host"`
	Urls     AppSettingsServerUrls `json:"urls"`
}

type AppSettingsServerUrls map[string]string

type AppSettingsTask struct {
	Items   []AppSettingsTaskItems `json:"items"`
	Filter  ArticleFilter          `json:"filter"`
	Timeout AppSettingsTimeout     `json:"timeout"`
}

type AppSettingsTaskItems struct {
	Region  string `json:"Region"`
	PageNum uint16 `json:"PageNum"`
}

type ArticleFilter struct {
	Id       map[string]uint64 `json:"Id"`
	RegionId map[string]uint8  `json:"RegionId"`
	Hash     map[string]uint8  `json:"Hash"`
	Link     map[string]uint8  `json:"Link"`
	Title    map[string]uint8  `json:"Title"`
	Message  map[string]uint8  `json:"Message"`
	Sex      map[string]uint8  `json:"Sex"`
	Age      map[string]uint8  `json:"Age"`
	Height   map[string]uint8  `json:"Height"`
	Weight   map[string]uint8  `json:"Weight"`
	Images   map[string]uint8  `json:"Images"`
	Phones   map[string]uint8  `json:"Phones"`
	Date     map[string]uint16 `json:"Date"`
}

type AppSettingsTimeout struct {
	LapInterval     string `json:"lapInterval"`
	ReqInterval     string `json:"reqInterval"`
	ReqTTL          string `json:"reqTTL"`
	ReqEconnRefused string `json:"reqEconnRefused"`
	ReqError        string `json:"reqError"`
	ReqEmptyList    string `json:"reqEmptyList"`
}

func MustLoadSettings(jsonFilename string) (settings *AppSettings) {
	file, err := os.Open(jsonFilename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &settings); err != nil {
		panic(err)
	}

	return
}

func (s AppSettings) String() string {
	return InspectData(s)
}
