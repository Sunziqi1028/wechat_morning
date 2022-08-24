/**
 * @Author: Sun
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2022/8/23 21:55
 */

package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"wechat_message/pkg/global"
	"wechat_message/pkg/model"
	"wechat_message/pkg/utils"
)

func GetAccessToken(appID, appSercet string) string {
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, appSercet)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get access token err:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var respBody model.ResponseAccessToken
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		fmt.Println("unmarshaler json err:", err)
		return ""
	}
	//global.AccessToken = respBody.AccessToken
	return respBody.AccessToken
}

func GetWeather(weatherKey, localtion string) (weather, temp, feelsLike, windDir string) {
	url := fmt.Sprintf("https://devapi.qweather.com/v7/weather/now?location=%s&key=%s", localtion, weatherKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get access token err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	weather = gojsonq.New().FromString(string(body)).Find("now.text").(string)        // 天气
	temp = gojsonq.New().FromString(string(body)).Find("now.temp").(string)           // 温度
	feelsLike = gojsonq.New().FromString(string(body)).Find("now.feelsLike").(string) // 体感温度
	windDir = gojsonq.New().FromString(string(body)).Find("now.windDir").(string)     // 风向
	//windSpeed = gojsonq.New().FromString(string(body)).Find("now.windSpeed").(string) // 风速
	//windScale = gojsonq.New().FromString(string(body)).Find("now.windScale").(string) // 风力等级

	return weather, temp, feelsLike, windDir
}

func GetConstellation(key, constellation string) {
	// http://api.tianapi.com/caihongpi/index
	url := fmt.Sprintf("http://api.tianapi.com/caihongpi/index?key=%s", key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	feelsLike := gojsonq.New().FromString(string(body)).Find("newslist.content") // 体感温度
	fmt.Println(string(body))
	fmt.Println(feelsLike)
}

func GetLove(key string) string {
	url := fmt.Sprintf("http://api.tianapi.com/saylove/index?key=%s", key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var love model.LoveResponse
	err = json.Unmarshal(body, &love)
	var loveContent []string
	for _, v := range love.Newslist {
		loveContent = append(loveContent, v.Content)
	}
	return loveContent[0]
}

func SendMessage(accessToken, loveDate, birthday, weather, temp, feelsLike, windDir, locationCN, loveContent string) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	tmpday := utils.GetLovedDay(loveDate, time.Now().Format("2006-01-02"))
	day := strconv.Itoa(int(tmpday))
	birth := utils.CalBirthDay(birthday)

	request := &model.Request{
		ToUser:     global.Config.Wechat.UserIDs,
		TopColor:   utils.RandomString(),
		TemplateId: global.Config.Wechat.TemplateID,
		Data: model.Data{
			Date: model.Date{
				Value: time.Now().Format("2006年1月02号 15:04:05"),
				Color: utils.RandomString(),
			},
			Region: model.Region{
				Value: locationCN,
				Color: utils.RandomString(),
			},
			Weather: model.Weather{
				Value: weather,
				Color: utils.RandomString(),
			},
			Temp: model.Temp{
				Value: temp,
				Color: utils.RandomString(),
			},
			WindDir: model.WindDir{
				Value: windDir,
				Color: utils.RandomString(),
			},
			LoveDay: model.LoveDay{
				Value: day,
				Color: utils.RandomString(),
			},
			Birthday: model.Birthday{
				Value: birth,
				Color: utils.RandomString(),
			},
			Morning: model.Morning{
				Value: loveContent,
				Color: utils.RandomString(),
			},
			FeelsLike: model.FeelsLike{
				Value: feelsLike,
				Color: utils.RandomString(),
			},
		},
	}
	data, _ := json.Marshal(request)
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return
}
