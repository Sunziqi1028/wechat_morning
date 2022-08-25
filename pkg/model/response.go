/**
 * @Author: Sun
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2022/8/23 22:08
 */

package model

type ResponseAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expire_in"`
}

type LoveResponse struct {
	Code     int           `json:"code"`
	Msg      string        `json:"msg"`
	Newslist []LoveNewList `json:"newslist"`
}

type LoveNewList struct {
	Content string `json:"content"`
}

type NewsListFighting struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	Fighting []Fighting `json:"newslist"`
}
type Fighting struct {
	Saying string
	Transl string
	Source string
}

type Constellation struct {
	Code      int          `json:"code"`
	Msg       string       `json:"msg"`
	NewsLists []ListCommon `json:"newslist"`
}

type ListCommon struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
