package database

import (
	"fmt"

	"gitlab.com/sholludev/sampoerna_notification/models"
	"gitlab.com/sholludev/sampoerna_notification/pkg/util/environment"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

type credential struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
	SSLMode  string
	Timezone string
}

func Init(driver string) {
	credential := &credential{
		Host:     environment.Get("DB_HOST"),
		Port:     environment.Get("DB_PORT"),
		User:     environment.Get("DB_USER"),
		Pass:     environment.Get("DB_PASS"),
		Name:     environment.Get("DB_NAME"),
		SSLMode:  environment.Get("DB_SSLMODE"),
		Timezone: environment.Get("DB_TZ"),
	}

	switch driver {
	case "mysql":
		credential.getMysql()
	case "postgres":
		credential.getPostgres()
	default:
		credential.getMysql()
	}
}

func DBManager() *gorm.DB {
	return db
}

func (c *credential) getMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", c.User, c.Pass, c.Host, c.Port, c.Name, c.Timezone)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func (c *credential) getPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s Timezone=%s", c.Host, c.Port, c.User, c.Pass, c.Name, c.SSLMode, c.Timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	db.AutoMigrate(
		&models.MKategori{},
		&models.TNotifikasi{},
	)
}

func TruncateTable(table string) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	db.Exec("TRUNCATE TABLE " + table)
	db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}
