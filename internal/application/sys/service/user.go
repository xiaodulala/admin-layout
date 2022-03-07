package service

import (
	"github.com/xiaodulala/admin-layout/internal/application/common/actions"
	cDto "github.com/xiaodulala/admin-layout/internal/application/common/dto"
	"github.com/xiaodulala/admin-layout/internal/application/common/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/internal/pkg/code"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"github.com/xiaodulala/admin-layout/pkg/utils"
	"gorm.io/gorm"
)

type SysUser struct {
	service.Service
}

// GetPage 获取SysUser列表
func (su *SysUser) GetPage(c *dto.SysUserGetPageReq, p *actions.DataPermission, list *[]models.SysUser, count *int64) error {
	var err error
	var data models.SysUser

	err = su.Orm.Debug().Preload("Dept").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Get 获取SysUser对象
func (su *SysUser) Get(d *dto.SysUserById, p *actions.DataPermission, model *models.SysUser) error {
	var data models.SysUser

	err := su.Orm.Model(&data).Debug().
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(code.ErrUserNotFound, "查看对象不存在或无权查看")
	}
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Insert 创建SysUser对象
func (su *SysUser) Insert(c *dto.SysUserInsertReq) error {
	var err error
	var data models.SysUser
	var i int64
	err = su.Orm.Model(&data).Where("username = ?", c.Username).Count(&i).Error
	if err != nil {

		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if i > 0 {
		return errors.WithCode(code.ErrUserAlreadyExist, "用户名已存在")
	}
	c.Generate(&data)
	err = su.Orm.Create(&data).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Update 修改SysUser对象
func (su *SysUser) Update(c *dto.SysUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := su.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	update := su.Orm.Model(&model).Where("user_id = ?", &model.UserId).Omit("password", "salt").Updates(&model)
	if err = update.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if update.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "update userinfo error")
	}
	return nil
}

// UpdateAvatar 更新用户头像
func (su *SysUser) UpdateAvatar(c *dto.UpdateSysUserAvatarReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := su.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	c.Generate(&model)
	err = su.Orm.Save(&model).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// UpdateStatus 更新用户状态
func (su *SysUser) UpdateStatus(c *dto.UpdateSysUserStatusReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := su.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")

	}
	c.Generate(&model)
	err = su.Orm.Save(&model).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// ResetPwd 重置用户密码
func (su *SysUser) ResetPwd(c *dto.ResetSysUserPwdReq, p *actions.DataPermission) error {
	var err error
	var model models.SysUser
	db := su.Orm.Scopes(
		actions.Permission(model.TableName(), p),
	).First(&model, c.GetId())
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权更新该数据")
	}
	c.Generate(&model)
	err = su.Orm.Save(&model).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	return nil
}

// Remove 删除SysUser
func (su *SysUser) Remove(c *dto.SysUserById, p *actions.DataPermission) error {
	var err error
	var data models.SysUser

	db := su.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "无权删除该数据")
	}
	return nil
}

// UpdatePwd 修改SysUser对象密码
func (su *SysUser) UpdatePwd(id int, oldPassword, newPassword string, p *actions.DataPermission) error {
	var err error

	if newPassword == "" {
		return nil
	}
	c := &models.SysUser{}

	err = su.Orm.Model(c).
		Scopes(
			actions.Permission(c.TableName(), p),
		).Select("UserId", "Password", "Salt").
		First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.WithCode(code.ErrDatabase, "无权更新该数据")
		}
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	var ok bool
	ok, err = utils.CompareHashAndPassword(c.Password, oldPassword)
	if err != nil {
		return errors.WithCode(code.ErrUserPasswordWrong, err.Error())
	}
	if !ok {
		return errors.WithCode(code.ErrUserPasswordWrong, "incorrect Password")
	}
	c.Password = newPassword
	db := su.Orm.Model(c).Where("user_id = ?", id).Select("Password", "Salt").Updates(c)
	if err = db.Error; err != nil {
		return errors.WithCode(code.ErrDatabase, err.Error())
	}
	if db.RowsAffected == 0 {
		return errors.WithCode(code.ErrDatabase, "set password error")
	}
	return nil
}

func (su *SysUser) GetProfile(c *dto.SysUserById, user *models.SysUser, roles *[]models.SysRole, posts *[]models.SysPost) error {
	err := su.Orm.Preload("Dept").First(user, c.GetId()).Error
	if err != nil {
		return err
	}
	err = su.Orm.Find(roles, user.RoleId).Error
	if err != nil {
		return err
	}
	err = su.Orm.Find(posts, user.PostIds).Error
	if err != nil {
		return err
	}

	return nil
}
