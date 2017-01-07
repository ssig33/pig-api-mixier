package client

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v0"
	"pig-api-mixier/conf"
	"sort"
	"strings"
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

type ByMtime []Pig

func (a ByMtime) Len() int           { return len(a) }
func (a ByMtime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMtime) Less(i, j int) bool { return a[i].Mtime_to_i > a[j].Mtime_to_i }

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

func getPigsAPI(urlInfo conf.UrlInfo, url string) []Pig {
	resty.SetBasicAuth(urlInfo.Basic[0], urlInfo.Basic[1])
	resp, _ := resty.R().Get(urlInfo.Url + url)
	pigs := parse(resp.String())
	for i := range pigs {
		pigs[i].Url = strings.Replace(pigs[i].Url, urlInfo.Rule[0], urlInfo.Rule[1], 1)
	}
	return pigs
}

func pigs(urlInfos []conf.UrlInfo, method string, query string) []Pig {
	slices := [][]Pig{}
	for _, urlInfo := range urlInfos {
		fmt.Println(urlInfo)
		pigs := getPigsAPI(urlInfo, "api/r/"+method+"?"+query)
		slices = append(slices, pigs)
	}
	pigs := merge(slices)
	sort.Sort(ByMtime(pigs))
	return pigs
}

func Latest(query string) []Pig {
	return pigs(conf.LatestUrlBases(), "latest", query)
}

func Search(query string) []Pig {
	return pigs(conf.SearchUrlBases(), "search", query)
}
