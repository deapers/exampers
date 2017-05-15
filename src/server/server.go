// batchdeal project main.go
package main

import (
	"database/sql"
	"flag"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	. "trade"

	"github.com/google/logger"
	_ "github.com/wendal/go-oci8"
)

func main() {

	// 从XML文件中加载系统配置项，包括日志、数据库、RPC等
	LoadSysConf()

	//日志文件配置，命令行启动增加参数“ -verbose”增加内容输出到控制台
	var verbose = flag.Bool("verbose", GLogVerbose, "print info level logs to stdout")
	flag.Parse()
	lf, err := os.OpenFile(GLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()
	logger.Init("talogger", *verbose, true, lf)

	//数据库连接池
	os.Setenv("NLS_LANG", "")
	GDBCon, err = sql.Open("oci8", GDBurl)
	if err != nil {
		logger.Fatal(err)
	}
	defer GDBCon.Close()

	//加载到全局的业务相关Map和全局变量中
	Init()

	//注册后台RPC的服务
	t := new(BatchDeal)
	rpc.Register(t)

	//RPC服务启动，并监听对应端口
	l, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		logger.Infof("Listener tcp err: %s\n", err)
		return
	}

	//死循环，一直监控RPC客户端的调用
	for {
		logger.Infoln("wating...")
		conn, err := l.Accept()
		if err != nil {
			logger.Infof("accept connection err: %s\n", conn)
		}
		go jsonrpc.ServeConn(conn)
	}
}
