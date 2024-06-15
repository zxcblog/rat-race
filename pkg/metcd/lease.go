package metcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 当服务启动时，将服务的信息注册到etcd中，同时绑定租约（lease）
// 并以续租约（keep leases alive）的方式检测服务是否正常运行
// 从而实现服务的健康检查

type Lease struct {
	Id          clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	mEtcd       *MEtcd
	ttl         int64

	key string
	val string
}

// 新建租约
func newLease(metcd *MEtcd, key, val string, ttl int64) (*Lease, error) {
	lease := &Lease{
		mEtcd: metcd,
		key:   key,
		val:   val,
		ttl:   ttl,
	}
	if err := lease.register(); err != nil {
		return nil, err
	}

	return lease, nil
}

func (l *Lease) register() error {
	ctx := context.Background()

	rsp, err := l.mEtcd.cli.Grant(ctx, l.ttl)
	if err != nil {
		return err
	}

	_, err = l.mEtcd.cli.Put(ctx, l.key, l.val, clientv3.WithLease(rsp.ID))
	if err != nil {
		return err
	}

	keepAlive, err := l.mEtcd.cli.KeepAlive(ctx, rsp.ID)
	if err != nil {
		return err
	}

	l.Id = rsp.ID
	l.keepAliveCh = keepAlive
	return nil
}

func (l Lease) keepAlive() {
	for res := range l.keepAliveCh {
		l.mEtcd.log.DebugF("%s租约续租成功,ID: %d, ttl:%d", l.key, res.ID, res.TTL)
	}
	l.mEtcd.log.DebugF("%s租约关闭", l.key)
}

// close 关闭租约
func (l Lease) close() error {
	_, err := l.mEtcd.cli.Revoke(context.Background(), l.Id)
	return err
}
