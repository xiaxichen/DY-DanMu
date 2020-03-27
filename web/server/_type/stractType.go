package _type

const (
	EMAILHTMLTYPE = "html"
)

type UserSearchStruct struct {
	UserName  string `json:"username"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
	EsIndex   string `j:"esIndex"`
	From      int    `json:"from"`
}

type BarrageCountStruct struct {
	StartTime int   `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type BarrageAllStruct struct {
	From    int    `json:"from"`
	EsIndex string `j:"esIndex"`
}

type QueryAllFieldStruct struct {
	From    int    `json:"from"`
	Query   string `json:"query"`
	EsIndex string `j:"esIndex"`
}
type StatisticsBarrageStruct struct {
	From      int   `json:"from"`
	StartTime int   `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

type EmailSendStruct struct {
	UserName string
	To       []string
	Body     string
	MailType string
}
