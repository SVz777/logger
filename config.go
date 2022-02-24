// Package log
// file    config.go
// @author
//  ___  _  _  ____
// / __)( \/ )(_   )
// \__ \ \  /  / /_
// (___/  \/  (____)
// (903943711@qq.com)
// @date    2022/2/24
// @desc

package logger

type Config []OutputConfig
type WriteConfig struct {
	// 日志文件路径
	Filename string `toml:"filename"`
	// 每个日志文件保存的最大尺寸 单位：M
	MaxSize int `toml:"max_size"`
	// 文件最多保存多少天
	MaxAge int `toml:"max_age"`
	// 日志文件最多保存多少个备份
	MaxBackups int `toml:"max_backups"`
	// 是否压缩
	Compress bool `toml:"compress"`
}

type OutputConfig struct {
	// 支持console,file
	Writer string `toml:"writer"`
	// file时配置
	WriteConfig WriteConfig `toml:"write_config"`
	// 支持console,json
	Formatter string `toml:"formatter"`
	Level     string `toml:"level"`
	// 支持 > >= < <= 不填默认>=
	LevelOp string `toml:"level_op"`
}
