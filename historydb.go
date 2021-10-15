package block7

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"

	ankadb "github.com/zhs007/ankadb"
	"github.com/zhs007/block7/block7pb"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type HistoryStageData struct {
	Nums          int         `json:"nums"`
	GameStateNums map[int]int `json:"gamestatenums"`
}

func (hsd *HistoryStageData) SetGameState(gs int) {
	if gs > 0 {
		_, isok := hsd.GameStateNums[gs]
		if isok {
			hsd.GameStateNums[gs]++
		} else {
			hsd.GameStateNums[gs] = 1
		}
	}
}

type HistoryDBStatsData struct {
	LatestHistoryID int64                     `json:"latesthistoryid"`
	HistoryNums     int                       `json:"historynums"`
	Stages          map[int]*HistoryStageData `json:"stages"`
	GameStateNums   map[int]int               `json:"gamestatenums"`
}

type HistoryDBDayStatsData struct {
	FirstHistoryID int64                     `json:"firsthistoryid"`
	HistoryNums    int                       `json:"historynums"`
	Stages         map[int]*HistoryStageData `json:"stages"`
	GameStateNums  map[int]int               `json:"gamestatenums"`
	UserIDNums     map[int64]int             `json:"useridnums"`
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

func getHistoryIDFromHistoryDBKey(key string) int64 {
	if len(key) > 2 {
		key1 := key[2:]

		i64, err := goutils.String2Int64(key1)
		if err == nil {
			return i64
		}

		goutils.Warn("getHistoryIDFromHistoryDBKey:String2Int64",
			zap.String("key", key),
			zap.Error(err))

		return 0
	}

	goutils.Warn("getHistoryIDFromHistoryDBKey",
		zap.String("key", key))

	return 0
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
	scene.Ts = goutils.GetCurTimestamp()

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
	// mapNums := make(map[int]int)
	stages := make(map[int]*HistoryStageData)
	gameStateNums := make(map[int]int)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, historydbname, "h:", func(key string, value []byte) error {
		historyNums++

		stage := &block7pb.Scene{}

		err = proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("HistoryDB.Stats:Unmarshal",
				zap.Error(err))
		}

		// _, isok := mapNums[int(stage.MapID2)]
		// if isok {
		// 	mapNums[int(stage.MapID2)]++
		// } else {
		// 	mapNums[int(stage.MapID2)] = 1
		// }

		if stage.StageID2 > 0 {
			_, isok := stages[int(stage.StageID2)]
			if isok {
				stages[int(stage.StageID2)].Nums++
				stages[int(stage.StageID2)].SetGameState(int(stage.GameState))
			} else {
				stages[int(stage.StageID2)] = &HistoryStageData{
					Nums:          1,
					GameStateNums: make(map[int]int),
				}

				stages[int(stage.StageID2)].SetGameState(int(stage.GameState))
			}
		}

		if stage.GameState > 0 {
			_, isok := gameStateNums[int(stage.GameState)]
			if isok {
				gameStateNums[int(stage.GameState)]++
			} else {
				gameStateNums[int(stage.GameState)] = 1
			}
		}

		// if stage.GameState > 0 {
		// 	_, isok = gameState[int(stage.GameState)]
		// 	if isok {
		// 		gameState[int(stage.GameState)].Nums++

		// 		gameState[int(stage.GameState)].SetGameState(int(stage.GameState))
		// 	} else {
		// 		gameState[int(stage.GameState)] = &HistoryStageData{
		// 			Nums:          1,
		// 			GameStateNums: make(map[int]int),
		// 		}
		// 	}
		// }

		return nil
	})
	db.mutexDB.Unlock()

	return &HistoryDBStatsData{
		LatestHistoryID: latestHistoryID,
		HistoryNums:     historyNums,
		Stages:          stages,
		GameStateNums:   gameStateNums,
	}, nil
}

