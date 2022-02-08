package common

import (
	"net/http"

	"auction/constant"
	"auction/logs"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constant.API_GET_LAST_RANK, GetLastRank)
	router.GET(constant.API_GET_CURRENT_RANK, GetCurrentRank)
	router.GET(constant.API_GET_STATISTICS_RECOMM, GetStatisticsRecomm)
	router.GET(constant.API_GET_REFRESHTIME, GetRefreshTime)
	router.POST(constant.API_POST_STATISTICS_RECOMM, PostStatisticsRecomm)

}

func GetLastRank(c *gin.Context) {
	holders, err := getLastRank()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, constant.CreateErrorResponse(constant.GET_RECORD_lIST_ERROR_MSG, err.Error()))
		return
	}
	c.JSON(http.StatusOK, constant.CreateSuccessResponse(holders))
}

func GetCurrentRank(c *gin.Context) {
	holders, err := getCurrentRank()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, constant.CreateErrorResponse(constant.GET_RECORD_lIST_ERROR_MSG, err.Error()))
		return
	}
	c.JSON(http.StatusOK, constant.CreateSuccessResponse(holders))
}

func GetStatisticsRecomm(c *gin.Context) {
	holderAddress := c.Param("holder")
	if len(holderAddress) != 42 {
		c.JSON(http.StatusBadRequest, constant.CreateErrorResponse(constant.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, constant.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	count, err := getStatisticsRecomm(holderAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, constant.CreateErrorResponse(constant.GET_RECORD_lIST_ERROR_MSG, err.Error()))
		return
	}
	c.JSON(http.StatusOK, constant.CreateSuccessResponse(count))
}

func PostStatisticsRecomm(c *gin.Context) {
	holderAddress := c.Param("holder")
	recommAddress := c.Param("recomm")
	if len(holderAddress) != 42 && len(recommAddress) != 42 {
		c.JSON(http.StatusBadRequest, constant.CreateErrorResponse(constant.HTTP_REQUEST_PARAM_VALUE_ERROR_CODE, constant.HTTP_REQUEST_PARAM_VALUE_ERROR_MSG))
		return
	}
	err := postStatisticsRecomm(holderAddress, recommAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, constant.CreateErrorResponse(constant.GET_RECORD_lIST_ERROR_MSG, err.Error()))
		return
	}
	c.JSON(http.StatusOK, constant.CreateSuccessResponse(nil))
}

func GetRefreshTime(c *gin.Context) {
	refreshTime, err := getRefreshTime()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, constant.CreateErrorResponse(constant.GET_RECORD_lIST_ERROR_MSG, err.Error()))
		return
	}
	c.JSON(http.StatusOK, constant.CreateSuccessResponse(refreshTime))
}
