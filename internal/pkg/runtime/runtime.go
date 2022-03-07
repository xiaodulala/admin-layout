package runtime

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/config"
	"gorm.io/gorm"
)

var RunTime = NewDefault()

// Resource 运行时资源句柄维护
type Resource struct {
	orm    *gorm.DB
	casbin *casbin.SyncedEnforcer
	engine *gin.Engine
	config *config.Config
}

func New(db *gorm.DB, e *casbin.SyncedEnforcer, engine *gin.Engine, config *config.Config) *Resource {
	return &Resource{orm: db, casbin: e, engine: engine, config: config}
}

func NewDefault() *Resource {
	return &Resource{}
}

func (r *Resource) SetOrm(orm *gorm.DB) {
	r.orm = orm
}

func (r *Resource) SetCasbin(casbin *casbin.SyncedEnforcer) {
	r.casbin = casbin
}

func (r *Resource) SetEngine(engine *gin.Engine) {
	r.engine = engine
}

func (r *Resource) SetConfig(config *config.Config) {
	r.config = config
}

func (r *Resource) Orm() *gorm.DB {
	return r.orm
}

func (r *Resource) Casbin() *casbin.SyncedEnforcer {
	return r.casbin
}

func (r *Resource) Engine() *gin.Engine {
	return r.engine
}

func (r *Resource) Config() *config.Config {
	return r.config
}
