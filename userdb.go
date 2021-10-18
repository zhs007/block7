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

type UserStageState struct {
	GameStateNums map[int]int `json:"gamestatenums"`
}

func (uss *UserStageState) SetGameState(gs int) {
	if gs > 0 {
		_, isok := uss.GameStateNums[gs]
		if isok {
			uss.GameStateNums[gs]++
		} else {
			uss.GameStateNums[gs] = 1
		}
	}
}

type UserDayStatsData struct {
	UserID   int64                   `json:"userid"`
	UserHash string                  `json:"userHash"`
	Stages   map[int]*UserStageState `json:"stages"`
}

func (udsd *UserDayStatsData) SetGameState(stage int, gs int) {
	if stage > 0 && gs > 0 {
		_, isok := udsd.Stages[stage]
		if isok {
			udsd.Stages[stage].SetGameState(gs)
		} else {
			udsd.Stages[stage] = &UserStageState{
				GameStateNums: make(map[int]int),
			}

			udsd.Stages[stage].SetGameState(gs)
		}
	}
}

type UserStageHistoryData struct {
	HistoryID      int64   `json:"history"`
	Ts             int64   `json:"ts"`
	ClickNums      int     `json:"clickNums"`
	AvgClickTime   float32 `json:"avgClickTime"`
	MaxClickTime   int     `json:"maxClickTime"`
	MinClickTime   int     `json:"minClickTime"`
	TotalClickTime int64   `json:"totalClickTime"`
	GameState      int     `json:"gamestate"`
	ClientVersion  string  `json:"clientVersion"`
	LastHP         int     `json:"lastHP"`
	LastCoin       int     `json:"lastCoin"`
	RefreshTimes   int     `json:"refreshTimes"`
	BackTimes      int     `json:"backTimes"`
	BombTimes      int     `json:"bombTimes"`
	RebirthTimes   int     `json:"rebirthTimes"`
}

type UserStageData struct {
	Historys []*UserStageHistoryData `json:"historys"`
	WinNums  int                     `json:"winnums"`
}

type UserCooking struct {
	Level    int  `json:"level"`
	Unlock   bool `json:"unlock"`
	StarNums int  `json:"starnum"`
}

type UserDBUserStatsData struct {
	UserID        int64                  `json:"userid"`
	UserHash      string                 `json:"userHash"`
	Stages        map[int]*UserStageData `json:"stages"`
	IPAddr        string                 `json:"ip"`
	CreateTime    string                 `json:"createTime"`
	LastLoginTime string                 `json:"lastLoginTime"`
	Name          string                 `json:"name"`
	Coin          int64                  `json:"coin"`
	Level         int                    `json:"level"`
	LevelArr      map[string]int         `json:"levelarr"`
	ToolsArr      map[string]int         `json:"toolsarr"`
	HomeScene     []int                  `json:"homeScene"`
	Cooking       []UserCooking          `json:"cooking"`
	Platform      string                 `json:"platform"`
	Version       int64                  `json:"version"`
	ClientVersion string                 `json:"clientVersion"`
	LastAwardTime string                 `json:"lastAwardTime"`
}

func (uusd *UserDBUserStatsData) AddHistory(historyID int64, pbHistory *block7pb.Scene) {
	if pbHistory.StageID2 > 0 && historyID > 0 {
		h, isok := uusd.Stages[int(pbHistory.StageID2)]
		if !isok {
			h = &UserStageData{}

			uusd.Stages[int(pbHistory.StageID2)] = h
		}

		tt := int64(0)
		maxt := -1
		mint := -1
		pt := -1
		for i := 0; i < len(pbHistory.History2)/4; i++ {
			ct := int(pbHistory.History2[i*4+3])
			if pt < 0 {
				pt = ct
			}

			ot := ct - pt

			if maxt < 0 || maxt < ot {
				maxt = ot
			}

			if (mint < 0 || mint > ot) && ot > 0 {
				mint = ot
			}

			tt += int64(ot)
		}

		hd := &UserStageHistoryData{
			HistoryID:     historyID,
			Ts:            pbHistory.Ts,
			ClickNums:     len(pbHistory.History2) / 4,
			GameState:     int(pbHistory.GameState),
			ClientVersion: pbHistory.ClientVersion,
			LastHP:        int(pbHistory.LastHP),
			LastCoin:      int(pbHistory.LastCoin),
			RefreshTimes:  int(pbHistory.RefreshTimes),
			BackTimes:     int(pbHistory.BackTimes),
			BombTimes:     int(pbHistory.BombTimes),
			RebirthTimes:  int(pbHistory.RefreshTimes),
		}

		if hd.ClickNums > 0 {
			hd.MaxClickTime = maxt
			hd.MinClickTime = mint
			hd.TotalClickTime = tt
			hd.AvgClickTime = float32(float64(tt) / float64(hd.ClickNums))
		}

		h.Historys = append(h.Historys, hd)

		if pbHistory.GameState == 1 {
			h.WinNums++
		}
	}
}

