package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sql/dao"
	"sql/model"
	"time"
)

const dsn = "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	var (
		err error
		db  *gorm.DB
		ctx = context.Background()
	)

	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	dao.SetDefault(db)

	create(ctx, dao.Q)
}

func create(ctx context.Context, query *dao.Query) {
	t := time.Date(2000, 2, 2, 0, 0, 0, 0, time.Local)
	users := []*model.User{
		{
			Phone:    "13512341234",
			Email:    "summer@gmail.com",
			Name:     "summer",
			Nickname: "夏天",
			Profile: model.UserProfile{
				Gender:   1,
				Level:    1,
				Birthday: &t,
			},
			Address: []model.Address{
				{Province: "", City: "上海", Address: "汤臣一品"},
				{Province: "", City: "北京", County: "朝阳区", Address: "朝阳区"},
			},
		},
		{
			Phone:    "13432584785",
			Email:    "13432584785@qq.com",
			Name:     "虾仁",
			Nickname: "虾仁猪心",
			Profile: model.UserProfile{
				Gender:    2,
				Level:     2,
				Signature: "大王叫我来巡山",
			},
			Address: []model.Address{
				{Province: "四川省", City: "成都市", County: "高新区", Address: "金融城"},
			},
		},
	}

	err := query.WithContext(ctx).User.Create(users...)
	if err != nil {
		log.Fatalln(err)
	}
}
