package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type UserInfo struct {
	Id    int64
	Name  string
	Token string
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	user := &UserInfo{
		Id:   123,
		Name: "hello",
	}
	token, err := createToken(user)
	if err != nil {
		log.Printf("err: %v", err)
		return
	}
	fmt.Println("token: ", token)
	parseToken(token)
}

func createToken(user *UserInfo) (string, error) {
	claim := jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte("token"))
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("token"), nil
	}
}

func parseToken(token string) (*UserInfo, error) {
	var user = new(UserInfo)
	parse, err := jwt.Parse(token, secret())
	if err != nil {
		log.Printf("parse err: %v", err)
		return nil, err
	}
	claim, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("err")
	}
	if !parse.Valid {
		return nil, fmt.Errorf("err")
	}
	user.Id = int64(claim["id"].(float64))
	user.Name = (claim["name"].(string))
	return user, nil
}
