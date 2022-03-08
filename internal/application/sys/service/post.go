package service

import (
	cDto "github.com/xiaodulala/admin-layout/internal/application/common/dto"
	"github.com/xiaodulala/admin-layout/internal/application/common/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/internal/pkg/code"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"gorm.io/gorm"
)

type SysPost struct {
	service.Service
}

// GetPage 获取SysPost列表
func (e *SysPost) GetPage(c *dto.SysPostPageReq, list *[]models.SysPost, count *int64) error {
	var err error
	var data models.SysPost

	err = e.Orm.Model(&data).
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

// Get 获取SysPost对象
func (e *SysPost) Get(d *dto.SysPostGetReq, model *models.SysPost) error {
	var err error
	var data models.SysPost

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrDatabase, "查看对象不存在或无权查看")
	}
	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	return nil
}

// Insert 创建SysPost对象
func (e *SysPost) Insert(c *dto.SysPostInsertReq) error {
	var err error
	var data models.SysPost
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	c.PostId = data.PostId
	return nil
}

// Update 修改SysPost对象
func (e *SysPost) Update(c *dto.SysPostUpdateReq) error {

	var model = models.SysPost{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)

	db := e.Orm.Save(&model)
	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")

	}
	return nil
}

// Remove 删除SysPost
func (e *SysPost) Remove(d *dto.SysPostDeleteReq) error {
	var data models.SysPost

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if db.Error != nil {
		return errors.WithCode(code.ErrDatabase, db.Error.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权删除该数据")
	}
	return nil
}
