package coin58

import (
	"testing"
	"net/http"
	. "../../GoEx/coin58"
	"github.com/nntaoli-project/GoEx"
	"time"
	// . "github.com/nntaoli-project/GoEx/coin58"
	"fmt"
)

var coin58 = New58Coin(http.DefaultClient, "0ac13122-7f1c-4a31-9bfe-abe65711d47f", "AD345B2F512EC626E97A1A920BD2154B")


func TestCoin58_GetTicker(t *testing.T) {
	t.Log(coin58.GetTicker(goex.NewCurrencyPair2("58b_btc")))
}

func TestCoin58_GetDepth(t *testing.T) {
	b := time.Now()
	dep , _ := coin58.GetDepth(1 , goex.NewCurrencyPair2("58b_ucc"))
	//dep, _ := coin58.GetDepth(2, goex.BTC_USD)
	t.Log(time.Now().Sub(b))
	t.Log(dep.BidList)
	t.Log(dep.AskList)
}

func TestCoin58_GetAccount(t *testing.T) {
	acc , err := coin58.GetAccount()
	fmt.Println("------------------")
	t.Log(err)
	t.Log(acc)
}

func TestCoin58_LimitSell(t *testing.T) {
	ord , err := coin58.LimitSell("11.2" , "0.17" , goex.NewCurrencyPair2("58b_ucc"))
	t.Log(err)
	t.Log(ord)
}

func TestCoin58_GetOneOrder(t *testing.T) {
	ord , _ := coin58.GetOneOrder("23613" , goex.NewCurrencyPair2("58b_ucc")) //23478
	t.Logf("%+v" , ord)
}

func TestCoin58_GetUnfinishOrders(t *testing.T) {
	t.Log(coin58.GetUnfinishOrders(goex.NewCurrencyPair2("58b_usdt")))
}

func TestCoin58_CancelOrder(t *testing.T) {
	t.Log(coin58.CancelOrder("72" , goex.NewCurrencyPair2("58b_usdt")))
}

