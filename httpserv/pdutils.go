package block7serv

import (
	"github.com/zhs007/block7/block7pb"
)

func LoginParams2PB(params *LoginParams) *block7pb.UserDeviceInfo {
	udi := &block7pb.UserDeviceInfo{
		UserHash:        params.UserHash,
		Game:            params.Game,
		Platform:        params.Platform,
		ADID:            params.ADID,
		GUID:            params.GUID,
		PlatformInfo:    params.PlatformInfo,
		GameVersion:     params.GameVersion,
		ResourceVersion: params.ResourceVersion,
		DeviceInfo:      params.DeviceInfo,
	}

	return udi
}

func UpdUserDataParams2PB(params *UpdUserDataParams, uds *UpdUserDataStatus) *block7pb.UserData {
	ud := &block7pb.UserData{
		Name:     params.Name,
		Platform: params.Platform,
		Version:  params.Version,
	}

	if uds.HasCoin {
		ud.Coin = params.Coin
	}

	if uds.HasCooking {
		for _, v := range params.Cooking {
			c := &block7pb.Cooking{
				Level:    int32(v.Level),
				Unlock:   v.Unlock,
				StarNums: int32(v.StarNums),
			}

			ud.Cooking = append(ud.Cooking, c)
		}
	}

	if uds.HasLevel {
		ud.Level = int32(params.Level)
	}

	if uds.HasHomeScene {
		for _, v := range params.HomeScene {
			ud.HomeScene = append(ud.HomeScene, int32(v))
		}
	}

	if uds.HasLevelArr {
		for k, v := range params.LevelArr {
			ud.LevelArr[int32(k)] = int32(v)
		}
	}

	if uds.HasToolsArr {
		for k, v := range params.ToolsArr {
			ud.ToolsArr[int32(k)] = int32(v)
		}
	}

	return ud
}
