#### 1 部署 ZTM HUB 服务

```bash
ztm run hub -l 0.0.0.0:8888 -n 127.0.0.1:8888 -d /tmp/.ztm.hub --permit /tmp/root.json
```

#### 2 部署 ZTM Home Agent 服务

```bash
ztm run agent -l 0.0.0.0:7771 -d /tmp/.ztm.home --permit /tmp/root.json --join k8s-mesh --join-as home
```

#### 3 部署 ZTM Office Agent 服务

```bash
ztm run agent -l 0.0.0.0:7772 -d /tmp/.ztm.office --permit /tmp/root.json --join k8s-mesh --join-as office
```

#### 4 测试

##### 4.1 组网

```
#CTR_AGENT=home make TestJoin
CTR_AGENT=home make TestListEndpoints

#CTR_AGENT=office make TestJoin
CTR_AGENT=office make TestListEndpoints
```

##### 4.2 ztm/tunnel 测试

```bash
CTR_AGENT=home   make TestStartApp
CTR_AGENT=office make TestStartApp

CTR_AGENT=home make TestOpenOutbound

export MY_HOST_IP=192.168.127.91

CTR_AGENT=office MY_HOST_IP=${MY_HOST_IP} make TestOpenInbound

curl ${MY_HOST_IP}:10081 -I
curl ${MY_HOST_IP}:10082 -I
```

##### 4.3 file 测试

```bash
CTR_AGENT=home   make TestPublishFile

CTR_AGENT=home   make TestDescribeFile
CTR_AGENT=office make TestDescribeFile

CTR_AGENT=home   make TestDownloadFile
CTR_AGENT=office make TestDownloadFile

CTR_AGENT=home   make TestEraseFile
CTR_AGENT=office make TestEraseFile

CTR_AGENT=home   make TestListFiles
CTR_AGENT=office make TestListFiles
```

