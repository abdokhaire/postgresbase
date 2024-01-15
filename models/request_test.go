package models_test

import (
	"testing"

	"github.com/AlperRehaYAZGAN/postgresbase/models"
)

func TestRequestTableName(t *testing.T) {
	m := models.Request{}
	if m.TableName() != "_requests" {
		t.Fatalf("Unexpected table name, got %q", m.TableName())
	}
}
