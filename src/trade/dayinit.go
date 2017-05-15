package trade

import (
	"log"
)

type DayInit struct {
	Service
	SysDate  string
	NextDate string
}

func (this *DayInit) New(code string, name string, sdate string, ndate string) *DayInit {
	this.TaskCode = code
	this.TaskName = name
	this.SysDate = sdate
	this.NextDate = ndate
	return this
}

func (this DayInit) Excute() error {
	log.Println("Begin DayInit...")
	log.Println("End DayInit.")

	return nil
}
