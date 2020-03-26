package _type

// BarrageCount:统计返回弹幕数
type BarrageCount struct {
	Count int `json:"count"`
}

// BarrageStatisticsCountResult:统计弹幕频率
type BarrageStatisticsCountResult struct {
	Count int    `json:"count"`
	Txt   string `json:"txt"`
}

// BarrageStatisticsUserCountResult:统计用户发送弹幕
type BarrageStatisticsUserCountResult struct {
	Count    int    `json:"count"`
	UserName string `json:"nn"`
	UserId   int    `json:"uid"`
}
