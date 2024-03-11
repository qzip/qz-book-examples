CREATE TABLE "cas" ( "w3cdid" TEXT NOT NULL UNIQUE, "namespace" TEXT NOT NULL DEFAULT 'in.innomon.b2b', "cas_data" JSON, "tmstamp" DATETIME NOT NULL, PRIMARY KEY("w3cdid") )
