package block7serv

import (
	"github.com/zhs007/block7"
	"go.uber.org/zap"
)

// BasicServ - basic server
type BasicServ struct {
}

func NewBasicServ() *BasicServ {
	return &BasicServ{}
}

// Login - login
func (serv *BasicServ) Login(params *LoginParams) (*LoginResult, error) {
	return nil, nil
}

// Mission - get mission
func (serv *BasicServ) Mission(params *MissionParams) (*MissionResult, error) {
	stage, err := block7.LoadStage("./cfg/level_0100.json")
	if err != nil {
		block7.Error("LoadStage",
			zap.Error(err))

		return nil, err
	}

	rng := block7.NewRngNormal()

	scene, err := block7.NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, block7.DefaultMaxBlockNums)
	if err != nil {
		block7.Error("NewScene",
			zap.Error(err))

		return nil, err
	}

	scene.IsOutputScene = true

	return &MissionResult{
		Scene:       scene,
		MissionHash: block7.GenHashCode(16),
	}, nil
}

// MissionData - upload mission data
func (serv *BasicServ) MissionData(params *MissionDataParams) (*MissionDataResult, error) {
	return &MissionDataResult{UserLevel: 100}, nil
}
