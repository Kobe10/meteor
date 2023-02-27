package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"meteor/global"
	"meteor/model"
	"meteor/model/request/url"
	"meteor/utils"
	"strconv"
	"time"
)

/**
短链业务处理
*/

type ShortToSourceUrl struct {
	model.ShortUrl
	SourceUrl string `json:"source_url" form:"source_url" gorm:"column:source_url;comment:长链地址"`
}

// 生成短链
func GenerateShortChains(sourceUrl *url.ShortUrlDTO, c *gin.Context) (shortUrl string, err error) {
	//1、查询redis缓存和db 是否已经生成短链 (查库加锁)
	url, duplicateErr := findShortChainsFormRedisAndDb(sourceUrl)
	if duplicateErr != nil {
		return "", duplicateErr
	}
	if len(url) != 0 {
		return url, nil
	}
	//2、生成短链 存入redis
	newUrl, generateErr := GenerateShortUrl(sourceUrl)
	if generateErr != nil {
		return "", duplicateErr
	}
	//存入redis
	timer := 60 * 60 * 24 * 7 * time.Second
	redisErr1 := global.GVA_REDIS.Set(sourceUrl.SourceUrlMD5, newUrl, timer).Err()
	redisErr2 := global.GVA_REDIS.Set(utils.MD5V([]byte(newUrl)), sourceUrl.SourceUrl, timer).Err()
	if redisErr1 != nil {
		return "", redisErr1
	}
	if redisErr2 != nil {
		return "", redisErr2
	}
	return newUrl, nil
}

// 查询已生成短链 通过源链接
func findShortChainsFormRedisAndDb(sourceUrl *url.ShortUrlDTO) (string, error) {
	// 查询redis
	//startTime := time.Now()
	cmd := global.GVA_REDIS.Get(sourceUrl.SourceUrlMD5)
	// 结束时间
	//endTime := time.Now()
	//syslog.GetLogger().WithFields(map[string]interface{}{
	//	"shortUrl": endTime.Sub(startTime).Milliseconds(),
	//}).Info(cmd.Args())

	shortUrl, err := cmd.Result()
	if err != redis.Nil {
		return shortUrl, cmd.Err()
	}

	if len(shortUrl) != 0 {
		return shortUrl, nil
	}
	// 查询DB
	var url model.ShortUrl
	err = global.GVA_DB.Where("source_url_md5 = ? and is_del = 0", sourceUrl.SourceUrlMD5).First(&url).Error
	if err != nil {
		return "", nil
	}
	return url.ShortUrl, nil
}

// 查询源url  通过短链
func FindSourceChainsFormRedisAndDb(shortUrl string) (string, error) {
	// 查询redis
	sourceUrl, err := global.GVA_REDIS.Get(utils.MD5V([]byte(shortUrl))).Result()
	if err != redis.Nil {
		return sourceUrl, nil
	}
	if len(sourceUrl) != 0 {
		return sourceUrl, nil
	}
	// 查询DB
	var shortUrlTemp model.ShortUrl
	var shortToSourceUrl ShortToSourceUrl
	dbErr := global.GVA_DB.Debug().Table(shortUrlTemp.TableName()).Select("short_url.* , short_url_config.source_url").Joins(
		"left join short_url_config on short_url.url_code = short_url_config.url_code").Find(&shortToSourceUrl, shortUrl).Error

	if dbErr != nil {
		return "", err
	}
	return shortToSourceUrl.SourceUrl, err
}

//创建短链url
func GenerateShortUrl(sourceUrl *url.ShortUrlDTO) (string, error) {
	// 生成唯一id，自增
	var incr model.Incr
	err := global.GVA_DB.Create(&incr).Error
	if err != nil {
		return "", errors.New("创建url失败")
	}
	urlCode := "SU-" + strconv.Itoa(int(incr.ID))
	// 获取短链地址
	url := utils.SpliceShortUrl(int(incr.ID))
	// 保存 shortUrl
	var shortUrlEntity model.ShortUrl
	shortUrlEntity.ShortUrl = url
	shortUrlEntity.Uid = sourceUrl.Uid
	shortUrlEntity.AppId = sourceUrl.AppId
	shortUrlEntity.SourceUrlMd5 = sourceUrl.SourceUrlMD5
	shortUrlEntity.UrlCode = urlCode
	shortUrlErr := global.GVA_DB.Create(&shortUrlEntity).Error
	if shortUrlErr != nil {
		return "", errors.New("创建短url失败")
	}
	// 保存shortUrlConfig
	var shortUrlConfigEntity model.ShortUrlConfig
	shortUrlConfigEntity.UrlCode = urlCode
	shortUrlConfigEntity.SourceUrl = sourceUrl.SourceUrl
	shortUrlConfigEntity.EncryptEnable = 0
	shortUrlConfigEntity.ValidEnable = 1

	shortUrlConfigErr := global.GVA_DB.Create(&shortUrlConfigEntity).Error
	if shortUrlConfigErr != nil {
		return "", errors.New("创建短url失败")
	}

	return url, nil
}
