package mpproxysql

import (
	"path/filepath"
	"testing"
)

func TestReadConfig(t *testing.T) {
	file := filepath.Join("..", "test", "proxysql.cnf")
	c := ReadConfig(file)
	expected := "admin:admin"
	if expected != c.adminCred {
		t.Errorf("adminCred does not match\nExpect: %v\nResult: %v", expected, c.adminCred)
	}
}
