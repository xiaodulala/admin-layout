package options

import (
	cliflag "github.com/xiaodulala/admin-layout/pkg/cliflag/flag"
	"github.com/xiaodulala/admin-layout/pkg/db/mysql"
	"github.com/xiaodulala/admin-layout/pkg/json"
	"github.com/xiaodulala/admin-layout/pkg/log"
)

type Options struct {
	ServerOptions *ServerOptions `json:"server" mapstructure:"server"`
	MysqlOptions  *mysql.Options `json:"mysql"    mapstructure:"mysql"`
	LogOptions    *log.Options   `json:"log" mapstructure:"log"`
}

func NewOptions() *Options {
	return &Options{
		ServerOptions: NewServerOptions(),
		MysqlOptions:  mysql.NewMySQLOptions(),
		LogOptions:    log.NewOptions(),
	}
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ServerOptions.AddFlags(fss.FlagSet("server"))
	o.MysqlOptions.AddFlags(fss.FlagSet("mysql"))
	o.LogOptions.AddFlags(fss.FlagSet("logs"))
	return
}

func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.ServerOptions.Validate()...)
	errs = append(errs, o.MysqlOptions.Validate()...)
	errs = append(errs, o.LogOptions.Validate()...)

	return errs
}

func (o *Options) String() string {
	data, _ := json.MarshalIndent(o, "", "")
	return string(data)
}
