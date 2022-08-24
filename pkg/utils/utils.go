/**
 * @Author: Sun
 * @Description:
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 2022/8/23 23:04
 */

package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomString() string {
	var tag = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	rand.Seed(time.Now().UnixNano())
	res := "#"
	for i := 1; i <= 6; i++ {
		res += tag[rand.Intn(len(tag))]
	}

	return res
}

func GetLovedDay(start_time, end_time string) int64 {
	var day int64
	t1, err := time.ParseInLocation("2006-01-02", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		day = diff / (3600 * 24)
		return day
	} else {
		return day
	}
}

func CalBirthDay(birthday string) string {
	tmpYear := time.Now().Year()
	year := strconv.Itoa(tmpYear)

	currentBirth := strings.Replace(birthday, "1997", year, -1)
	birth, err := time.ParseInLocation("2006-01-02", currentBirth, time.Local)
	if err == nil {
		diff := birth.Unix() - time.Now().Unix() //
		fmt.Println(diff)
		if diff > 0 {
			tmpDay := diff / (3600 * 24)
			day := strconv.Itoa(int(tmpDay))
			return day
		} else {
			tmpDay := (-diff) / (3600 * 24)
			t := tmpYear % 4
			if t == 0 {
				day := 366 - tmpDay
				day1 := strconv.Itoa(int(day))
				return day1
			} else {
				day := 365 - tmpDay
				day1 := strconv.Itoa(int(day))
				return day1
			}
		}
	}
	return ""
}
