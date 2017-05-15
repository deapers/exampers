package trade

import (
	"database/sql"
	"errors"
	"strconv"
	"sync"

	"github.com/google/logger"
)

const (
	MAXPARTITION    = 10
	MAXSUBPARTITION = 10
)

type TRequest struct {
	sfundacco       string
	stradeacco      string
	sfundcode       string
	ssharetype      string
	sagencyno       string
	snetno          string
	schildnetno     string
	sotheracco      string
	sothertrade     string
	sotheragency    string
	sothernetno     string
	sothercustno    string
	srequestno      string
	srequestdate    string
	srequesttime    string
	smachinetime    string
	ssysdate        string
	scdate          string
	scserialno      string
	fbalance        float64
	fshares         float64
	fconfirmbalance float64
	fconfirmshares  float64
	sbusinflag      string
	soutbusinflag   string
	fnetvalue       float64
	stastatus       string
	staresult       string
	stacause        string
	sispresub       string
	frealshares     float64
	frealincome     float64
	fpret1subshare  float64
	fpresubshare    float64
	fothershares    float64
	ftotalshares    float64
	fothertsubshare float64
	fothersubshare  float64
	sist0subcanrdm  string
	sreserve        string
	shopedate       string
	slastsubdate    string
	scdate1         string
	scfmdate        string
	sadjustrule     string
	scpflag         string
	slastsubmodify  string
	sothersubmodify string
	scapitalmode    string
	sbankacco       string
	sregistrole     string
	dacceptdate     string
	sotherserialno  string
	slineid         string
	ipartition      int32
}

type TradeConfirm struct {
	Service
}

func (this TradeConfirm) Excute() error {
	var waitPart sync.WaitGroup

	logger.Infoln("Begin TradeConfirm...")
	for i := 0; i < MAXPARTITION; i++ {
		waitPart.Add(1)
		go DealRequestPart(GDBCon, &waitPart, int32(i))
	}
	waitPart.Wait()

	logger.Infoln("End TradeConfirm.")
	return nil
}

func DealRequestPart(db *sql.DB, wp *sync.WaitGroup, prt int32) error {
	var waitSubPart sync.WaitGroup
	logger.Infof("Begin DealRequest[%d]... ", prt)

	for i := 0; i < MAXSUBPARTITION; i++ {
		waitSubPart.Add(1)
		go DealRequestCfm(db, &waitSubPart, prt, int32(i))
	}
	waitSubPart.Wait()
	wp.Done()

	logger.Infof("End DealRequest[%d] ", prt)
	return nil
}

func DealRequestCfm(db *sql.DB, wsp *sync.WaitGroup, prt int32, subprt int32) error {
	var rtn error
	var sql string
	logger.Infof("Begin DealRequestCfm[%d,%d]... ", prt, subprt)
	sql = "select vc_fundacco, nvl(vc_tradeacco, '-'), vc_fundcode," +
		"       nvl(c_sharetype, 'A') c_sharetype, '207' vc_branchcode, '207' vc_netno," +
		"       vc_netno vc_childnetno," +
		"       decode(c_businflag, '098', nvl(vc_otheracco, '-'), vc_otheracco) vc_otheracco," +
		"       decode(c_businflag,'098',nvl(vc_othertradeacco, '-'),vc_othertradeacco) vc_othertradeacco," +
		"       null c_otheragencyno, null vc_othernetno," +
		"       nvl(vc_requestno, '*') vc_requestno," +
		"       nvl(to_char(vc_requestdate,'YYYYMMDD'), '20010101') vc_requedate," +
		"       nvl(vc_requesttime, '*') vc_requesttime, '20010101000101' vc_machinedate," +
		"       nvl(to_char(vc_sysdate, 'YYYYMMDD'), '20010101') vc_sysdate, null vc_confirmdate," +
		"       en_balance, en_share, 0 en_confirmbala, 0 en_confirmshare," +
		"       c_businflag,substr(vc_acceptdate, 1, 8) vc_accepdate, vc_otherserialno," +
		"       l_partition" +
		"  from trequest_cds t" +
		" where nvl(c_tastatus, 0) = '0'" +
		"   and c_confirmflag = '9' " +
		"   and c_state = '2' " +
		"   and c_businflag in ('022', '098') " +
		"   and l_partition = " + strconv.Itoa(int(prt)) +
		"   and l_subpartion = decode(" + strconv.Itoa(int(subprt)) + ", null, l_subpartion, " + strconv.Itoa(int(subprt)) + ") " +
		"   and vc_requestdate = to_date('" + GReqDate + "', 'YYYYMMDD') " +
		" order by t.vc_tradeacco, t.vc_acceptdate, t.c_businflag "
	logger.Infoln(sql)
	rows, err := db.Query(sql)
	if err != nil {
		logger.Fatal(prt, subprt, err)
		rtn = errors.New("db error")
	}

	for rows.Next() {
		logger.Infof("Deal trequest_cds[%d][%d]", prt, subprt)
	}
	rows.Close()

	wsp.Done()
	logger.Infof("End DealRequestCfm[%d,%d] ", prt, subprt)
	return rtn
}
