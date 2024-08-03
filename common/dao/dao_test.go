package dao

import (
	"fmt"
	"testing"
)

func init() {
	var err error

	// mysql分库 test test-1 test-2 test-3
	DatabasePool = make(map[string]*MySQLConn)
	mysqlConn, err := MySQLConnInit("test")
	if err != nil {
		panic(err)
	}
	DatabasePool["test"] = mysqlConn
	for i := 1; i <= 3; i++ {
		dbname := fmt.Sprintf("test-%v", i)
		mysqlConn, err := MySQLConnInit(dbname)
		if err != nil {
			panic(err)
		}
		DatabasePool[dbname] = mysqlConn
	}

	err = CacheInit()
	if err != nil {
		panic(err)
	}

	DefaultUser.Init()
}

func TestInit(t *testing.T) {

}
