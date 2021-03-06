package verflag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/xiaodulala/admin-layout/pkg/version"
)

// VersionValue 自定义flag type
type VersionValue int

// Define some const.
const (
	VersionFalse VersionValue = 0
	VersionTrue  VersionValue = 1
	VersionRaw   VersionValue = 2
)

const strRawVersion string = "raw"

// IsBoolFlag comment
func (v *VersionValue) IsBoolFlag() bool {
	return true
}

// Get comment
func (v *VersionValue) Get() interface{} {
	return v
}

// Set comment
func (v *VersionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

// String 格式化字符串
func (v *VersionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Type The type of the flag as required by the pflag.Value interface.
func (v *VersionValue) Type() string {
	return "sys-version"
}

// VersionVar defines a flag with the specified name and usage string.
func VersionVar(p *VersionValue, name string, value VersionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	// "--sys-version" will be treated as "--sys-version=true"
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version wraps the VersionVar function.
func Version(name string, value VersionValue, usage string) *VersionValue {
	p := new(VersionValue)
	VersionVar(p, name, value, usage)
	return p
}

const versionFlagName = "sys-version"

var versionFlag = Version(versionFlagName, VersionFalse, "Print sys-version information and quit.")

// AddFlags registers this package's flags on arbitrary FlagSets, such that they point to the
// same value as the global flags.
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested will check if the -sys-version flag was passed
// and, if so, print the sys-version and exit.
func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
