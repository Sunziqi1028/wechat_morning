/**
 * @Author: Sun
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/8/23 21:57
 */

package model

type Config struct {
	Wechat      Wechat
	Information Information
}

type Wechat struct {
	AppID      string
	AppSecret  string
	TemplateID string
	UserIDs    string
}

type Information struct {
	WeatherKey    string
	Location      string
	Birthday      string
	LoveDate      string
	Constellation string
	TianApiKey    string
	LocationCN    string
}
