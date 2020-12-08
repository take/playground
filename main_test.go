package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	jinzhu := User{Name: "jinzhu", Languages: []Language{{Code: "JA", Name: "Japanese"}}}

	DB.Create(&jinzhu)

	query := User{Languages: []Language{{Code: "JA"}}}
	var user User

	DB.Preload("Languages").Where(&query).Find(&user)
}
