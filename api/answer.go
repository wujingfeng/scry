package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scry/model"
)

// 关于面试文档中各问题的回答

// Question1
//  @Description: 问题 1：从一个空文件开始，你会如何创建一个 Go 程序的基本结构？
//  @return string
func Question1() string {
	return "合理的功能逻辑拆分.  以及关键逻辑的注释"
}

// Question2
//  @Description: 问题 2：使用代码举一个数组和切片声明的例子。
func Question2() {
	// 数组 固定长度
	var intArr [2]int64
	// 切片  自动扩容
	var intSlice []int64

	intArr[0] = 1
	intArr[1] = 2
	//intArr[2] = 3 // 直接报错

	intSlice[0] = 0
	intSlice[1] = 1
	intSlice[2] = 2
	intSlice[3] = 3
}

// Question3
//  @Description: 问题 3：使用代码举一个匿名结构的例子。
func Question3() {
	var w http.ResponseWriter
	var req *http.Request
	// 定义匿名结构体
	user := struct {
		Name string `json:"name"`
	}{}

	// 将 http 请求中的内容解析到匿名结构体重
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// Question4
//  @Description: 问题 4：举一个defer语句延迟函数调用的例子。
func Question4() {
	// 调用方负责捕获异常
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Errorf("here has some panic, and panic message is: %v", r))
		}
	}()

	doPanic()

	fmt.Println("test defer")
}

// doPanic
//  @Description: 发生异常的方法
func doPanic() {
	panic("something wrong")
}

// Question5
//  @Description: 问题 5：编写一个 Golang 程序，声明一个字符串变量，打印变量的地址，声明另一个 int 变量，以及指向它的指针。
func Question5() {
	s := "string"
	// 打印 s 内存地址
	fmt.Println(&s)

	i := 10
	// 打印 i 内存地址
	fmt.Println(&i)

	// 指针地址赋值
	ptrI := &i
	// 打印指针地址
	fmt.Println(ptrI)
	// 输出指针对应的值
	fmt.Println(*ptrI)
}

// Question6
//  @Description: 问题 6：创建一个定义命名类型和该类型的方法（接收器函数）的 Go 程序。
func Question6() {
	consumeDetail := model.ConsumeDetail{}
	// 指针对象方法的调用
	fmt.Println(consumeDetail.TableName())
	// 普通对象方法的调用
	fmt.Println(model.ConsumeDetail{}.DB())
}

// Question7
//  @Description: 问题 7：编写一个使用 goroutine 和 channel 的简单 Golang 程序。
func Question7() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			select {
			case s := <-ch1:
				fmt.Println("ch1", s)
			case s := <-ch2:
				fmt.Println("ch2", s)
			}
		}
	}()

	ch1 <- "ch1 test1"
	ch2 <- "ch2 test1"
	ch2 <- "ch2 test2"
	ch1 <- "ch1 test2"
}

// Question8
//  @Description: 二、编程题：以下测试题二选一
//1. 会员管理（您也可以选择以写一个区块链自动化合约方式来表达这个业务需求）
//某电商，计划在其公众号上实现会员管理。
//具体需求：
//（1）老会员可以邀请新会员加入，形成上下级关系。每一个会员可以有多个下级，但是只有一个上级。
//（2）每个会员，都有多个消费记录，记录其消费明细。
//（3）每个月底，统计每个会员及其下级会员的消费总额
//请完成下列任务：

//（1）设计这个会员管理功能所需的数据库表（PostgreSQL）
func Question2_1() {
	// =================================================================
	// ==========未使用过 PostgreSQL, 因此一下演示基于 mysql================
	// =================================================================

	// 用户基础信息表
	//CREATE TABLE `user` (
	//	`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户 id',
	//	`openid` varchar(32) NOT NULL DEFAULT '' COMMENT '微信授权 openid',
	//	`unionid` varchar(32) NOT NULL DEFAULT '' COMMENT '微信授权 unionid',
	//	`nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '用户昵称',
	//	`phone` varchar(11) NOT NULL DEFAULT '' COMMENT '用户电话',
	//	`code` varchar(32) NOT NULL DEFAULT '' COMMENT '用户内码',
	//	PRIMARY KEY (`id`),
	//	UNIQUE KEY `uidx_code` (`code`),
	//	UNIQUE KEY `uidx_openid` (`openid`),
	//	UNIQUE KEY `uidx_unionid` (`unionid`)
	//) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='用户基础信息表';

	// 邀请记录表
	//CREATE TABLE `invite_record` (
	//	`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '邀请记录 id',
	//	`user_id` int(11) NOT NULL DEFAULT 0 COMMENT '邀请人 id',
	//	`invitee_user_id` int(11) NOT NULL DEFAULT 0 COMMENT '被邀请人 id',
	//	`created_time` int(11) NOT NULL DEFAULT 0 COMMENT '邀请时间',
	//	PRIMARY KEY (`id`),
	//	KEY `idx_ user_id` (`user_id`),
	//	KEY `idx_invitee_user_id` (`invitee_user_id`)
	//) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='邀请记录表';

	// 用户消费记录明细表
	//CREATE TABLE `consume_detail` (
	//	`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '消费记录 id',
	//	`user_id` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '用户 id',
	//	`money` int(11) NOT NULL DEFAULT 0 COMMENT '消费金额(单位:分)',
	//	`created_time` int(11) NOT NULL DEFAULT 0 COMMENT '消费时间',
	//	PRIMARY KEY (`id`),
	//	KEY `user_id` (`user_id`)
	//) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='用户消费记录明细表';

}

// Question2_2
//  @Description: （2）用你所熟悉的go语言的一种ORM，描述所设计的数据表
func Question2_2() {
	fmt.Println(model.BaseModel{})
	fmt.Println(model.InviteRecord{})
	fmt.Println(model.ConsumeDetail{})
}

// Question2_3
//  @Description: （3）编写计算某个会员，在某个月消费总额：其自身消费和所有下级会员的消费的总和
func Question2_3() {
	Consume{}.CountConsume(1, 0, 0)
}
