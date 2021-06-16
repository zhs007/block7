package block7serv

import "github.com/zhs007/block7"

// MissionParams - mission parameters
type MissionParams struct {
	MissionID int `json:"missionid"`
}

// MissionResult - mission result
type MissionResult struct {
	Scene *block7.Scene `json:"scene"`
}

// MissionDataParams - missionData parameters
type MissionDataParams struct {
	History [][]int `json:"history"`
}

// MissionDataResult - missionData result
type MissionDataResult struct {
	UserLevel int `json:"userLevel"`
}
