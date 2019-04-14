# 测试

## 自测

> 自己本地些的 API如何测试

- 本地启动相关服务，使用 Postman 接口自测

## 单元测试

> 对代码的检测

- [参考](https://github.com/wuxiaoxiaoshen/go-how-to-write-test)

## 性能测试

- 基准测试
- 结合相关图表展示，对有问题的地方进行优化


- pprof
```text
import _ "net/http/pprof"


go tool pprof ***/debug/pprof/profile

top20

top -cum

```

## 指标

- Codecov 统计代码覆盖率和相关文件的分析