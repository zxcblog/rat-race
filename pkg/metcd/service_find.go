package metcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
)

type ClientInfo struct {
	client     *clientv3.Client
	serverList map[string]string
	lock       sync.Mutex
}

func NewClientInfo(client *clientv3.Client) *ClientInfo {
	return &ClientInfo{
		client:     client,
		serverList: map[string]string,
	}
}

func (s *ClientInfo) GetService(prefix string)([]string, error) {
	addrs, err := s.getServiceByName(prefix);
	if err != nil {
		panic(err)
	}

	go s.watcher(prefix)
	return addrs, nil
}

func(s *ClientInfo) watcher(prefix string) {
	rch := s.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT : // 写入
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				s.DelServiceList(string(ev.Kv.Key))
			}	
		}
	}
}

func(s *ClientInfo) getServiceByName(prefix string) ([]string, error) {
	resp, err := s.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	addres := s.extractAddrs(resp)
	return addres, nil
}


func (s *ClientInfo)extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			s.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}


func (s *ClientInfo)SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = val
}

func(s *ClientInfo) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)

	if newRes, err := s.getServiceByName(key); err != nil {
		log.Panic(err)
	} else {
		log.Panic("新的服务列表：%s", newRes)
	}
}

func (s *ClientInfo) SerList2Array() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, v)

	for _, v := range s.serverList {
		addrs =append(addrs, v)
	}
	return addrs
}

1. 服务发现和注册
2. 没用框架
3. protobuf 协议，dns 服务发现， nginx 负载均衡
4. 文件分片传输

5. 通过前缀进行认证
6. 参数校验 go-validatior
7.
8.

9. 分布式
9.1.







