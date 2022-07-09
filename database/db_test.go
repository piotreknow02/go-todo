package database_test

import (
	"go-todo/database"
	"testing"
)

func TestConnectToDabatase(t *testing.T) {
	_, err := database.GetDatabase()
	if err != nil {
		t.Errorf("Error while connecting to database: %v", err)
		return
	}
	t.Logf("PASSED Connected to database")
}
