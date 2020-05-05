package bikeshareapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ApiClient struct {
	Client *http.Client
}

const urlRoot = "https://hanetwi.ddns.net/bikeshare/api/v1/"
const urlGetPlaces = urlRoot + "places?"
const urlGetCounts = urlRoot + "counts?"
const urlGetDistances = urlRoot + "distances?"

//NewApiClient コンストラクタ
func NewApiClient() ApiClient {
	var api ApiClient
	api.Client = &http.Client{}
	return api
}

//GetPlaces 駐輪場検索
func (api ApiClient) GetPlaces(option SearchPlacesOption) (JPlacesBody, error) {
	//リクエストBody作成
	values := url.Values{}
	values.Set("area", option.Area)
	values.Set("spot", option.Spot)
	values.Set("q", option.Query)

	req, err := http.NewRequest(
		"GET",
		urlGetPlaces,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return JPlacesBody{}, err
	}
	// リクエストHead作成
	ContentLength := strconv.FormatInt(req.ContentLength, 10)
	req.Header.Set("Content-Length", ContentLength)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := api.Client.Do(req)
	if err != nil {
		return JPlacesBody{}, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	var data JPlacesBody
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return JPlacesBody{}, err
	}

	return data, nil
}

func test() {
	api := NewApiClient()
	api.GetPlaces(struct {
		Area  string
		Spot  string
		Query string
	}{Area: "D1"})
}
