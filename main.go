package main

import (
	"encoding/json"
	"fmt"
	"read-excel-go-oracle/db"
	"read-excel-go-oracle/lib"
	"regexp"

	"strconv"
	"strings"

	_ "github.com/godror/godror"
	"github.com/gofrs/uuid"
	"github.com/schollz/progressbar/v3"

	structexcel "github.com/douyacun/go-struct-excel"
)

const (
	filename = "example.xlsx"
	sheet    = "all"
)

func lastString(s string, separator string) string {
	last := s[strings.LastIndex(s, separator)+1:]
	return last
}

func trans(e lib.SitefolderEntry) lib.Inventory {

	item := lib.Inventory{
		AppType:      e.Category,
		ClusterName:  e.ClusterName,
		ClusterGroup: e.ClusterGroup,
		Biz:          e.Biz,
		Purpose:      e.Purpose,
		Region:       e.Region,
		RackLocation: e.RackLocation,
		Brand:        e.Brand,
		Model:        e.Model,
		RamInfo:      e.Ram,
		DiskInfo:     e.Disk,
		CpuInfo:      e.CpuModel,
		RaidInfo:     e.RaidPlan,
		Bandwidth:    e.Bandwidth,
		NicInfo:      e.NicInfo,
		Hostname:     e.Hostname,
		SerialNumber: e.DeviceSN,
		ServiceIp:    e.ServiceIp,
		BmcMgmtIp:    e.BmcMgmtIp,
		BmcMgmtUser:  e.BmcMgmtUser,
		MonitorIp:    e.MonitorIp,
		// Status:       e.Status,
		DeviceFrom: e.DevFrom,
		CreateTime: e.CreateTime,
		QuitTime:   e.QuitTime,
		OsVersion:  e.OsVersion,
		Remarks1:   e.Remarks1,
		Remarks2:   e.Remarks2,
		Domain:     e.Domain,
	}
	u, _ := uuid.NewV4()
	item.UUID = strings.ToUpper(u.String())

	// item.Purpose = strings.Split(e.Purpose, ",")

	// NicNum
	if strings.Contains(e.NicInfo, "口") {

		nicNumStr := strings.Split(e.NicInfo, "口")[0]
		NicNum, err := strconv.Atoi(nicNumStr)
		if err != nil {
			item.NicNum = -1
		} else {
			item.NicNum = NicNum
		}

	} else {
		item.NicNum = -1
	}

	// status
	switch e.Status {
	case "待退网":
		item.Status = lib.StatusWaitQuiting
	case "故障":
		item.Status = lib.StatusWaitQuiting
	case "退网":
		item.Status = lib.StatusQuit
	case "移交":
		item.Status = lib.StatusMoved
	case "正常":
		item.Status = lib.StatusOK
	case "":
		item.Status = lib.StatusUnknown
	}

	// NicBondMode
	if strings.Contains(e.NicInfo, "bond") {
		const bondRegex = `bond\d`
		r := regexp.MustCompile(bondRegex)
		match := r.FindAllString(e.NicInfo, -1)

		item.NicBondMode = match[0]

	}

	// fmt.Println(e.Hostname)
	if strings.Contains(e.Hostname, "-") {

		short_hostname := lastString(e.Hostname, "-")

		item.ShortHostname = strings.ToLower(short_hostname)
	} else {
		item.ShortHostname = strings.ToLower(e.Hostname)
	}

	// fqdn = short + "." + fqdn

	// domain

	// IsCtrl
	if e.IsCtrl == "yes" {
		item.IsCtrl = 1

	} else if e.IsCtrl == "no" {
		item.IsCtrl = 0

	} else if e.IsCtrl == "na" {
		item.IsCtrl = -1

	} else {
		item.IsCtrl = -1
	}

	return item
}

func main() {

	excel, err := structexcel.OpenExcel(filename)
	if err != nil {
		fmt.Println(err)
	}

	sheet, err := excel.OpenSheet(sheet)
	if err != nil {
		panic(err)
	}
	if data, err := sheet.ReadData(lib.SitefolderEntry{}); err != nil {
		fmt.Println(err)
	} else if d, ok := data.([]*lib.SitefolderEntry); ok {
		b, err := json.Marshal(d)

		if err != nil {
			fmt.Println(err)
		}

		SitefolderData := []lib.SitefolderEntry{}

		json.Unmarshal(b, &SitefolderData)

		fmt.Println(len(SitefolderData))

		db.SyncTable()

		bar := progressbar.Default(int64(len(SitefolderData)))

		for _, item := range SitefolderData {

			new_item := trans(item)

			err = db.AddInventoryOne(new_item)

			if err != nil {
				fmt.Println(err) //exit
			}

			bar.Add(1)

		}

		fmt.Println("Done!")

	}
}
