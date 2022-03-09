package sys_version

import (
	"github.com/xiaodulala/admin-layout/internal/appctl/cmd/migrate/migration"
	"runtime"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683660Tables)
}

func _1599190683660Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
