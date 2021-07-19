package block7serv

import block7game "github.com/zhs007/block7/game"

// MissionParams - mission parameters
type MissionParams struct {
	UserHash  string `json:"userHash"`
	MissionID int    `json:"missionid"`
}

// MissionResult - mission result
type MissionResult struct {
	Scene   *block7game.Scene `json:"scene"`
	SceneID int64             `json:"mission"`
}

// MissionDataParams - missionData parameters
type MissionDataParams struct {
	UserHash string  `json:"userHash"`
	SceneID  int64   `json:"mission"`
	History  [][]int `json:"history"`
}

// MissionDataResult - missionData result
type MissionDataResult struct {
	UserLevel int `json:"userLevel"`
}

// LoginParams - login parameters
type LoginParams struct {
	UserHash        string `json:"userHash"`
	Game            string `json:"game"`
	Platform        string `json:"platform"`
	ADID            string `json:"adid"`
	GUID            string `json:"guid"`
	PlatformInfo    string `json:"platformInfo"`
	GameVersion     string `json:"gameVersion"`
	ResourceVersion string `json:"resVersion"`
	DeviceInfo      string `json:"deviceInfo"`
}

// LoginResult - login result
type LoginResult struct {
	UserID   int64  `json:"uid"`
	UserHash string `json:"userHash"`
}