type UserDBStatsData struct {
	LatestUserID int64 `json:"latestuserid"`
	UserNums     int   `json:"usernums"`
	UserDataNums int   `json:"userdatanums"`
}

type UserDBDayStatsData struct {
	FirstUserID       int64                       `json:"firstuserid"`
	NewUserNums       int                         `json:"newusernums"`
	NewUserDataNums   int                         `json:"newuserdatanums"`
	FirstUserDataUID  int64                       `json:"firstuserdatauid"`
	AliveUserNums     int                         `json:"aliveusernums"`
	AliveUserDataNums int                         `json:"aliveuserdatanums"`
	Users2            map[int64]*UserDayStatsData `json:"users2"`
}

const userdbname = "userdb"
const userHashKeyPrefix = "h:"
const userIDKey = "curuserid"

func makeUserDBKey(userid int64) string {
	return fmt.Sprintf("u:%v", userid)
}

func getUserIDFromUserDBKey(key string) int64 {
	if len(key) > 2 {
		key1 := key[2:]

		i64, err := goutils.String2Int64(key1)
		if err == nil {
			return i64
		}

		goutils.Warn("getUserIDFromUserDBKey:String2Int64",
			zap.String("key", key),
			zap.Error(err))

		return 0
	}

	goutils.Warn("getUserIDFromUserDBKey",
		zap.String("key", key))

	return 0
}

func makeUserHashDBKey(uasehash string) string {
	return goutils.AppendString(userHashKeyPrefix, uasehash)
}

func makeUserDataDBKey(name string, platform string) string {
	return fmt.Sprintf("d:%v:%v", platform, name)
}

// UserDB - database
type UserDB struct {
	AnkaDB    ankadb.AnkaDB
	mutexDB   sync.Mutex
	historyDB *HistoryDB
}

// NewUserDB - new UserDB
func NewUserDB(dbpath string, httpAddr string, engine string, historydb *HistoryDB) (*UserDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   userdbname,
		Engine: engine,
		PathDB: userdbname,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		return nil, err
	}

	db := &UserDB{
		AnkaDB:    ankaDB,
		historyDB: historydb,
	}

	return db, err
}

// _setCurUserID - set current userID
func (db *UserDB) _setCurUserID(ctx context.Context, userid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, userid)
	if err != nil {
		goutils.Error("UserDB._setCurUserID:binary.Write",
			zap.Error(err))

		return err
	}

	// db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, userIDKey, buf.Bytes())
	// db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// GetCurUserID - get current userID
func (db *UserDB) GetCurUserID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	defer db.mutexDB.Unlock()

	buf, err := db.AnkaDB.Get(ctx, userdbname, userIDKey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			err = db._setCurUserID(ctx, 1)
			if err != nil {
				goutils.Error("UserDB.GetCurUserID:_setCurUserID",
					zap.Error(err))

				return 0, err
			}

			return 1, nil
		}

		goutils.Warn("UserDB.GetCurUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		goutils.Error("UserDB.GetCurUserID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db._setCurUserID(ctx, userid+1)
	if err != nil {
		goutils.Error("UserDB.GetCurSceneID:_setCurUserID",
			zap.Error(err))

		return 0, err
	}

	return userid, nil
}

