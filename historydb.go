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

type HistoryDBStatsData struct {
	LatestHistoryID int64       `json:"latesthistoryid"`
	HistoryNums     int         `json:"historynums"`
	MapNums         map[int]int `json:"mapnums"`
	StageNums       map[int]int `json:"stagenums"`
}

const historydbname = "historydb"
const historyIDKey = "curhistoryid"

func makeHistoryDBKey(historyid int64) string {
	return fmt.Sprintf("h:%v", historyid)
}

func makeUserHistoryDBKey(mapid int32, userid int64, sceneid int64, historyid int64) string {
	return fmt.Sprintf("u:%v:%v:%v:%v", userid, mapid, sceneid, historyid)
}

func makeSceneHistoryDBKey(mapid int32, userid int64, sceneid int64, historyid int64) string {
	return fmt.Sprintf("s:%v:%v:%v:%v", mapid, sceneid, userid, historyid)
}

// HistoryDB - database
type HistoryDB struct {
	AnkaDB  ankadb.AnkaDB
	mutexDB sync.Mutex
}

// NewStageDB - new HistoryDB
func NewHistoryDB(dbpath string, httpAddr string, engine string) (*HistoryDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   historydbname,
		Engine: engine,
		PathDB: historydbname,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		return nil, err
	}

	db := &HistoryDB{
		AnkaDB: ankaDB,
	}

	return db, err
}

// _setCurHistoryID - set current historyID
func (db *HistoryDB) _setCurHistoryID(ctx context.Context, historyid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, historyid)
	if err != nil {
		goutils.Error("HistoryDB._setCurHistoryID:binary.Write",
			zap.Error(err))

		return err
	}

	// db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, historyIDKey, buf.Bytes())
	// db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// GetCurHistoryID - get historyID
func (db *HistoryDB) GetCurHistoryID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	defer db.mutexDB.Unlock()

	buf, err := db.AnkaDB.Get(ctx, historydbname, historyIDKey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			err = db._setCurHistoryID(ctx, 1)
			if err != nil {
				goutils.Error("HistoryDB.GetCurHistoryID:_setCurHistoryID",
					zap.Error(err))

				return 0, err
			}

			return 1, nil
		}

		return 0, err
	}

	var historyid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &historyid)
	if err != nil {
		goutils.Error("HistoryDB.GetCurHistoryID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db._setCurHistoryID(ctx, historyid+1)
	if err != nil {
		goutils.Error("HistoryDB.GetCurHistoryID:_setCurHistoryID",
			zap.Error(err))

		return 0, err
	}

	return historyid, nil
}

// updHistory - update history
func (db *HistoryDB) updHistory(ctx context.Context, scene *block7pb.Scene) error {
	buf, err := proto.Marshal(scene)
	if err != nil {
		goutils.Warn("HistoryDB.updHistory:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeHistoryDBKey(scene.HistoryID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Warn("HistoryDB.updHistory:Set",
			zap.Error(err))

		return err
	}

	return nil
}

// GetHistory - get history
func (db *HistoryDB) GetHistory(ctx context.Context, historyid int64) (*block7pb.Scene, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, historydbname, makeHistoryDBKey(historyid))
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
		goutils.Warn("HistoryDB.GetHistory:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return stage, nil
}

// SaveHistory - save history
func (db *HistoryDB) SaveHistory(ctx context.Context, scene *block7game.Scene) (*block7pb.Scene, error) {
	hid, err := db.GetCurHistoryID(ctx)
	if err != nil {
		goutils.Warn("HistoryDB.SaveHistory:GetCurHistoryID",
			zap.Error(err))

		return nil, err
	}

	pbhistory, err := scene.ToHistoryPB()
	if err != nil {
		goutils.Warn("HistoryDB.SaveHistory:ToHistoryPB",
			zap.Error(err))

		return nil, err
	}

	pbhistory.HistoryID = hid

	err = db.updHistory(ctx, pbhistory)
	if err != nil {
		goutils.Warn("HistoryDB.SaveHistory:updHistory",
			zap.Error(err))

		return nil, err
	}

	db.setHistoryID(ctx, pbhistory.MapID2, pbhistory.UserID, pbhistory.SceneID, pbhistory.HistoryID)

	return pbhistory, nil
}

// setHistoryID - set historyID
func (db *HistoryDB) setHistoryID(ctx context.Context, mapid int32, userid int64, sceneid int64, historyid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, historyid)
	if err != nil {
		goutils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeUserHistoryDBKey(mapid, userid, sceneid, historyid), buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.String("key", makeUserHistoryDBKey(mapid, userid, sceneid, historyid)),
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeSceneHistoryDBKey(mapid, userid, sceneid, historyid), buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.String("key", makeSceneHistoryDBKey(mapid, userid, sceneid, historyid)),
			zap.Error(err))

		return err
	}

	return nil
}

// GetLatestHistoryID - get latest historyID
func (db *HistoryDB) GetLatestHistoryID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, historydbname, historyIDKey)
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		goutils.Error("HistoryDB.GetLatestHistoryID:Get",
			zap.Error(err))

		return 0, err
	}

	var historyid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &historyid)
	if err != nil {
		goutils.Error("HistoryDB.GetLatestHistoryID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return historyid, nil
}

// Stats - statistics
func (db *HistoryDB) Stats(ctx context.Context) (*HistoryDBStatsData, error) {
	latestHistoryID, err := db.GetLatestHistoryID(ctx)
	if err != nil {
		goutils.Error("HistoryDB.Stats:GetLatestHistoryID",
			zap.Error(err))

		return nil, err
	}

	historyNums := 0
	mapNums := make(map[int]int)
	stageNums := make(map[int]int)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, historydbname, "h:", func(key string, value []byte) error {
		historyNums++

		stage := &block7pb.Scene{}

		err = proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("HistoryDB.Stats:Unmarshal",
				zap.Error(err))
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

	return &HistoryDBStatsData{
		LatestHistoryID: latestHistoryID,
		HistoryNums:     historyNums,
		MapNums:         mapNums,
		StageNums:       stageNums,
	}, nil
}
