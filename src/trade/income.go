package trade

import (
	"log"
)

type IncomeSplit struct {
	Service
}

func (this IncomeSplit) Excute() error {
	log.Println("Begin IncomeSplit...")
	log.Println(this)
	log.Println("End IncomeSplit.")

	return nil
}

type IncomeInvest struct {
	Service
}

func (this IncomeInvest) Excute() error {
	log.Println("Begin IncomeInvest...")
	log.Println(this)
	log.Println("End IncomeInvest.")

	return nil
}
