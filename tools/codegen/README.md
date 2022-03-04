# 功能

1. 根据声名的业务错误码自动注册到错误包(`github.com/xiaodulala/component-tools/tree/main/pkg/errors`)中
2. 自动生成错误码md格式文档

# 安装

```bash
go install github.com/xiaodulala/component-tools/tools/codegen@latest
```

# 使用

```bash
codegen -h 
# 注册业务码
codegen -type=int
# 生成文档
codegen -type=int -docs -output ./base_code_generated.md
```
