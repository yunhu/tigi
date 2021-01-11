## 初始化
```
git clone https://github.com/yunhu/tigi.git
cd tigi
make init
```
## tigi

### 构建
```
make 
```
### 启动
```
make run
```

### 停止
```
make stop
```

### 配置
1. 是否开发模式是通过全局变量GODEVELOPMODE=dev来设置的，其它情况(默认)都为线上，会加载对应的配置
```shell
export GODEVELOPMODE=dev
```
