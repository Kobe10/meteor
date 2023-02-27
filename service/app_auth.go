package service

import (
	"meteor/global"
	"meteor/model"
)

/**
通过appId 查询app实体
*/
//@author: fuzq
//@function: getAppByAppId
//@description: 通过appId 查询app实体
//@param: appId string
//@return: model.App
func FindAppByAppId(appId string) (app model.App, err error) {
	err = global.GVA_DB.Where("app_id = ? and is_del = 0 and status = 1", appId).First(&app).Error
	return app, err
}
