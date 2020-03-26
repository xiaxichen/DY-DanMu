package route

import (
	"DY-DanMu/web/server/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//gin framewoke。包括Logger,Recovery
	router := gin.Default()
	router.POST("/search/user", handler.Wrapper(handler.SearchUserBarrage))
	router.POST("/search/all", handler.Wrapper(handler.SearchBarrageAll))
	router.POST("/search", handler.Wrapper(handler.SearchAllField))
	router.POST("/search/Count", handler.Wrapper(handler.SearchBarrageCount))
	router.POST("/search/word_cloud", handler.Wrapper(handler.StatisticsBarrageForTime))
	router.POST("/search/user_count_top", handler.Wrapper(handler.StatisticsUserBarrageForTime))
	//router.POST("/search/count", handler.SearchBarrageCount)
	return router
}
