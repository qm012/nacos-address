# VIP address service center of Nacos single machine / cluster

See the [中文文档](https://github.com/qm012/nacos-address/blob/master/README-zh.md) for Chinese readme.

## Introduction

&emsp;&emsp;&emsp;Based on Nacos（official site:http://nacos.io ）The additional web server is designed for addressing the addresses of server and client, 
reducing changes (server and client projects) and facilitating dynamic expansion and management。
It is suitable for single machine or cluster management of self built Nacos，
[阿里云的MSE微服务引擎托管](https://cn.aliyun.com/product/aliware/mse) It doesn't need to be considered. The government has dealt with it。<br/><br/>
&emsp;&emsp;&emsp;When the company needed to use the configuration center last year, through research and selection, 
Nacos was finally used as the configuration center and registration center。In the process of using, 
we Some problems were also found in the client project (spring cloud) and other clients (springboot, go, Node.js , python... Etc.)
When we configure the server address `spring.cloud.nacos.config.serverAddr=127.0.0.1:8848,127.0.0.2:8848,127.0.0.3:8848`，
If we have 100 + clients and the address changes, do we need to change all 100 clients? This is not flexible。
The official documents and blogs don't explain it in detail. Basically, it's only to configure 'serveraddr'. In the server cluster mode, 
new server machines are offline or online，We need to be in` cluster.conf `Address changes, there will be similar scenarios,
 so go to the source code and research, resulting in this project。<br/><br/>

&emsp;&emsp;&emsp;It adapts to the client and server, providing usage and deployment documents, 
unified API management of address lists with different storage methods, and support for docker. At the same time, 
we are also welcome to use, suggest and contribute. If it is helpful or helpful in the future, welcome to star。

## Storage support

-[x] redis (single machine + cluster)

-[x] cache (standalone)

-[x] File (standalone)

## Function

-[x] docker support

## Parameters

  * nacos-address system parameters(environment)
     
 | Parameter name | meaning | Optional value | Default value |
 | ------------ | ------------ | ------------ | ------------ |
 | APP_MODE         | application mode         | cluster/standalone | true |
 | ACCOUNT_USERNAME | Operation API username   | NULL               | nacos |
 | ACCOUNT_PASSWORD | Operation API password   | NULL               | nacos |
 | REDIS_HOST       | redis address            | NULL               | ""    |
 | REDIS_PASSWORD   | redis password           | NULL               | ""    |
 | REDIS_DB         | redis select             | NULL               | 0     |
 
  * nacos-address Data source parameters
  
    * By default, the 'application. Conf' file will prevail

    * The priority of system parameters is higher than that of file configuration,
      which is read from configuration file by default
      
 ### Mode explain 
 
 1. standalone (environment or (yaml)app.model=standalone) default
    1. set redis:   redis store data
    2. no set redis: 
        1. `cluster.conf`file exists     :Read from file (real time read at two second intervals)
        2. `cluster.conf`file not exists :Cache stores data (after the lifecycle of data ends in the current process)
2. cluster(environment or (yaml)app.model=cluster)must set redis
    1. The data is stored in redis
    2. No matter whether it runs in stand-alone or cluster mode, 
    redis is equivalent to a cluster. Cluster mode only makes a mandatory verification,
    and must be connected to redis

 ## Use
 
 ### Download compressed package
 
In the 'releases' page, select the corresponding operating system compressed package, unzip and configure it to run.
 
 ### docker
 
1. Search for and install [nacos-address](https://registry.hub.docker.com/repository/docker/qm012/nacos-address) in dockerhub

2. [Docker-compose](https://github.com/qm012/nacos-address/blob/v1.1.3/docker-compose.yml) mode can also be used

## Configuration and deployment of Nacos server and client

### Using agents(nginx)

```
server {
    # Listening port, which corresponds to the connection point port specified by the client and server to initiate the request. 
        The default value is 8080, which does not need to be changed
    listen                 8080;
    # The custom domain name value corresponds to the domain name configured by the client and server
    server_name            nacos-address.aliyun.com;
    location / {
        # nacos vip server ip:port,You can also use the upstream proxy here
        proxy_pass         http://127.0.0.1:8849/;
        proxy_set_header   Host             $host:$server_port;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        proxy_set_header   Access-Control-Allow-Origin  *;
    }
}
```

### Nacos server

1. Stand alone
    1. start-up
        ```
        sh startup.sh -m standalone
        ```
2. cluster
    1. start-up
        ```
        sh startup.sh
        ```
    2. configure
        1. The latest version provides a way to configure the cluster list
            1. change cluster.conf
            2. application.properties nacos.member.list=192.168.16.101:8847?raft_port=8807，192.168.16.101?raft_port=8808，192.168.16.101:8849?raft_port=8809
            3. AddressServer vip mode (Recommended)
        2. Configuration for mode 3
            1. Read from environment variable
                1. windows Computer > right click Properties > Advanced System Settings > environment variables > 
                new environment variable name: Address_ server_ Domain variable value: 
                custom domain name value (for example: Nacos- vip.aliyun.com (corresponding to the server of nginx proxy_ name))
                2. linux In the same way
            2. Read from Nacos system parameters (recommended)
                * Note: (external / built-in) please refer to Nacos（官方网站:http://nacos.io）
                * Server code address:AddressServerMemberLookup.initAddressSys()
                ```
                 # Initialization addressing mode needs to be changed
                 nacos.core.member.lookup.type=address-server
                 # The server corresponding to nginx proxy needs to be changed (the domain name can be set by yourself)_ name
                 address.server.domain=nacos-address.aliyun.com 
                 # There is no need to change the port of Nacos' get address request. The default port is 8080
                 # address.server.port=8080
                 # To change the access address of Nacos, the link of the request is /nacos/serverlist. However, 
                   the protocol of getting list between the client and the server is not compatible,
                   so we need to open a new interface to be compatible with the server
                 address.server.url=/nacos/server/serverlist
                ```
         3. If we start normally, we can operate the service list of Nacos by using the following open API method, 
         and the server can discover it normally (it is recommended to initialize the service list data through the 'open API' operation after the Nacos VIP service is started)


### Nacos client

1. to configure(spring cloud)
   ```
   spring:
      cloud:
        nacos:
          config:
            file-extension: yaml
            prefix: ${spring.application.name}
            # Connect to the connection point specified by Nacos server. 
              We put the domain name on the corresponding endpoint.
            endpoint: nacos-address.aliyun.com
          discovery:
            endpoint: nacos-address.aliyun.com
   ```
   
2. How domain names specify addresses
    1. Local hosts configuration
        ```
        127.0.0.1 nacos-address.aliyun.com
        ```
    2. Using DNS to resolve domain name in LAN
    3. Load SLB of alicloud intranet 
      (free. When clients deploy to alicloud servers, there is no need to specify a domain name. 
      The only drawback is that the quota may be insufficient in the region at some time.)
    
3. Similar to other clients(SpringBoot，Go，Node.js，Python...Etc.)The way of access is the same

4. For the client, whether the server is a single machine or a cluster, we can use this method to facilitate future expansion.
   Once configured for permanent use (unless the domain name is changed), 
   there is no need to worry about the need to update all services for the expansion of the registration center and service center.

## Open API Guide

### Nacos-address console link: http://127.0.0.1:8849/index  Follow the prompts on the page

###  Login get Token
1. Request example
```shell
curl http://127.0.0.1:8849/login \
 -X POST \
 -H "Content-Type:application/json" \
 -d '{"username":"nacos","password":"nacos"}'
```
2. Return to example
```
{
    "code": 200,
    "message": "success",
    "data": "token value"
}
```

### Query Nacos service address list - client

1. Request example
```shell
curl http://127.0.0.1:8849/nacos/serverlist
```
2. Return to example
```
127.0.0.1
127.0.0.2
127.0.0.3

```    
   
### Query Nacos service address list - server

1. Request example
```shell
curl http://127.0.0.1:8849/nacos/server/serverlist
```
2. Return to example
```
{
    "code": 200，
    "message": null，
    "data": "127.0.0.1\n127.0.0.2\n127.0.0.3\n"
}
```
    
### Add Nacos server address

1. Request example
```shell
curl http://127.0.0.1:8849/nacos/serverlist \
 -X POST \
 -H "Content-Type:application/json" \
 -H "Authorization:Bearer token-value" \
 -d '{"clusterIps": ["127.0.0.1","127.0.0.2","127.0.0.3"]}' -v
```
2. Return to example
```
{
   "code": 200，
   "message": null，
   "data": null
}
```

### Remove Nacos server address

1. Request example
```shell
curl http://127.0.0.1:8849/nacos/serverlist \
 -X DELETE \
 -H "Content-Type:application/json" \
 -H "Authorization:Bearer token-value" \
 -d '{"clusterIps": ["127.0.0.1"]}' -v
```

2. Return to example
```
{
   "code": 200，
   "message": null，
   "data": null
}
```
       
### Clear Nacos server address

1. Request example
```shell
curl http://127.0.0.1:8849/nacos/serverlist/all \
 -X DELETE \
 -H "Content-Type:application/json" \
 -H "Authorization:Bearer token-value" -v
```
2. Return to example
```
{
   "code": 200，
   "message": null，
   "data": null
}
```