// UpdUser - update user
func (db *UserDB) UpdUser(ctx context.Context, user *block7pb.UserInfo) error {
	buf, err := proto.Marshal(user)
	if err != nil {
		goutils.Warn("UserDB.UpdUser:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, makeUserDBKey(user.UserID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Warn("UserDB.UpdUser:Set",
			zap.Error(err))

		return err
	}

	return nil
}

// GetUser - get user
func (db *UserDB) GetUser(ctx context.Context, userid int64) (*block7pb.UserInfo, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, userdbname, makeUserDBKey(userid))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		goutils.Warn("UserDB.GetUser:Get",
			zap.Error(err))

		return nil, err
	}

	user := &block7pb.UserInfo{}

	err = proto.Unmarshal(buf, user)
	if err != nil {
		goutils.Warn("UserDB.GetUser:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return user, nil
}

// HasUserHash - has userhash
func (db *UserDB) HasUserHash(ctx context.Context, userhash string) (bool, error) {
	db.mutexDB.Lock()
	_, err := db.AnkaDB.Get(ctx, userdbname, makeUserHashDBKey(userhash))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return false, nil
		}

		goutils.Warn("UserDB.HasUserHash:Get",
			zap.Error(err))

		return false, err
	}

	return true, nil
}

// UpdUserHash - update userhash
func (db *UserDB) UpdUserHash(ctx context.Context, userhash string, userid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, userid)
	if err != nil {
		goutils.Error("UserDB.UpdUserHash:binary.Write",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, makeUserHashDBKey(userhash), buf.Bytes())
	db.mutexDB.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// GetUserID - get userID
func (db *UserDB) GetUserID(ctx context.Context, userhash string) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, userdbname, makeUserHashDBKey(userhash))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		goutils.Warn("UserDB.GetUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		goutils.Error("UserDB.GetUserID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return userid, nil
}

// DelUserHash - del userhash
func (db *UserDB) DelUserHash(ctx context.Context, userhash string) error {
	db.mutexDB.Lock()
	err := db.AnkaDB.Delete(ctx, userdbname, makeUserHashDBKey(userhash))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil
		}

		goutils.Warn("UserDB.DelUserHash:Get",
			zap.Error(err))

		return err
	}

	return nil
}

// genUserHash - generator a user hash
func (db *UserDB) genUserHash(ctx context.Context) (string, error) {
	for {
		userhash := goutils.GenHashCode(16)
		hasuh, err := db.HasUserHash(ctx, userhash)
		if err != nil {
			goutils.Error("UserDB.genUserHash:HasUserHash",
				zap.Error(err))

			return "", err
		}

		if !hasuh {
			return userhash, nil
		}
	}
}

// NewUser - new a userinfo
func (db *UserDB) NewUser(ctx context.Context, udi *block7pb.UserDeviceInfo) (*block7pb.UserInfo, error) {
	if udi.UserHash != "" {
		db.DelUserHash(ctx, udi.UserHash)
	}

	uid, err := db.GetCurUserID(ctx)
	if err != nil {
		goutils.Error("UserDB.NewUser:GetCurUserID",
			zap.Error(err))

		return nil, err
	}

	ui := &block7pb.UserInfo{
		UserID: uid,
	}

	userhash, err := db.genUserHash(ctx)
	if err != nil {
		goutils.Error("UserDB.NewUser:genUserHash",
			zap.Error(err))

		return nil, err
	}

	udi.UserHash = userhash
	udi.CreateTs = goutils.GetCurTimestamp()
	udi.LastLoginTs = udi.CreateTs
	udi.LoginTimes++
	ui.Data = append(ui.Data, udi)

	err = db.UpdUser(ctx, ui)
	if err != nil {
		goutils.Error("UserDB.NewUser:UpdUser",
			zap.Error(err))

		return nil, err
	}

	err = db.UpdUserHash(ctx, udi.UserHash, ui.UserID)
	if err != nil {
		goutils.Error("UserDB.NewUser:UpdUserHash",
			zap.Error(err))

		return nil, err
	}

	return ui, nil
}

