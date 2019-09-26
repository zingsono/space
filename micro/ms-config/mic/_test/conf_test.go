package _test

import (
	"testing"

	"mic/conf"
)

func TestNowFile(t *testing.T) {
	config := conf.NowFile("D:\\Projects\\space\\micro\\ms-config\\mic\\config.json")
	t.Logf(config.Application.Name)
}
