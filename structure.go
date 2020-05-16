package bikeshareapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//JsonTimeLayout 時刻フォーマット
const JsonTimeLayout = "2006/01/02 15:04"

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//  JSONマーシャリング構造体
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//JCountsBody JSONマーシャリング用
type JCountsBody struct {
	Area        string `json:"area"`
	Spot        string `json:"spot"`
	Description string `json:"description"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Name        string `json:"name"`
	Counts      []struct {
		Count    string `json:"count"`
		Datetime string `json:"datetime"`
		Day      string `json:"day"`
		Hour     string `json:"hour"`
		Minute   string `json:"minute"`
		Month    string `json:"month"`
		Year     string `json:"year"`
	} `json:"counts"`
}

//GetSpotInfo GetSpotInfo構造体を返す
func (body JCountsBody) GetSpotInfo() SpotInfo {
	var spotinfo SpotInfo
	spotinfo.Area = body.Area
	spotinfo.Spot = body.Spot
	spotinfo.Name = body.Name
	spotinfo.Description = body.Description
	if lat, err := strconv.ParseFloat(body.Lat, 64); err == nil {
		spotinfo.Lat = lat
	}
	if lon, err := strconv.ParseFloat(body.Lon, 64); err == nil {
		spotinfo.Lon = lon
	}
	for _, item := range body.Counts {
		if s, err := NewBikeCount(item.Datetime, item.Count); err == nil {
			spotinfo.Counts = append(spotinfo.Counts, s)
		}
	}
	return spotinfo
}

//JPlacesBody JSONマーシャリング用
type JPlacesBody struct {
	Num   int `json:"num"`
	Items []struct {
		Area        string `json:"area"`
		Spot        string `json:"spot"`
		Description string `json:"description"`
		Lat         string `json:"lat"`
		Lon         string `json:"lon"`
		Name        string `json:"name"`
		Recent      Recent `json:"recent"`
	} `json:"items"`
}

//GetSpotInfoList GetSpotInfo構造体を返す
func (body JPlacesBody) GetSpotInfoList() []SpotInfo {
	var spotinfoList []SpotInfo
	for _, item := range body.Items {
		var spotinfo SpotInfo
		spotinfo.Area = item.Area
		spotinfo.Spot = item.Spot
		spotinfo.Name = item.Name
		spotinfo.Description = item.Description
		if lat, err := strconv.ParseFloat(item.Lat, 64); err == nil {
			spotinfo.Lat = lat
		}
		if lon, err := strconv.ParseFloat(item.Lon, 64); err == nil {
			spotinfo.Lon = lon
		}
		if count, err := NewBikeCount(item.Recent.Datetime, item.Recent.Count); err == nil {
			spotinfo.Counts = append(spotinfo.Counts, count)
		}
		spotinfoList = append(spotinfoList, spotinfo)
	}
	return spotinfoList
}

//Recent 最新の台数情報を格納する
type Recent struct {
	Count    string `json:"count"`
	Datetime string `json:"datetime"`
}

//JDistancesBody JSONマーシャリング用
type JDistancesBody struct {
	Num   int `json:"num"`
	Items []struct {
		Area        string `json:"area"`
		Spot        string `json:"spot"`
		Description string `json:"description"`
		Lat         string `json:"lat"`
		Lon         string `json:"lon"`
		Name        string `json:"name"`
		Distance    string `json:"distance"`
		Recent      Recent `json:"recent"`
	} `json:"items"`
}

//GetSpotInfoList GetSpotInfo構造体を返す
func (body JDistancesBody) GetSpotInfoList() []SpotInfo {
	var spotinfoList []SpotInfo
	for _, item := range body.Items {
		var spotinfo SpotInfo
		spotinfo.Area = item.Area
		spotinfo.Spot = item.Spot
		spotinfo.Name = item.Name
		spotinfo.Description = item.Description
		if lat, err := strconv.ParseFloat(item.Lat, 64); err == nil {
			spotinfo.Lat = lat
		}
		if lon, err := strconv.ParseFloat(item.Lon, 64); err == nil {
			spotinfo.Lon = lon
		}
		if count, err := NewBikeCount(item.Recent.Datetime, item.Recent.Count); err == nil {
			spotinfo.Counts = append(spotinfo.Counts, count)
		}
		spotinfoList = append(spotinfoList, spotinfo)
	}
	return spotinfoList
}

//GetDistanceList GetSpotInfo構造体を返す
func (body JDistancesBody) GetDistanceList() []string {
	var distances []string
	for _, item := range body.Items {
		distances = append(distances, item.Distance)
	}
	return distances
}

//JAllPlacesBody JSONマーシャリング用
type JAllPlacesBody struct {
	Num   int `json:"num"`
	Items []struct {
		Area string `json:"area"`
		Spot string `json:"spot"`
		Name string `json:"name"`
	} `json:"items"`
}

//JGraphResponse Graphリクエストの返信用
type JGraphResponse struct {
	Title  string `json:"title"`
	Width  string `json:"width"`
	Height string `json:"height"`
	URL    string `json:"url"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//   共通構造体
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//BikeCount 台数
type BikeCount struct {
	Time  time.Time
	Count int
}

//NewBikeCount JSONから変換するため
func NewBikeCount(datetime string, count string) (BikeCount, error) {
	t, err := time.Parse(JsonTimeLayout, datetime)
	if err != nil {
		return BikeCount{}, err
	}
	c, err := strconv.Atoi(count)
	if err != nil {
		return BikeCount{}, err
	}
	return BikeCount{Time: t, Count: c}, nil
}

//SpotInfo 駐輪場情報
type SpotInfo struct {
	Area, Spot, Name, Description string
	Lat, Lon                      float64
	Counts                        []BikeCount
}

//DistanceInfo 指定地点からの距離情報
type DistanceInfo struct {
	BaseLat, BaseLon float64
	Spots            []struct {
		SpotInfo SpotInfo
		Distance string
	}
}

//SpotName 駐輪場の名前
type SpotName struct {
	Area, Spot, Name string
}

//GraphInfo グラフ画像の情報
type GraphInfo struct {
	Title  string
	Width  int
	Height int
	URL    string `json:"url"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//  リクエストパラメータ構造体
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//SearchPlacesOption 駐輪場検索
type SearchPlacesOption struct {
	Area, Spot, Query string
}

//GetQuery 検索条件作成
func (option SearchPlacesOption) GetQuery() string {
	var params []string
	if option.Area != "" {
		param := fmt.Sprintf("area=%s", option.Area)
		params = append(params, param)
	}
	if option.Spot != "" {
		param := fmt.Sprintf("spot=%s", option.Spot)
		params = append(params, param)
	}
	if option.Query != "" {
		param := fmt.Sprintf("q=%s", option.Query)
		params = append(params, param)
	}
	return strings.Join(params, `&`)
}

//SearchCountsOption 台数検索
type SearchCountsOption struct {
	Area, Spot, Day string
}

//GetQuery 検索条件作成
func (option SearchCountsOption) GetQuery() string {
	var params []string
	if option.Area != "" {
		param := fmt.Sprintf("area=%s", option.Area)
		params = append(params, param)
	}
	if option.Spot != "" {
		param := fmt.Sprintf("spot=%s", option.Spot)
		params = append(params, param)
	}
	if option.Day != "" {
		param := fmt.Sprintf("day=%s", option.Day)
		params = append(params, param)
	}
	return strings.Join(params, `&`)
}

//SearchDistanceOption 近いスポット検索
type SearchDistanceOption struct {
	Lat, Lon float64
}

//GetQuery 検索条件作成
func (option SearchDistanceOption) GetQuery() string {
	var params []string
	if option.Lat != 0 {
		param := fmt.Sprintf("lat=%v", option.Lat)
		params = append(params, param)
	}
	if option.Lon != 0 {
		param := fmt.Sprintf("lon=%v", option.Lon)
		params = append(params, param)
	}
	return strings.Join(params, `&`)
}

//SearchGraphOption グラフ検索
type SearchGraphOption struct {
	Area, Spot  string
	Property    string
	Days        []string
	DrawTitle   bool
	UploadImgur bool
}

//GetQuery 検索条件作成
func (option SearchGraphOption) GetQuery() string {
	var params []string
	if option.Area != "" {
		param := fmt.Sprintf("area=%s", option.Area)
		params = append(params, param)
	}
	if option.Spot != "" {
		param := fmt.Sprintf("spot=%s", option.Spot)
		params = append(params, param)
	}
	if option.Property != "" {
		param := fmt.Sprintf("property=%s", option.Property)
		params = append(params, param)
	}
	if len(option.Days) > 0 {
		param := fmt.Sprintf("days=%s", strings.Join(option.Days, `,`))
		params = append(params, param)
	}
	if option.DrawTitle {
		params = append(params, "title=yes")
	}
	if option.UploadImgur {
		params = append(params, "imgur=yes")
	}
	return strings.Join(params, `&`)
}
