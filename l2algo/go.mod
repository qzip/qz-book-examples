module l2algo

go 1.20

require merkle v0.0.0

require github.com/cometbft/cometbft v0.37.2 // indirect

replace (
	merkle => ../quartz/merkle
	qz => ../quartz/qz
)
