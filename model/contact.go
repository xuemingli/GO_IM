package model

import "time"

const (
	//用户加用户好友，则Cate为0x01
	CONCAT_CATE_USER = 0x01
	//用户加群，则Cate为0x02
	CONCAT_CATE_COMUNITY = 0x02
)

//好友和群都存在这个表里面
//可根据具体业务做拆分
type Contact struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//自己的Id
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 记录是谁的
	//对方的Id
	Dstobj int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"` // 对端信息
	//
	Cate int    `xorm:"int(11)" form:"cate" json:"cate"`      // 什么类型
	Memo string `xorm:"varchar(120)" form:"memo" json:"memo"` // 备注
	//
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
