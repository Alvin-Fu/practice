package adapter

import "fmt"

var (
	GameTypeReadAlert int32= 1
	GameTypeWarcraft int32 = 2
)

// 新的玩游戏接口
type PlayGames interface {
	PlayGame(gameType int32)
}
// 老的玩游戏接口
type OldPlayGames struct {

}

func (*OldPlayGames) RedAlert(){
	fmt.Println("32 bit Red Alert")
}

func (*OldPlayGames) Warcraft(){
	fmt.Println("32 bit Warcraft ")
}

// 适配器
type PlayGameOSAdapter struct {
	OldPlayGames
}

func NewOSAdapter()*PlayGameOSAdapter {
	return &PlayGameOSAdapter{}
}

func (p *PlayGameOSAdapter) PlayGame(gameType int32){
	switch gameType {
	case GameTypeReadAlert:
		p.OldPlayGames.RedAlert()
	case GameTypeWarcraft:
		p.OldPlayGames.Warcraft()
	default:
		fmt.Println("unsupported types")
	}
}
