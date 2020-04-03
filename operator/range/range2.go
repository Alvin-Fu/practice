package main

import "fmt"

type ID struct {
	id int
}

func main(){
	id1 := ID{1}
	id2 := ID{2}
	ids1 := []ID{id1, id2}
	ids2 := []*ID{}
	for _, id := range ids1{
		ids2 = append(ids2, &id)
	}
	for _, id := range ids2{
		fmt.Println((*id).id)
	}
}
