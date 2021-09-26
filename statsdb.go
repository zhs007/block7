package block7

import (
	"sync"

	ankadb "github.com/zhs007/ankadb"
)

const statsdbname = "statsdb"

// StatsDB - database
type StatsDB struct {
	AnkaDB  ankadb.AnkaDB
	mutexDB sync.Mutex
}

// NewStatsDB - new StatsDB
func NewStatsDB(dbpath string, httpAddr string, engine string) (*StatsDB, error) {
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
		AnkaDB: ankaDB,
	}

	return db, err
}
