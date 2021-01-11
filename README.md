
# tigi





### 初始化
```
git clone https://github.com/yunhu/tigi.git
cd tigi
make init
```
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


### 基本结构
```
├── main.go                 #入口文件
├── client/                 #资源连接
├── common/                 #公共函数等
├── config/                 #资源连接
│   └── config.go           #配置结构 
├── controller/             #Api入口
├── cron/                   #计划任务
├── log/                    #日志目录
├── model/                  #model层
├── output/                 #编译打包转出目录
├── router/                 #路由配置
├── st/                     #数据结构定义
├── worker/                 #协程池
```