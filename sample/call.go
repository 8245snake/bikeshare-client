package main

import (
	"fmt"

	bikeshareapi "github.com/8245snake/bikeshare-client"
)

//API クライアント
var API bikeshareapi.ApiClient

//GetPlaces D1エリア（新宿区あたり）の駐輪場を検索
func GetPlaces() {
	spots, err := API.GetPlaces(bikeshareapi.SearchPlacesOption{Area: "D1"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", spots)
}

//GetCounts 千代田区役所の2020年5月5日の台数を検索
func GetCounts() {
	counts, err := API.GetCounts(bikeshareapi.SearchCountsOption{Area: "A1", Spot: "01", Day: "20200505"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", counts)
}

//GetDistances 新宿駅から近いスポットを１０件取得
func GetDistances() {
	distances, err := API.GetDistances(bikeshareapi.SearchDistanceOption{Lat: 35.689274, Lon: 139.700646})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", distances)
}

//GetAllSpotNames 現在有効な駐輪場の名前とコードの一覧を取得
func GetAllSpotNames() {
	names, err := API.GetAllSpotNames()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", names)
}

//GetGraph 千代田区役所
func GetGraph() {
	option := bikeshareapi.SearchGraphOption{
		Area:      "A1",
		Spot:      "01",
		DrawTitle: true,
	}
	graph, err := API.GetGraph(option)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", graph)
}

//GetUsers 千代田区役所
func GetUsers() {
	users, err := API.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", users)
}

func main() {
	//APIクライアント初期化
	API = bikeshareapi.NewApiClient()
	GetPlaces()
	GetAllSpotNames()
	GetCounts()
	GetDistances()
	GetGraph()
	GetUsers()
}
