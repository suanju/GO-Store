package service

func (ls ServicesLists) ConversionRsp() []string {
	var list []string
	for _, v := range ls {
		list = append(list, v.Service.Name)
	}
	return list

}
