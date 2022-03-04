package options

import (
	"fmt"
	"github.com/spf13/pflag"
)

type ServerOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	BindAddress string   `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int      `json:"bind-port"    mapstructure:"bind-port"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Mode:        "debug",
		BindAddress: "127.0.0.1",
		BindPort:    9099,
		Middlewares: []string{},
	}
}

func (o *ServerOptions) Validate() []error {
	var errs []error

	if o.BindPort < 0 || o.BindPort > 65535 {
		errs = append(errs,
			fmt.Errorf("--dss.bind-port %v must be between 0 and 65535, inclusive. 0 for turning off insecure (HTTP) port", o.BindPort))
	}

	return errs
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Mode, "server.mode", o.Mode, ""+
		"Start the server in a specified server mode. Supported server mode: debug, test, release.")
	fs.StringVar(&o.BindAddress, "server.bind-address", o.BindAddress, "set to 0.0.0.0 for all IPv4 interfaces.")
	fs.IntVar(&o.BindPort, "server.bind-port", o.BindPort, "set port with bind.address.")
	fs.StringSliceVar(&o.Middlewares, "server.middlewares", o.Middlewares, ""+
		"List of allowed middlewares for server, comma separated. If this list is empty default middlewares will be used.")
}
