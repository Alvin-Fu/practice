package main

import (
	"practice/operator/pb"
	"fmt"
	"time"
)

func main(){
	t := time.Now()
	defer fmt.Println("first: ",time.Since(t))
	defer func() {
		fmt.Println("func: ",time.Since(t))
	}()
	time.Sleep(2 * time.Second)
	//req := &mypb.GetMatchinfoReq{
	//	Uid: 122,
	//}
	//resp := new(mypb.GetMatchinfoResp)
	//test(req, resp)
	//fmt.Printf("resp: %s\n", resp.String())
}

func test(req *mypb.GetMatchinfoReq, resp *mypb.GetMatchinfoResp){
	fmt.Printf("req: %s\n", req.String())
	defer fmt.Printf("resp 1: %v\n", resp)
	resp.Errmsg = "ok"
	resp.Errcode = 0
	resp.Data = append(resp.Data, &mypb.Matchinfo{
		UserRank: 2,
		UserReward: true,
		StartTime: time.Now().Unix(),
		MatchId: "asljlsk",
		GameId: 12,
	})
	return
}
