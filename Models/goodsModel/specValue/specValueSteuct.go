package spec

func (l *ValueList) Response() ValueList {
	var list []Value

	for _, v := range *l {
		//var image string
		//if len(v.Image) <= 0 {
		//	image = v.Image
		//} else {
		//	image = conversion.FormattingSrc(v.Image)
		//}
		list = append(list, Value{
			ID:      v.ID,
			GoodsId: v.GoodsId,
			SpecId:  v.SpecId,
			Value:   v.Value,
			//Image:   image,
		})
	}
	return list

}
