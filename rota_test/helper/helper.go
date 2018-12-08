package helper

import (
	"encoding/binary"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"time"
)

func TestTeamMembers() []byte {
	_, filename, _, _ := runtime.Caller(1)
	basePath := filepath.Dir(filename)
	data, err := ioutil.ReadFile(basePath + "/test-team-members.yml")
	if err == nil {
		return data
	}
	panic(err)
}

func Uint16ToBytes(intVal uint16) []byte{
	byteVal := make([]byte, 2)
	binary.BigEndian.PutUint16(byteVal, intVal)
	return byteVal
}

func BytesToUint16(byteVal []byte) uint16 {
	return binary.BigEndian.Uint16(byteVal)
}

func Yesterday() string {
	return time.Now().AddDate(0,0,-1).Format("02-01-2006")
}

func DayBeforeYesterday() string {
	return time.Now().AddDate(0,0,-2).Format("02-01-2006")
}

func DaysBeforeToday(num int) string {
	return time.Now().AddDate(0,0,-num).Format("02-01-2006")
}