package dial

import (
	clientV3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

var E *clientV3.Client

func EtcdClient(etcdAddr string,timeOut int) (err error) {
	addr := strings.Split(etcdAddr,",")
	client,err := clientV3.New(clientV3.Config{
		Endpoints: addr,
		DialTimeout: time.Duration(timeOut) * time.Second,
	})
	E = client
	return
}

/*
func SetEtcd(key string,dataMgr []interface{}) (err error) {
	data ,err := json.Marshal(dataMgr)
	if err != nil {
		loggo.ErrorFormat("set etcd key-value marshal dataMgr to json err:%v\n",err)
		return
	}
	ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)
	_,err = E.Put(ctx,key,string(data))
	cancel()
	if err != nil {
		loggo.ErrorFormat("put failed err:%v", err)
		return
	}
	//UpdateData(dataMgr)
	return
}

func GetEtcd() {
	key := fmt.Sprintf("%s%s",Config.EtcdKeyPrefix,productKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp,err := dial.E.Get(ctx,key)
	cancel()
for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}


func UpdateDataMgr(dataMgr []interface{}) (err error) {
	var temp = make(map[int]SecProduct,1024)
	for _,v := range secProductArr {
		temp[v.ProductId] = v
	}
	SecConf.RWProductLock.Lock()
	SecConf.SecProductMap = temp
	SecConf.RWProductLock.Unlock()
	return
}


// watch key value
func WatchEtcd(key string,etcdAddr string,timeOut int) {
	addr := strings.Split(etcdAddr,",")
	client,err := clientV3.New(clientV3.Config{
		Endpoints: addr,
		DialTimeout: time.Duration(timeOut) * time.Second,
	})
	if err != nil {
		loggo.Error("watch etcd client err:",err)
		return
	}
	loggo.Stat("watch etcd key:%s run",key)
	for {
		rch := client.Watch(context.Background(),key)
		var secProductArr []SecProduct
		var getConfSucc = true

		for wresp := range rch {

			for _,ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Info("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value,&secProductArr)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", secProductArr)
				UpdateSecConf(secProductArr)
			}

		}

	}
}
*/
