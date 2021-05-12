package main

import (
	"fmt"
	"math/rand"
	"time"
)

//分布式系统，gc等
func main() {
	var tn1, tn21, tn22, tn23 float64
	for i := 0; i < 10; i++ {
		nums := randNum()
		rue1, t1 := deleteNegative1(nums)
		rue21, t21 := deleteNegative2(nums, 20)
		rue22, t22 := deleteNegative2(nums, 100)
		rue23, t23 := deleteNegative2(nums, 10000)
		fmt.Println(nums[:100])
		fmt.Println("rue1: ", len(rue1), rue1[:100])
		fmt.Println("rue2 20: ", len(rue21), rue21[:100])
		fmt.Println("rue2 100: ", len(rue22), rue22[:100])
		fmt.Println("rue2 200: ", len(rue23), rue23[:100])
		fmt.Printf("rate: %v, t1: %d, t21: %d\n", float64(t21)/float64(t1), t1, t21)
		fmt.Printf("rate: %v, t1: %d, t22: %d\n", float64(t22)/float64(t1), t1, t22)
		fmt.Printf("rate: %v, t1: %d, t23: %d\n", float64(t23)/float64(t1), t1, t23)
		tn1 += float64(t1)
		tn21 += float64(t21)
		tn22 += float64(t22)
		tn23 += float64(t23)
	}
	fmt.Printf("rate1: %v,rate2: %v, rate3: %v, tn21: %f, tn22: %f, tn23: %f tn1: %f\n", (tn21 / tn1),
		tn22/tn1, tn23/tn1, tn21, tn22, tn23, tn1)
}

func randNum() []int {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan []int, 10)
	for i := 0; i < 10; i++ {
		go func() {
			tmp := make([]int, 100000)
			for i := 0; i < 100000; i++ {
				n := 1
				if rand.Int()%2 == 0 {
					n = -1
				}
				tmp[i] = rand.Int() * n
			}
			ch <- tmp
		}()
	}
	nums := make([]int, 1000000)
	for i := 0; i < 10; i++ {
		copy(nums[i*100000:(i+1)*100000], <-ch)
	}
	return nums
}

func deleteNegative1(nums []int) ([]int, int) {
	t := time.Now()
	rue := make([]int, 0)
	for _, n := range nums {
		if n >= 0 {
			rue = append(rue, n)
		}
	}
	return rue, int(time.Since(t))
}

func deleteNegative2(nums []int, gcount int) ([]int, int) {
	t := time.Now()
	ch := make(chan []int, gcount)
	gnum := len(nums) / gcount
	for i := 0; i < gcount; i++ {
		go func(i int) {
			tmp := make([]int, gnum+1)
			tmp[0] = i
			index := 1
			for _, n := range nums[i*gnum : (i+1)*gnum] {
				if n >= 0 {
					tmp[index] = n
					index++
				}
			}
			ch <- tmp[:index]
		}(i)
	}
	temp := make(map[int][]int)
	count := 0
	for i := 0; i < gcount; i++ {
		r := <-ch
		temp[r[0]] = r[1:]
		count += len(r[1:])
	}
	rue := make([]int, count)
	n := 0
	for i := 0; i < gcount; i++ {
		copy(rue[n:], temp[i])
		n += len(temp[i])
	}
	//fmt.Println(time.Since(t).String())
	return rue, int(time.Since(t))
}
