package svc

import (
	"server/internal/config"
	"server/internal/middleware"
	"server/model/admin_dic"
	"server/model/admin_dic_detail"
	"server/model/admin_operation"
	"server/model/anthority_model"
	"server/model/menu"
	"server/model/user_auth"
	"server/model/user_model"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       user_model.SysUserModel
	MenuModel       menu.SysBaseMenusModel
	Authority       anthority_model.SysAuthoritiesModel
	AuthUser        user_auth.SysUserAuthorityModel
	DicModel        admin_dic.SysDictionariesModel
	DicDetail       admin_dic_detail.SysDictionaryDetailsModel
	Operation       admin_operation.SysOperationRecordsModel
	Rdb             *redis.Client
	OperationRecord rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: user_model.NewSysUserModel(conn),
		MenuModel: menu.NewSysBaseMenusModel(conn),
		Authority: anthority_model.NewSysAuthoritiesModel(conn),
		AuthUser:  user_auth.NewSysUserAuthorityModel(conn),
		DicModel:  admin_dic.NewSysDictionariesModel(conn),
		DicDetail: admin_dic_detail.NewSysDictionaryDetailsModel(conn),
		Operation: admin_operation.NewSysOperationRecordsModel(conn),
		Rdb: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Host + ":" + c.Redis.Port,
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
		OperationRecord: middleware.NewOperationRecordMiddleware().Handle,
	}
}
