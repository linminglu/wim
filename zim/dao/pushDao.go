package dao

type PushDao struct {
	Cmd   string `json:"cmd" db:"cmd"`     //action
	Fuid  string `json:"fuid" db:"fuid"`   //Sender Uid
	Fname string `json:"fname" db:"fname"` //Sender Uid
	To    string `json:"to" db:"to"`       //Receiver标签（u/*、g/*、t/*,b/*）
	Tuid  string `json:"tuid" db:"tuid"`   //tuid
	Stime int64  `json:"stime" db:"stime"` //Server发送时间
	seq   int64  `json:"seq"  db:"seq"`    //seq_svr
}

func NewPushDao() (p *PushDao) {
	p = &PushDao{}
	return
}

//cmd:push force online offline
func (p *PushDao) Assert(cmd string, sd *SendDao) {
	p.Cmd = cmd
	p.Fuid, p.Fname, p.Stime, p.seq, p.To, p.Tuid = sd.Fuid, sd.Fname, sd.Stime, sd.Seq, sd.To, sd.Tuid
}