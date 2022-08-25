/**
 * @Author: Sun
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/8/23 21:50
 */

package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"wechat_message/pkg/config"
	"wechat_message/pkg/global"
	"wechat_message/pkg/http"
)

func init() {
	var err error
	global.Config, err = config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	// gocron.Every(30).Second().Do(task)
	gocron.Every(1).Day().At("08:45").Do(task)

	<-gocron.Start()
}

func task() {
	accessToken := http.GetAccessToken(global.Config.Wechat.AppID, global.Config.Wechat.AppSecret) // 获取微信token

	weather, temp, feelsLike, windDir := http.GetWeather(global.Config.Information.WeatherKey, global.Config.Information.Location) // 获取天气

	fighting := http.GetFighting(global.Config.Information.TianApiKey) // 获取每日一句

	resp1, resp2, resp3 := http.GetConstellation(global.Config.Information.TianApiKey, global.Config.Information.Constellation) // 获取星座运势

	http.SendMessage(accessToken, global.Config.Information.LoveDate, global.Config.Information.Birthday,
		weather, temp, feelsLike, windDir, global.Config.Information.LocationCN, fighting.Saying, fighting.Transl, fighting.Source,
		resp1.Content, resp2.Content, resp3.Content)
}