// UpdUserDeviceInfo - Update userinfo with userdeviceinfo
func (db *UserDB) UpdUserDeviceInfo(ctx context.Context, udi *block7pb.UserDeviceInfo) (*block7pb.UserInfo, error) {
	if udi.UserHash == "" {
		return db.NewUser(ctx, udi)
	}

	uid, err := db.GetUserID(ctx, udi.UserHash)
	if err != nil {
		goutils.Error("UserDB.UpdUserDeviceInfo:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		return db.NewUser(ctx, udi)
	}

	ui, err := db.GetUser(ctx, uid)
	if err != nil {
		goutils.Error("UserDB.UpdUserDeviceInfo:GetUser",
			zap.Error(err))

		return nil, err
	}

	for _, cudi := range ui.Data {
		if cudi.UserHash == udi.UserHash {
			cudi.ADID = udi.ADID
			cudi.GUID = udi.GUID
			cudi.PlatformInfo = udi.PlatformInfo
			cudi.GameVersion = udi.GameVersion
			cudi.ResourceVersion = udi.ResourceVersion
			cudi.DeviceInfo = udi.DeviceInfo

			cudi.LastLoginTs = goutils.GetCurTimestamp()
			cudi.LoginTimes++

			err = db.UpdUser(ctx, ui)
			if err != nil {
				goutils.Error("UserDB.UpdUserDeviceInfo:UpdUser",
					zap.Error(err))

				return nil, err
			}

			return ui, nil
		}
	}

	udi.LastLoginTs = goutils.GetCurTimestamp()
	udi.LoginTimes++

	ui.Data = append(ui.Data, udi)

	err = db.UpdUser(ctx, ui)
	if err != nil {
		goutils.Error("UserDB.UpdUserDeviceInfo:UpdUser",
			zap.Error(err))

		return nil, err
	}

	return ui, nil
}

// GetUserData - get userData
func (db *UserDB) GetUserData(ctx context.Context, name string, platform string) (*block7pb.UserData, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, userdbname, makeUserDataDBKey(name, platform))
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		goutils.Warn("UserDB.GetUserData:Get",
			zap.Error(err))

		return nil, err
	}

	ud := &block7pb.UserData{}

	err = proto.Unmarshal(buf, ud)
	if err != nil {
		goutils.Warn("UserDB.GetUserData:Unmarshal",
			zap.Error(err))

		return nil, err
	}

	return ud, nil
}

// UpdUserData - update userData
func (db *UserDB) UpdUserData(ctx context.Context, ud *block7pb.UserData, uds *UpdUserDataStatus) (int64, int64, error) {
	ud0, err := db.GetUserData(ctx, ud.Name, ud.Platform)
	if err != nil {
		goutils.Warn("UserDB.UpdUserData:GetUserData",
			zap.Error(err))

		return 0, 0, err
	}

	nud := MergeUserData(ud0, ud, uds)
	nud.LastTs = goutils.GetCurTimestamp()

	// goutils.Debug("UserDB.UpdUserData",
	// 	goutils.JSON("ud", ud),
	// 	goutils.JSON("uds", uds),
	// 	goutils.JSON("ud0", ud0),
	// 	goutils.JSON("nud", nud))

	buf, err := proto.Marshal(nud)
	if err != nil {
		goutils.Warn("UserDB.UpdUserData:Marshal",
			zap.Error(err))

		return 0, 0, err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, makeUserDataDBKey(ud.Name, ud.Platform), buf)
	db.mutexDB.Unlock()
	if err != nil {
		goutils.Warn("UserDB.UpdUserData:Set",
			zap.Error(err))

		return 0, 0, err
	}

	if ud0 != nil {
		return ud0.Version, nud.Version, nil
	}

	return 0, nud.Version, nil
}

// GetLatestUserID - get latest userID
func (db *UserDB) GetLatestUserID(ctx context.Context) (int64, error) {
	db.mutexDB.Lock()
	buf, err := db.AnkaDB.Get(ctx, userdbname, userIDKey)
	db.mutexDB.Unlock()
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return 0, nil
		}

		goutils.Warn("UserDB.GetLatestUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		goutils.Error("UserDB.GetLatestUserID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	return userid, nil
}

