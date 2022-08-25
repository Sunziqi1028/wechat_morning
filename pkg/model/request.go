/**
 * @Author: Sun
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2022/8/24 21:31
 */

package model

type Request struct {
	ToUser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	TopColor   string `json:"topcolor"`
	Data       Data   `json:"data"`
}

type Data struct {
	Date     Date     `json:"date"`
	Region   Region   `json:"region"`
	Weather  Weather  `json:"weather"`
	Temp     Temp     `json:"temp"`
	WindDir  WindDir  `json:"wind_dir"`
	LoveDay  LoveDay  `json:"love_day"`
	Birthday Birthday `json:"birthday"`
	//Morning   Morning   `json:"morning"`
	FeelsLike FeelsLike `json:"feelslike"`
	Saying    Saying    `json:"saying"`
	Transl    Transl    `json:"transl"`
	Source    Source    `json:"source"`
	Color     Color     `json:"color"`
	Num       Num       `json:"num"`
	Overview  Overview  `json:"overview"`
}

type Date struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Region struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Weather struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Temp struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type WindDir struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type LoveDay struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Birthday struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Morning struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type FeelsLike struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Saying struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Transl struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Source struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Color struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Num struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Overview struct {
	Value string `json:"value"`
	Color string `json:"color"`
}
