package _test

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	connectionString := "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	fmt.Println(dbName)
}
