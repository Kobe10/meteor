#### 日志包

##### 使用说明
    日志包是对logrus包的进一步封装
    包含功能
        1.日志分类记录 error warn info
        2.日志自动切割
        3.记录请求ID - 通过传参 context.Context
        4.gin框架日志记录中间件
            a.初始化请求ID
            b.记录请求info

##### 文件说明
|文件|简介|状态|作者|备注|
| :----: | :----: | :----: |:----:|:----:|
|logrus.go|日志初始化|done|-|-|
|logger.go|logger接口定义及实现|done|-|-|
|log_handler.go|组合此接口便拥有了日志能力|done|-|-|
|gin_logger_middleware.go|gin框架日志中间件|done|-|-|


##### 示例代码

###### 1.初始化日志模块
    syslog.Setup(syslog.LoggerConfig{
        LogPath: config.GetString("log_path"),
        LogFile: config.GetString("log_file"),
    })
    	
###### 2.初始化日志模块

    gin.SetMode(config.GetString("gin_mode"))
    //fmt.Println("gin_mode=", gin.Mode())

    r := gin.New()

    //gin中间件
    //r.Use(gin.Logger())

    //自定义日志中间件
    r.Use(syslog.LoggerToFile())
    
###### 3.结构体组合logHandler
    //组合
    type alternativeListModel struct {
        syslog.LogHandler
    }
    
    func NewAlternativeListModel(ctx context.Context) *alternativeListModel {
        a := new(alternativeListModel)
        a.SetLogger(ctx)
        return a
    }
    
    //实例化
    alternativeList := models.NewAlternativeListModel(c.Request.Context())
    
    //使用
    func (a *alternativeListModel) GetAlternativeTotal(maps interface{}) (int, error) {
    	var count int
    	if err := DB.Model(&AlternativeList{}).Where(maps).Count(&count).Error; err != nil {
    		a.GetLogger().Error(err.Error())
    		return 0, err
    	}
    
    	return count, nil
    }
###### 4.直接使用日志功能
    //使用此方法便于记录请求ID - 建议 
    syslog.NewLogger(ctx)
    
    //直接使用不记录请求ID - 不建议
    syslog.GetLogger()


    