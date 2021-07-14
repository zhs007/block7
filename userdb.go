package block7

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	ankadb "github.com/zhs007/ankadb"
	"github.com/zhs007/block7/block7pb"
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
	return AppendString(userHashKeyPrefix, uasehash)
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
		Error("UserDB.setCurUserID:binary.Write",
			zap.Error(err))

		return err
	}

	err = db.AnkaDB.Set(ctx, userdbname, userIDKey, buf.Bytes())
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
				Error("UserDB.GetCurUserID:setCurUserID",
					zap.Error(err))

				return 0, err
			}

			return 1, nil
		}

		Warn("UserDB.GetCurUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		Error("UserDB.GetCurUserID:binary.Read",
			zap.Error(err))

		return 0, err
	}

	err = db.setCurUserID(ctx, userid+1)
	if err != nil {
		Error("UserDB.GetCurSceneID:setCurSceneID",
			zap.Error(err))

		return 0, err
	}

	return userid, nil
}

// UpdUser - update user
func (db *UserDB) UpdUser(ctx context.Context, user *block7pb.UserInfo) error {
	buf, err := proto.Marshal(user)
	if err != nil {
		Warn("UserDB.UpdUser:Marshal",
			zap.Error(err))

		return err
	}

	db.mutexDB.Lock()
	err = db.AnkaDB.Set(ctx, userdbname, makeUserDBKey(user.UserID), buf)
	db.mutexDB.Unlock()
	if err != nil {
		Warn("UserDB.UpdUser:Set",
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

		Warn("UserDB.GetUser:Get",
			zap.Error(err))

		return nil, err
	}

	user := &block7pb.UserInfo{}

	err = proto.Unmarshal(buf, user)
	if err != nil {
		Warn("UserDB.GetUser:Unmarshal",
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

		Warn("UserDB.HasUserHash:Get",
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
		Error("UserDB.UpdUserHash:binary.Write",
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

		Warn("UserDB.GetUserID:Get",
			zap.Error(err))

		return 0, err
	}

	var userid int64
	reader := bytes.NewReader(buf)
	err = binary.Read(reader, binary.LittleEndian, &userid)
	if err != nil {
		Error("UserDB.GetUserID:binary.Read",
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

		Warn("UserDB.DelUserHash:Get",
			zap.Error(err))

		return err
	}

	return nil
}

// genUserHash - generator a user hash
func (db *UserDB) genUserHash(ctx context.Context) (string, error) {
	for {
		userhash := GenHashCode(16)
		hasuh, err := db.HasUserHash(ctx, userhash)
		if err != nil {
			Error("UserDB.genUserHash:HasUserHash",
				zap.Error(err))

			return "", err
		}

		if !hasuh {
			return userhash, nil
		}
	}
}

// GetUserID - get userID
func (db *UserDB) NewUser(ctx context.Context, udi *block7pb.UserDeviceInfo) (*block7pb.UserInfo, error) {
	if udi.UserHash != "" {
		db.DelUserHash(ctx, udi.UserHash)
	}

	uid, err := db.GetCurUserID(ctx)
	if err != nil {
		Error("UserDB.NewUser:GetCurUserID",
			zap.Error(err))

		return nil, err
	}

	ui := &block7pb.UserInfo{
		UserID: uid,
	}

	userhash, err := db.genUserHash(ctx)
	if err != nil {
		Error("UserDB.NewUser:genUserHash",
			zap.Error(err))

		return nil, err
	}

	udi.UserHash = userhash

	return ui, nil
}
