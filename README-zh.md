# Nacos 单机/集群的vip地址服务中心

## 简介

&emsp;&emsp;&emsp;基于Nacos（官方网站:http://nacos.io ）的额外web服务器，针对`服务端`和`客户端`地址寻址，减少改动(服务端和客户端项目)，方便动态扩容和管理。
适用于自建Nacos的单机或集群管理，[阿里云的MSE微服务引擎托管](https://cn.aliyun.com/product/aliware/mse)则不需要考虑，官方已经处理好。<br/><br/>
&emsp;&emsp;&emsp;在去年公司需要使用配置中心时，通过调研和选型，最终使用Nacos来作为配置中心和注册中心。在使用的过程中，我们
也发现了一些问题，在客户端项目中(spring cloud)和其他客户端(SpringBoot，Go，Node.js，Python...等等)我们在配置服务器地址时 `spring.cloud.nacos.config.serverAddr=127.0.0.1:8848,127.0.0.2:8848,127.0.0.3:8848`，
如果我们有100+个客户端，地址发生变化，是不是需要对100个客户端都进行更改，这样的方式并不灵活。官方文档和博客上也没有讲的很细，基本只有配置`serverAddr`的方式，而服务端集群模式下，下线或上线新的服务端机器，
我们需要在`cluster.conf`进行地址变更，也会出现类似的场景，所以去看源码和调研，产生了这个项目。<br/><br/>
&emsp;&emsp;&emsp;适配客户端和服务端，提供使用方式和部署文档、对不同存储方式的地址列表的统一API管理、对Docker的支持等等，同时也欢迎大家使用、建议、并贡献。如果对大家有所帮助或将来有所帮助，欢迎`Star`一下哦

## 存储支持

- [x] Redis (单机+集群) 

- [x] Cache (单机)

- [x] File (单机)

## 功能

- [x] Docker支持 

## 参数

  * nacos-address系统参数(environment)
     
 | 参数名 | 含义 | 可选值 | 默认值 |
 | ------------ | ------------ | ------------ | ------------ |
 | APP_MODE         | 应用模式  | cluster/standalone | true  |
 | ACCOUNT_USERNAME | 操作api用户名 | NULL           | nacos |
 | ACCOUNT_PASSWORD | 操作api密码   | NULL           | nacos |

  * nacos-address数据源参数
    * 默认以`application.conf`文件为准
    * 系统参数优先级大于文件配置,默认从配置文件读取

## 模式说明

 1. 单机 (environment或app.model=standalone)默认
    1. 配置redis:   redis存储数据
    2. 不配置redis: 
        1. `cluster.conf`文件存在   :从文件中读取操作(两秒钟间隔实时读取)
        2. `cluster.conf`文件不存在 :Cache存储数据(数据的生命周期到当前进程结束后)
2. 集群(environment或app.model=cluster)必须配置redis
    1. 数据在redis中存储
    2. 无论是否单机或集群模式运行,配置了redis,就相当于一个集群。集群模式只是做了个强制校验必须连接redis

## 使用

### 下载压缩包使用

在`releases`页面选择对应的操作系统压缩包，解压后配置运行即可。

### docker方式使用

1. 在`dockerhub`中搜索nacos-address使用并安装
2. 也可以使用docker-compose方式
       
## nacos服务端和客户端配置部署

### 代理(nginx)

```
server {
    # 监听端口，对应客户端和服务端发起请求所指定的连接点端口 默认值8080，无需更改
    listen                 8080;
    # 自定义域名值 对应客户端和服务端配置的域名
    server_name            nacos-address.aliyun.com;
    location / {
        # nacos address 服务的ip:port,此处也可以使用upstream 代理
        proxy_pass         http://127.0.0.1:8849/;
        proxy_set_header   Host             $host:$server_port;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        proxy_set_header   Access-Control-Allow-Origin  *;
    }
}
```

### Nacos服务端

1. 单机
    1. 启动
        ```
        sh startup.sh -m standalone
        ```
2. 集群
    1. 启动
        ```
        sh startup.sh
        ```
    2. 配置
        1. 当前最新版本提供配置集群列表的方式
            1. 解压目录nacos/的conf目录下，有配置文件cluster.conf更改 ip:port 列表
            2. application.properties配置系统参数获取 nacos.member.list=192.168.16.101:8847?raft_port=8807，192.168.16.101?raft_port=8808，192.168.16.101:8849?raft_port=8809
            3. AddressServer vip模式的寻址方式 (推荐使用)
        2. 针对第3种方式的配置
            1. 从环境变量读取
                1. windows 我得电脑->右键属性->高级系统设置->环境变量->新建环境变量 变量名:address_server_domain 变量值: 自定义域名值(例如:nacos-address.aliyun.com(对应nginx代理的 server_name))
                2. linux 同理
            2. 从Nacos系统参数读取(推荐)
                * 注:(外置/内置)数据源配置方式请参考Nacos（官方网站:http://nacos.io）
                * 服务端代码地址:AddressServerMemberLookup.initAddressSys()
                ```
                 # 需要更改 初始化寻址模式
                 nacos.core.member.lookup.type=address-server
                 # 需要更改(自己设定的域名即可) 对应nginx代理的 server_name 
                 address.server.domain=nacos-address.aliyun.com 
                 # 无需更改 nacos的获取地址请求的端口默认8080
                 # address.server.port=8080
                 # 需要更改 nacos的获取地址请求的链接是 /nacos/serverlist ，但是客户端和服务端的获取列表协议并不兼容，所以我们需要新开一个接口去兼容服务端
                 address.server.url=/nacos/server/serverlist
                ```
         3. 我们正常启动，使用下面的Open api方式就可以操作nacos的服务列表(也可使用操作页面的形式)了，服务端就可以正常发现了(建议在nacos-address服务启动后通过`Open API`操作初始化好服务列表数据)


### Nacos客户端

1. 配置(spring cloud)
   ```
   spring:
      cloud:
        nacos:
          config:
            file-extension: yaml
            prefix: ${spring.application.name}
            # 连接Nacos Server指定的连接点 我们将域名放入对应的 endpoint 上
            endpoint: nacos-address.aliyun.com
          discovery:
            endpoint: nacos-address.aliyun.com
   ```
   
2. 域名指定地址的方式
    1. 本地hosts配置
        ```
        127.0.0.1 nacos-address.aliyun.com
        ```
    2. 局域网内 使用 DNS 解析域名
    3. 阿里云内网的负载SLB(免费，客户端部署到阿里云服务器时就无需指定域名了，唯一的缺点可能某些时刻在大区下配额不足.)
    
3. 类似于其他客户端(SpringBoot，Go，Node.js，Python...等等)接入的方式同理

4. 针对于客户端，无论服务端是单机还是集群，我们都可以使用这种方式，方便以后的扩展，一次配置永久使用(除非域名更改的情况下) 在也不用为注册中心和服务中心扩容需要更新所有的服务而担心了。

## Open API 指南

### 界面操作地址: http://127.0.0.1:8849/index，根据页面提示操作

###  登录获取Token

1. 请求示例
```
curl http://127.0.0.1:8849/login -X POST -H "Content-Type:application/json" -d '{"username":"nacos","password":"nacos"}'
```
2. 返回示例
```
{
    "code": 200,
    "message": "success",
    "data": "token value"
}
```    

### 查询nacos服务地址列表-客户端

1. 请求示例
```
curl http://127.0.0.1:8849/nacos/serverlist
```
2. 返回示例
```
127.0.0.1
127.0.0.2
127.0.0.3

```    
   
### 查询nacos服务地址列表-服务端

1. 请求示例
```
curl http://127.0.0.1:8849/nacos/server/serverlist
```
2. 返回示例
```
{
    "code": 200，
    "message": null，
    "data": "127.0.0.1\n127.0.0.2\n127.0.0.3\n"
}
```
    
### 增加nacos服务端地址

1. 请求示例
```
curl http://127.0.0.1:8849/nacos/serverlist -X POST -H "Content-Type:application/json" -H "Authorization:Bearer token-value" -d '{"clusterIps": ["127.0.0.1","127.0.0.2","127.0.0.3"]}' -v
```
2. 返回示例
```
{
   "code": 200，
   "message": null，
   "data": null
}
```

### 移除nacos服务端地址

1. 请求示例
```
curl http://127.0.0.1:8849/nacos/serverlist -X DELETE -H "Content-Type:application/json" -H "Authorization:Bearer token-value" -d '{"clusterIps": ["127.0.0.1"]}' -v
```

2. 返回示例
```
{
   "code": 200，
   "message": null，
   "data": null
}
```
       
### 清空nacos服务端地址

1. 请求示例
```
curl http://127.0.0.1:8849/nacos/serverlist/all -X DELETE -H "Content-Type:application/json" -H "Authorization:Bearer token-value" -v
```
2. 返回示例
```
{
   "code": 200，
   "message": null，
   "data": null
}
```