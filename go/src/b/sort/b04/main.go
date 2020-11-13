package main

import (
	"fmt"
	"sort"
)

// HeroKind type.
type HeroKind int

// Const declare.
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// Hero struct.
type Hero struct {
	Name string
	Cate HeroKind
}

// Heros slice.
type Heros []*Hero

// Len func.
func (s Heros) Len() int {
	return len(s)
}

// Less func.
func (s Heros) Less(i, j int) bool {
	if s[i].Cate != s[j].Cate {
		return s[i].Cate < s[j].Cate
	}
	return s[i].Name < s[j].Name
}

// Swap func.
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
		&Hero{"测试", None},
	}
	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v\n", v)
	}
}
