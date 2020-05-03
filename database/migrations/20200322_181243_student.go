package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Student_20200322_181243 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Student_20200322_181243{}
	m.Created = "20200322_181243"

	migration.Register("Student_20200322_181243", m)
}

// Run the migrations
func (m *Student_20200322_181243) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE student(`id` int(11) DEFAULT NULL,`name` varchar(128) NOT NULL,`sex` varchar(128) NOT NULL,`age` int(11) DEFAULT NULL,`employed` int(11) DEFAULT NULL,`offerid` int(11) DEFAULT NULL)")
}

// Reverse the migrations
func (m *Student_20200322_181243) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `student`")
}
