package flags

import "github.com/urfave/cli/v2"

// CustomFlags : 自定义命令行参数
var CustomFlags = []cli.StringFlag{
	cli.StringFlag{
		Name: "dbhost",
		//Value: "120.24.177.228:3306",
		Value: "10.10.10.44:3306",
		Usage: "database address",
	},

	cli.StringFlag{
		Name:  "username",
		Value: "xtec-fa",
		//Value: "root",
		Usage: "database username",
	},

	cli.StringFlag{
		Name:  "password",
		Value: "16170hc2!",
		//Value: "root",
		Usage: "database password",
	},

	cli.StringFlag{
		Name:  "dbname",
		Value: "yamazen",
		Usage: "database",
	},
}
