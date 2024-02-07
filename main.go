package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cardFaceNumber = 0
const cardZone = 1
const (
	deck = iota
	bin
)
const cardPos = 2

var state []*[]int

func zeroTransformation(state []*[]int, cards []*[]int, transform bool) (transformed bool, length int) {
	card := cards[0]
	length = 1
	if (*card)[cardZone] == deck && (*card)[cardPos] == maxPos(state, deck) {
		if transform {
			(*card)[cardZone] = bin
			(*card)[cardPos] = maxPos(state, bin) + 1
		}
		transformed = true
		return
	}

	transformed = false
	return
}


func permutations(arr []*[]int)[][]*[]int{
    var helper func([]*[]int, int)
    res := [][]*[]int{}

    helper = func(arr []*[]int, n int){
        if n == 1{
            tmp := make([]*[]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}




//recursive function to generate all subsets
func generateSubsets(nums []*[]int, index int, subsets [][]*[]int) [][]*[]int {
	if index == len(nums) {
		return subsets
	}
	if index == 0 {
		subsets = append(subsets, []*[]int{})
	}
	n := len(subsets)
	for i := 0; i < n; i++ {
		set := make([]*[]int, len(subsets[i]))
		copy(set, subsets[i])
		set = append(set, nums[index])
		subsets = append(subsets, set)
	}
	return generateSubsets(nums, index+1, subsets)
}

	
func canTransformation(state []*[]int, transformation func([]*[]int, []*[]int, bool) (bool, int)) (transformable [][]*[]int) {
	var _, length = transformation(state, state, false)	
	subsets := generateSubsets(state, 0, nil)
	for _, subset := range subsets {
		if len(subset) == length {
			for _,  subset1:= range permutations(subset) {
				transformed, _ := transformation(state, subset1, false)
				if transformed {
					transformable = append(transformable, subset1)
				}
			}
		}
	}
		
	

    return
}


func maxPos(state []*[]int, zone int) (pos int) {
	pos = -1
	for _, card := range state {
		if (*card)[cardZone] == zone && (*card)[cardPos] > pos {
			pos = (*card)[cardPos]
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(2)
		card := []int{
			randNum,
			0,
			maxPos(state, 0) + 1,
		}
		state = append(state, &card)
	}
	t := canTransformation(state, zeroTransformation)
	fmt.Println("Subsets:")
	for _, subset := range t {
		for _, card := range subset {
			fmt.Print(*card, " ")
		}
		fmt.Println()
	}
	
}
