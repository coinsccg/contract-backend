package models

import (
	"auction/constant"
	"auction/db"
	"auction/logs"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type Holder struct {
	ID              int64  `json:"id"`
	Stage           int64  `json:"stage"`            // 期数
	HolderAddress   string `json:"holder_address"`   // 持币地址
	LastReward      string `json:"last_reward"`      // 最新奖励
	LastQuantity    string `json:"last_quantity"`    // 上一次持币量
	CurrentQuantity string `json:"current_quantity"` // 当前持币量
	RefreshTime     int64  `json:"refresh_time"`
}

type Recomm struct {
	ID            int64  `json:"id"`
	HolderAddress string `json:"holder_address"`
	RecommAddress string `json:"recomm_address"`
}

type Resp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []List `json:"result"`
}

type List struct {
	TokenHolderAddress  string `json:"TokenHolderAddress"`
	TokenHolderQuantity string `json:"TokenHolderQuantity"`
}

type RefreshTime struct {
	Stage       int64 `json:"stage"`
	RefreshTime int64 `json:"refresh_time"`
}

func InsertHolder(resp *Resp) error {
	var (
		startHolder Holder
		stage       int64 = 1
	)
	db := database.GetDB()

	if err := db.Table("holders").Limit(1).Find(&startHolder).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return err
	}
	if startHolder.Stage != 0 {
		stage = startHolder.Stage
	}

	for _, v := range resp.Result {
		var n = 0
		for _, v1 := range constant.ADDRESS_LIST {
			if v.TokenHolderAddress == strings.ToLower(v1) {
				n++
			}
		}
		if n > 0 {
			continue
		}
		var holder = Holder{
			Stage:           stage,
			HolderAddress:   v.TokenHolderAddress,
			CurrentQuantity: v.TokenHolderQuantity,
			LastReward:      "0",
			LastQuantity:    "0",
			RefreshTime:     time.Now().Unix(),
		}
		var count int64
		err := db.Table("holders").Where("holder_address = ?", v.TokenHolderAddress).Count(&count).Error
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if count == 0 {
			// insert
			if err := db.Table("holders").Create(&holder).Error; err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		} else {
			// update
			if err := db.Table("holders").Where("holder_address = ?", v.TokenHolderAddress).
				Update("current_quantity", v.TokenHolderQuantity).Error; err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}

	return nil
}

func FindLastHoldersRank() ([]*Holder, error) {
	var holders []*Holder

	db := database.GetDB()
	if err := db.Table("holders").Order("last_reward desc").Limit(20).Find(&holders).Error; err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return holders, nil
}

func FindCurrentHoldersRank() ([]*Holder, error) {
	var holders []*Holder

	db := database.GetDB()
	if err := db.Table("holders").
		Select("id, holder_address, (current_quantity - last_quantity) as last_reward, last_quantity, current_quantity, refresh_time, stage").
		Order("last_reward desc").Limit(20).Find(&holders).Error; err != nil && err != gorm.ErrRecordNotFound {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return holders, nil
}

func UpdateLastReward(holders []*Holder) (err error) {
	db := database.GetDB()
	for _, v := range holders {
		if err = db.Table("holders").Where("holder_address = ?", v.HolderAddress).
			Updates(map[string]interface{}{
				"stage":         gorm.Expr("stage + ?", 1),
				"last_reward":   v.LastReward,
				"last_quantity": v.CurrentQuantity,
			}).Error; err != nil {
			logs.GetLogger().Error(err)
			return
		}
	}
	return
}

func UpdateRefreshTime(refresh int64) (err error) {
	db := database.GetDB()
	var (
		count       int64
		refreshTime = RefreshTime{
			Stage:       1,
			RefreshTime: refresh,
		}
	)
	err = db.Table("refresh_times").Count(&count).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if count == 0 {
		// insert
		if err = db.Table("refresh_times").Create(&refreshTime).Error; err != nil {
			logs.GetLogger().Error(err)
			return
		}
	} else {
		if err = db.Table("refresh_times").Where("stage = ?", 1).
			Update("refresh_time", refresh).Error; err != nil {
			logs.GetLogger().Error(err)
			return
		}
	}

	return
}

func InsertRecomm(holderAddress, recommAddress string) (err error) {
	db := database.GetDB()
	var count int64
	if err = db.Table("recomms").Where("holder_address = ? AND recomm_address = ?", holderAddress, recommAddress).
		Count(&count).Error; err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if count == 0 {
		// insert
		var recomm = Recomm{
			HolderAddress: holderAddress,
			RecommAddress: recommAddress,
		}
		if err = db.Table("recomms").Create(&recomm).Error; err != nil {
			logs.GetLogger().Error(err)
			return
		}
	}
	return
}

func FindRecomm(holderAddress string) (int64, error) {
	db := database.GetDB()
	var count int64
	if err := db.Table("recomms").Where("recomm_address = ?", holderAddress).
		Count(&count).Error; err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return count, nil
}

func FindRefreshTime() (int64, error) {
	db := database.GetDB()
	var refreshTime RefreshTime
	if err := db.Table("refresh_times").Where("stage = ?", 1).
		Find(&refreshTime).Error; err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return refreshTime.RefreshTime, nil
}
