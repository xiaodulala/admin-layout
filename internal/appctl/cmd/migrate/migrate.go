package migrate

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiaodulala/admin-layout/internal/appctl/cmd/migrate/migration"
	_ "github.com/xiaodulala/admin-layout/internal/appctl/cmd/migrate/migration/bus-version"
	"github.com/xiaodulala/admin-layout/internal/appctl/cmd/migrate/migration/models"
	_ "github.com/xiaodulala/admin-layout/internal/appctl/cmd/migrate/migration/sys-version"
	"github.com/xiaodulala/admin-layout/pkg/db/mysql"
	"gorm.io/gorm"
)

var (
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "go-admin migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	fmt.Println(color.GreenString("running..."))
	initConfig()
	initDB()
}

func initConfig() {
	fmt.Println(color.GreenString("init config..."))
	// viper 加载配置文件
	viper.SetConfigName("appctl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println(color.GreenString("init db..."))
	db, err := mysql.New(&mysql.Options{
		Host:                  viper.GetString("mysql.host"),
		Username:              viper.GetString("mysql.username"),
		Password:              viper.GetString("mysql.password"),
		Database:              viper.GetString("mysql.database"),
		MaxIdleConnections:    viper.GetInt("mysql.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("mysql. max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("mysql.max-connection-life-time"),
		LogLevel:              viper.GetInt("mysql.log-level"),
	})
	if err != nil {
		panic(err)
	}
	if err = migrateModel(db); err != nil {
		panic(err)
	}
}

func migrateModel(db *gorm.DB) error {

	err := db.Debug().AutoMigrate(&models.Migration{})
	if err != nil {
		panic(err)
	}

	migration.Migrate.SetDb(db)
	migration.Migrate.Migrate()

	fmt.Println(color.GreenString("init db success..."))
	return nil
}
