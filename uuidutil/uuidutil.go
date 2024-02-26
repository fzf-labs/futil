package uuidutil

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
	"github.com/teris-io/shortid"
)

// GenUUID 生成随机字符串，eg: 76d27e8c-a80e-48c8-ad20-e5562e0f67e4
func GenUUID() string {
	return uuid.NewString()
}

// GenShortID 生成一个id eq: 4ugh9poIR
func GenShortID() (string, error) {
	return shortid.Generate()
}

// KSUID 生成一个ksuid eq: 2clqvQKAruN0RG5olPEsALuJ1al
func KSUID() string {
	return ksuid.New().String()
}

// KSUIDByTime 生成一个基于时间的ksuid eq: 2clqzfqkaw6oexqtcsn8eakck3w
func KSUIDByTime() string {
	s, _ := ksuid.NewRandomWithTime(time.Now())
	return strings.ToLower(s.String())
}
