# nacos address lookup server

## Quick Start
`
docker run -d --name nacos-address -v nacos-address-config:/home/nacos-address/config -v nacos-address-logs:/home/nacos-address/logs  -p 8849:8849 -e APP_MODE=standalone qm012/nacos-address
`

## Quick Redis Start
`
docker run -d --name nacos-address -v nacos-address-config:/home/nacos-address/config -v nacos-address-logs:/home/nacos-address/logs  -p 8849:8849 -e APP_MODE=standalone --link redis qm012/nacos-address
`
## Quick Compose Start

```shell
version: '3'

services:
  redis:
    image: redis:latest
    networks:
      - nacos-address
  nacos-address:
    image: qm012/nacos-address:latest
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

volumes:
  nacos-address-config:
  nacos-address-logs:

networks:
  nacos-address:
    driver: bridge
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

### Refer to application.yaml for other startup parameters

### Open the Nacos-address console in your browser

* link：http://127.0.0.1:8849/index/