# nacos address lookup server

## Tips

### It is not recommended mounting the disk in volume mode. Please mount the disk in named volume mode

## Quick Start
```shell
docker run -d --name nacos-address \
 --restart=always \
 -v nacos-address-config:/home/nacos-address/config \
 -v nacos-address-logs:/home/nacos-address/logs \
 -p 8849:8849 \
 -e APP_MODE=standalone \
 qm012/nacos-address
```

## Quick Redis Start
```shell
docker run -d --name nacos-address \
 --restart=always \
 -v nacos-address-config:/home/nacos-address/config \
 -v nacos-address-logs:/home/nacos-address/logs \
 -p 8849:8849 \
 -e APP_MODE=standalone \
 -e REDIS_HOST="redis:6379" \
 -e REDIS_PASSWORD="" \
 -e REDIS_DB=0 \
 --link redis \
 qm012/nacos-address
```
## Quick Compose Start

```shell
version: '3'

services:
  redis:
    image: redis:latest
    restart: always
    networks:
      - nacos-address
  nacos-address:
    image: qm012/nacos-address:latest
    restart: always
    ports:
      - 8849:8849
    depends_on:
      - redis
    links:
      - redis
    networks:
      - nacos-address
    volumes:
      - nacos-address-config:/home/nacos-address/config
      - nacos-address-logs:/home/nacos-address/logs
    environment:
      APP_MODE: standalone
      ACCOUNT_USERNAME: nacos
      ACCOUNT_PASSWORD: nacos
      REDIS_HOST: "redis:6379"
      REDIS_PASSWORD: ""
      REDIS_DB: 0

volumes:
  nacos-address-config:
  nacos-address-logs:

networks:
  nacos-address:
```

`
docker-compose up
`

## Common property configuration

| Parameter name | meaning | Optional value | Default value |
 | ------------ | ------------ | ------------ | ------------ |
 | APP_MODE         | application mode         | cluster/standalone | true |
 | ACCOUNT_USERNAME | Operation API username   | NULL               | nacos |
 | ACCOUNT_PASSWORD | Operation API password   | NULL               | nacos |
 | REDIS_HOST       | redis address            | NULL               | ""    |
 | REDIS_PASSWORD   | redis password           | NULL               | ""    |
 | REDIS_DB         | redis select             | NULL               | 0     |
 
### Refer to application.yaml for other startup parameters

### Open the Nacos-address console in your browser

* link：http://127.0.0.1:8849/index/