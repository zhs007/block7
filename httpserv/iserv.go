package block7serv

import "github.com/zhs007/block7"

// IService - service
type IService interface {
	// GetConfig - get configuation
	GetConfig() *Config
	// Login - login
	Login(params *LoginParams) (*LoginResult, error)
	// Mission - get mission
	Mission(params *MissionParams) (*MissionResult, error)
	// MissionData - upload mission data
	MissionData(params *MissionDataParams) (*MissionDataResult, error)
	// GetUserData - get UserData
	GetUserData(params *UserDataParams) (*UserDataResult, error)
	// UpdUserData - update UserData
	UpdUserData(ud *UpdUserDataParams, uds *block7.UpdUserDataStatus) (*UpdUserDataResult, error)
	// Stats - statistics
	Stats(params *StatsParams) (*StatsResult, error)
	// Start - start
	Start()
	// Stop - stop
	Stop()
}
