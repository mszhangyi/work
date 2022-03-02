package base

import (
	"bytes"
	"fmt"
	"github.com/mszhangyi/work/udpLog"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

type HttpStarter struct {
	udpLog.BaseStarter
}

func (s *HttpStarter) Init() {
	httpClient = creteHttpClient()
}

func HttpPost(method, url string, str string) {
	//fmt.Println(str)
	lIndex := strings.Index(str, "level")
	mIndex := strings.Index(str, "msg")
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(`{"@t":"`+str[6:lIndex-2]+`","@l":"`+str[lIndex+6:mIndex-1]+`","@m":"`+str[mIndex+5:len(str)-2]+`"}`)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if resp != nil {
		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body), "--------")
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
	}
}

func creteHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1000,
			DisableKeepAlives:   true,
			DisableCompression:  false,
		},
		Timeout: 30 * time.Second,
	}
	return client
}
