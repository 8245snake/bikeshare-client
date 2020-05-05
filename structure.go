package bikeshareapi

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

//JAllPlacesBody JSONマーシャリング用
type JAllPlacesBody struct {
	Num   int `json:"num"`
	Items []struct {
		Area string `json:"area"`
		Spot string `json:"spot"`
		Name string `json:"name"`
	} `json:"items"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//  リクエストパラメータ構造体
/////////////////////////////////////////////////////////////////////////////////////////////////////////

//SearchPlacesOption 駐輪場検索
type SearchPlacesOption struct {
	Area, Spot, Query string
}

//SearchCountsOption 台数検索
type SearchCountsOption struct {
	Area, Spot, Day string
}

//SearchDistanceOption 近いスポット検索
type SearchDistanceOption struct {
	Lat, Lon float32
}
