package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type StoreService struct {
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

func (s *StoreService) GetOneStudentName() string {
	return s.GetName()

}

func (s *StoreService) GetName() string {
	names := []string{
		"赵领宇", "刘其美", "张子涵",
		"赵格祥", "吴吕诺", "乔浩宇",
		"郝雨萱", "邓思瑶", "文浩山",
		"杨浩翔", "段韶环",
	}
	res, err := rand.Int(rand.Reader, big.NewInt(int64(len(names))))
	if err != nil {
		fmt.Println(err)
	}
	return names[int(res.Int64())]
}
