#### 1 部署 ZTM CA 服务

```bash
ztm run ca --listen 0.0.0.0:9999 --database /opt/ztm.db
#CA
curl 10.101.1.5:9999/api/certificates/ca
#PRIVATE KEY
curl -X POST 10.101.1.5:9999/api/certificates/root
#CERTIFICATE
curl 10.101.1.5:9999/api/certificates/root
#DELETE
curl -X DELETE 10.101.1.5:9999/api/certificates/root
```

#### 2 部署 ZTM HUB 服务

```bash
ztm run hub --listen 0.0.0.0:8888 --ca 192.168.226.62:9999
```

#### 3 部署 ZTM Home Agent 服务

```bash
ztm run agent --listen 0.0.0.0:7777 --database /opt/ztm.db
```

#### 4 部署 ZTM Office Agent 服务

```bash
ztm run agent --listen 0.0.0.0:7777 --database /opt/ztm.db
```
