package flags

import (
	"flag"
	"fmt"
	"os"
)

type FlagStruct struct {
	CnfigFlag string
	HostFlag  string
	UserFlag  string
	PwdFlag   string
	PortFlag  int
	DBFlag    string
}

func GetFlages() (*FlagStruct, error) {
	helpFlagLong := flag.Bool("help", false, "Show help message")
	hostFlag := flag.String("h", "", "DB Host")
	cnfigFlag := flag.String("c", "", "Config file path")
	userFlag := flag.String("u", "", "DB User")
	pwdFlag := flag.String("p", "", "DB Password")
	portFlag := flag.Int("P", 3306, "DB Port")
	dbFlag := flag.String("d", "", "DB Name")

	flag.Usage = func() {
		fmt.Println("Usage: ")
		fmt.Println("  -h, --host string    DB Host")
		fmt.Println("  -c, --config string  Config file path")
		fmt.Println("  -u, --user string    DB User")
		fmt.Println("  -p, --pwd string     DB Password")
		fmt.Println("  -P, --port int       DB Port (default: 3306)")
		fmt.Println("  -d, --db string      DB Name")
	}

	flag.Parse()

	// 如果用户没有输入任何参数或者输入 -help，显示帮助信息并退出
	if flag.NFlag() == 0 || *helpFlagLong {
		flag.Usage()
		os.Exit(0)
	}

	// 规则 1：-c 和 -h 不能同时提供
	if *cnfigFlag != "" && *hostFlag != "" {
		return nil, fmt.Errorf("cannot use -c (config) and -h (host) at the same time")
	}

	// 规则 2：如果提供了 -h，则必须提供 -u、-p、-d
	if *hostFlag != "" {
		if *userFlag == "" || *pwdFlag == "" || *dbFlag == "" {
			return nil, fmt.Errorf("when using -h (host), you must also provide -u (user), -p (password), and -d (dbname)")
		}
	}

	// 组装 FlagStruct
	flagstr := FlagStruct{}
	if *cnfigFlag != "" {
		flagstr.CnfigFlag = *cnfigFlag
	} else {
		flagstr.HostFlag = *hostFlag
		flagstr.UserFlag = *userFlag
		flagstr.PwdFlag = *pwdFlag
		flagstr.PortFlag = *portFlag
		flagstr.DBFlag = *dbFlag
	}

	return &flagstr, nil
}
