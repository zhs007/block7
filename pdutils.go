package block7

import (
	"github.com/zhs007/block7/block7pb"
	"github.com/zhs007/goutils"
)

func MergeUserData(ud0 *block7pb.UserData, ud1 *block7pb.UserData, uds *UpdUserDataStatus) *block7pb.UserData {
	if ud0 == nil {
		ud1.CreateTs = goutils.GetCurTimestamp()

		return ud1
	}

	if ud1.Version <= ud0.Version {
		return ud0
	}

	ud := &block7pb.UserData{
		Name:        ud0.Name,
		Platform:    ud0.Platform,
		Version:     ud1.Version,
		LastAwardTs: ud1.LastAwardTs,
	}

	if uds.HasCoin {
		ud.Coin = ud1.Coin
	} else {
		ud.Coin = ud0.Coin
	}

	if uds.HasCooking {
		for _, v := range ud1.Cooking {
			c := &block7pb.Cooking{
				Level:    int32(v.Level),
				Unlock:   v.Unlock,
				StarNums: int32(v.StarNums),
			}

			ud.Cooking = append(ud.Cooking, c)
		}
	} else {
		for _, v := range ud0.Cooking {
			c := &block7pb.Cooking{
				Level:    int32(v.Level),
				Unlock:   v.Unlock,
				StarNums: int32(v.StarNums),
			}

			ud.Cooking = append(ud.Cooking, c)
		}
	}

	if uds.HasLevel {
		ud.Level = ud1.Level
	} else {
		ud.Level = ud0.Level
	}

	if uds.HasHomeScene {
		ud.HomeScene = append(ud.HomeScene, ud1.HomeScene...)
	} else {
		ud.HomeScene = append(ud.HomeScene, ud0.HomeScene...)
	}

	ud.LevelArr = make(map[string]int32)

	for k, v := range ud0.LevelArr {
		ud.LevelArr[k] = v
	}

	if uds.HasLevelArr {
		for k, v := range ud1.LevelArr {
			ud.LevelArr[k] = v
		}
	}

	ud.ToolsArr = make(map[string]int32)

	for k, v := range ud0.ToolsArr {
		ud.ToolsArr[k] = v
	}

	if uds.HasToolsArr {
		for k, v := range ud1.ToolsArr {
			ud.ToolsArr[k] = v
		}
	}

	if ud1.UserID > 0 {
		ud.UserID = ud1.UserID
	} else {
		ud.UserID = ud0.UserID
	}

	if ud1.ClientVersion != "" {
		ud.ClientVersion = ud1.ClientVersion
	} else {
		ud.ClientVersion = ud0.ClientVersion
	}

	return ud
}

func SetHistoryInDayStatsData(dsd *block7pb.DayStatsData, hds *HistoryDBDayStatsData) {
	dsd.HistoryNums = int32(hds.HistoryNums)
	dsd.HistoryGameStateNums = goutils.MapII2MapI32I32(hds.GameStateNums)
	dsd.HistoryUserIDNums = goutils.MapI64I2MapI64I32(hds.UserIDNums)

	if len(hds.Stages) > 0 {
		dsd.HistoryStages2 = make(map[int32]*block7pb.HistoryStageData)

		for k, v := range hds.Stages {
			dsd.HistoryStages2[int32(k)] = &block7pb.HistoryStageData{
				Nums:          int32(v.Nums),
				GameStateNums: goutils.MapII2MapI32I32(v.GameStateNums),
			}
		}
	}
}

func SetUserInDayStatsData(dsd *block7pb.DayStatsData, uds *UserDBDayStatsData) {
	dsd.NewUserNums = int32(uds.NewUserNums)
	dsd.AliveUserNums = int32(uds.AliveUserNums)
	dsd.FirstUserID = uds.FirstUserID
	dsd.NewUserDataNums = int32(uds.NewUserDataNums)
	dsd.AliveUserDataNums = int32(uds.AliveUserDataNums)
	dsd.FirstUserDataUID = uds.FirstUserDataUID

	if len(uds.Users2) > 0 {
		dsd.Users2 = make(map[int64]*block7pb.UserDayStatsData)

		for k, v := range uds.Users2 {
			udsd := &block7pb.UserDayStatsData{
				UserHash: v.UserHash,
				Stages:   make(map[int32]*block7pb.UserStageData),
			}

			for k1, v1 := range v.Stages {
				cusd := &block7pb.UserStageData{
					GameStateNums: make(map[int32]int32),
				}

				for k2, v2 := range v1.GameStateNums {
					cusd.GameStateNums[int32(k2)] = int32(v2)
				}

				udsd.Stages[int32(k1)] = cusd
			}

			dsd.Users2[k] = udsd
		}
	}
}
