package trade

import (
	"database/sql"
	"errors"

	"github.com/google/logger"
)

var (
	GDBCon      *sql.DB
	GLogPath    string
	GLogVerbose bool
	GRPCPort    string
	GDBurl      string

	GMap_Fund      map[string]TFundinfo
	GMap_Money     map[string]TMoneyProduct
	GMap_Param     map[string]TSysParameter
	GMap_ParamReal map[string]TSysParameter
	GCfmDate       string
	GReqDate       string
)

type TFundinfo struct {
	sfundcode    string
	sfundname    string
	smoneytype   string
	fissueprice  float64
	smanagername string
	strusteecode string
	sissuedate   string
	ssetupdate   string
	fmaxbala     float64
	fmaxshares   float64
	fminbala     float64
	fminshares   float64
	ffactcollect float64
}

func Init() {
	var err error

	err = LoadSysParameter(GDBCon)
	if err != nil {
		logger.Fatal(err)
	}

	err = LoadSysParameterReal(GDBCon)
	if err != nil {
		logger.Fatal(err)
	}

	err = LoadFundInfo(GDBCon)
	if err != nil {
		logger.Fatal(err)
	}

	err = LoadMoneyProduct(GDBCon)
	if err != nil {
		logger.Fatal(err)
	}

	GCfmDate, err = GetSysParameter("SysDate")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("GCfmDate=[%s]", GCfmDate)

	GReqDate = GetRequestdate(GCfmDate + "145959")
	logger.Infof("GReqDate=[%s]", GReqDate)

}

func LoadFundInfo(db *sql.DB) error {
	var fnd TFundinfo
	var rtn error

	rows, err := db.Query(
		"    select a.c_fundcode,a.c_fundname,a.c_moneytype,a.f_issueprice,a.c_managername,a.c_trusteecode, " +
			"       to_char(a.d_issuedate,'yyyymmdd'),to_char(a.d_setupdate,'yyyymmdd'),a.f_maxbala,a.f_maxshares,a.f_minbala,a.f_minshares,a.f_factcollect " +
			"  from tfundinfo a " +
			" where a.c_fundstatus not in ('6','9') " +
			" order by a.c_fundcode ")
	if err != nil {
		logger.Fatal(err)
		rtn = errors.New("db error")
	}

	//创建map或清空原有的map
	if GMap_Fund != nil {
		for k, _ := range GMap_Fund {
			delete(GMap_Fund, k)
		}
	} else {
		GMap_Fund = make(map[string]TFundinfo)
	}

	for rows.Next() {
		rows.Scan(&fnd.sfundcode, &fnd.sfundname, &fnd.smoneytype, &fnd.fissueprice, &fnd.smanagername, &fnd.strusteecode,
			&fnd.sissuedate, &fnd.ssetupdate, &fnd.fmaxbala, &fnd.fmaxshares, &fnd.fminbala, &fnd.fminshares, &fnd.ffactcollect)

		GMap_Fund[fnd.sfundcode] = fnd
	}
	logger.Infoln(GMap_Fund)
	rows.Close()

	return rtn
}

func GetFundInfo(fundcode string) (*TFundinfo, error) {
	fundinfo, ok := GMap_Fund[fundcode]
	if !ok {
		return nil, errors.New("Not Found!")
	}
	return &fundinfo, nil
}

type TMoneyProduct struct {
	sfundcode          string
	sassgintype        string
	sincometailtype    string
	sinvestcircle      string
	sinvestmonth       string
	sinvestday         string
	sincomeinvesttype  string
	fmaxredeembyallot  float64
	ssubendtime        string
	sbankperiod        string
	ssubincometype     string
	smanageracco       string
	fsingleminsub      float64
	fsingleminrdm      float64
	fsinglemaxrdm      float64
	fmaxtotalredeem    float64
	fmaxtotalnetredeem float64
	fredeemratio       float64
	fmaxallot          float64
	fmaxnetallot       float64
	faccomaxallot      float64
	faccomaxnetallot   float64
	faccomaxrdm        float64
	fsinglemaxsub      float64
	sbatchtime         string
	sist0subcanrdm     string
	sadjustrule        string
	fpermitnum         float64
	ft1rdmmaxratio     float64
	faccomaxshares     float64
	faccomaxt1rdm      float64
	fsinglemint1rdm    float64
	fsinglemaxt1rdm    float64
	sincometypebyrdm   string
	sregistrole        string
	sincomedeadline    string
}

