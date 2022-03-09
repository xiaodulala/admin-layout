package service

import (
	"github.com/casbin/casbin/v2"
	cDto "github.com/xiaodulala/admin-layout/internal/application/common/dto"
	"github.com/xiaodulala/admin-layout/internal/application/common/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/internal/pkg/code"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SysRole struct {
	service.Service
}

// GetPage 获取SysRole列表
func (e *SysRole) GetPage(c *dto.SysRoleGetPageReq, list *[]models.SysRole, count *int64) error {
	var err error
	var data models.SysRole

	err = e.Orm.Model(&data).Preload("SysMenu").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Get 获取SysRole对象
func (e *SysRole) Get(d *dto.SysRoleGetReq, model *models.SysRole) error {
	var err error
	db := e.Orm.First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, "查看对象不存在或无权查看")
	}
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	model.MenuIds, err = e.GetRoleMenuId(model.RoleId)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Insert 创建SysRole对象
func (e *SysRole) Insert(c *dto.SysRoleInsertReq, cb *casbin.SyncedEnforcer) error {
	var err error
	var data models.SysRole
	var dataMenu []models.SysMenu
	err = e.Orm.Preload("SysApi").Where("menu_id in ?", c.MenuIds).Find(&dataMenu).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	c.SysMenu = dataMenu
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}

	for _, menu := range dataMenu {
		for _, api := range menu.SysApi {
			_, err = cb.AddNamedPolicy("p", data.RoleKey, api.Path, api.Action)
		}
	}
	_ = cb.SavePolicy()
	return nil
}

// Update 修改SysRole对象
func (e *SysRole) Update(c *dto.SysRoleUpdateReq, cb *casbin.SyncedEnforcer) error {
	var err error
	tx := e.Orm
	var model = models.SysRole{}
	var mlist = make([]models.SysMenu, 0)
	tx.Preload("SysMenu").First(&model, c.GetId())
	tx.Preload("SysApi").Where("menu_id in ?", c.MenuIds).Find(&mlist)
	err = tx.Model(&model).Association("SysMenu").Delete(model.SysMenu)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	c.Generate(&model)
	model.SysMenu = &mlist
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)

	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")
	}

	_, err = cb.RemoveFilteredPolicy(0, model.RoleKey)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	mp := make(map[string]interface{}, 0)
	polices := make([][]string, 0)
	for _, menu := range mlist {
		for _, api := range menu.SysApi {
			if mp[model.RoleKey+"-"+api.Path+"-"+api.Action] != "" {
				mp[model.RoleKey+"-"+api.Path+"-"+api.Action] = ""
				//_, err = cb.AddNamedPolicy("p", model.RoleKey, api.Path, api.Action)
				polices = append(polices, []string{model.RoleKey, api.Path, api.Action})
			}
		}
	}
	_, err = cb.AddNamedPolicies("p", polices)
	if err != nil {
		return err
	}
	_ = cb.SavePolicy()
	return nil
}

// Remove 删除SysRole
func (e *SysRole) Remove(c *dto.SysRoleDeleteReq) error {

	tx := e.Orm

	var model = models.SysRole{}
	tx.Preload("SysMenu").Preload("SysDept").First(&model, c.GetId())
	db := tx.Select(clause.Associations).Delete(&model)

	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")
	}
	return nil
}

// GetRoleMenuId 获取角色对应的菜单ids
func (e *SysRole) GetRoleMenuId(roleId int) ([]int, error) {
	menuIds := make([]int, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		menuIds = append(menuIds, l[i].MenuId)
	}
	return menuIds, nil
}

func (e *SysRole) UpdateDataScope(c *dto.RoleDataScopeReq) error {
	var err error
	tx := e.Orm
	var dlist = make([]models.SysDept, 0)
	var model = models.SysRole{}
	tx.Preload("SysDept").First(&model, c.RoleId)
	tx.Where("id in ?", c.DeptIds).Find(&dlist)
	err = tx.Model(&model).Association("SysDept").Delete(model.SysDept)
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	c.Generate(&model)
	// fixme 无法插入角色部门关联表
	model.SysDept = &dlist
	db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")
	}
	return nil
}

// UpdateStatus 修改SysRole对象status
func (e *SysRole) UpdateStatus(c *dto.UpdateStatusReq) error {
	tx := e.Orm
	var model = models.SysRole{}
	tx.First(&model, c.GetId())
	c.Generate(&model)
	db := tx.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")
	}
	return nil
}

// GetById 获取SysRole对象
func (e *SysRole) GetById(roleId int) ([]string, error) {
	permissions := make([]string, 0)
	model := models.SysRole{}
	model.RoleId = roleId
	if err := e.Orm.Model(&model).Preload("SysMenu").First(&model).Error; err != nil {
		return nil, err
	}
	l := *model.SysMenu
	for i := 0; i < len(l); i++ {
		permissions = append(permissions, l[i].Permission)
	}
	return permissions, nil
}
