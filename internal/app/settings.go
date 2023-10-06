package app

import (
	"encoding/json"
	"io"
	"os"
	"yu/golang/internal/app/validation"
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
	Region  string `json:"region"`
	PageNum uint16 `json:"pageNum"`
}

type AppSettingsTimeout struct {
	LapInterval     string `json:"lapInterval"`
	ReqInterval     string `json:"reqInterval"`
	ReqTTL          string `json:"reqTTL"`
	ReqEconnRefused string `json:"reqEconnRefused"`
	ReqError        string `json:"reqError"`
	ReqEmptyList    string `json:"reqEmptyList"`
}

type ArticleFilter struct {
	Id       map[string]uint64 `json:"id"`
	RegionId map[string]uint8  `json:"regionId"`
	Hash     map[string]uint8  `json:"hash"`
	Link     map[string]uint8  `json:"link"`
	Title    map[string]uint8  `json:"title"`
	Message  map[string]uint8  `json:"message"`
	Sex      map[string]uint8  `json:"sex"`
	Age      map[string]uint8  `json:"age"`
	Height   map[string]uint8  `json:"height"`
	Weight   map[string]uint8  `json:"weight"`
	Images   map[string]uint8  `json:"images"`
	Phones   map[string]uint8  `json:"phones"`
	Date     map[string]uint16 `json:"date"`
}

func (f *ArticleFilter) IsValid(anyStruct any) bool {
	flag, _ := validation.CheckIsValid(f, anyStruct, true)
	return flag
}

func (f *ArticleFilter) Validate(anyStruct any) []string {
	_, wrongFields := validation.CheckIsValid(f, anyStruct, false)
	return wrongFields
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
