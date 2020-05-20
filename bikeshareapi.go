package bikeshareapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/8245snake/bikeshare_api/src/lib/static"
)

//ApiClient クライアント
type ApiClient struct {
	Client  *http.Client
	CertKey string
}

const (
	urlRoot         = "https://hanetwi.ddns.net/bikeshare/api/v1/"
	urlGetPlaces    = urlRoot + "places?"
	urlGetCounts    = urlRoot + "counts?"
	urlGetDistances = urlRoot + "distances?"
	urlGetAllPlaces = urlRoot + "all_places"
	urlGetUser      = urlRoot + "private/users"
	urlGetGraph     = "https://hanetwi.ddns.net/bikeshare/graph?"
)

//NewApiClient コンストラクタ
func NewApiClient() ApiClient {
	var api ApiClient
	api.Client = &http.Client{}
	return api
}

//SetCertKey キーを設定
func (api *ApiClient) SetCertKey(certKey string) {
	api.CertKey = certKey
}

//SendGetRequest GETリクエストを送信してレスポンスのバイト配列を得る
func (api *ApiClient) SendGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("cert", api.CertKey)
	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}

//GetPlaces 駐輪場検索
func (api *ApiClient) GetPlaces(option SearchPlacesOption) ([]SpotInfo, error) {
	url := urlGetPlaces + option.GetQuery()
	byteArray, err := api.SendGetRequest(url)
	if err != nil {
		return []SpotInfo{}, err
	}
	var data static.JPlacesBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		return []SpotInfo{}, err
	}
	return GetSpotInfoListByPlaces(data), nil
}

//GetCounts 台数検索
func (api ApiClient) GetCounts(option SearchCountsOption) (SpotInfo, error) {
	url := urlGetCounts + option.GetQuery()
	byteArray, err := api.SendGetRequest(url)
	if err != nil {
		return SpotInfo{}, err
	}
	var data static.JCountsBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return SpotInfo{}, err
	}
	return GetSpotInfoByJCount(data), nil
}

//GetDistances 近いスポット検索
func (api ApiClient) GetDistances(option SearchDistanceOption) (DistanceInfo, error) {
	url := urlGetDistances + option.GetQuery()
	byteArray, err := api.SendGetRequest(url)
	if err != nil {
		return DistanceInfo{}, err
	}
	var data static.JDistancesBody
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return DistanceInfo{}, err
	}
	spotinfoList := GetSpotInfoListByDistance(data)
	distanceInfo := DistanceInfo{BaseLat: option.Lat, BaseLon: option.Lon}
	distances := GetDistanceList(data)
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
	byteArray, err := api.SendGetRequest(urlGetAllPlaces)
	if err != nil {
		return []SpotName{}, err
	}
	var data static.JAllPlacesBody
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
	url := urlGetGraph + option.GetQuery()
	byteArray, err := api.SendGetRequest(url)
	if err != nil {
		return GraphInfo{}, err
	}
	var data static.JGraphResponse
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return GraphInfo{}, err
	}
	return GetGraphInfoByJGraphResponse(data), nil
}

//GetUsers ユーザ情報取得
func (api ApiClient) GetUsers() ([]Users, error) {
	byteArray, err := api.SendGetRequest(urlGetUser)
	if err != nil {
		return nil, err
	}
	var data static.JUsers
	if err := json.Unmarshal(([]byte)(byteArray), &data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	var users []Users
	for _, jUser := range data.Users {
		users = append(users,
			Users{
				LineID:    jUser.LineID,
				SlackID:   jUser.SlackID,
				Favorites: jUser.Favorites,
				Histories: jUser.Histories,
				Notifies:  jUser.Notifies,
			},
		)
	}
	return users, nil
}
