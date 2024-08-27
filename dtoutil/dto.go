package dto

import (
	"database/sql"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// Copy copier的封装,包含字段映射
func Copy(toValue, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
		Converters: []copier.TypeConverter{
			TimeToTimeStampPb(),
			TimeToString(),
			GormDeletedAtToTimeStampPb(),
			GormDeletedAtToString(),
			SQLNullTimeToTimeStampPb(),
			SQLNullTimeToString(),
		},
	})
}

// TimeToTimeStampPb time.Time to *timestamppb.Timestamp
func TimeToTimeStampPb() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(time.Time)
			return timestamppb.New(t), nil
		},
	}
}

// TimeToString time.Time to string
func TimeToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: time.Time{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(time.Time)
			return t.Format("2006-01-02 15:04:05"), nil
		},
	}
}

// GormDeletedAtToTimeStampPb gorm.DeletedAt to *timestamppb.Timestamp
func GormDeletedAtToTimeStampPb() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: gorm.DeletedAt{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(gorm.DeletedAt)
			return timestamppb.New(t.Time), nil
		},
	}
}

// GormDeletedAtToString gorm.DeletedAt to *timestamppb.Timestamp
func GormDeletedAtToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: gorm.DeletedAt{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(gorm.DeletedAt)
			return t.Time.Format("2006-01-02 15:04:05"), nil
		},
	}
}

// SQLNullTimeToTimeStampPb sql.NullTime to *timestamppb.Timestamp
func SQLNullTimeToTimeStampPb() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: sql.NullTime{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(sql.NullTime)
			return timestamppb.New(t.Time), nil
		},
	}
}

// SQLNullTimeToString sql.NullTime to *timestamppb.Timestamp
func SQLNullTimeToString() copier.TypeConverter {
	return copier.TypeConverter{
		SrcType: sql.NullTime{},
		DstType: &timestamppb.Timestamp{},
		Fn: func(src interface{}) (interface{}, error) {
			t := src.(sql.NullTime)
			return t.Time.Format("2006-01-02 15:04:05"), nil
		},
	}
}
