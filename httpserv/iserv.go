package block7serv

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
}
