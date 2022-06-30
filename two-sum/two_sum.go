package main

import "github.com/emirpasic/gods/maps/hashmap"

func TwoSumHashTableImplementation(nums []int, target int) []int {
	m := hashmap.New()

	for i := 0; i < len(nums); i++ {
		value, found := m.Get(nums[i])
		if found {
			m.Put(nums[i], append(value.([]int), i))
		}

		m.Put(nums[i], []int{i})
	}

	for i := 0; i < len(nums); i++ {
		number1 := nums[i]
		number2 := target - number1

		indexes, found := m.Get(number2)
		if found {
			number2Index := indexes.([]int)[0]
			if number2Index == i {
				if len(indexes.([]int)) > 1 {
					return []int{i, indexes.([]int)[1]}
				}

				continue
			}

			return []int{i, number2Index}
		}
	}

	return nil
}
