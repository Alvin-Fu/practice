package main

import "fmt"

func main(){
    arr := []int{2, 1, 3, 1, 5, 4, 6, 8 }
    fmt.Println(sortArrayByParityII(arr))
}

func sortArrayByParityII(A []int) []int {
    index0 := make([]int, 0)
    index1 := make([]int, 0)
    for i, a := range A{
        n, m := a % 2, i %2
        if n != m {
            if n == 0  {
                if len(index1) != 0 {
                    A[i], A[index1[0]] = A[index1[0]], A[i]
                    index1 = index1[1:]
                } else {
                    index0 = append(index0, i)
                }
            } else {
                if len(index0) != 0 {
                    A[i], A[index0[0]] = A[index0[0]], A[i]
                    index0 = index0[1:]
                } else {
                    index1 = append(index1, i)
                }
            }
        }

    }
    return A
}
