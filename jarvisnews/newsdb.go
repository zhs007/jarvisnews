package jarvisnews

import (
	"context"

	"github.com/zhs007/ankadb"
	"github.com/zhs007/jarviscore/base"
	"go.uber.org/zap"
)

// newsDB - news database
type newsDB struct {
	ankaDB ankadb.AnkaDB
}

// newNewsDB - new news db
func newDTDataDB(dbpath string, httpAddr string, engine string) (*newsDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath

	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   WebSitesDBName,
		Engine: engine,
		PathDB: WebSitesDBName,
	})

	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   NewsDBName,
		Engine: engine,
		PathDB: NewsDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		jarvisbase.Error("newDTDataDB", zap.Error(err))

		return nil, err
	}

	jarvisbase.Info("newDTDataDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &newsDB{
		ankaDB: ankaDB,
	}

	return db, nil
}

// addURL -
func (db *newsDB) addURL(ctx context.Context, url string) error {
	curkey := makeURLKey(url)

	err := db.ankaDB.Set(ctx, WebSitesDBName, curkey, []byte("nodata"))
	if err != nil {
		jarvisbase.Warn("newsDB.addURL:Set", zap.Error(err))

		return err
	}

	return nil
}

// getURL -
func (db *newsDB) getURL(ctx context.Context, url string) (string, error) {
	curkey := makeURLKey(url)

	buf, err := db.ankaDB.Get(ctx, WebSitesDBName, curkey)
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return "", nil
		}

		jarvisbase.Warn("newsDB.getURL:Get", zap.Error(err))

		return "", err
	}

	return string(buf), nil
}
