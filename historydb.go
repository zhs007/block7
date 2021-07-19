package block7

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	ankadb "github.com/zhs007/ankadb"
	"github.com/zhs007/block7/block7pb"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const historydbname = "historydb"
const historyIDKey = "curhistoryid"

func makeHistoryDBKey(historyid int64) string {
	return fmt.Sprintf("h:%v", historyid)
}

func makeUserHistoryDBKey(mapid string, userid int64, sceneid int64, historyid int64) string {
	return fmt.Sprintf("u:%v:%v:%v:%v", userid, mapid, sceneid, historyid)
}

func makeSceneHistoryDBKey(mapid string, userid int64, sceneid int64, historyid int64) string {
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

// setCurHistoryID - set current historyID
func (db *HistoryDB) setCurHistoryID(ctx context.Context, historyid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, historyid)
	if err != nil {
		block7utils.Error("HistoryDB.setCurHistoryID:binary.Write",
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
			err = db.setCurHistoryID(ctx, 1)
			if err != nil {
				block7utils.Error("HistoryDB.GetCurHistoryID:setCurHistoryID",
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
		block7utils.Error("HistoryDB.GetCurHistoryID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db.setCurHistoryID(ctx, historyid+1)
	if err != nil {
		block7utils.Error("HistoryDB.GetCurHistoryID:setCurHistoryID",
			zap.Error(err))

		return 0, err
	}

	return historyid, nil
}

// updHistory - update history
func (db *HistoryDB) updHistory(ctx context.Context, scene *block7pb.Scene) error {
	buf, err := proto.Marshal(scene)
	if err != nil {
		block7utils.Warn("HistoryDB.updHistory:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeHistoryDBKey(scene.HistoryID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		block7utils.Warn("HistoryDB.updHistory:Set",
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
		block7utils.Warn("HistoryDB.GetHistory:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return stage, nil
}

// SaveHistory - save history
func (db *HistoryDB) SaveHistory(ctx context.Context, scene *Scene) (*block7pb.Scene, error) {
	hid, err := db.GetCurHistoryID(ctx)
	if err != nil {
		block7utils.Warn("HistoryDB.SaveHistory:GetCurHistoryID",
			zap.Error(err))

		return nil, err
	}

	pbhistory, err := scene.ToHistoryPB()
	if err != nil {
		block7utils.Warn("HistoryDB.SaveHistory:ToHistoryPB",
			zap.Error(err))

		return nil, err
	}

	pbhistory.HistoryID = hid

	err = db.updHistory(ctx, pbhistory)
	if err != nil {
		block7utils.Warn("HistoryDB.SaveHistory:updHistory",
			zap.Error(err))

		return nil, err
	}

	db.setHistoryID(ctx, pbhistory.MapID, pbhistory.UserID, pbhistory.SceneID, pbhistory.HistoryID)

	return pbhistory, nil
}

// setHistoryID - set historyID
func (db *HistoryDB) setHistoryID(ctx context.Context, mapid string, userid int64, sceneid int64, historyid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, historyid)
	if err != nil {
		block7utils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeUserHistoryDBKey(mapid, userid, sceneid, historyid), buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		block7utils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.String("key", makeUserHistoryDBKey(mapid, userid, sceneid, historyid)),
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, historydbname, makeSceneHistoryDBKey(mapid, userid, sceneid, historyid), buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		block7utils.Error("HistoryDB.setHistoryID:binary.Write",
			zap.String("key", makeSceneHistoryDBKey(mapid, userid, sceneid, historyid)),
			zap.Error(err))

		return err
	}

	return nil
}
