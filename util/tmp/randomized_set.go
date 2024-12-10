package tmp

import "math/rand"

// 实现RandomizedSet 类：
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

type RandommizeSet struct {
	nums    []int
	indices map[int]int
}

func BuildRandomizeSet() RandommizeSet {
	return RandommizeSet{[]int{}, map[int]int{}}
}

func (rs *RandommizeSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}

	id := len(rs.nums)
	rs.nums = append(rs.nums, val)
	rs.indices[val] = id
	return true
}

func (rs *RandommizeSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.nums = rs.nums[:last]
	rs.indices[rs.nums[id]] = id
	delete(rs.indices, val)
	return true
}

func (rs *RandommizeSet) Random() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
