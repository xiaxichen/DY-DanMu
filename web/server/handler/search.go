package handler

import (
	"DY-DanMu/DMconfig/config"
	_type2 "DY-DanMu/dbConn/_type"
	"DY-DanMu/persistServer/server"
	"DY-DanMu/web/client"
	"DY-DanMu/web/server/_type"
	"DY-DanMu/web/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SearchUserBarrage:查询用户发送弹幕
func SearchUserBarrage(ctx *gin.Context) error {
	var data _type.UserSearchStruct
	err := ctx.BindJSON(&data)
	if err != nil || data.UserName == "" {
		return ParameterError("Post Data Err")
	}
	if data.EndTime == 0 {
		data.EndTime = int(time.Now().UnixNano() / 1e6)
	}
	result := server.UserBarrageResult{
		ResultList: nil,
		Hits:       0,
	}
	data.EsIndex = config.DYWebConfig.ElasticIndex
	err = client.ClientRPC.Call(config.DYWebConfig.UserSearch, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.UserSearch, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": result, "msg": ""})
	return nil
}

// SearchBarrageCount:获取指定时间内的弹幕总数
func SearchBarrageCount(ctx *gin.Context) error {
	var data _type.BarrageCountStruct
	err := ctx.BindJSON(&data)
	if err != nil || data.EndTime == 0 {
		ctx.JSON(200, gin.H{"code": 400, "msg": "Post Data Err"})
		return ParameterError("Post Data Err")
	}
	result := _type2.BarrageCount{
		Count: 0,
	}
	err = client.ClientRPC.Call(config.DYWebConfig.IndexBarrageCount, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.IndexBarrageCount, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": result, "msg": ""})
	return nil
}

// SearchBarrageAll:获取全部弹幕时间倒排序 每页10条
func SearchBarrageAll(ctx *gin.Context) error {
	var data _type.BarrageAllStruct
	err := ctx.BindJSON(&data)
	if err != nil {
		return ParameterError("Post Data Err")
	}
	result := server.UserBarrageResult{
		ResultList: nil,
		Hits:       0,
	}
	data.EsIndex = config.DYWebConfig.ElasticIndex
	err = client.ClientRPC.Call(config.DYWebConfig.BarrageAll, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.BarrageAll, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": result, "msg": ""})
	return nil
}

// SearchAllField:检索所有字段 每页10条
func SearchAllField(ctx *gin.Context) error {
	var data _type.QueryAllFieldStruct
	err := ctx.BindJSON(&data)
	if err != nil {
		return ParameterError("Post Data Err")
	}
	result := server.UserBarrageResult{
		ResultList: nil,
		Hits:       0,
	}
	data.EsIndex = config.DYWebConfig.ElasticIndex
	err = client.ClientRPC.Call(config.DYWebConfig.SearchAllField, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.SearchAllField, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": result, "msg": ""})
	return nil
}

//StatisticsBarrageForTime:根据时间统计弹幕频率
func StatisticsBarrageForTime(ctx *gin.Context) error {
	var data _type.StatisticsBarrageStruct
	err := ctx.BindJSON(&data)
	if err != nil || data.From == 0 {
		return ParameterError("Post Data Err")
	}
	result := []_type2.BarrageStatisticsCountResult{}
	err = client.ClientRPC.Call(config.DYWebConfig.StatisticsBarrage, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.StatisticsBarrage, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"cod": 200, "data": result, "msg": ""})
	return nil
}

//StatisticsUserBarrageForTime:根据时间统计用户发送弹幕
func StatisticsUserBarrageForTime(ctx *gin.Context) error {
	var data _type.StatisticsBarrageStruct
	err := ctx.BindJSON(&data)
	if err != nil || data.From == 0 {
		return ParameterError("Post Data Err")
	}
	result := []_type2.BarrageStatisticsUserCountResult{}
	err = client.ClientRPC.Call(config.DYWebConfig.StatisticsUserBarrage, data, &result)
	if err != nil {
		err = util.RpcClientShutDownErrorhandler(err)
		if err == nil {
			err = client.ClientRPC.Call(config.DYWebConfig.StatisticsUserBarrage, data, &result)
		}
	}
	if err != nil {
		return ServerError()
	}
	ctx.JSON(http.StatusOK, gin.H{"cod": 200, "data": result, "msg": ""})
	return nil
}
