package flags_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zsp108/dbbmsql/internal/flags"
)

func Test(t *testing.T) {
	t.Log("Test")
	flagstr, err := flags.GetFlages()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	t.Logf("config file path: %s, host: %s, user: %s, password: %s, port: %d, db: %s", flagstr.CnfigFlag, flagstr.HostFlag, flagstr.UserFlag, flagstr.PwdFlag, flagstr.PortFlag, flagstr.DBFlag)
}
