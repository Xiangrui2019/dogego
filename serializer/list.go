package serializer

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

func BuildDataList(items interface{}, total uint) *DataList {
	return &DataList{
		Items: items,
		Total: total,
	}
}
