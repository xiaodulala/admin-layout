package flag

import (
	goflag "flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// WordSepNormalizeFunc 标准化标识集名称 下划线修改为中划线
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc 带下划线的选项增加警告日志
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.ReplaceAll(name, "_", "-")
		fmt.Printf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, nname)
		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// InitFlags 初始化标识集
func InitFlags(flags *pflag.FlagSet) {
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}

// PrintFlags 打印标识集所有选项
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(f *pflag.Flag) {
		fmt.Printf("FLAG: --%s=%q\n", f.Name, f.Value)
	})
}
