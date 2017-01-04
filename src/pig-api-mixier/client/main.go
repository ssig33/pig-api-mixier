package client

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v0"
	"pig-api-mixier/conf"
)

type Pig struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Mtime        string `json:"mtime"`
	Size         int    `json:"size"`
	Url          string `json:"url"`
	Mtime_to_i   int    `json:"mtime_to_i"`
	Vdr          string `json:"vdr"`
	Custom_links string `json:"custom_links"`
	Metadata     bool   `json:"metadata"`
	Src          bool   `json:"src"`
}

//["key", "name", "mtime", "size", "url", "type", "mtime_to_i", "vdr", "custom_links", "metadata", "srt"]

func main() {
	fmt.Println("vim-go")
}

func parse(str string) []Pig {
	var pigs []Pig
	json.Unmarshal([]byte(str), &pigs)
	return pigs
}

func merge(slices [][]Pig) []Pig {
	result := []Pig{}
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}

func Latest() []Pig {
	slices := [][]Pig{}
	for _, urlInfo := range conf.LatestUrlBases() {
		fmt.Println(urlInfo)
		resty.SetBasicAuth(urlInfo.Basic[0], urlInfo.Basic[1])
		resp, _ := resty.R().Get(urlInfo.Url + "api/r/latest")
		slices = append(slices, parse(resp.String()))
	}
	pigs := merge(slices)
	return pigs
}
