// json-rpc project json-rpc.go
package main

// first we create a simple golang rpc server based on socket
import (
	"batchtask"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9998", 1000*1000*1000*30)
	if err != nil {
		fmt.Printf("create client err:%s\n", err)
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var reply int
	task := &batchtask.BatchTask{
		TaskCode: "BatchDeal",
		TaskName: "批处理1",
	}
	err = client.Call("BatchTask.Excute", task, &reply)

	fmt.Printf("reply: %s[%s] excuted, return %s, err: %s\n", task.TaskCode, task.TaskName, reply, err)

}
