package base

import (
	"context"
	"github.com/mszhangyi/work/udpLog"
	"github.com/mszhangyi/work/udpLog/utils"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	props   *systemConf
	EtcdKey string
)

type PropsStarter struct {
	udpLog.BaseStarter
}

type systemConf struct {
	Addr string `json:"addr"`
}

func (p *PropsStarter) Init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"81.68.243.67:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("connect failed, err:", err)
		return
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		log.Println("get failed, err:", err)
		return
	}
	props = &systemConf{}
	utils.ByteJsonByData(resp.Kvs[0].Value, props)
}
