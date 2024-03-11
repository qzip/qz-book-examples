CREATE TABLE `cas` (
  `w3cdid` varchar(512) NOT NULL,
  `namespace` varchar(512) DEFAULT ('one.vyapar.cas'),
  `cas_data` json,
  `tmstamp` timestamp,
  PRIMARY KEY (`w3cdid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin

