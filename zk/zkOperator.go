package zk

import(
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type ZkClient struct{
	conn *zk.Conn
}

type ZkPathStat struct{
	PathStat *zk.Stat
}

type ZkEventChannal struct{
	EventChannel <-chan zk.Event
}

type ZkEvent struct{
	event zk.Event
}

var ZKACL = map[string][]zk.ACL{
	"PERMALL" : zk.WorldACL(zk.PermAll),
}
var ZKFLAG = map[string]int32{
	"PER" : 0,
	"TMP" : zk.FlagEphemeral,
	"SEQ" : zk.FlagSequence,
}

func NewZkClient(hosts []string, sessionTime time.Duration) (*ZkClient, error){
	conn, _, err := zk.Connect(hosts, sessionTime)
	if err != nil{
		return &ZkClient{}, err
	}
	return &ZkClient{conn}, nil
}

func CloseZkClient(zkclient *ZkClient){
	if zkclient.conn != nil{
		zkclient.conn.Close()
	}
}

func (zkclient ZkClient)CreatePath(path string, data string, flag int32, acls []zk.ACL) (string, error) {
	ret, err_create := zkclient.conn.Create(path, []byte(data), flag, acls)
	if err_create != nil{
		return ret, err_create
	}
	return ret, nil
}

func (zkclient ZkClient)SetPath(path string, data string, version int32) (*ZkPathStat, error){
	ret, err_set := zkclient.conn.Set(path, []byte(data), version)
	if err_set!= nil{
		return &ZkPathStat{}, err_set
	}
	return &ZkPathStat{ret}, nil
}

func (zkclient ZkClient)GetPath(path string) (string, *ZkPathStat, error){
	ret, pathStat, err_get := zkclient.conn.Get(path)
	if err_get != nil{
		return "", &ZkPathStat{}, err_get
	}
	return string(ret), &ZkPathStat{pathStat}, nil
}

func (zkclient ZkClient)DeletePath(path string, version int32) error{
	err_delete := zkclient.conn.Delete(path, version)
	if err_delete != nil{
		return err_delete
	}
	return nil
}

func (zkclient ZkClient)WatchPath(path string, Watchfunc func(*ZkEventChannal)) error {
	for {
		_, _, ch, err_watch := zkclient.conn.GetW(path)
		if err_watch != nil{
			return err_watch
		}
		Watchfunc(&ZkEventChannal{ch})
	}
}

func (zkclient ZkClient)WatchChildren(path string, Watchfunc func(*ZkEventChannal)) error {
	for{
		_, _, ch, err_Child := zkclient.conn.ChildrenW(path)
		if err_Child != nil{
			return err_Child
		}
		Watchfunc(&ZkEventChannal{ch})
	}
}