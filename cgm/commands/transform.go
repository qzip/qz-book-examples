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
	UserId  string `json:"userid"`
}

func csv2db(param interface{}) error {
	if cp, err := getParams(param); err != nil {
		return err
	} else if err = cp.Process(); err != nil {
		return err
	}

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

func (cd *CSV2DbParam) Process() error {

	return nil
}

const csv2dbHelp = `
 csv2db: Converts CSV to cgm database.
 Sqlite database is used.


`
