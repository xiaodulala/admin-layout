package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiaodulala/admin-layout/pkg/version"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "admin-layout sys-version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(version.Get().String())
	return nil
}
