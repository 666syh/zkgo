package service

import(
	"context"
	"zkgo/utils"
	// "fmt"
)

type Service interface{
	RefreshKey(ctx context.Context, target interface{}) error
}

type zkservice struct{
	retChan chan interface{}
}

func NewZkService(retChan chan interface{}) Service{
	return &zkservice{retChan}
}

func (zs zkservice) RefreshKey(ctx context.Context, target interface{}) error{
	params := utils.ToMapString(target)
	zs.retChan <- params
	return nil
}