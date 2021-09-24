package block7serv

import block7game "github.com/zhs007/block7/game"

// MissionParams - mission parameters
type MissionParams struct {
	UserHash  string `json:"userHash"`
	MissionID int    `json:"missionid"`
	SceneID   int64  `json:"mission"`
	HistoryID int64  `json:"history"`
}

// MissionResult - mission result
type MissionResult struct {
	Scene   *block7game.Scene `json:"scene"`
	SceneID int64             `json:"mission"`
}

// MissionDataParams - missionData parameters
type MissionDataParams struct {
	UserHash      string                     `json:"userHash"`
	SceneID       int64                      `json:"mission"`
	History       [][]int                    `json:"history"`
	HistoryID     int64                      `json:"srcHistory"`
	RngData       []int64                    `json:"rngdata"`
	GameState     int32                      `json:"gamestate"`
	InitArr       [][][]int                  `json:"initArr"`
	BlockNums     int                        `json:"blockNums"`
	StageType     int                        `json:"stageType"`
	SpecialLayers []*block7game.SpecialLayer `json:"specialLayers"`
	FirstItem     int                        `json:"firstItem"`
	MissionID     int                        `json:"missionid"`
}

// MissionDataResult - missionData result
type MissionDataResult struct {
	UserLevel int   `json:"userLevel"`
	HistoryID int64 `json:"history"`
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
