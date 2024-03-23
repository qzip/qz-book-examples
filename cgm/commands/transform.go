package commands

import (
	"encoding/json"
	"fmt"
)

func init() {
	RegisterCommand("csv2db", csv2db, csv2dbHelp)
}

type CSV2DbParam struct {
	CSVfile string `json:"csvFile"`
	DbPath  string `json:"inFileName"`
}

func csv2db(param interface{}) error {

	return nil
}

func getParams(param interface{}) (cp *CSV2DbParam, err error) {
	cp = &CSV2DbParam{}
	if param == nil {
		err = fmt.Errorf("csv2db.getParams: nil param")
		return
	}
	by, err := json.Marshal(param)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, &cp)

	return
}

const csv2dbHelp = `
 csv2db: Converts CSV to cgm database.

`
