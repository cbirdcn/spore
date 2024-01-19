package service

import (
	"fmt"
	"sync"
	"time"
)

// 动态业务Map
var AccountsMap sync.Map

// 同步状态Map
var SaveMap sync.Map
var LoadMap sync.Map

var UnSyncMap = make(map[int]interface{})

type Accounts struct {
	Id          int    `gorm:"id,increment" json:"Id"`
	AccountName string `gorm:"unique" json:"AccountName"`
	//Password    string
	//DeviceId    string `gorm:"unique" json:"DeviceId"`
	//DeviceName  string
	//IsBanned    int `gorm:"index" json:"IsBanned"`
	//Os          int `gorm:"default:0" json:"Os"`
	//SdkUid      string
	//SdkToken    string
	//CreatedAt   string `gorm:"index" json:"CreatedAt"`
	//LastLogin   string `gorm:"default:2006-01-02 15:04:05"`
	//LastSvrId   int `gorm:"default:0"`
}

func testCoError(){
	go testCo()
	go testCo()

	time.Sleep(1*time.Second)

	SaveMap.Range(func(key, value interface{}) bool{
		fmt.Println("iterate:", key, value)
		return true
	})
}

func testCo(){
	for i := 0;i <= 1; i++ {
		SaveMap.Store(1, 1)
	}
}

// 玩家登录时写入，下线后再登录重新写入
func Load(){
	id := 1
	//fmt.Println(fmt.Sprintf("%+v", SyncMap))
	//fmt.Println(fmt.Sprintf("%+v", account))

	// 判断SyncMap是否有正在/正要同步的数据
	mapKey := "account_id:"+string(rune(1))
	if _, ok := SaveMap.Load(mapKey); ok {
		// todo：强制优先同步数据到db
	} else {
		if _, ok := LoadMap.Load(mapKey); ok {
			return
		} else {
			// load from db
			detail := loadAccounts(id)
			LoadMap.Store(mapKey, detail)

			if _, ok := AccountsMap.Load(id); ok {
				return
			} else {
				AccountsMap.Store(id, detail)
			}
		}
	}

	AccountsMap.Range(func(key, value interface{}) bool{
		fmt.Println("iterate:", key, value)
		return true
	})

	fmt.Println(fmt.Sprintf("%+v", SaveMap))
	fmt.Println(fmt.Sprintf("%+v", LoadMap))
	fmt.Println(fmt.Sprintf("%+v", AccountsMap))

	AccountsMap.Store(id, 2)
	AccountsMap.Range(func(key, value interface{}) bool{
		fmt.Println("iterate:", key, value)
		return true
	})
	fmt.Println(fmt.Sprintf("%+v", AccountsMap))

}

func loadAccounts(id int) Accounts{
	// 略：通过账号id到db获取到结构体对象
	account := Accounts{
		Id: 1,
		AccountName: "andy",
	}
	return account
}