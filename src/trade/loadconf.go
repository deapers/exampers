package trade

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/google/logger"
)

type SysConf struct {
	XMLName    xml.Name   `xml:"resources"`
	LogConf    LogConf    `xml:"log"`
	ServerConf ServerConf `xml:"server"`
	SchemaConf SchemaConf `xml:"schema"`
}

type LogConf struct {
	XMLName  xml.Name   `xml:"log"`
	Property []Property `xml:"property"`
}

type ServerConf struct {
	XMLName  xml.Name   `xml:"server"`
	Property []Property `xml:"property"`
}

type SchemaConf struct {
	XMLName  xml.Name   `xml:"schema"`
	Property []Property `xml:"property"`
}

type Property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}

func LoadSysConf() {
	var conf SysConf

	xmlfile, err := ioutil.ReadFile("../../conf/server.xml")
	if err != nil {
		logger.Fatal(err)
	}

	err = xml.Unmarshal(xmlfile, &conf)
	if err != nil {
		logger.Fatal(err)
	}
	for _, line := range conf.LogConf.Property {
		switch line.Name {
		case "LogPath":
			GLogPath = line.Value
			fmt.Println("GLogPatch = " + GLogPath)
		case "LogVerbose":
			GLogVerbose, err := strconv.ParseBool(line.Value)
			if err != nil {
				logger.Fatal(err)
			}
			fmt.Println("GLogVerbose = " + strconv.FormatBool(GLogVerbose))
		default:
			fmt.Println("[" + line.Name + "] is an invalid property of log! ")
		}
	}
	for _, line := range conf.ServerConf.Property {
		switch line.Name {
		case "RPCPort":
			GRPCPort = ":" + line.Value
			fmt.Println("GRPCPort = " + GRPCPort)
		default:
			fmt.Println("[" + line.Name + "] is an invalid property of server! ")
		}
	}
	for _, line := range conf.SchemaConf.Property {
		switch line.Name {
		case "DBurl":
			GDBurl = line.Value
			fmt.Println("GDBurl = " + GDBurl)
		default:
			fmt.Println("[" + line.Name + "] is an invalid property of server! ")
		}
	}

}
