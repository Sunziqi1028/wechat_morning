/**
 * @Author: Sun
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/8/23 21:58
 */

package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"wechat_message/pkg/model"
)

func NewConfig() (*model.Config, error) {
	filepathTmp, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(filepathTmp)
	var cfg *model.Config
	configPath := filepath.Join(filepathTmp, "config/config.toml")
	if _, err = toml.DecodeFile(configPath, &cfg); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cfg, err
}
