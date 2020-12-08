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

	var users []User

	// ↓本当はこんなことがしたい
	// query := User{Languages: []Language{{Code: "JA"}}}
	// DB.Preload("Languages").Where(&query).Find(&users)

	DB.
		Table("users").
		Joins("left join user_speaks on user_speaks.user_id = users.id").
		Where("user_speaks.language_code = ?", "JA").
		Scan(&users)

	if users[0].Name != jinzhu.Name {
		t.Errorf("failed")
	}
}
