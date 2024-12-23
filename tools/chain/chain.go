package chain

import "github.com/gilperopiola/god/tools"

var _ tools.ChainMapper[int, string] = &Mapper{}

type Mapper map[int]string

func (cmap Mapper) Set(i int, s string) tools.ChainMapper[int, string] {
	cmap[i] = s
	return cmap
}
