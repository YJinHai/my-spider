package base

type MapFloat64Sorter []MapFloat64Item

type MapFloat64Item struct {
	Key string

	Val float64
}

func NewMapSorter(m map[string]float64) MapFloat64Sorter {

	ms := make(MapFloat64Sorter, 0, len(m))

	for k, v := range m {

		ms = append(ms, MapFloat64Item{k, v})

	}

	return ms

}

func (ms MapFloat64Sorter) Len() int {

	return len(ms)

}

func (ms MapFloat64Sorter) Less(i, j int) bool {

	return ms[i].Val < ms[j].Val // 按值排序

	//return ms[i].Key < ms[j].Key // 按键排序

}

func (ms MapFloat64Sorter) Swap(i, j int) {

	ms[i], ms[j] = ms[j], ms[i]

}
