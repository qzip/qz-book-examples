CREATE TABLE `cas` (
  `id` varchar(36) default(uuid()) primary key
  `w3cdid` varchar(512) NOT NULL,
  `namespace` varchar(512) DEFAULT ('one.vyapar.cas'),
  `cas_data` json,
  `tmstamp` timestamp,
   KEY (`w3cdid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin

