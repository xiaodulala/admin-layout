package ginauth

import (
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/utils"
	"gorm.io/gorm"
)

type Login struct {
	Username string ` json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
	//Code     string ` json:"code" binding:"required"`
	//UUID     string ` json:"uuid" binding:"required"`
}

func (u *Login) GetUser(tx *gorm.DB) (user models.SysUser, role models.SysRole, err error) {
	err = tx.Table("sys_user").Where("username = ?  and status = 2", u.Username).First(&user).Error
	if err != nil {
		log.Errorf("get user error, %s", err.Error())
		return
	}
	_, err = utils.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		log.Errorf("user login error, %s", err.Error())
		return
	}
	err = tx.Table("sys_role").Where("role_id = ? ", user.RoleId).First(&role).Error
	if err != nil {
		log.Errorf("get role error, %s", err.Error())
		return
	}
	return
}
