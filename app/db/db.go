package db

import (
	"database/sql"
	"fmt"
	"github.com/poster-keisuke/sample-clearn-architecture/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const retrySleepPeriod = 5 * time.Second

type DB struct {
	db *sql.DB
}

func NewDB(config config.DBConfig) (*gorm.DB, error) {
	db, err := connect(config, 10)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//func GetStmt(ctx context.Context, query string) (*sql.Stmt, error) {
//	stmt, err := sql.DB.PrepareContext(ctx, query)
//	if err != nil {
//		return nil, xerrors.Errorf(": %w", err)
//	}
//
//	return stmt, nil
//}

func connect(config config.DBConfig, retryCount int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if retryCount > 0 {
			time.Sleep(retrySleepPeriod)
			return connect(config, retryCount-1)
		}
		return nil, fmt.Errorf("failed to connect database")
	}

	return db, nil

}
