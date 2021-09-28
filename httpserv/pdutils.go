package block7serv

import (
	"github.com/zhs007/block7"
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
		IPAddr:          params.IPAddr,
	}

	return udi
}

func UpdUserDataParams2PB(params *UpdUserDataParams, uds *block7.UpdUserDataStatus) *block7pb.UserData {
	ud := &block7pb.UserData{
		Name:          params.Name,
		Platform:      params.Platform,
		Version:       params.Version,
		ClientVersion: params.ClientVersion,
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
		ud.LevelArr = make(map[string]int32)

		for k, v := range params.LevelArr {
			ud.LevelArr[k] = int32(v)
		}
	}

	if uds.HasToolsArr {
		ud.ToolsArr = make(map[string]int32)

		for k, v := range params.ToolsArr {
			ud.ToolsArr[k] = int32(v)
		}
	}

	return ud
}

func PB2UserDataResult(ud *block7pb.UserData) *UserDataResult {
	udr := &UserDataResult{
		Name:          ud.Name,
		Coin:          ud.Coin,
		Level:         int(ud.Level),
		Platform:      ud.Platform,
		Version:       ud.Version,
		ClientVersion: ud.ClientVersion,
		LastAwardTs:   ud.LastAwardTs,
	}

	for _, v := range ud.Cooking {
		c := &Cooking{
			Level:    int(v.Level),
			Unlock:   v.Unlock,
			StarNums: int(v.StarNums),
		}

		udr.Cooking = append(udr.Cooking, c)
	}

	for _, v := range ud.HomeScene {
		udr.HomeScene = append(udr.HomeScene, int(v))
	}

	for k, v := range ud.LevelArr {
		udr.LevelArr[k] = int(v)
	}

	for k, v := range ud.ToolsArr {
		udr.ToolsArr[k] = int(v)
	}

	return udr
}
