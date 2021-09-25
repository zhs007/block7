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

type Cooking struct {
	Level    int  `json:"level"`
	Unlock   bool `json:"unlock"`
	StarNums int  `json:"starnum"`
}

// UpdUserDataParams - update userdata parameters
type UpdUserDataParams struct {
	Name      string      `json:"name"`
	Coin      int64       `json:"coin"`
	Level     int         `json:"level"`
	LevelArr  map[int]int `json:"levelarr"`
	ToolsArr  map[int]int `json:"toolsarr"`
	HomeScene []int       `json:"homeScene"`
	Cooking   []*Cooking  `json:"cooking"`
	Platform  string      `json:"platform"` // it's like android, iphone
	Version   int64       `json:"version"`
}

// UpdUserDataResult - update userdata result
type UpdUserDataResult struct {
	OldVersion int64 `json:"oldVersion"`
	NewVersion int64 `json:"newVersion"`
}

// UserDataParams - userdata parameters
type UserDataParams struct {
	Name     string `json:"name"`
	Platform string `json:"platform"` // it's like android, iphone
}

// UserDataResult - userdata parameters
type UserDataResult struct {
	Name      string      `json:"name"`
	Coin      int64       `json:"coin"`
	Level     int         `json:"level"`
	LevelArr  map[int]int `json:"levelarr"`
	ToolsArr  map[int]int `json:"toolsarr"`
	HomeScene []int       `json:"homeScene"`
	Cooking   []*Cooking  `json:"cooking"`
	Platform  string      `json:"platform"` // it's like android, iphone
	Version   int64       `json:"version"`
}
