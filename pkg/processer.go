package pkg

import(
	zkop "zkgo/zk"
	"time"
	log "zkgo/utils/logger"
)

const(
	ROOT = "/server"
)
type Chans struct {
	RetChan chan interface{}
}

func NewChans() *Chans{
	return &Chans{make(chan interface{})}
}

func Register(retCh *Chans){
	for target := range retCh.RetChan{
		go doRegister(target)
	}
}

func doRegister(target interface{}){
		hosts := []string{"127.0.0.1:2181"}
		v := target.(map[string]string)
		conn, err := zkop.NewZkClient(hosts, 10*time.Second)
		if err != nil{
			// fmt.Println(1, err)
			log.Error.Println(err)
			return 
		}
		_, err = conn.CreatePath(ROOT+"/"+v["action"]+"|"+v["host"]+"|"+v["port"], "", zkop.ZKFLAG["TMP"], zkop.ZKACL["PERMALL"])
		if err != nil{
			// fmt.Println(2, err)
			log.Error.Println(err)
			zkop.CloseZkClient(conn)
			return
		}
		// fmt.Println("regist success")
		log.Info.Println("regist success")
		// fmt.Println(target)
}