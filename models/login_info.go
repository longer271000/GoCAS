package models

type LoginInfo struct {
	ID int64 "xorm:pk autoincr 'id' json:'id'"
	LoginName string "xorm:'varchar(32)' json:'loginName'"
	LoginPwd string "xorm:'varchar(32)' json:'loginPwd'"
	Status int64 "xorm:'default 0' json:'status'"
	CreateTime int64 "xorm:'numeric(0,50)' json:'create_time'"
}
