package dto

import (
	"database/sql"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func TestCopy(t *testing.T) {
	type Source struct {
		TimeField      time.Time
		DeletedAtField gorm.DeletedAt
		NullTimeField  sql.NullTime
	}

	type Destination struct {
		TimeField      *timestamppb.Timestamp
		DeletedAtField *timestamppb.Timestamp
		NullTimeField  *timestamppb.Timestamp
	}

	src := Source{
		TimeField:      time.Now(),
		DeletedAtField: gorm.DeletedAt{Time: time.Now(), Valid: true},
		NullTimeField:  sql.NullTime{Time: time.Now(), Valid: true},
	}

	var dest Destination
	err := Copy(&dest, &src)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	if dest.TimeField == nil || dest.DeletedAtField == nil || dest.NullTimeField == nil {
		t.Error("Copy did not convert fields correctly")
	}
}
