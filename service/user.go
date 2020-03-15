package service

import (
	"GO_IM/model"
	"GO_IM/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

//注册函数
func (s *UserService) Register(
	mobile, //手机号
	plainpwd, //明文密码
	nickname, //昵称
	avatar, //头像
	sex string) (user model.User, err error) {
	//检测手机号是否存在
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("This mobile already register.")
	}
	//否则插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Sex = sex
	tmp.Nickname = nickname
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngin.InsertOne(&tmp)

	return tmp, err
}

//登录函数
func (s *UserService) Login(mobile, passwd string) (user model.User, err error) {
	//检测手机号是否存在
	tmp := model.User{}
	DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("This user don't exist!")
	}
	if !util.ValidatePasswd(passwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("Passwd error!")
	}
	//刷新Token,安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	tmp := model.User{}
	DbEngin.ID(userId).Get(&tmp)

	return tmp
}
