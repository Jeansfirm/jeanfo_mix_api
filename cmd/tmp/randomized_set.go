package main

import (
	"fmt"
	util_tmp "jeanfo_mix/util/tmp"
)

func main() {
	randomizedSet := util_tmp.BuildRandomizeSet()
	randomizedSet.Insert(3)
	randomizedSet.Insert(4)
	randomizedSet.Remove(2)
	randomizedSet.Remove(3)
	randomizedSet.Insert(3)
	randomizedSet.Insert(1)
	randomizedSet.Insert(5)

	fmt.Println("random -- ", randomizedSet.Random())
	fmt.Println("random -- ", randomizedSet.Random())
	fmt.Println("random -- ", randomizedSet.Random())
	fmt.Println("random -- ", randomizedSet.Random())
}
