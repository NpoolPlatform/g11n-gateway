//nolint:dupl
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

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"

	intlconst "github.com/NpoolPlatform/internationalization/pkg/message/const"

	entintl "github.com/NpoolPlatform/internationalization/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/google/uuid"
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
var langs = []*entintl.Lang{}
var appLangs = []*entintl.AppLang{}

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

func migrateLang(ctx context.Context, cli *entintl.Client) error {
	_langs, err := cli.
		Lang.
		Query().
		All(ctx)
	if err != nil {
		logger.Sugar().Errorw("migrateLang", "error", err)
		return err
	}

	langs = _langs

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		infos, err := tx.
			Lang.
			Query().
			Limit(1).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(infos) > 0 {
			return nil
		}

		for _, lang := range langs {
			_, err = tx.
				Lang.
				Create().
				SetID(lang.ID).
				SetLang(lang.Lang).
				SetLogo(lang.Logo).
				SetName(lang.Name).
				SetShort(lang.Short).
				Save(_ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func migrateAppCountry(ctx context.Context) error {
	offset := int32(0)
	const limit = int32(100)

	for {
		apps, _, err := appmwcli.GetApps(ctx, offset, limit)
		if err != nil {
			return err
		}
		if len(apps) == 0 {
			return nil
		}

		checkExist := false

		err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			acs, err := tx.
				AppCountry.
				Query().
				Limit(1).
				All(_ctx)
			if err != nil {
				return err
			}
			if len(acs) > 0 && !checkExist {
				checkExist = true
				return nil
			}

			for _, app := range apps {
				for _, country := range countries {
					_, err := tx.
						AppCountry.
						Create().
						SetAppID(uuid.MustParse(app.ID)).
						SetCountryID(country.ID).
						Save(_ctx)
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			return err
		}

		offset += limit
	}
}

func migrateAppLang(ctx context.Context, cli *entintl.Client) error {
	_appLangs, err := cli.
		AppLang.
		Query().
		All(ctx)
	if err != nil {
		logger.Sugar().Errorw("migrateAppLang", "error", err)
		return err
	}

	appLangs = _appLangs

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		acs, err := tx.
			AppLang.
			Query().
			Limit(1).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(acs) > 0 {
			return nil
		}

		for _, appLang := range appLangs {
			_, err := tx.
				AppLang.
				Create().
				SetAppID(appLang.AppID).
				SetLangID(appLang.LangID).
				SetMain(appLang.MainLang).
				Save(_ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func migrateMessage(ctx context.Context, cli *entintl.Client) error {
	msgs, err := cli.
		Message.
		Query().
		All(ctx)
	if err != nil {
		logger.Sugar().Errorw("migrateMessage", "error", err)
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		acs, err := tx.
			Message.
			Query().
			Limit(1).
			All(_ctx)
		if err != nil {
			return err
		}
		if len(acs) > 0 {
			return nil
		}

		for _, msg := range msgs {
			_, err := tx.
				Message.
				Create().
				SetID(msg.ID).
				SetAppID(msg.AppID).
				SetLangID(msg.LangID).
				SetMessageID(msg.MessageID).
				SetMessage(msg.Message).
				SetGetIndex(0).
				SetDisabled(false).
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

	if err := migrateCountry(ctx, cli); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	if err := migrateLang(ctx, cli); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	if err := migrateAppCountry(ctx); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	if err := migrateAppLang(ctx, cli); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	if err := migrateMessage(ctx, cli); err != nil {
		logger.Sugar().Errorw("Migrate", "error", err)
		return err
	}

	return nil
}
