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
	gocron.Every(1).Day().At("08:30").Do(task)

	<-gocron.Start()
}

func task() {
	accessToken := http.GetAccessToken(global.Config.Wechat.AppID, global.Config.Wechat.AppSecret)
	weather, temp, feelsLike, windDir := http.GetWeather(global.Config.Information.WeatherKey, global.Config.Information.Location)
	loveContent := http.GetLove(global.Config.Information.TianApiKey)
	http.SendMessage(accessToken, global.Config.Information.LoveDate, global.Config.Information.Birthday, weather, temp, feelsLike, windDir, global.Config.Information.LocationCN, loveContent)
}
