package bikeshareapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ApiClient struct {
	Client *http.Client
}

const (
	urlRoot         = "https://hanetwi.ddns.net/bikeshare/api/v1/"
	urlGetPlaces    = urlRoot + "places?"
	urlGetCounts    = urlRoot + "counts?"
	urlGetDistances = urlRoot + "distances?"
	urlGetAllPlaces = urlRoot + "all_places"
	urlGetGraph     = "https://hanetwi.ddns.net/bikeshare/graph?"
)

//NewApiClient コンストラクタ
func NewApiClient() ApiClient {
	var api ApiClient
	api.Client = &http.Client{}
	return api
}

//GetPlaces 駐輪場検索
func (api ApiClient) GetPlaces(option SearchPlacesOption) ([]SpotInfo, error) {
	url := urlGetPlaces + option.GetQuery()
	resp, err := api.Client.Get(url)
	if err != nil {
		return []SpotInfo{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	var data JPlacesBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return []SpotInfo{}, err
	}
	return data.GetSpotInfoList(), nil
}

//GetCounts 台数検索
func (api ApiClient) GetCounts(option SearchCountsOption) (SpotInfo, error) {
	url := urlGetCounts + option.GetQuery()
	resp, err := api.Client.Get(url)
	if err != nil {
		return SpotInfo{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	var data JCountsBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return SpotInfo{}, err
	}
	return data.GetSpotInfo(), nil
}

//GetDistances 近いスポット検索
func (api ApiClient) GetDistances(option SearchDistanceOption) (DistanceInfo, error) {
	url := urlGetDistances + option.GetQuery()
	resp, err := api.Client.Get(url)
	if err != nil {
		return DistanceInfo{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	var data JDistancesBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return DistanceInfo{}, err
	}
	spotinfoList := data.GetSpotInfoList()
	distanceInfo := DistanceInfo{BaseLat: option.Lat, BaseLon: option.Lon}
	distances := data.GetDistanceList()
	for i, item := range spotinfoList {
		distanceInfo.Spots = append(distanceInfo.Spots, struct {
			SpotInfo SpotInfo
			Distance string
		}{SpotInfo: item, Distance: distances[i]})
	}
	return distanceInfo, nil
}

//GetAllSpotNames すべてのスポットの名前だけ検索
func (api ApiClient) GetAllSpotNames() ([]SpotName, error) {
	resp, err := api.Client.Get(urlGetAllPlaces)
	if err != nil {
		return []SpotName{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	var data JAllPlacesBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return []SpotName{}, err
	}

	var names []SpotName
	for _, item := range data.Items {
		names = append(names, SpotName{Area: item.Area, Spot: item.Spot, Name: item.Name})
	}
	return names, nil
}

//GetGraph グラフ検索
func (api ApiClient) GetGraph(option SearchGraphOption) (GraphInfo, error) {
	var data JGraphResponse
	var graph GraphInfo
	url := urlGetGraph + option.GetQuery()
	resp, err := api.Client.Get(url)
	if err != nil {
		return graph, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return graph, err
	}
	height, _ := strconv.Atoi(data.Height)
	width, _ := strconv.Atoi(data.Width)
	graph = GraphInfo{Title: data.Title, Height: height, Width: width, URL: data.URL}
	return graph, nil
}
