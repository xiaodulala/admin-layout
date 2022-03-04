package globalflag

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// AddGlobalFlags 显式注册库(log, verflag等)的全局标志。
// 我们这样做是为了防止不需要的标志泄漏到组件的标志集。
func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}

// normalize 用连字符替换下划线，我们在注册组件标志时应该总是使用连字符而不是下划线
func normalize(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

// Register 注册globalName
func Register(local *pflag.FlagSet, globalName string) {
	if f := flag.CommandLine.Lookup(globalName); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		pflagFlag.Name = normalize(pflagFlag.Name)
		local.AddFlag(pflagFlag)
	} else {
		panic(fmt.Sprintf("failed to find flag in global flagset (flag): %s", globalName))
	}
}
