package commands

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func init() {
	RegisterCommand("csv2db", csv2db, csv2dbHelp)
}

type CSV2DbParam struct {
	CSVfile string `json:"csvFile"`
	DbPath  string `json:"dbPath"`
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

const stmt = "REPLACE INTO cgm (dtm, glucose, userid) VALUES(?,?,?)"

func (cd *CSV2DbParam) Process() error {
	db, err := sql.Open("sqlite3", cd.DbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	fileIn, err := os.Open(cd.CSVfile)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	pstmt, err := db.Prepare(stmt)
	if err != nil {
		return err
	}

	rd := bufio.NewReader(fileIn)
	i := 1
	for {

		ln, err := rd.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		ln = strings.TrimSpace(ln)
		if ln != "" {
			if i > 2 {
				row := strings.Split(ln, ",")
				if len(row) < 2 {
					return fmt.Errorf("CSV2DbParam.Process: expecting 2 column at %v row", i)
				}
				if _, err = pstmt.Exec(row[0], row[1], cd.UserId); err != nil {
					return err
				}
			}
			i++

		}
	}

	return nil
}

const csv2dbHelp = `
 csv2db: Converts CSV to cgm database.
 Sqlite database is used.

 CREATE TABLE "cgm" (
	"dtm"	DATETIME NOT NULL,
	"glucose"	INTEGER,
	"userid"	TEXT,
	PRIMARY KEY("dtm")
)

CREATE INDEX "uid" ON "cgm" (
	"userid"	ASC
)

The CSV file is assumed as having 2 column :
 1st row headers
 2nd row onwards: datetime, cgm values mg/DL


`
