package base

import (
	"github.com/mszhangyi/work/udpLog"
	"github.com/mszhangyi/work/udpLog/utils"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
)

var (
	udpConn *net.UDPConn
)

type UdpStarter struct {
	udpLog.BaseStarter
}

func (t *UdpStarter) Init() {
	var err error
	udpConn, err = net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 12201})
	if err != nil {
		panic("resolve tcp addr err: " + err.Error())
		return
	}
}

func (t *UdpStarter) Start() {
	go func() {
		for {
			//解包  获取data大小
			dataLen := make([]byte, 4)
			_, err := udpConn.Read(dataLen)
			if err != nil {
				logrus.Error("读取数据失败!", err)
				continue
			}
			//第四步  读出data数据
			data := make([]byte, utils.Uint32(dataLen))
			_, err = udpConn.Read(data)
			if err != nil {
				logrus.Error("读取数据失败!", err)
				continue
			}
			go HttpPost("POST", "http://"+props.Addr+"/api/events/raw?clef", string(data))
		}
	}()
	logrus.Error("启动服务端成功,端口：", 12201)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("关闭服务 ")
}

func (t *UdpStarter) Stop() {
	logrus.Info("程序退出")
	udpConn.Close()
}
