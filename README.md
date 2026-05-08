# api-service-template

## 初始化
```
# 使用当前目录名作为新 module（例如目录叫 my-api → module 变为 my-api）
./rename_module.sh
# 显式指定 module 路径（推荐：与远端仓库 import 路径一致）
./rename_module.sh github.com/you/my-api
# 只看会改哪些文件，不写盘
./rename_module.sh --dry-run github.com/you/my-api
```

## 运行
```
make serve
```

## 构建
```
make build
```
