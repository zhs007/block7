package block7serv

import "errors"

var (
	// ErrInvalidUserHash - invalid userhash
	ErrInvalidUserHash = errors.New("invalid userhash")
	// ErrInvalidMissionID - invalid MissionID
	ErrInvalidMissionID = errors.New("invalid MissionID")
	// ErrInvalidSceneID - invalid SceneID
	ErrInvalidSceneID = errors.New("invalid SceneID")
	// ErrInvalidUserDataName - invalid UserData Name
	ErrInvalidUserDataName = errors.New("invalid UserData Name")
	// ErrInvalidToken - invalid Token
	ErrInvalidToken = errors.New("invalid Token")
)
