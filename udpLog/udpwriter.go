// Copyright 2012 SocialCode. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package udpLog

import (
	"fmt"
	"net"
	"os"
)

type UDPWriter struct {
	conn   net.Conn	`json:"-"`
	Hostname string `json:"hostname"`
	Facility string  `json:"facility"`
	Msg string	`json:"msg"`
}

// What compression type the writer should use when sending messages
// to the graylog2 server
type CompressType int


// New returns a new GELF Writer.  This writer can be used to send the
// output of the standard Go log functions to a central GELF server by
// passing it to log.SetOutput()
func NewUDPWriter(addr ,tag string) (*UDPWriter, error) {
	var err error
	w := new(UDPWriter)
	if w.conn, err = net.Dial("udp", addr); err != nil {
		return nil, err
	}
	if w.Hostname, err = os.Hostname(); err != nil {
		return nil, err
	}
	w.Facility = tag
	return w, nil
}


func (w *UDPWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	//第二步  获取数据长度 写入一个定长4个字节的uint32 中
	dataLen := make([]byte, 4)
	PutUint32(dataLen, uint32(len(p)))
	w.conn.Write(dataLen)
	//第四步  写数据和Method
	n, err = w.conn.Write(p)
	if err != nil {
		return
	}
	return len(p), nil
}

func (w *UDPWriter) Close() error {
	if w.conn == nil {
		return nil
	}
	return w.conn.Close()
}

//写入固定格式的uint32
func PutUint32(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

//解出固定格式uint32
func Uint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
