package util

type DefaultMap struct {
	Map          map[int]int
	DefaultValue int
}

func (m DefaultMap) Get(key int) int {
	if v, ok := m.Map[key]; ok {
		return v
	}
	return m.DefaultValue
}
