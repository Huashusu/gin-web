package core

import (
	"gin-web/global"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotateLog = new(fileWriteLog)

const (
	Day = time.Hour * 24
	_   = iota
	KB  = 1 << (10 * iota)
	MB  = 1 << (10 * iota)
	GB  = 1 << (10 * iota)
)

type fileWriteLog struct{}

// GetWriterSyncer 根据日志不同级别获取写入的文件流
func (f *fileWriteLog) GetWriterSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWrite, err := rotatelogs.New(
		path.Join(global.CONFIG.Log.Dir, "%Y-%m-%d", level+".log"),
		rotatelogs.ForceNewFile(),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.CONFIG.Log.MaxRetentionDays)*Day),
		rotatelogs.WithRotationTime(Day*1),
		rotatelogs.WithRotationSize(MB*100),
	)
	if global.CONFIG.Log.ConsoleOutput {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWrite)), err
	}
	return zapcore.AddSync(fileWrite), err
}
