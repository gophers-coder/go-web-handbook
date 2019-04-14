# 文档

> 使用 Markdown

- 需求文档
    - 需求文档，由产品经理给出，反复确认
- 设计文档
    - 包括需求分析
    - 模型设计
    - 检验标准等
- API文档
    - Swagger API 文档
    - vuepress
    - docsify


## 针对开源的库

还需要提供代码文档：[godoc](https://godoc.org)

侧重在下面三点：
- Package: 项目说明文档，一般使用 doc.go 文件来说明，支持一套规范 [https://blog.golang.org/godoc-documenting-go-code](https://blog.golang.org/godoc-documenting-go-code)
- function: 函数注释和说明
- Const: 常量注释和说明

本地启动文档服务：

``` 
godoc -http="6060"
```

开源的库上传至 Github 等托管平台可以在线访问 godoc 搜索到。


## 非开源(业务)

- 需求文档
    - 实现什么功能
    - 设计原型
- 设计文档
    - 模型设计文档
        - 数据库设计
        - 表结构
        - 以及如何实现功能
    - 错误处理文档
        - 状态码信息
    - 日志处理文档
        - 日志的格式
        - 日志的路径
- API 文档
    - 路由
    - 参数
    - 响应
