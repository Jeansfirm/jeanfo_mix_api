package main

// 解析任务之间的依赖关系，并输出一种可以执行的任务顺序

import "fmt"

func parse(task string, depts []string, input map[string][]string, res *[]string, done map[string]any) {
	if _, ok := done[task]; ok {
		return
	}
	for i := range depts {
		dt := depts[i]
		if _, ok := done[dt]; ok {
			continue
		}
		if _, ok := input[dt]; ok {
			parse(dt, input[dt], input, res, done)
		} else {
			*res = append(*res, dt)
			done[dt] = nil
		}
	}

	*res = append(*res, task)
	done[task] = nil

}

// func main() {
func _() {
	input := map[string][]string{
		"a": []string{"b", "c"},
		"c": []string{"d"},
	}

	res := &[]string{}
	done := map[string]any{}

	for task, depts := range input {
		parse(task, depts, input, res, done)
	}

	fmt.Println(*res)
}
