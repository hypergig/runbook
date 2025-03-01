package schema

type Step struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	// possible modules
	Exec  *Exec  `json:"exec,omitempty"`
	Steps *Steps `json:"steps,omitempty"`
}

// modules
type Steps []*Step

type Exec struct {
	Cmd  string   `json:"cmd,omitempty"`
	Args []string `json:"args,omitempty"`
}
