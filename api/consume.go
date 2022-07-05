package api

import (
	"fmt"
	"scry/model"
)

type Consume struct {
}

// CountConsume
//  @Description: 统计一段时间内 某个会员及其名下所有用户的总消费额
//  @receiver c
//  @param userid	会员 id
//  @param startTime	开始时间
//  @param endTime 结束时间
//  @return int64
//  @author	wujingfeng 2022-07-05 16:00:53
func (c Consume) CountConsume(userid int64, startTime int64, endTime int64) int64 {

	// 先获取名下所有的会员 id
	var allUserIds []int64
	allUserIds, err := model.InviteRecord{}.GetAllChildByUserId(userid, allUserIds)
	if err != nil {
		fmt.Println(err)
	}

	// 统计所有会员的消费金额
	var money int64
	money, err = model.ConsumeDetail{}.CountByUserIds(allUserIds, startTime, endTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(money)
	return money
}
