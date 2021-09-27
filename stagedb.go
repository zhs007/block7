package block7

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	ankadb "github.com/zhs007/ankadb"
	"github.com/zhs007/block7/block7pb"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type StageDBStatsData struct {
	LatestSceneID int64       `json:"latestsceneid"`
	SceneNums     int         `json:"scenenums"`
	MapNums       map[int]int `json:"mapnums"`
	StageNums     map[int]int `json:"stagenums"`
}

const stagedbname = "stagedb"
const sceneIDKey = "cursceneid"

func makeStageDBKey(sceneid int64) string {
	return fmt.Sprintf("s:%v", sceneid)
}

// StageDB - database
type StageDB struct {
	AnkaDB  ankadb.AnkaDB
	mutexDB sync.Mutex
}

// NewStageDB - new StageDB
func NewStageDB(dbpath string, httpAddr string, engine string) (*StageDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   stagedbname,
		Engine: engine,
		PathDB: stagedbname,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		return nil, err
	}

	db := &StageDB{
		AnkaDB: ankaDB,
	}

	return db, err
}

// _setCurSceneID - set current sceneID
func (db *StageDB) _setCurSceneID(ctx context.Context, sceneid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, sceneid)
	if err != nil {
		goutils.Error("StageDB._setCurSceneID:binary.Write",
			zap.Error(err))

		return err
	}

	// db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, stagedbname, sceneIDKey, buf.Bytes())
	// db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// GetSymbol - get symbol
func (db *StageDB) GetCurSceneID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	defer db.mutexDB.Unlock()

	buf, err := db.AnkaDB.Get(ctx, stagedbname, sceneIDKey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			err = db._setCurSceneID(ctx, 1)
			if err != nil {
				goutils.Error("StageDB.GetCurSceneID:_setCurSceneID",
					zap.Error(err))

				return 0, err
			}

			return 1, nil
		}

		return 0, err
	}

	var sceneid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &sceneid)
	if err != nil {
		goutils.Error("StageDB.GetCurSceneID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db._setCurSceneID(ctx, sceneid+1)
	if err != nil {
		goutils.Error("StageDB.GetCurSceneID:_setCurSceneID",
			zap.Error(err))

		return 0, err
	}

	return sceneid, nil
}

// UpdStage - update stage
func (db *StageDB) UpdStage(ctx context.Context, scene *block7pb.Scene) error {
	buf, err := proto.Marshal(scene)
	if err != nil {
		goutils.Warn("StageDB.UpdStage:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, stagedbname, makeStageDBKey(scene.SceneID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Warn("StageDB.UpdStage:Set",
			zap.Error(err))

		return err
	}

	return nil
}

// GetStage - get stage
func (db *StageDB) GetStage(ctx context.Context, sceneid int64) (*block7pb.Scene, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, stagedbname, makeStageDBKey(sceneid))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	stage := &block7pb.Scene{}

	err = proto.Unmarshal(buf, stage)
	if err != nil {
		goutils.Warn("StageDB.GetStage:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return stage, nil
}

// SaveStage - save stage
func (db *StageDB) SaveStage(ctx context.Context, scene *block7game.Scene) (*block7pb.Scene, error) {
	sid, err := db.GetCurSceneID(ctx)
	if err != nil {
		goutils.Warn("StageDB.SaveStage:GetCurSceneID",
			zap.Error(err))

		return nil, err
	}

	pbscene, err := scene.ToScenePB()
	if err != nil {
		goutils.Warn("StageDB.SaveStage:ToScenePB",
			zap.Error(err))

		return nil, err
	}

	pbscene.SceneID = sid

	err = db.UpdStage(ctx, pbscene)
	if err != nil {
		goutils.Warn("StageDB.SaveStage:UpdStage",
			zap.Error(err))

		return nil, err
	}

	return pbscene, nil
}

// GetLatestSceneID - get latest sceneID
func (db *StageDB) GetLatestSceneID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, stagedbname, sceneIDKey)
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		goutils.Error("StageDB.GetLatestSceneID:Get",
			zap.Error(err))

		return 0, err
	}

	var sceneid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &sceneid)
	if err != nil {
		goutils.Error("StageDB.GetLatestSceneID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return sceneid, nil
}

// Stats - statistics
func (db *StageDB) Stats(ctx context.Context) (*StageDBStatsData, error) {
	lastSceneID, err := db.GetLatestSceneID(ctx)
	if err != nil {
		goutils.Error("StageDB.Stats:GetLatestSceneID",
			zap.Error(err))

		return nil, err
	}

	sceneNums := 0
	mapNums := make(map[int]int)
	stageNums := make(map[int]int)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, stagedbname, "s:", func(key string, value []byte) error {
		sceneNums++

		stage := &block7pb.Scene{}
		err = proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("StageDB.Stats:Unmarshal",
				zap.Error(err))

			return nil
		}

		_, isok := mapNums[int(stage.MapID2)]
		if isok {
			mapNums[int(stage.MapID2)]++
		} else {
			mapNums[int(stage.MapID2)] = 1
		}

		if stage.StageID2 > 0 {
			_, isok = stageNums[int(stage.StageID2)]
			if isok {
				stageNums[int(stage.StageID2)]++
			} else {
				stageNums[int(stage.StageID2)] = 1
			}
		}

		return nil
	})
	db.mutexDB.Unlock()

	return &StageDBStatsData{
		LatestSceneID: lastSceneID,
		SceneNums:     sceneNums,
		MapNums:       mapNums,
		StageNums:     stageNums,
	}, nil
}
