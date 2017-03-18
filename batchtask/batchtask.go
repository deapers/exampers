// batchtask project batchtask.go
package batchtask

import (
	"fmt"
)

type BatchTask struct {
	TaskCode string
	TaskName string
}

func (this *BatchTask) Excute(bTask BatchTask, rtn *int) error {
	fmt.Println(bTask)
	*rtn = 0
	return nil
}
