package common

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp() *Resp {
	return &Resp{}
}

func (r *Resp) SetCode(code int) *Resp {
	r.Code = code
	return r
}

func (r *Resp) SetMsg(msg string) *Resp {
	r.Msg = msg
	return r
}

func (r *Resp) SetData(data interface{}) *Resp {
	r.Data = data
	return r
}
