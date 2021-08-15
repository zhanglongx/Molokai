// Copyright 2020 Longxiao Zhang <zhanglongx@gmail.com>.
// All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package core

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
	"github.com/zhanglongx/Molokai/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Test_holdings(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	tables := []interface{}{new(Holding)}
	db := &holdings{
		DB: initTestDB(t, tables...),
	}

	for _, tc := range []struct {
		name string
		test func(*testing.T, *holdings)
	}{
		{"Create", test_holding_Create},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				err := clearTables(t, db.DB, tables...)
				if err != nil {
					t.Fatal(err)
				}
			})
		})
		tc.test(t, db)
	}
}

func test_holding_Create(t *testing.T, db *holdings) {
	holding, err := db.Create("000001.SZ", date.TodayUTC().String(), 1.1, "")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, common.Symbol("000001.SZ"), holding.Symbol)
	assert.Equal(t, float64(1.1), holding.Cost)
	assert.Equal(t, "", holding.Reminders)
}

func initTestDB(t *testing.T, tables ...interface{}) *gorm.DB {
	t.Helper()

	dsn := "root:123@tcp(localhost:3306)/testMolokai?charset=utf8mb4&parseTime=true"
	now := time.Now().UTC().Truncate(time.Second)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return now
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}

		if t.Failed() {
			t.Logf("testMolokai left intact")
			return
		}
	})

	err = db.Migrator().AutoMigrate(tables...)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func clearTables(t *testing.T, db *gorm.DB, tables ...interface{}) error {
	if t.Failed() {
		return nil
	}

	for _, t := range tables {
		err := db.Where("TRUE").Delete(t).Error
		if err != nil {
			return err
		}
	}

	return nil
}
