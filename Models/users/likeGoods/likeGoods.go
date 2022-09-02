package likeGoods

import (
	"GO-Store/Databases/Mysql"
	"time"
)

//LikeGoods 用户喜欢商品表
type LikeGoods struct {
	ID        uint      `json:"id" gorm:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UserID    uint      `json:"user_id" gorm:"user_id"`
	GoodsId   uint      `json:"goods_id" gorm:"goods_id"`
}

func (LikeGoods) TableName() string {
	return "ml_users_like_goods"
}

func (lg *LikeGoods) Like(uid uint, goodId uint) bool {
	lg.UserID = uid
	lg.GoodsId = goodId
	//判断是否已经点赞
	err := Mysql.Db.Where(&lg).Find(&lg).Error
	if err != nil {
		return false
	}
	if lg.ID > 0 {
		Mysql.Db.Where(&lg).Delete(&lg)
		return true
	}
	err = Mysql.Db.Create(&lg).Error
	if err != nil {
		return false
	}
	return true

}

func (lg *LikeGoods) IsLike() bool {
	if lg.ID > 0 {
		return true
	}
	return false
}
