package ch32_diy_obj_pool

import (
	"errors"
	"time"
)

type ReuseableObj struct {

}

type ObjPool struct {
	bufChan chan *ReuseableObj
}

func NewPool(numOfObj int) *ObjPool {
	pool := ObjPool{}
	pool.bufChan = make(chan * ReuseableObj, numOfObj)
	for i:=0; i<numOfObj; i++ {
		pool.bufChan <- &ReuseableObj{}
	}
	return &pool
}

func (p *ObjPool)GetObj(timeout time.Duration)  (*ReuseableObj, error){
	select {
	case ret := <-p.bufChan:
			return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")

	}
}

func (p *ObjPool) ReleaseObj(obj *ReuseableObj) error{
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}