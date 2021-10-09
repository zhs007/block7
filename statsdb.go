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
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type StatsDBStatsData struct {
	FirstTime  string                            `json:"firstTime"`
	LatestTime string                            `json:"latestTime"`
	DayStats   map[string]*block7pb.DayStatsData `json:"dayStats"`
}

const statsdbname = "statsdb"
const latestStatsKey = "lateststats"
const firstStatsKey = "firststats"

func makeDayStatsDataDBKey(ts int64) string {
	return fmt.Sprintf("d:%v", ts)
}

// StatsDB - database
type StatsDB struct {
	AnkaDB  ankadb.AnkaDB
	mutexDB sync.Mutex
	ticker  *time.Ticker
	userDB  *UserDB
	stageDB *StageDB
}

// NewStatsDB - new StatsDB
func NewStatsDB(dbpath string, httpAddr string, engine string, userdb *UserDB, stagedb *StageDB) (*StatsDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   statsdbname,
		Engine: engine,
		PathDB: statsdbname,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		return nil, err
	}

	db := &StatsDB{
		AnkaDB:  ankaDB,
		userDB:  userdb,
		stageDB: stagedb,
	}

	return db, err
}

// GetLatestStatsTs - get latest timeatamp
func (db *StatsDB) GetLatestStatsTs(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, statsdbname, latestStatsKey)
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		return 0, err
	}

	var lateststats int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &lateststats)
	if err != nil {
		goutils.Error("StatsDB.GetLatestStatsTs:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return lateststats, nil
}

// GetFirstStatsTs - get first timeatamp
func (db *StatsDB) GetFirstStatsTs(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, statsdbname, firstStatsKey)
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		return 0, err
	}

	var firststats int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &firststats)
	if err != nil {
		goutils.Error("StatsDB.GetFirstStatsTs:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return firststats, nil
}

// UpdDayStats - update DayStats
func (db *StatsDB) UpdDayStats(ctx context.Context, dsd *block7pb.DayStatsData) error {
	buf, err := proto.Marshal(dsd)
	if err != nil {
		goutils.Warn("StatsDB.UpdDayStats:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, statsdbname, makeDayStatsDataDBKey(dsd.Ts), buf)
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Warn("StatsDB.UpdDayStats:Set",
			zap.Error(err))

		return err
	}

	return nil
}

// GetDayStats - get DayStats
func (db *StatsDB) GetDayStats(ctx context.Context, ts int64) (*block7pb.DayStatsData, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, statsdbname, makeDayStatsDataDBKey(ts))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		goutils.Warn("StatsDB.GetDayStats:Get",
			zap.Error(err))

		return nil, err
	}

	dsd := &block7pb.DayStatsData{}

	err = proto.Unmarshal(buf, dsd)
	if err != nil {
		goutils.Warn("StatsDB.GetDayStats:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return dsd, nil
}

// genDayStats - genarate DayStats
func (db *StatsDB) genDayStats(ctx context.Context, cdt time.Time, lastUID int64) (*block7pb.DayStatsData, error) {
	// nt := time.Now()
	// curts := goutils.FormatUTCDayTs(nt)
	// cdt := time.Unix(curts, 0)

	uds, err := db.userDB.StatsDay(ctx, cdt, lastUID)
	if err != nil {
		goutils.Warn("StatsDB.GetDayStats:StatsDay:userDB",
			zap.Error(err))

		return nil, err
	}

	sds, err := db.stageDB.StatsDay(ctx, cdt)
	if err != nil {
		goutils.Warn("StatsDB.GetDayStats:StatsDay:stageDB",
			zap.Error(err))

		return nil, err
	}

	return &block7pb.DayStatsData{
		Ts:                cdt.Unix(),
		NewUserNums:       int32(uds.NewUserNums),
		AliveUserNums:     int32(uds.AliveUserNums),
		FirstUserID:       uds.FirstUserID,
		NewUserDataNums:   int32(uds.NewUserDataNums),
		AliveUserDataNums: int32(uds.AliveUserDataNums),
		FirstUserDataUID:  uds.FirstUserDataUID,
		FirstSceneID:      sds.FirstSceneID,
		SceneNums:         int32(sds.SceneNums),
		MapNums:           goutils.MapII2MapI32I32(sds.MapNums),
		StageNums:         goutils.MapII2MapI32I32(sds.StageNums),
	}, nil
}

// Start - start
func (db *StatsDB) Start() {
	db.ticker = time.NewTicker(5 * time.Minute)

	go db.onTimer()
}

// onTimer - on Timer
func (db *StatsDB) Stop() {
	if db.ticker != nil {
		db.ticker.Stop()
	}
}

// onTimer - on Timer
func (db *StatsDB) onTimer() {
	firstts, err := db.GetFirstStatsTs(context.Background())
	if err != nil {
		goutils.Warn("StatsDB.onTimer:GetFirstStatsTs",
			zap.Error(err))

		return
	}

	latestts, err := db.GetLatestStatsTs(context.Background())
	if err != nil {
		goutils.Warn("StatsDB.onTimer:GetLatestStatsTs",
			zap.Error(err))

		return
	}

	latestdayts := int64(0)

	if latestts > 0 {
		latestdayts = goutils.FormatUTCDayTs(time.Unix(latestts, 0))
	}

	// firstuserid := int64(0)
	dayfirstuid := int64(0)

	goutils.Info("StatsDB.onTimer",
		zap.String("first", time.Unix(firstts, 0).Format("2006-01-02_15:04:05")),
		zap.String("latest", time.Unix(latestts, 0).Format("2006-01-02_15:04:05")),
		zap.String("latestday", time.Unix(latestdayts, 0).Format("2006-01-02_15:04:05")),
	)

	for range db.ticker.C {
		nt := time.Now()
		curts := goutils.FormatUTCDayTs(nt)
		cdt := time.Unix(curts, 0)

		// new day
		if curts != latestdayts {
			uid, err := db.userDB.findTodayFirstUserID(context.Background(), cdt, dayfirstuid)
			if err != nil {
				goutils.Warn("StatsDB.onTimer:findTodayFirstUserID:newday",
					zap.Error(err))
			}

			dayfirstuid = uid

			latestdayts = curts
		}

		if dayfirstuid == 0 {
			uid, err := db.userDB.findTodayFirstUserID(context.Background(), cdt, 0)
			if err != nil {
				goutils.Warn("StatsDB.onTimer:findTodayFirstUserID",
					zap.Error(err))
			}

			dayfirstuid = uid
		}

		dsd, err := db.genDayStats(context.Background(), cdt, dayfirstuid)
		if err != nil {
			goutils.Warn("StatsDB.onTimer:GenDayStats",
				zap.Error(err))
		}

		dsd.Ts = nt.Unix()

		err = db.UpdDayStats(context.Background(), dsd)
		if err != nil {
			goutils.Warn("StatsDB.onTimer:UpdDayStats",
				zap.Error(err))
		}

		if firstts == 0 {
			firstts = nt.Unix()
			db.setFirstStatsTs(context.Background(), firstts)
		}

		latestts = nt.Unix()
		db.setLatestStatsTs(context.Background(), latestts)
	}
}

// setLatestStatsTs - set latest timeatamp
func (db *StatsDB) setLatestStatsTs(ctx context.Context, ts int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, ts)
	if err != nil {
		goutils.Error("StatsDB.setLatestStatsTs:binary.Write",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, statsdbname, latestStatsKey, buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// setFirstStatsTs - set first timeatamp
func (db *StatsDB) setFirstStatsTs(ctx context.Context, ts int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, ts)
	if err != nil {
		goutils.Error("StatsDB.setFirstStatsTs:binary.Write",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, statsdbname, firstStatsKey, buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// Stats - statistics
func (db *StatsDB) Stats(ctx context.Context) (*StatsDBStatsData, error) {
	firstts, err := db.GetFirstStatsTs(context.Background())
	if err != nil {
		goutils.Warn("StatsDB.Stats:GetFirstStatsTs",
			zap.Error(err))

		return nil, err
	}

	latestts, err := db.GetLatestStatsTs(context.Background())
	if err != nil {
		goutils.Warn("StatsDB.Stats:GetLatestStatsTs",
			zap.Error(err))

		return nil, err
	}

	mapDayStats := make(map[string]*block7pb.DayStatsData)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, statsdbname, "d:", func(key string, value []byte) error {
		dsd := &block7pb.DayStatsData{}
		err = proto.Unmarshal(value, dsd)
		if err != nil {
			goutils.Warn("StatsDB.Stats:Unmarshal",
				zap.Error(err))

			return nil
		}

		nt := time.Unix(dsd.Ts, 0)
		curts := goutils.FormatUTCDayTs(nt)
		cdt := time.Unix(curts, 0)

		dsd.Time = time.Unix(dsd.Ts, 0).Format("2006-01-02_15:04:05")
		dsd.Ts = 0

		mapDayStats[cdt.Format("2006-01-02")] = dsd

		return nil
	})
	db.mutexDB.Unlock()

	return &StatsDBStatsData{
		FirstTime:  time.Unix(firstts, 0).Format("2006-01-02_15:04:05"),
		LatestTime: time.Unix(latestts, 0).Format("2006-01-02_15:04:05"),
		DayStats:   mapDayStats,
	}, nil
}
