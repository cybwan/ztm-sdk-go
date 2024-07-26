#### 1 部署 ZTM HUB 服务

```bash
ztm run hub -l 0.0.0.0:8888 -n 127.0.0.1:8888 -d /tmp/.ztm.hub --permit /tmp/root.json
```

#### 2 部署 ZTM Home Agent 服务

```bash
ztm run agent -l 0.0.0.0:7771 -d /tmp/.ztm.home
```

#### 3 部署 ZTM Office Agent 服务

```bash
ztm run agent -l 0.0.0.0:7772 -d /tmp/.ztm.office
```

#### 4 测试

```bash
CTR_AGENT=home make TestJoin
CTR_AGENT=home make TestListEndpoints

CTR_AGENT=office make TestJoin
CTR_AGENT=office make TestListEndpoints

CTR_AGENT=home make TestStartApp
CTR_AGENT=office make TestStartApp

CTR_AGENT=home make TestOpenOutbound

export MY_HOST_IP=192.168.127.91

CTR_AGENT=office MY_HOST_IP=${MY_HOST_IP} make TestOpenInbound

curl ${MY_HOST_IP}:10081 -I
curl ${MY_HOST_IP}:10082 -I
```

