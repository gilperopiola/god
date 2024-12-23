package tools

/* -~-~-~-~ God's Tools -~-~-~-~ */

/* -~-~-~-~ Chain -~-~-~-~ */

type ChainMapper[KT comparable, VT any] interface {
	Set(key KT, val VT) ChainMapper[KT, VT]
}
