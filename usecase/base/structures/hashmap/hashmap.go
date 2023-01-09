package hashmap

type Hashmap struct {
	Entries map[string]interface{} `json:"entries"`
}

func NewHashmap() *Hashmap {
	return &Hashmap{
		Entries: make(map[string]interface{}),
	}
}

func (h *Hashmap) Set(key string, value interface{}) {
	h.Entries[key] = value
}

func (h *Hashmap) Get(key string) interface{} {
	return h.Entries[key]
}

func (h *Hashmap) Delete(key string) {
	delete(h.Entries, key)
}
func (h *Hashmap) ClearMap() {
	h.Entries = make(map[string]interface{})
}
func (h *Hashmap) ReturnCompleteMap() interface{} {
	return h.Entries
}
