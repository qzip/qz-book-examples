CREATE TABLE "cas" (
	"id"	INTEGER NOT NULL,
	"w3cdid"	TEXT NOT NULL,
	"namespace"	TEXT NOT NULL DEFAULT 'in.innomon.b2b',
	"cas_data"	JSON,
	"tmstamp"	DATETIME NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
)
