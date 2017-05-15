// trade project svchannel.go
package trade

import (
	"github.com/google/logger"
)

type Service struct {
	TaskCode string
	TaskName string
}

func (this *Service) New(code string, name string) error {
	this.TaskCode = code
	this.TaskName = name
	return nil
}

func (this Service) Excute() error {
	logger.Infoln("Service start...")
	logger.Infoln(this)
	logger.Infoln("Service end!")
	return nil
}

type BatchDeal int

func (this *BatchDeal) TaskExcute(sv Service, rtn *int) error {
	switch sv.TaskCode {
	case "DAYINIT":
		var d DayInit
		d.New(sv.TaskCode, sv.TaskName, "20170421", "20170422")
		d.Excute()
	case "TRADEACONFIRM_CTA":
		var d TradeConfirm
		d.New(sv.TaskCode, sv.TaskName)
		d.Excute()
	default:
		logger.Errorln("An invalid Service")
	}

	*rtn = 0
	return nil
}
