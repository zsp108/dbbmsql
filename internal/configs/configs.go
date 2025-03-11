package configs

import (
	"fmt"
	"os"

	"github.com/zsp108/dbbmsql/internal/flags"
	"github.com/zsp108/dbbmsql/internal/options"
	"github.com/zsp108/dbbmsql/pkg/utils"
)

var opts options.Options

/*
通过yaml 文件获取配置文件
*/
func getConfByPath(confPath string) (*options.Options, error) {
	// projectPath, err := utils.GetProjectRoot()

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %s", confPath)
	}

	// 通过yaml 文件获取配置存入 Options 结构体中

	err := utils.UnmarshalYAMLFile(confPath, &opts)
	if err != nil {
		return nil, err
	}
	return &opts, nil
}

func getConfByFlags(flags *flags.FlagStruct) (*options.Options, error) {
	// 通过命令行参数获取配置存入 Options 结构体中
	opts.MySQLOptions.Host = flags.HostFlag
	opts.MySQLOptions.Username = flags.UserFlag
	opts.MySQLOptions.Password = flags.PwdFlag
	opts.MySQLOptions.Port = flags.PortFlag
	opts.MySQLOptions.Database = flags.DBFlag
	return &opts, nil
}

func Ztest() {
	fmt.Println("test")
	flages, err := flags.GetFlages()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(flages.CnfigFlag)
	opts, err := getConfByPath(flages.CnfigFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(opts.MySQLOptions.Host)
}
