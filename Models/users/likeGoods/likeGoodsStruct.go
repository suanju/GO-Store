package likeGoods

func (lg *LikeGoods) IsLike() bool {
	if lg.ID > 0 {
		return true
	}
	return false
}
