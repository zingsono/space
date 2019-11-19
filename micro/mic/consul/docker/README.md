# consul

官网： https://www.consul.io        
Github： https://github.com/hashicorp/consul.git      
Docker： https://hub.docker.com/_/consul   

# Consul docker container

下载镜像 `docker pull consul:1.6`

consul 镜像关键信息：
```
VOLUME /consul/data
EXPOSE 8300
EXPOSE 8301 8301/udp 8302 8302/udp
EXPOSE 8500 8600 8600/udp
CMD ["agent", "-dev", "-client", "0.0.0.0"]

```

运行容器：
```
docker volume create consul1
docker rm -f consul1
docker run --name consul1 -p 8500:8500 -v consul1:/consul/data -d --restart=always consul:1.6 agent -node=consul1 -server -ui -bind=127.0.0.1 -client=0.0.0.0 -bootstrap
```


# Consul Cli 


# Consul API 






















