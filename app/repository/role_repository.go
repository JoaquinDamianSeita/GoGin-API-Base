package repository

import (
	"GoGin-API-Base/app/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole() ([]dao.Role, error)
	FindRoleById(id int) (dao.Role, error)
	Save(Role *dao.Role) (dao.Role, error)
	DeleteRoleById(id int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (u RoleRepositoryImpl) FindAllRole() ([]dao.Role, error) {
	var roles []dao.Role

	var err = u.db.Preload("Role").Find(&roles).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return roles, nil
}

func (u RoleRepositoryImpl) FindRoleById(id int) (dao.Role, error) {
	role := dao.Role{
		ID: id,
	}
	err := u.db.Preload("Role").First(&role).Error
	if err != nil {
		log.Error("Got and error when find Role by id. Error: ", err)
		return dao.Role{}, err
	}
	return role, nil
}

func (u RoleRepositoryImpl) Save(role *dao.Role) (dao.Role, error) {
	var err = u.db.Save(role).Error
	if err != nil {
		log.Error("Got an error when save Role. Error: ", err)
		return dao.Role{}, err
	}
	return *role, nil
}

func (u RoleRepositoryImpl) DeleteRoleById(id int) error {
	err := u.db.Delete(&dao.Role{}, id).Error
	if err != nil {
		log.Error("Got an error when delete Role. Error: ", err)
		return err
	}
	return nil
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
