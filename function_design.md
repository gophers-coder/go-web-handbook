# 单功能开发

> 编写符合 Restful 风格的 API

- 确保你已经正确的理解了需求
- 确保你已经正确的设计了模型



## 路由设计

- 版本设计: (v1、v2、v3...)
- 资源设计:
- 路径参数:

``` 
https://api.github.com/users/{username}
```

- users 资源，通常使用名词
- username 路径参数

## 参数检验

通常配合定义结构体的 Tag 来进行设置：[validator](https://godoc.org/gopkg.in/go-playground/validator.v9)


```go 
type CreateParams struct {
    Name string `form:"name" json:"name" binding:"required"`
    Number int `form:"number" json:"name" binding:"max=100"`
    Phone string `form:"phone" json:"phone" binding:"len=11"`
}
```

这样做的好处是，你不需要在处理逻辑中编写相应的函数来检验这些参数。

更多用法：[validator](https://godoc.org/gopkg.in/go-playground/validator.v9)


## 资源操作

- 新增资源: POST
- 更新资源: PATCH
- 删除资源: DELETE
- 获取资源: GET

注意几点：
- 选择正确的方法（POST，PATCH，DELETE，GET）
- 设置相应的错误处理码：遇到什么样的场景返回相应的状态码和具体信息
- 设置正确的响应信息：一般对 model 进行 序列化，返回指定的信息

