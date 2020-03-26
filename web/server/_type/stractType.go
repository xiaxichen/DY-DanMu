package _type

type UserSearchStract struct {
	UserName  string `json:"username"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	EsIndex   string `j:"esIndex"`
	From      int    `json:"from"`
}

type BarrageCountStract struct {
	StartTime int   `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type BarrageAllStract struct {
	From    int    `json:"from"`
	EsIndex string `j:"esIndex"`
}

type QueryAllFieldStract struct {
	From    int    `json:"from"`
	Query   string `json:"query"`
	EsIndex string `j:"esIndex"`
}
type StatisticsBarrageStract struct {
	From      int   `json:"from"`
	StartTime int   `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}
