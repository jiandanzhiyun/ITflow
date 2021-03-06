package response

import (
	"encoding/json"
	"fmt"
	"itflow/bug/model"
)

type Onefd struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	ModDate  int64  `json:"date"`
	IsFile   bool   `json:"isfile"`
	IsOwner  bool   `json:"isowner"`
	HasWrite bool   `json:"haswrite"`
	Ru       bool   `json:"readuser"`
	Rname    string `json:"readname"`
	Wu       bool   `json:"writeuser"`
	Wname    string `json:"writename"`
}

type Response struct {
	Id         int64                     `json:"id,omitempty" type:"int" need:"否" default:"" information:"返回的id，某些端口需要"`
	AffectId   int64                     `json:"affectid,omitempty" type:"int" need:"否" default:"" information:"返回的行数，某些端口需要"`
	Code       int                       `json:"code" type:"int" need:"是" default:"" information:"返回的状态码， 0为成功，非0失败"`
	Msg        string                    `json:"message,omitempty" type:"string" need:"否" default:"" information:"错误信息， 状态码非肯定有"`
	Path       *Onefd                    `json:"path,omitempty" type:"object" need:"否" default:"" information:"返回路径，共享文件接口用到"`
	Filename   string                    `json:"filename,omitempty" type:"string" need:"否" default:"" information:"文件名，共享文件接口用到"`
	Data       []byte                    `json:"data,omitempty" type:"bytes" need:"否" default:"" information:"返回数据，某些接口用到"`
	UpdateTime int64                     `json:"updatetime,omitempty" type:"string" need:"否" default:"" information:"更新时间， 共享文件接口用到"`
	Size       int64                     `json:"size,omitempty" type:"int" need:"否" default:"" information:"文件大小，共享文件接口用到"`
	HeaderList []*model.Table_headerlist `json:"headerlist,omitempty" type:"array" need:"否" default:"" information:"用到的时候再标识， 一下想不起来"`
}

func (es *Response) ErrorE(err error) []byte {
	es.Msg = err.Error()
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Errorf(format string, args ...interface{}) []byte {
	es.Msg = fmt.Sprintf(format, args...)
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Error(msg string) []byte {
	es.Msg = msg
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) Success() []byte {
	es.Msg = "success"
	send, _ := json.Marshal(es)
	return send
}

func (es *Response) ErrorNoPermission() []byte {
	es.Msg = "没有权限"
	send, _ := json.Marshal(es)
	return send
}
