package app

import (
	cliflag "github.com/xiaodulala/admin-layout/pkg/cliflag/flag"
)

// CliOptions 命令行选项参数接口
type CliOptions interface {
	// Flags 填充应用程序选项参数。并返回整个集合。
	Flags() (fss cliflag.NamedFlagSets)
	// Validate 检测给定的选项参数合法性
	Validate() []error
}

// ConfigurableOptions 从配置文件配置选项参数
type ConfigurableOptions interface {
	ApplyFlags() []error
}

// CompleteableOptions 选项补全接口
type CompleteableOptions interface {
	Complete() error
}

// PrintableOptions 选项参数打印接口
type PrintableOptions interface {
	String() string
}
