#### 日志包

##### 使用说明
    通用响应包
    包含功能
        1.json解析 - 优化版本数据兼容性好

##### 文件说明
|文件|简介|状态|作者|备注|
| :----: | :----: | :----: |:----:|:----:|
|json_parser.go|json解析|done|-|-|


##### 示例代码

###### 1.解析http响应json到结构体
	if err := responses.GetJsonParser().NewDecoder(resp.Body).Decode(&response); err != nil {

		p.GetLogger().Error("Json Decode Err:", err)
		return nil, err
	}
    	


    