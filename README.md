# dsl

构建一个简单的解释器

参考 https://ruslanspivak.com/lsbasi-part1/




# 场景举例

1. 希望提供一个API，需要根据特定规则获取用户总数，规则示例如下

```
最近7天登陆过 并且（ 
    设备为IOS  
    或者（ 
            设备为Android   
            并且 版本在 （1.1 / 1.2 / 1.3）中
        )
    )
    并且 最近一次登录地区为海外 

```

常规方案

```
POST /user/count

需要构造这样的查询JSON数据结构，很复杂
```

DSL自定义规则方案
```
POST /user/count

{
    "rule":"last_login_time > 当前时间减去7天的时间戳 AND (device_type = \"IOS\" OR (device_type = \"ANDROID\" AND app_version in (\"1.0\",\"1.1\",\"1.2\"))) AND last_login_area NOT mainland"
}
//美化一下语法
last_login_time > 当前时间减去7天的时间戳 
AND (
    device_type = "IOS" 
    OR (
        device_type = "ANDROID" 
        AND 
        app_version in ("1.0","1.1","1.2")
        )
    ) 
AND last_login_area NOT mainland"
```

