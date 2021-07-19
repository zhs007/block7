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

const userdbname = "userdb"
const userHashKeyPrefix = "h:"
const userIDKey = "curuserid"

func makeUserDBKey(userid int64) string {
	return fmt.Sprintf("u:%v", userid)
}

func makeUserHashDBKey(uasehash string) string {
	return block7utils.AppendString(userHashKeyPrefix, uasehash)
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

// setCurUserID - set current userID
func (db *UserDB) setCurUserID(ctx context.Context, userid int64) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, userid)
	if err != nil {
		block7utils.Error("UserDB.setCurUserID:binary.Write",
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
			err = db.setCurUserID(ctx, 1)
			if err != nil {
				block7utils.Error("UserDB.GetCurUserID:setCurUserID",
					zap.Error(err))

				return 0, err
			}

			return 1, nil
		}

		block7utils.Warn("UserDB.GetCurUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		block7utils.Error("UserDB.GetCurUserID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db.setCurUserID(ctx, userid+1)
	if err != nil {
		block7utils.Error("UserDB.GetCurSceneID:setCurSceneID",
			zap.Error(err))

		return 0, err
	}

	return userid, nil
}

// UpdUser - update user
func (db *UserDB) UpdUser(ctx context.Context, user *block7pb.UserInfo) error {
	buf, err := proto.Marshal(user)
	if err != nil {
		block7utils.Warn("UserDB.UpdUser:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, makeUserDBKey(user.UserID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		block7utils.Warn("UserDB.UpdUser:Set",
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

		block7utils.Warn("UserDB.GetUser:Get",
			zap.Error(err))

		return nil, err
	}

	user := &block7pb.UserInfo{}

	err = proto.Unmarshal(buf, user)
	if err != nil {
		block7utils.Warn("UserDB.GetUser:Unmarshal",
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

		block7utils.Warn("UserDB.HasUserHash:Get",
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
		block7utils.Error("UserDB.UpdUserHash:binary.Write",
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

		block7utils.Warn("UserDB.GetUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		block7utils.Error("UserDB.GetUserID:binary.Read",
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

		block7utils.Warn("UserDB.DelUserHash:Get",
			zap.Error(err))

		return err
	}

	return nil
}

// genUserHash - generator a user hash
func (db *UserDB) genUserHash(ctx context.Context) (string, error) {
	for {
		userhash := block7utils.GenHashCode(16)
		hasuh, err := db.HasUserHash(ctx, userhash)
		if err != nil {
			block7utils.Error("UserDB.genUserHash:HasUserHash",
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
		block7utils.Error("UserDB.NewUser:GetCurUserID",
			zap.Error(err))

		return nil, err
	}

	ui := &block7pb.UserInfo{
		UserID: uid,
	}

	userhash, err := db.genUserHash(ctx)
	if err != nil {
		block7utils.Error("UserDB.NewUser:genUserHash",
			zap.Error(err))

		return nil, err
	}

	udi.UserHash = userhash
	udi.CreateTs = GetCurTimestamp()
	udi.LastLoginTs = udi.CreateTs
	udi.LoginTimes++
	ui.Data = append(ui.Data, udi)

	err = db.UpdUser(ctx, ui)
	if err != nil {
		block7utils.Error("UserDB.NewUser:UpdUser",
			zap.Error(err))

		return nil, err
	}

	err = db.UpdUserHash(ctx, udi.UserHash, ui.UserID)
	if err != nil {
		block7utils.Error("UserDB.NewUser:UpdUserHash",
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
		block7utils.Error("UserDB.UpdUserDeviceInfo:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		return db.NewUser(ctx, udi)
	}

	ui, err := db.GetUser(ctx, uid)
	if err != nil {
		block7utils.Error("UserDB.UpdUserDeviceInfo:GetUser",
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

			cudi.LastLoginTs = GetCurTimestamp()
			cudi.LoginTimes++

			err = db.UpdUser(ctx, ui)
			if err != nil {
				block7utils.Error("UserDB.UpdUserDeviceInfo:UpdUser",
					zap.Error(err))

				return nil, err
			}

			return ui, nil
		}
	}

	udi.LastLoginTs = GetCurTimestamp()
	udi.LoginTimes++

	ui.Data = append(ui.Data, udi)

	err = db.UpdUser(ctx, ui)
	if err != nil {
		block7utils.Error("UserDB.UpdUserDeviceInfo:UpdUser",
			zap.Error(err))

		return nil, err
	}

	return ui, nil
}