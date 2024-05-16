package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 191,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	dao.SetDefault(db)

	//create(ctx, dao.Q)
	//find(ctx, dao.Q)
	//update(ctx, dao.Q)
	delete(ctx, dao.Q)
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

func find(ctx context.Context, query *dao.Query) {
	address := query.Address

	user, err := query.WithContext(ctx).User.
		Preload(query.User.Profile).
		Preload(query.User.Address.On(address.City.Eq("上海")).Select(address.ID, address.UserID, address.Province, address.City, address.County, address.Address)).
		Where(query.User.Name.Like("%s%")).
		First()
	if err != nil {
		log.Fatalln(err)
	}

	bytes, _ := json.Marshal(&user)
	fmt.Println(string(bytes))
}

func update(ctx context.Context, query *dao.Query) {
	info, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(1)).Updates(map[string]interface{}{
		"phone": "19999999999",
		"email": "summer@163.com",
	})
	if err != nil {
		fmt.Printf("update user fail, err:%v\n", err)
		return
	}

	fmt.Printf("RowsAffected:%v\n", info.RowsAffected)
}

func delete(ctx context.Context, query *dao.Query) {
	info, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(2)).Delete()
	if err != nil {
		fmt.Printf("delete user fail, err:%v\n", err)
		return
	}
	fmt.Printf("RowsAffected:%v\n", info.RowsAffected)
}