func LoadMoneyProduct(db *sql.DB) error {
	var pro TMoneyProduct
	var rtn error

	rows, err := db.Query(
		"    select a.c_fundcode, a.c_assgintype, a.c_incometailtype, a.c_investcircle, " +
			"		a.c_investmonth, a.c_investday, a.c_incomeinvesttype, a.f_maxredeembyallot, " +
			" 		a.c_subendtime, a.c_bankperiod, a.c_subincometype, a.c_manageracco, " +
			" 		a.f_singleminsub, a.f_singleminrdm, a.f_singlemaxrdm, a.f_maxtotalredeem, " +
			" 		a.f_maxtotalnetredeem, a.f_redeemratio, a.f_maxallot, a.f_maxnetallot, " +
			" 		a.f_accomaxallot, a.f_accomaxnetacclot, a.f_accomaxrdm, a.f_singlemaxsub, " +
			" 		a.c_batchtime, a.c_ist0subcanrdm, a.c_adjustrule, a.f_permitnum, " +
			" 		a.f_t1rdmmaxratio, a.f_accomaxshares, a.f_accomaxt1rdm, a.f_singlemint1rdm, " +
			" 		a.f_singlemaxt1rdm, a.c_incometypebyrdm,a.c_registrole,a.c_incomedeadline" +
			"  from tmoneyproduct a" +
			" order by a.c_fundcode ")
	if err != nil {
		logger.Fatal(err)
		rtn = errors.New("db error")
	}

	//创建map或清空原有的map
	if GMap_Money != nil {
		for k, _ := range GMap_Money {
			delete(GMap_Money, k)
		}
	} else {
		GMap_Money = make(map[string]TMoneyProduct)
	}

	for rows.Next() {
		rows.Scan(&pro.sfundcode, &pro.sassgintype, &pro.sincometailtype, &pro.sinvestcircle,
			&pro.sinvestmonth, &pro.sinvestday, &pro.sincomeinvesttype, &pro.fmaxredeembyallot,
			&pro.ssubendtime, &pro.sbankperiod, &pro.ssubincometype, &pro.smanageracco,
			&pro.fsingleminsub, &pro.fsingleminrdm, &pro.fsinglemaxrdm, &pro.fmaxtotalredeem,
			&pro.fmaxtotalnetredeem, &pro.fredeemratio, &pro.fmaxallot, &pro.fmaxnetallot,
			&pro.faccomaxallot, &pro.faccomaxnetallot, &pro.faccomaxrdm, &pro.fsinglemaxsub,
			&pro.sbatchtime, &pro.sist0subcanrdm, &pro.sadjustrule, &pro.fpermitnum,
			&pro.ft1rdmmaxratio, &pro.faccomaxshares, &pro.faccomaxt1rdm, &pro.fsinglemint1rdm,
			&pro.fsinglemaxt1rdm, &pro.sincometypebyrdm, &pro.sregistrole, &pro.sincomedeadline)

		GMap_Money[pro.sfundcode] = pro
	}
	logger.Infoln(GMap_Money)
	rows.Close()

	return rtn
}

func GetMoneyProduct(fundcode string) (*TMoneyProduct, error) {
	moneypro, ok := GMap_Money[fundcode]
	if !ok {
		logger.Errorf("fundcode[%s] not Found!", fundcode)
		return nil, errors.New("Not Found!")
	}
	return &moneypro, nil
}

type TSysParameter struct {
	sclass    string
	sitem     string
	svalue    string
	sdescribe string
}

func LoadSysParameter(db *sql.DB) error {
	var sys TSysParameter
	var rtn error

	rows, err := db.Query(
		"    select a.c_class, a.c_item, a.c_value, a.c_describe " +
			"  from tsysparameter a " +
			" order by a.l_order ")
	if err != nil {
		logger.Fatal(err)
		rtn = errors.New("db error")
	}

	//创建map或清空原有的map
	if GMap_Param != nil {
		for k, _ := range GMap_Param {
			delete(GMap_Param, k)
		}
	} else {
		GMap_Param = make(map[string]TSysParameter)
	}

	for rows.Next() {
		rows.Scan(&sys.sclass, &sys.sitem, &sys.svalue, &sys.sdescribe)
		GMap_Param[sys.sitem] = sys
	}
	logger.Infoln(GMap_Param)
	rows.Close()

	return rtn
}

func GetSysParameter(item string) (string, error) {
	param, ok := GMap_Param[item]
	if !ok {
		return "", errors.New("Not Found!")
	}
	return param.svalue, nil
}

func LoadSysParameterReal(db *sql.DB) error {
	var sys TSysParameter
	var rtn error

	rows, err := db.Query(
		"    select a.c_class, a.c_item, a.c_value, a.c_describe " +
			"  from tsysparameter_real a " +
			" order by a.l_order ")
	if err != nil {
		logger.Fatal(err)
		rtn = errors.New("db error")
	}

	//创建map或清空原有的map
	if GMap_ParamReal != nil {
		for k, _ := range GMap_ParamReal {
			delete(GMap_ParamReal, k)
		}
	} else {
		GMap_ParamReal = make(map[string]TSysParameter)
	}

	for rows.Next() {
		rows.Scan(&sys.sclass, &sys.sitem, &sys.svalue, &sys.sdescribe)
		GMap_ParamReal[sys.sitem] = sys
	}
	logger.Infoln(GMap_ParamReal)
	rows.Close()

	return rtn
}

func GetSysParameterReal(item string) (string, error) {
	param, ok := GMap_ParamReal[item]
	if !ok {
		return "", errors.New("Not Found!")
	}
	return param.svalue, nil
}

func GetRequestdate(acceptdate string) string {
	var reqdate string

	sSysdate, err := GetSysParameterReal("SysDate")
	if err != nil {
		logger.Fatal(err)
	}

	if acceptdate < sSysdate+"150000" {
		reqdate = sSysdate
	} else {
		reqdate, err = GetSysParameterReal("NextDate")
		if err != nil {
			logger.Fatal(err)
		}
	}

	return reqdate
}
