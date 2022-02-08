package common

import (
	"auction/logs"
	"auction/models"
)

func getLastRank() ([]*models.Holder, error) {
	holders, err := models.FindLastHoldersRank()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return holders, nil
}

func getCurrentRank() ([]*models.Holder, error) {
	holders, err := models.FindCurrentHoldersRank()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return holders, nil
}

func getStatisticsRecomm(holderAddress string) (int64, error) {
	count, err := models.FindRecomm(holderAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return count, nil
}

func postStatisticsRecomm(holderAddress, recommAddress string) error {
	err := models.InsertRecomm(holderAddress, recommAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

func getRefreshTime() (int64, error) {
	refreshTime, err := models.FindRefreshTime()
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return refreshTime, nil
}
