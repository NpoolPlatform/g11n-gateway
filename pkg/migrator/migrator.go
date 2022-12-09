package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"

	"github.com/NpoolPlatform/g11n-manager/pkg/db"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"

	intlconst "github.com/NpoolPlatform/internationalization/pkg/message/const"

	entintl "github.com/NpoolPlatform/internationalization/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

const (
	keyUsername = "username"
	keyPassword = "password"
	keyDBName   = "database_name"
	maxOpen     = 10
	maxIdle     = 10
	MaxLife     = 3
)

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Infow("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

var countries = []*entintl.Country{}

func migrateCountry(ctx context.Context, cli *entintl.Client) error {
	_countries, err := cli.
		Country.
		Query().
		All(ctx)
	if err != nil {
		logger.Sugar().Errorw("migrateCountry", "error", err)
		return err
	}

	countries = _countries

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		infos, err := tx.
			Country.
			Query().
			Limit(1).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(infos) > 0 {
			return nil
		}

		for _, country := range countries {
			_, err = tx.
				Country.
				Create().
				SetID(country.ID).
				SetCountry(country.Country).
				SetCode(country.Code).
				SetShort(country.Short).
				SetFlag(country.Flag).
				Save(_ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func Migrate(ctx context.Context) error {
	if err := db.Init(); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	conn, err := open(intlconst.ServiceName)
	if err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}
	defer conn.Close()

	cli := entintl.NewClient(entintl.Driver(entsql.OpenDB(dialect.MySQL, conn)))

	// TODO: migrate country
	if err := migrateCountry(ctx, cli); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	// TODO: migrate lang
	// TODO: migrate appcountry
	// TODO: migrate applang
	// TODO: migrate message
	return nil
}
