package metcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// ServiceReg 创建租约注册服务
type ServiceReg struct {
	client        *clientv3.Client
	lease         clientv3.Lease
	leaseResp     *clientv3.LeaseGrantResponse
	cancelFunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
}

func NewServiceReg(client *clientv3.Client, timeNum int64) (*ServiceReg, error) {
	ser := &ServiceReg{}

	if err := ser.setLease(timeNum); err != nil {
		return nil, err
	}
	go ser.ListenLeaseRespChan()

	return ser, nil
}

func (s *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(s.client)

	//	设置租约时间
	leaseResp, err := lease.Grant(context.Background(), timeNum)
	if err != nil {
		return err
	}

	// 设置续租
	ctx, cancelFunc := context.WithCancel(context.Background())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		cancelFunc()
		return err
	}

	s.lease = lease
	s.leaseResp = leaseResp
	s.cancelFunc = cancelFunc
	s.keepAliveChan = leaseRespChan
	return nil
}

func (s *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-s.keepAliveChan:
			if leaseKeepResp == nil {
				//	log.Println("关闭续租功能")
			} else {
				//log.Println("续租成功")
			}
		}
	}
}

// PutService 通过租约注册服务
func (s *ServiceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(s.client)
	_, err := kv.Put(context.Background(), key, val, clientv3.WithLease(s.leaseResp.ID))
	return err
}

func (s *ServiceReg) RevokeLease() error {
	s.cancelFunc()
	time.Sleep(2 * time.Second)
	_, err := s.lease.Revoke(context.Background(), s.leaseResp.ID)
	return err
}
