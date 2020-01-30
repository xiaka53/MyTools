package app

import (
	"MyTools/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Upgrade struct {
	Version string
	latest  map[string]interface{}
}

func showUpgrade() {
	upgrade := Upgrade{Version: Version}
	upgrade.exec()
}

func (upgrade *Upgrade) exec() {
	var waitGroutp = sync.WaitGroup{}
	islock := true
	go func() {
		utils.Outputs("正在检测最新版本")
		for {
			if !islock {
				utils.OutputLn("")
				waitGroutp.Done()
				break
			}
			utils.Outputs(".")
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		upgrade.getLatesVersion()
		time.Sleep(time.Second * 3)
		islock = false
		waitGroutp.Done()
	}()

	waitGroutp.Add(2)
	waitGroutp.Wait()

	utils.OutputLn("当前版本：" + upgrade.Version)
	latesVersion := upgrade.latest["version"].(string)
	ret := upgrade.compreVersion(latesVersion, upgrade.Version)
	if ret <= 0 {
		utils.OutputLn("感谢您的支持，当前已是最新版本。")
		return
	}

	utils.OutputLn("检测到新版本：" + latesVersion)
	url := upgrade.downloadUrl()
	if url == "" {
		utils.OutputLn("暂不支持" + runtime.GOOS + "系统自动更新，请邮件联系陶然。")
		return
	}
	filename := path.Base(url)
	savePath := os.TempDir() + filename
	err := upgrade.downloadFile(url, savePath, func(length, downLen int64) {
		process := float64(downLen) / float64(length) * 100
		utils.Outputs("\rdownloading " + fmt.Sprintf("%.2f", process) + "%")
	})
	if err != nil {
		utils.OutputLn("下载失败：" + err.Error())
		return
	}
	utils.OutputLn("downloading 100%")
}

//获取最新版本数据
func (upgrade *Upgrade) getLatesVersion() {
	resp, err := http.Get("http://admin.cocofan.cn/api/Demo/test1")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var releaseInfo map[string]interface{}
	err = json.Unmarshal(body, &releaseInfo)
	if err != nil {
		panic(err)
	}
	upgrade.latest = releaseInfo["data"].(map[string]interface{})
}

//版本比较
// lim=1 src版本最新   lim=0 版本相同 lim=-1 other版本新
func (Upgrade) compreVersion(src string, other string) (lim int) {
	src = strings.Trim(src, "v")
	other = strings.Trim(other, "v")
	v1 := strings.Split(src, ".")
	v2 := strings.Split(other, ".")

	if len(v1) > len(v2) {
		lim = len(v1)
	} else {
		lim = len(v2)
	}

	for {
		if len(v1) >= lim {
			break
		}
		v1 = append(v1, "0")
	}
	for {
		if len(v2) >= lim {
			break
		}
		v2 = append(v2, "0")
	}
	lim = 0
	for i := 0; i < lim; i++ {
		num1, _ := strconv.Atoi(v1[i])
		num2, _ := strconv.Atoi(v1[i])

		if num1 > num2 {
			lim = 1
			break
		}
		if num1 > num2 {
			lim = -1
			break
		}
	}
	return
}

//获取下载URL
func (upgrade *Upgrade) downloadUrl() (url string) {
	sysOs := runtime.GOOS
	if sysOs == "darwin" {
		sysOs = "macOs"
	}
	filename := sysOs + "-" + runtime.GOARCH
	for _, val := range upgrade.latest["assets"].([]interface{}) {
		asset := val.(map[string]interface{})
		if strings.Index(asset["name"].(string), filename) != -1 {
			url = asset["download_url"].(string)
			break
		}
	}
	return
}

//下载文件
func (Upgrade) downloadFile(url, downloadPath string, fb func(length, downLen int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64) //读取服务器，返回文件大小
	if err != nil {
		return err
	}
	file, err := os.Create(downloadPath) //新建文件
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	for {
		nr, er := resp.Body.Read(buf) //读取获得bytes
		if nr > 0 {
			nw, ew := file.Write(buf[0:nr]) //写入bytes
			if nw > 0 {                     //数据长度大于0
				written += int64(nw)
			}
			if ew != nil { //写入错误
				err = ew
				break
			}
			if nr != nw { //读取长度不等于写入长度
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
		fb(fsize, written) //没有错误 使用callback
	}
	return err
}
