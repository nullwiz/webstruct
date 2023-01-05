package entity

type Operation struct {
	SessionID string `json:"sessionID"`
	ID        string `json:"ID"`
	Entity    string `json:"entity"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type Session struct {
	ID         string       `json:"ID"`
	Operations []*Operation `json:"operations"`
	Timestamp  int64        `json:"timestamp"`
	// map of entities results
	// key is the entity name
	// value is the entity result
	Results map[string]interface{} `json:"results"`
}

// Only used for local
func (s *Session) AddOperation(o *Operation) {
	s.Operations = append(s.Operations, o)
}

func (s *Session) ClearOperations() {
	s.Operations = []*Operation{}
}

func (s *Session) SetResult(entity string, result interface{}) {
	s.Results[entity] = result
}

func (s *Session) NullResult(entity string) {
	s.Results[entity] = nil
}
