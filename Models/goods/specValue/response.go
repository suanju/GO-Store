package spec

func (l *ValueList) Response() ValueList {
	var list []Value

	for _, v := range *l {

		list = append(list, Value{
			ID:      v.ID,
			GoodsId: v.GoodsId,
			SpecId:  v.SpecId,
			Value:   v.Value,
		})
	}
	return list

}
