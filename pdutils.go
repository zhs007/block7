package block7

import (
	"github.com/zhs007/block7/block7pb"
)

func MergeUserData(ud0 *block7pb.UserData, ud1 *block7pb.UserData, uds *UpdUserDataStatus) *block7pb.UserData {
	if ud0 == nil {
		return ud1
	}

	if ud1.Version <= ud0.Version {
		return ud0
	}

	ud := &block7pb.UserData{
		Name:     ud0.Name,
		Platform: ud0.Platform,
		Version:  ud1.Version,
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

	for k, v := range ud0.LevelArr {
		ud.LevelArr[k] = v
	}

	if uds.HasLevelArr {
		for k, v := range ud1.LevelArr {
			ud.LevelArr[k] = v
		}
	}

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
