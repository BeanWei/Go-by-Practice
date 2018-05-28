package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"

	"github.com/modood/table"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//天气接口API
const (
	APIURL = "http://www.sojson.com/open/api/weather/json.shtml?city="
)

//Weather 定义天气结构
type Weather struct {
	Date    string
	Sunrise string
	High    string
	Low     string
	Sunset  string
	Aqi     string
	FX      string
	FL      string
	Type    string
	Notice  string
}

// Response 接口的响应
type Response struct {
	Date    string `json:"date"`    // 日期
	Message string `json:"message"` // 请求响应消息
	Status  int    `json:"status"`  // 状态码：成功状态为200 ，失败为非200
	City    string `json:"city"`    // 城市
	Count   int32  `json:"count"`   // count
	Data    Data   `json:"data"`    // 数据
}

// Data 数据
type Data struct {
	Shidu     string        `json:"shidu"`     // 湿度
	Pm25      float64       `json:"pm25"`      // pm25
	Pm10      float64       `json:"pm10"`      // pm10
	Quality   string        `json:"quality"`   // 质量
	WenDu     string        `json:"wendu"`     // 温度
	Notice    string        `json:"ganmao"`    // 温馨提示：感冒指数
	Yesterday WeatherInfo   `json:"yesterday"` // 昨天
	Forecast  []WeatherInfo `json:"forecast"`  // 预测
}

//WeatherInfo 构造json数据的结构
type WeatherInfo struct {
	Date    string  `json:"date"`    // 日期
	Sunrise string  `json:"sunrise"` // 日出
	High    string  `json:"high"`    // 最高温
	Low     string  `json:"low"`     // 最低温
	Sunset  string  `json:"sunset"`  // 日落
	Aqi     float32 `json:"aqi"`     // AQI
	Fx      string  `json:"fx"`      // 风向
	Fl      string  `json:"fl"`      // 风力
	Type    string  `json:"type"`    // 天气
	Notice  string  `json:"notice"`  // 温馨提示
}

//WeatherInfo 获取天气信息
func getWeather(r Response) {
	wi := []WeatherInfo{}
	for _, item := range r.Data.Forecast {
		w := WeatherInfo{}
		w.Date = item.Date
		w.Sunrise = item.Sunrise
		w.High = item.High
		w.Low = item.Low
		w.Sunset = item.Sunset
		w.Aqi = item.Aqi
		w.Fx = item.Fx
		w.Fl = item.Fl
		w.Type = item.Type
		w.Notice = item.Notice
		wi = append(wi, w)
	}
	table.Output(wi)
}

//ColorPrint 设置终端字体颜色
func ColorPrint(s string, i int) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

// //使用方法，直接调用即可输出带颜色的文本
// ColorPrint("[OK];", 2|8) //亮绿色

//reqAPI 访问接口
func reqAPI(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Weather-Cli"
	app.Usage = "天气助手"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "city, c",
			Value: "北京",
			Usage: "城市中文名",
		},
	}

	app.Action = func(c *cli.Context) error {
		city := c.String("city")
		tipsInfo := fmt.Sprintf("查询的城市：%s\n", city)
		ColorPrint(tipsInfo, 6|8) //

		var body, err = reqAPI(APIURL + city)
		if err != nil {
			log.Fatal(err)
			return nil
		}

		var r Response
		err = json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		if r.Status != 200 {
			log.Fatalf("错误请求：%s", r.Status)
			return nil
		}
		getWeather(r)
		return nil
	}
	app.Run(os.Args)
}
