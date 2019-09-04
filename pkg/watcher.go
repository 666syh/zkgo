package pkg

import(
	"sync"
	"strings"
	"fmt"
	log "zkgo/utils/logger"
	"crypto/rand"
	"math/big"
	"time"
	zkop "zkgo/zk"
)

var EvaHosts []string
var Mux sync.RWMutex
var(
	urlprefix = "http://"
	urlsuffix = "/test"
)
const(
	ZKRoot = "/server"
	actionName = "/test"
)
func DoWatch() error{
	hosts := []string{"127.0.0.1:2181"}
	conn, err := zkop.NewZkClient(hosts, time.Second*5)
	if err != nil {
		return err
	}
	go watcher(conn, ZKRoot)
	go initEvaHosts(conn, ZKRoot)
	return nil
}

func watcher(conn *zkop.ZkClient, prefix string){
	conn.WatchChildren(prefix, func(ech *zkop.ZkEventChannal){
		for _ = range ech.EventChannel{
			resetHosts(conn, prefix)
		}
	})
}

func resetHosts(conn *zkop.ZkClient, prefix string){
	EvaHosts = []string{}
	pre := fmt.Sprintf("%v", actionName)
	children, _, err := conn.GetChildren(prefix)
	if err != nil{
		log.Error.Println(err)
		return
	}
	Mux.Lock()
	defer Mux.Unlock()
	for _, child := range children {
		if strings.Index(child, pre) >= 0 {
			temp := strings.Split(child[:len(child)-10], "|")
			evahost := temp[1]+":"+temp[2]
			EvaHosts = append(EvaHosts, evahost)
		}
	}
	log.Info.Println(EvaHosts)
}

func GetRandomHost() string{
	Mux.RLock()
	defer Mux.RUnlock()
	if len(EvaHosts) > 0{
		max := big.NewInt(int64(len(EvaHosts)))
		randIndex, err := rand.Int(rand.Reader, max)
		if err!=nil{
			log.Error.Print("rand index error", err)
			return ""
		}
		log.Info.Println(randIndex.Uint64())
		return urlprefix+EvaHosts[randIndex.Uint64()]+urlsuffix
	}
	return ""
}

func initEvaHosts(conn *zkop.ZkClient, prefix string){
	resetHosts(conn, prefix)
}