// json-rpc project json-rpc.go
package main

// first we create a simple golang rpc server based on socket
import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	. "trade"

	_ "github.com/wendal/go-oci8"
)

func main() {
	var sFlowCode, sFlowName, sPFlowCode, sTaskTime, sDependsql, sMulti, sState, sRowid string
	var reply int
	var s Service

	/* rpc客户端初始化 */
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9998", 1000*1000*1000*30)
	if err != nil {
		fmt.Printf("create client err:%s\n", err)
		return
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)

	/* 连接批处理配置库 */
	os.Setenv("NLS_LANG", "")
	db, err := sql.Open("oci8", "cds1/oracle@127.0.0.1/orcl.js.local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		"    select a.vc_workflowname,a.vc_workflowcode,a.vc_pworkflowcode," +
			"       a.vc_tasktime,a.vc_dependsql,nvl(a.c_multi,'0'),a.c_state,a.rowid " +
			"  from tworkflow a," +
			"      (select c_value d_sysdate " +
			"         from tsysparameter_real a " +
			"        where a.c_item = 'SysDate') b " +
			" where a.vc_taskdate = b.d_sysdate " +
			" order by a.l_workflowgroup, a.l_executeorder")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&sFlowName, &sFlowCode, &sPFlowCode, &sTaskTime, &sDependsql, &sMulti, &sState, &sRowid)
		/*if sState == "0" {*/
		fmt.Printf("Info: [%s] start excute, waiting...\n", sFlowCode)
		s.New(sFlowCode, sFlowName)
		err = client.Call("BatchDeal.TaskExcute", s, &reply)
		if err != nil {
			fmt.Printf("ERROR:service exit with errors!!!\n")
			break
		}
		fmt.Printf("Info: [%s] is excuted success!\n", sFlowCode)
		/*} else {
			fmt.Printf("Info: [%s] is excuted, skipped.\n", sFlowCode)
		}*/

	}

	rows.Close()

}