// Stats - statistics
func (db *UserDB) Stats(ctx context.Context) (*UserDBStatsData, error) {
	latestUserID, err := db.GetLatestUserID(ctx)
	if err != nil {
		goutils.Error("UserDB.Stats:GetLatestUserID",
			zap.Error(err))

		return nil, err
	}

	userNums := 0
	userDataNums := 0

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "u:", func(key string, value []byte) error {
		userNums++

		return nil
	})
	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "d:", func(key string, value []byte) error {
		userDataNums++

		return nil
	})
	db.mutexDB.Unlock()

	return &UserDBStatsData{
		LatestUserID: latestUserID,
		UserNums:     userNums,
		UserDataNums: userDataNums,
	}, nil
}

// statsUserDay - statistics
func (db *UserDB) statsUserDay(ctx context.Context, t time.Time, ui *block7pb.UserInfo) (*UserDayStatsData, error) {
	udsd := &UserDayStatsData{
		UserID: ui.UserID,
		Stages: make(map[int]*UserStageState),
	}

	if len(ui.Data) > 0 {
		udsd.UserHash = ui.Data[0].UserHash
	}

	// db.historyDB.statsDayUser(ctx, t, udsd)

	return udsd, nil
}

// StatsDay - statistics
func (db *UserDB) StatsDay(ctx context.Context, t time.Time, lastUserID int64) (*UserDBDayStatsData, error) {
	firstuid := int64(0)
	newusers := 0
	loginusers := 0

	users2 := make(map[int64]*UserDayStatsData)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "u:", func(key string, value []byte) error {
		cuid := getUserIDFromUserDBKey(key)
		if cuid > 0 && cuid <= lastUserID {
			return nil
		}

		user := &block7pb.UserInfo{}

		err := proto.Unmarshal(value, user)
		if err != nil {
			goutils.Warn("UserDB.StatsDay:Unmarshal",
				zap.String("key", key),
				zap.Error(err))

			return nil
		}

		if len(user.Data) > 0 {
			isui := false
			rt := time.Unix(user.Data[0].CreateTs, 0)
			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
				newusers++

				if firstuid == 0 {
					firstuid = user.UserID
				} else if firstuid > user.UserID {
					firstuid = user.UserID
				}

				isui = true
			}

			lt := time.Unix(user.Data[0].LastLoginTs, 0)
			if t.Year() == lt.Year() && t.YearDay() == lt.YearDay() {
				loginusers++

				isui = true
			}

			if isui {
				udsd, err := db.statsUserDay(ctx, t, user)
				if err != nil {
					goutils.Warn("UserDB.StatsDay:statsUserDay",
						zap.String("key", key),
						zap.Error(err))

					return nil
				}

				if udsd != nil {
					users2[user.UserID] = udsd
				}
			}
		}

		return nil
	})

	firstuduid := int64(0)
	newuds := 0
	loginuds := 0
	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "d:", func(key string, value []byte) error {
		// cuid := getUserIDFromUserDBKey(key)
		// if cuid > 0 && cuid <= lastUserID {
		// 	return nil
		// }

		ud := &block7pb.UserData{}

		err := proto.Unmarshal(value, ud)
		if err != nil {
			goutils.Warn("UserDB.StatsDay:Unmarshal",
				zap.String("key", key),
				zap.Error(err))

			return nil
		}

		rt := time.Unix(ud.CreateTs, 0)
		if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
			newuds++

			if ud.UserID > 0 {
				if firstuduid == 0 {
					firstuduid = ud.UserID
				} else if firstuduid > ud.UserID {
					firstuduid = ud.UserID
				}
			}
		}

		lt := time.Unix(ud.LastTs, 0)
		if t.Year() == lt.Year() && t.YearDay() == lt.YearDay() {
			loginuds++
		}

		return nil
	})
	db.mutexDB.Unlock()

	return &UserDBDayStatsData{
		FirstUserID:       firstuid,
		NewUserNums:       newusers,
		AliveUserNums:     loginusers,
		FirstUserDataUID:  firstuduid,
		NewUserDataNums:   newuds,
		AliveUserDataNums: loginuds,
		Users2:            users2,
	}, nil
}

