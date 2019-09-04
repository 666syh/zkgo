package main

import(
	"github.com/go-kit/kit/log"
	"zkgo/kit/service"
	"zkgo/kit/endpoints"
	"zkgo/kit/transport"
	zkop "zkgo/zk"
	processer "zkgo/pkg"
	"os"
	"net/http"
	"time"
	"fmt"
	"strconv"
)

func main(){
	retCh := processer.NewChans()
	err := zkfunc(retCh)
	if err != nil{
		close(retCh.RetChan)
		os.Exit(0)
	}
	kitfunc(retCh)
}
func zkfunc(retCh *processer.Chans)error{
	hosts := []string{"localhost:2181"}
	conn, err := zkop.NewZkClient(hosts, time.Second*10)
	if err != nil{
		fmt.Println(err)
		return err
	}else{
		fmt.Println("success")
		// defer zkop.CloseZkClient(conn)
	}
	hostName, _ :=  os.Hostname()
	timeUnix := time.Now().Unix()
	_, err = conn.CreatePath("/"+hostName+strconv.FormatInt(timeUnix, 10), "connect", zkop.ZKFLAG["TMP"], zkop.ZKACL["PERMALL"])
	if err != nil{
		fmt.Println(err)
		return err
	}

	go conn.WatchChildren("/iamServer", func (ech *zkop.ZkEventChannal){
		ch := ech.EventChannel
		for event := range ch{
			fmt.Println("=================")
			fmt.Println("path:", event.Path)
			fmt.Println("type:", event.Type)
			fmt.Println("state", event.State)
			fmt.Println("=================")
			fmt.Println("")
		}
	})
	fmt.Println("starting watching")
	go processer.Register(retCh)
	return nil
}

func kitfunc(retCh *processer.Chans){
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewZkService(retCh.RetChan)
	epoint := endpoint.MakeEndpoints(service, logger)

	httpHandler := transport.NewHttpHandler(epoint, logger)
	http.ListenAndServe(":7777", httpHandler)
}