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
