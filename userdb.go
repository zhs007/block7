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

type UserDBStatsData struct {
	LatestUserID int64 `json:"latestuserid"`
	UserNums     int   `json:"usernums"`
	UserDataNums int   `json:"userdatanums"`
}

type UserDBDayStatsData struct {
	FirstUserID       int64 `json:"firstuserid"`
	NewUserNums       int   `json:"newusernums"`
	NewUserDataNums   int   `json:"newuserdatanums"`
	FirstUserDataUID  int64 `json:"firstuserdatauid"`
	AliveUserNums     int   `json:"aliveusernums"`
	AliveUserDataNums int   `json:"aliveuserdatanums"`
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
	AnkaDB  ankadb.AnkaDB
	mutexDB sync.Mutex
}

// NewUserDB - new UserDB
func NewUserDB(dbpath string, httpAddr string, engine string) (*UserDB, error) {
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
		AnkaDB: ankaDB,
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

// StatsDay - statistics
func (db *UserDB) StatsDay(ctx context.Context, t time.Time, lastUserID int64) (*UserDBDayStatsData, error) {
	firstuid := int64(0)
	newusers := 0
	loginusers := 0

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
			rt := time.Unix(user.Data[0].CreateTs, 0)
			if t.Year() == rt.Year() && t.YearDay() == rt.YearDay() {
				newusers++

				if firstuid == 0 {
					firstuid = user.UserID
				} else if firstuid > user.UserID {
					firstuid = user.UserID
				}
			}

			lt := time.Unix(user.Data[0].LastLoginTs, 0)
			if t.Year() == lt.Year() && t.YearDay() == lt.YearDay() {
				loginusers++
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
	}, nil
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
