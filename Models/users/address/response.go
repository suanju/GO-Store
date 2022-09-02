package address

//FormattedOutput 格式化输出
func (rl RegionList) FormattedOutput() OutputRegionList {
	data := make(OutputRegionList, len(rl))
	for k, v := range rl {
		data[k].CityId = v.ID
		data[k].Name = v.Name
	}
	return data
}
