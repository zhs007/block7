package block7serv

import "github.com/zhs007/block7"

// MissionParams - mission parameters
type MissionParams struct {
	MissionID int `json:"missionid"`
}

// MissionResult - mission result
type MissionResult struct {
	Scene       *block7.Scene `json:"scene"`
	MissionHash string        `json:"mission"`
}

// MissionDataParams - missionData parameters
type MissionDataParams struct {
	MissionHash string  `json:"mission"`
	History     [][]int `json:"history"`
}

// MissionDataResult - missionData result
type MissionDataResult struct {
	UserLevel int `json:"userLevel"`
}

// VersionParams - version parameters
type VersionParams struct {
	History [][]int `json:"history"`
}

// VersionResult - version result
type VersionResult struct {
	UserLevel int `json:"userLevel"`
}
