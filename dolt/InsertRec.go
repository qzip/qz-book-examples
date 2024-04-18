package main

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// see: https://github.com/dolthub/driver/blob/main/example/main.go
	//_ "github.com/dolthub/driver"
	//https://pkg.go.dev/github.com/cyberphone/json-canonicalization#section-readme
	canjson "github.com/docker/go/canonical/json"
	//canjson "github.com/cyberphone/json-canonicalization"
)

const (
	// https://www.dolthub.com/repositories/innomon/Harvest2Invoice/data/main/cas
	DefaultW3Cprefix = "did:doltdb:innomon:Harvest2Invoice:cas:"
	DefaultNamespace = "in.innomon.ns.default"
	stmt             = "REPLACE INTO cas (w3cdid, namespace, cas_data, tmstamp) VALUES(?,?,?,?)"
)

type CasRecCfg struct {
	DbConnection string      `json:"dbConnection"`
	DbType       string      `json:"dbtype,omitempty"`
	W3Cprefix    string      `json:"w3cPrefix,omitempty"`
	Namespace    string      `json:"namespace,omitempty"`
	Data         interface{} `json:"data"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage %v <json config file>\n", os.Args[0])
		return
	}
	cfg, err := readCfg(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("main.readCfg: %s", err.Error()))
		return
	}
	json_data, err := json.Marshal(cfg.Data)
	if err != nil {
		log.Fatal(err)
		return
	}
	can_data, err := canjson.MarshalCanonical(json_data)
	if err != nil {
		log.Fatal(err)
		return
	}
	var db *sql.DB
	if cfg.DbType == "" {
		cfg.DbType = "mysql"
	}
	log.Printf("dbtype [%v], connection [%v]\n", cfg.DbType, cfg.DbConnection)
	if db, err = sql.Open(cfg.DbType, cfg.DbConnection); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	pstmt, err := db.PrepareContext(context.Background(), stmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pstmt.Close()
	h := hash(can_data)
	key := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(h)
	w3cdid := cfg.W3Cprefix + key
	log.Printf("w3cDID[%v] ns [%v]\n", w3cdid, cfg.Namespace)
	tmstamp := time.Now()
	if _, err = pstmt.ExecContext(context.Background(), w3cdid, cfg.Namespace, can_data, tmstamp); err == nil {
		log.Println("success")
	} else {
		log.Fatal(err)
	}

}

func readCfg(fname string) (*CasRecCfg, error) {
	cfg := &CasRecCfg{
		W3Cprefix: DefaultW3Cprefix,
		Namespace: DefaultNamespace,
	}
	jsonb, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonb, &cfg)
	return cfg, err
}

// Hash returns SHA-2 SHA-512 hash
func hash(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	return hash.Sum(nil)
}
