package block7serv

// IService - service
type IService interface {
	// Mission - get mission
	Mission(params *MissionParams) (*MissionResult, error)
	// MissionData - upload mission data
	MissionData(params *MissionDataParams) (*MissionDataResult, error)
}
