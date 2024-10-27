package timeutil

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon(t *testing.T) {
	assert.Equal(t, true, Carbon().IsZero())
}

func TestTime2Carbon(t *testing.T) {
	time2Carbon := TimeToCarbon(time.Now())
	assert.Equal(t, false, time2Carbon.IsZero())
}
func TestStrToCarbon(t *testing.T) {
	strToCarbon := StrToCarbon("2021-01-01 00:00:00")
	assert.Equal(t, false, strToCarbon.IsZero())
}

func TestNowCarbon(t *testing.T) {
	assert.Equal(t, false, NowCarbon().IsZero())
}
func TestNowTime(t *testing.T) {
	assert.Equal(t, false, NowTime().IsZero())
}

func TestNowUnix(t *testing.T) {
	assert.Equal(t, true, NowUnix() != 0)
}
func TestNowString(t *testing.T) {
	assert.Equal(t, true, NowString() != "")
}

func TestNowMillisecondString(t *testing.T) {
	assert.Equal(t, true, NowMillisecondString() != "")
}

func TestNowMicrosecondString(t *testing.T) {
	assert.Equal(t, true, NowMicrosecondString() != "")
}

func TestNowSQLNullTime(t *testing.T) {
	assert.Equal(t, true, NowSQLNullTime().Valid)
}

func TestTimeToDateTimeString(t *testing.T) {
	assert.Equal(t, true, TimeToDateTimeString(time.Now()) != "")
}

func TestTimeToSQLNullTime(t *testing.T) {
	now := time.Now()
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want sql.NullTime
	}{
		{
			name: "case1",
			args: args{
				t: now,
			},
			want: sql.NullTime{
				Time:  now,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToSQLNullTime(tt.args.t), "TimeToSQLNullTime(%v)", tt.args.t)
		})
	}
}

func TestTimeToHuman(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: time.Now().Add(time.Minute),
			},
			want: "1分钟后",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToHuman(tt.args.t), "TimeToHuman(%v)", tt.args.t)
		})
	}
}

func TestTimeToHumanShow(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: time.Now(),
			},
			want: "刚刚发布",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToHumanShow(tt.args.t), "TimeToHumanShow(%v)", tt.args.t)
		})
	}
}

func TestTimeToDiffForHumans(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				t: time.Now(),
			},
			want: "刚刚",
		},
		{
			name: "case 2",
			args: args{
				t: time.Now().Add(time.Hour * 86400),
			},
			want: "9 年后",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeToDiffForHumans(tt.args.t), "TimeToDiffForHumans(%v)", tt.args.t)
		})
	}
}

func TestStrToTime(t *testing.T) {
	strToTime := StrToTime("2021-01-01 00:00:00")
	assert.Equal(t, false, strToTime.IsZero())
}

func TestRFC3339(t *testing.T) {
	type args struct {
		tt time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				tt: time.Now(),
			},
			want: time.Now().Format(time.RFC3339),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RFC3339(tt.args.tt), "RFC3339(%v)", tt.args.tt)
		})
	}
}
