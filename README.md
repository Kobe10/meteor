## 流星

### 介绍

**这是一个用go开发的短链系统，特点就是快，短；所以取名流星，就像流星一样，既快又短暂；**



### 目录说明

```java
conf 环境配置文件
config 集成中间件配置文件
controllers 是控制器目录
core 核心组件
docker docker文件
docs swagger文件
global 全局变量配置
httptrack http通用组件包
initialize 组件初始化
middleware 是中间件目录 
models 是模型层，数据库相关定义
resources 是资源层，聚合数据使用
routers 是路由配置目录
service 业务逻辑代码
runtime 是程序运行时临时目录或者是日志目录
untils 是一些基础工具包
config.yaml  配置文件
go.mod    项目依赖
main.go   项目启动入口
README.md 项目必读文档
```


##功能开发(开发中)
### 基础组件	
| 服务        | 使用技术    |  进度  |  进度  |
| :--------:   | :-----:   | :----: | :----: |
|   日志   |    zap    |  ✅   |                            |
| 配置管理 |   viper   |  ✅   |                            |
| 文档管理 | Swagger2  |  ✅   |                            |
|   orm    |   gorm    |  ✅   |                            |
|   http   | httptrack |  ✅   |          自己封装          |
|  cache   |   redis   |  ✅   |                            |
|  敏感词  |  eunomia  |  ✅   |                            |
|   限流   | sentinel  |  🏗   |                            |
| 单测框架 | GoConvey  |  ✅   |                            |
| 监控告警 |  Lambda   |  🏗   |                            |
| 操作审计 |           |  🏗   | 系统关键操作日志记录和查询 |
   

### 基础功能
| 服务        | 使用技术    |  进度  |  进度  |
| :--------:   | :-----:   | :----: | :----: |
| 短链生成算法        |       |  🏗     | |
| 生成短链接口        |       |   🏗     | |
| 短链查询接口        |       |   🏗     | |
| 接口鉴权开发        |       |   🏗     | 基于OAuth2 协议  |


### 后台管理功能
| 服务        | 使用技术    |  进度  |  进度  |
| :--------:   | :-----:   | :----: | :----: |
| 短链列表        |       |  🏗     | |
| 应用管理        |       |   🏗     | |
| 短链配置管理        |       |   🏗     | |




### 项目启动

#### go 环境搭建   推荐1.15.4版本    

mac环境参考：https://mojotv.cn/2019/07/30/go-install

win环境参考：https://mojotv.cn/2019/07/30/go-install



**bash_profile配置参考**

```shell
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/Users/fuzhiqiang/go
export GOPROXY=https://goproxy.io
```



#### IDE

推荐GoLand

GoLand配置go环境

![](https://tva1.sinaimg.cn/large/0081Kckwly1gkk8cf9z3kj315a0u00xp.jpg)

GoLand 代理配置，防止下载依赖失败

![image-20201110180033069](https://tva1.sinaimg.cn/large/0081Kckwly1gkk8dr6n5aj31600u0agz.jpg)





#### gin 安装

```
go get -u github.com/gin-gonic/gin
```

中文开发文档： https://github.com/skyhee/gin-doc-cn


#### 安装项目依赖

```
go get -d -v ./...
```

如果想安装其他依赖，手动安装即可，安装命令：

```
go get -u 依赖路径名称  例如  go get -u github.com/go-redis/redis
```



### 启动项目

```
go run main.go
```

### 开发规范

#### 代码格式化规范(自动检测代码)

- goland 可以使用 go  fmt； 配置教程：https://www.cnblogs.com/chnmig/p/12446979.html

#### 单测规范(goconvey 框架)

- go的单测详细介绍：https://www.cnblogs.com/zhulipeng-1998/p/12863736.html
- goland单测使用快捷 ： http://xytschool.com/resource/297.html
- goconvey单测框架基本使用：https://www.cnblogs.com/WayneZeng/p/9290711.html

### 未完待续。。。

