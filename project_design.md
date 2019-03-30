# 项目组织

良好的项目组织，便于维护代码，便于扩展功能。

我们可以按功能划分对项目进行组织，我们可以借鉴开源项目的组织方式。

``` 
# cd project
cmd/*
configs/*
deployments/*
docs/*
examples/*
init/*
pkg/*
scripts/*
src/*
test/*
tools/*
vendor/*
.gitignore
Makefile
README.md
```

- cmd : 提供命令行工具
- configs: 配置文件目录
- deployments: 部署相关，比如 Dockerfile
- docs: 文档，比如 Swagger 文档，也可以项目的说明
- examples: 库的使用示例
- init: 初始化操作，比如读取配置文件，启动数据库等
- pkg: 中间件功能，提供服务辅助功能
- scripts: 脚本信息，通常指 shell 脚本
- src: 项目的核心处理逻辑，一般划分为三个板块：controller 控制逻辑, router 路由,assistance 辅助函数 
- test: 项目单元测试
- vendor: 第三方库
- .gitignore: 版本管理忽略文件
- Makefile: 项目构建命令文件
- README.md: 项目说明


后续根据功能划分把相应的代码放在指定的目录下。