package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20190718_171856 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20190718_171856{}
	m.Created = "20190718_171856"

	migration.Register("User_20190718_171856", m)
}

// Run the migrations
func (m *User_20190718_171856) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20190718_171856) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
