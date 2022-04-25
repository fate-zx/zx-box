package logger

import "time"

// Config 日志配置对象
type Config struct {
	Debug        bool          // 是否以debug模式运行
	Name         string        // 日志写入文件名
	Level        uint32        // 日志写入级别（0:panic/1:fatal/2:error/3:warn/4:info/5:debug/6:trace）
	Caller       bool          // 是否打印调用者
	Color        bool          // 是否打印级别色彩
	MaxAge       time.Duration // 日志保留时间
	RotationTime time.Duration // 日志切割时间
}
