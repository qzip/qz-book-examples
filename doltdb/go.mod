module doltdb

go 1.21.6

require {
	qz v0.0.0
)

replace (
	merkle => ../../quartz/merkle
	qz => ../../quartz/qz
)
