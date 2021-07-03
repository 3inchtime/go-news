# go-news

## Consul

```bash
docker pull consul

docker run -d -p 8500:8500 -v /data/consul:/consul/data -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_1 consul:latest agent -server -bootstrap -ui -node=1 -client='0.0.0.0'
```

- `agent` : 表示启动 Agent 进程。
- `-server`：表示启动 Consul Server 模式。
- `-client`：表示启动 Consul Cilent 模式。
- `-bootstrap`：表示这个节点是 `Server-Leader` ，每个数据中心只能运行一台服务器。技术角度上讲 Leader 是通过 Raft 算法选举的，但是集群第一次启动时需要一个引导 Leader，在引导群集后，建议不要使用此标志。
- `-ui`：表示启动 Web UI 管理器，默认开放端口 `8500`，所以上面使用 Docker 命令把 8500 端口对外开放。
- `-node`：节点的名称，集群中必须是唯一的。
- `-client`：表示 Consul 将绑定客户端接口的地址，`0.0.0.0` 表示所有地址都可以访问。
- `-join`：表示加入到某一个集群中去。 如：`-json=192.168.1.23`