// StatsDay - statistics
func (db *HistoryDB) StatsDay(ctx context.Context, t time.Time) (*HistoryDBDayStatsData, error) {
	// latestHistoryID, err := db.GetLatestHistoryID(ctx)
	// if err != nil {
	// 	goutils.Error("HistoryDB.Stats:GetLatestHistoryID",
	// 		zap.Error(err))

	// 	return nil, err
	// }

	firstHistoryID := int64(0)
	historyNums := 0
	// mapNums := make(map[int]int)
	stages := make(map[int]*HistoryStageData)
	gameStateNums := make(map[int]int)
	userIDNums := make(map[int64]int)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, historydbname, "h:", func(key string, value []byte) error {
		// historyNums++

		stage := &block7pb.Scene{}

		err := proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("HistoryDB.StatsDay:Unmarshal",
				zap.Error(err))
		}

		if stage.Ts > 0 {
			rt := time.Unix(stage.Ts, 0)
			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
				historyNums++

				if stage.HistoryID > 0 {
					if firstHistoryID == 0 || firstHistoryID > stage.HistoryID {
						firstHistoryID = stage.HistoryID
					}
				}

				// _, isok := mapNums[int(stage.MapID2)]
				// if isok {
				// 	mapNums[int(stage.MapID2)]++
				// } else {
				// 	mapNums[int(stage.MapID2)] = 1
				// }

				if stage.StageID2 > 0 {
					_, isok := stages[int(stage.StageID2)]
					if isok {
						stages[int(stage.StageID2)].Nums++
						stages[int(stage.StageID2)].SetGameState(int(stage.GameState))
					} else {
						stages[int(stage.StageID2)] = &HistoryStageData{
							Nums:          1,
							GameStateNums: make(map[int]int),
						}

						stages[int(stage.StageID2)].SetGameState(int(stage.GameState))
					}
				}

				if stage.GameState > 0 {
					_, isok := gameStateNums[int(stage.GameState)]
					if isok {
						gameStateNums[int(stage.GameState)]++
					} else {
						gameStateNums[int(stage.GameState)] = 1
					}
				}

				if stage.UserID > 0 {
					_, isok := userIDNums[stage.UserID]
					if isok {
						userIDNums[stage.UserID]++
					} else {
						userIDNums[stage.UserID] = 1
					}
				}
			}
		}

		return nil
	})
	db.mutexDB.Unlock()

	return &HistoryDBDayStatsData{
		FirstHistoryID: firstHistoryID,
		HistoryNums:    historyNums,
		Stages:         stages,
		GameStateNums:  gameStateNums,
		UserIDNums:     userIDNums,
	}, nil
}

// statsDayUser - statistics
func (db *HistoryDB) statsDayUser(ctx context.Context, t time.Time, udsd *UserDayStatsData) error {
	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, historydbname, "h:", func(key string, value []byte) error {
		stage := &block7pb.Scene{}

		err := proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("HistoryDB.statsDayUser:Unmarshal",
				zap.Error(err))
		}

		if stage.UserID == udsd.UserID && stage.Ts > 0 {
			rt := time.Unix(stage.Ts, 0)
			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
				if stage.StageID2 > 0 && stage.GameState > 0 {
					udsd.SetGameState(int(stage.StageID2), int(stage.GameState))
				}
			}
		}

		return nil
	})
	db.mutexDB.Unlock()

	return nil
}

// statsUser - statistics
func (db *HistoryDB) statsUser(ctx context.Context, uusd *UserDBUserStatsData) error {
	// goutils.Info("HistoryDB.statsUser",
	// 	goutils.JSON("uusd", uusd))

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, historydbname, "h:", func(key string, value []byte) error {
		stage := &block7pb.Scene{}

		err := proto.Unmarshal(value, stage)
		if err != nil {
			goutils.Warn("HistoryDB.statsDayUser:Unmarshal",
				zap.Error(err))
		}

		// goutils.Info("HistoryDB.statsUser",
		// 	zap.String("key", key),
		// 	zap.Int64("uid", stage.UserID),
		// 	zap.Int32("stageID", stage.StageID2))

		if stage.UserID == uusd.UserID && stage.StageID2 > 0 {
			if stage.HistoryID == 0 {
				stage.HistoryID = getHistoryIDFromHistoryDBKey(key)
			}

			if stage.HistoryID == 0 {
				goutils.Warn("HistoryDB.statsUser:non-historyID",
					zap.String("key", key))
			}

			uusd.AddHistory(stage.HistoryID, stage)
			// goutils.Info("HistoryDB.statsUser",
			// 	zap.String("key", key),
			// 	zap.Int("stageID", int(stage.StageID2)),
			// 	zap.Int64("HistoryID", stage.HistoryID))
		}

		return nil
	})
	db.mutexDB.Unlock()

	return nil
}