// UserStats - statistics
func (db *UserDB) UserStats(ctx context.Context, uid int64) (*UserDBUserStatsData, error) {
	user, err := db.GetUser(ctx, uid)
	if err != nil {
		goutils.Error("UserDB.UserStats:GetUser",
			zap.Int64("uid", uid),
			zap.Error(err))

		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	uusd := &UserDBUserStatsData{
		UserID: user.UserID,
		Stages: make(map[int]*UserStageData),
	}

	if len(user.Data) > 0 {
		uusd.UserHash = user.Data[0].UserHash
		uusd.IPAddr = user.Data[0].IPAddr
		uusd.CreateTime = time.Unix(user.Data[0].CreateTs, 0).Format("2006-01-02_15:04:05")
	}

	ud, err := db.GetUserData(ctx, uusd.UserHash, "android")
	if err != nil {
		goutils.Error("UserDB.UserStats:GetUserData",
			zap.Int64("uid", uid),
			zap.Error(err))

		return nil, err
	}

	if ud != nil {
		uusd.Name = ud.Name
		uusd.Platform = ud.Platform
		uusd.Coin = ud.Coin
		uusd.ClientVersion = ud.ClientVersion
		uusd.LastAwardTime = time.Unix(ud.LastAwardTs, 0).Format("2006-01-02_15:04:05")
		uusd.LastLoginTime = time.Unix(ud.LastTs, 0).Format("2006-01-02_15:04:05")
		uusd.Level = int(ud.Level)
		uusd.Version = ud.Version

		for _, v := range ud.Cooking {
			uusd.Cooking = append(uusd.Cooking, UserCooking{
				Level:    int(v.Level),
				Unlock:   v.Unlock,
				StarNums: int(v.StarNums),
			})
		}

		uusd.LevelArr = make(map[string]int)
		for k1, v1 := range ud.LevelArr {
			uusd.LevelArr[k1] = int(v1)
		}

		uusd.ToolsArr = make(map[string]int)
		for k2, v2 := range ud.ToolsArr {
			uusd.ToolsArr[k2] = int(v2)
		}

		for _, v3 := range ud.HomeScene {
			uusd.HomeScene = append(uusd.HomeScene, int(v3))
		}
	}

	db.historyDB.statsUser(ctx, uusd)

	return uusd, nil
}

// // countTodayUsers - count users today
// func (db *UserDB) countTodayUsers(ctx context.Context, t time.Time, lastUserID int64) (int, int, error) {
// 	newusers := 0
// 	loginusers := 0
// 	// ct := time.Unix(ts, 0)

// 	db.mutexDB.Lock()
// 	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "u:", func(key string, value []byte) error {
// 		user := &block7pb.UserInfo{}

// 		err := proto.Unmarshal(value, user)
// 		if err != nil {
// 			goutils.Warn("UserDB.CountTodayNewUsers:Unmarshal",
// 				zap.Error(err))

// 			return nil
// 		}

// 		if len(user.Data) > 0 {
// 			rt := time.Unix(user.Data[0].CreateTs, 0)
// 			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
// 				newusers++
// 			}

// 			lt := time.Unix(user.Data[0].LastLoginTs, 0)
// 			if t.Year() == lt.Year() && t.YearDay() == lt.YearDay() {
// 				loginusers++
// 			}
// 		}

// 		return nil
// 	})
// 	db.mutexDB.Unlock()

// 	return newusers, loginusers, nil
// }

// findTodayFirstUserID - find first userID today
func (db *UserDB) findTodayFirstUserID(ctx context.Context, t time.Time, lastUserID int64) (int64, error) {
	firstuserid := int64(0)
	// ct := time.Unix(ts, 0)

	db.mutexDB.Lock()
	db.AnkaDB.ForEachWithPrefix(ctx, userdbname, "u:", func(key string, value []byte) error {
		cuid := getUserIDFromUserDBKey(key)
		if cuid > 0 && cuid <= lastUserID {
			return nil
		}

		user := &block7pb.UserInfo{}

		err := proto.Unmarshal(value, user)
		if err != nil {
			goutils.Warn("UserDB.findTodayFirstUserID:Unmarshal",
				zap.Error(err))

			return nil
		}

		if len(user.Data) > 0 {
			rt := time.Unix(user.Data[0].CreateTs, 0)
			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
				if firstuserid == 0 || user.UserID < firstuserid {
					firstuserid = user.UserID
				}
			}
		}

		return nil
	})
	db.mutexDB.Unlock()

	return firstuserid, nil
}
