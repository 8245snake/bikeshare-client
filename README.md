# bikeshare-client

## 概要
docomoシェアサイクル台数検索APIを叩くGo言語ライブラリ  
以下のAPI仕様の通りにリクエストを送信する  
https://hanetwi.ddns.net/bikeshare/reference.html

## 導入
以下のコマンドでインストールする  
```
go get -u github.com/8245snake/bikeshare_api/src/lib/static  
go get -u github.com/8245snake/bikeshare-client
```  

## 使い方
`/sample/call.go` の実装を参考にする  
  
以下のコードで初期化する

```
// APIクライアント初期化
API = bikeshareapi.NewApiClient()
// 秘密文字列を設定（非公開API用）
API.SetCertKey("your cert key") //your cert key
 // デバッグ用エンドポイント設定
API.SetEndpoint("http://localhost:5001/")
```  
  
いろいろな情報を取得  
```
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
	counts, err := API.GetCounts(bikeshareapi.SearchCountsOption{Area: "A1", Spot: "01", Day: "20200522"})
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

//CheckStatus システム正常性チェック
func CheckStatus() {
	status, err := API.GetStatus()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", status)
}
```
