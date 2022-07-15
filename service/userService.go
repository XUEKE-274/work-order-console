package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
	"time"
	"work-order-console/dao"
	"work-order-console/domain/entity"
	"work-order-console/domain/enum"
	"work-order-console/domain/enum/errorCodeEnum"
	"work-order-console/domain/enum/roleEnum"
	"work-order-console/exception"
	"work-order-console/logger"
	"work-order-console/utils"
	"work-order-console/web/request"
)


type UserServiceApi interface {
	AddUser(username string, password string, role enum.RoleEnum, cxt context.Context) error
	ModifyUser(id string, role enum.RoleEnum) error
	DelById(id string) error
	QueryById(id string) (*entity.UserEntity, error)
	QueryByUsername(username string) (*entity.UserEntity, error)
	ListUser(params *request.UserListRequest) (*[]entity.UserEntity, int64, error)
	GetByUsername(username string) (*entity.UserEntity,error)
	MatchCredential(pwd string, dbPwd string) bool
	GetByUsernameMysql(username string) (*entity.UserEntity,error)
}


type userService struct {
	userDao dao.UserDaoApi
	logger logger.NewLogger
	db *gorm.DB
}



var regUserService = fx.Provide(func(userDao dao.UserDaoApi, logger logger.NewLogger, db *gorm.DB) UserServiceApi {
	return &userService{
		userDao,
		logger,
		db,
	}
})

func (mine *userService)GetByUsername(username string) (*entity.UserEntity,error) {
	r,e := mine.userDao.QueryByUsername(username)
	if e != nil {
		return &entity.UserEntity{} , e
	}
	return r, nil
}

func (mine *userService)GetByUsernameMysql(username string) (*entity.UserEntity,error) {

	var u entity.UserEntity


	mine.db.First(&u).Where("username = {}", username)



	return &u, nil
}


func (mine *userService) MatchCredential(pwd string, dbPwd string) bool  {
	// dbPwd = hash (16) + salt(16)
	hash := dbPwd[0:64]
	salt := dbPwd[64:96]

	source := pwd + salt

	h := sha256.New()
	h.Write([]byte(source))
	sum := h.Sum(nil)
	userPwdHash := hex.EncodeToString(sum)
	return userPwdHash == hash
}

func (mine *userService) AddUser(username string, password string, role enum.RoleEnum, cxt context.Context) error {
	// 分布式锁 todo
	log := mine.logger.NewInstance(cxt)
	_, e := mine.QueryByUsername(username)
	if e!= nil {
		return doAdd(username, password, role, mine.userDao)
	}

	log.Info("username has exist , username = ", username)
	return exception.NewException(errorCodeEnum.USERNAME_HAS_EXIST)

}

func doAdd(username string, password string, role enum.RoleEnum, userDao dao.UserDaoApi) error{
	var user = entity.UserEntity{}
	//user.Id = primitive.NewObjectID()
	user.Username = username
	user.EnPwd = buildEnPwd(password)
	user.Role = role
	user.CreateBy = ""
	user.ModifyBy = ""
	user.CreateTime = time.Now()
	user.ModifyTime = time.Now()
	return userDao.Save(&user)
}

func buildEnPwd(password string) string {
	salt := utils.NewUuid()
	target := password + salt
	h := sha256.New()
	h.Write([]byte(target))
	sum := h.Sum(nil)
	userPwdHash := hex.EncodeToString(sum)
	return userPwdHash + salt
}

func (mine *userService) ModifyUser(id string, role enum.RoleEnum) error{
	u, e := mine.QueryById(id)
	if e != nil {
		return e
	}
	// 系统管理源不能被编辑
	if u.Role == roleEnum.SYSTEM {
		return exception.NewException(errorCodeEnum.SYS_USER_CAN_NOT_OPEAT)
	}
	// 不能将普通用户修改为系统管理员
	if role == roleEnum.SYSTEM {
		return exception.NewException(errorCodeEnum.NOT_PERMISSION)
	}
	// do update
	return mine.userDao.UpdateById(id, role)
}

func (mine *userService) DelById(id string) error {
	u, e := mine.QueryById(id)
	if e != nil {
		return e
	}
	// 系统管理源不能被删除
	if u.Role == roleEnum.SYSTEM {
		return exception.NewException(errorCodeEnum.SYS_USER_CAN_NOT_OPEAT)
	}
	// do delete
	return mine.userDao.RemoveById(id)
}

func (mine *userService) QueryById(id string) (*entity.UserEntity, error) {
	return mine.userDao.QueryById(id)
}

func (mine *userService) ListUser(params *request.UserListRequest) (*[]entity.UserEntity, int64, error) {

	return mine.userDao.PageQuery(params)
}

func (mine *userService) QueryByUsername(username string) (*entity.UserEntity, error) {
	return mine.userDao.QueryByUsername(username)

}
