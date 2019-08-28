package gcloud

type Gcloud struct {
	Version     string                   `json:"CLI_VERSION"`
	FlagLists   []*Flag                  `json:"SERIALIZED_FLAG_LIST"`
	VERSION     string                   `json:"VERSION"`
	Capsule     string                   `json:"capsule,omitempty"`
	Commands    map[string]*Command      `json:"commands,omitempty"`
	Constraints *Constraints             `json:"constraints,omitempty"`
	Flags       *Flag                    `json:"flags,omitempty"`
	IsGlobal    bool                     `json:"is_global,omitempty"`
	IsGroup     bool                     `json:"is_group,omitempty"`
	IsHidden    bool                     `json:"is_hidden,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Path        []string                 `json:"path,omitempty"`
	Positionals []map[string]interface{} `json:"positionals,omitempty"`
	Release     string                   `json:"release,omitempty"`
	Sections    *Sections                `json:"sections,omitempty"`
}

type Flag struct {
	Attr         *FlagAttr     `json:"attr,omitempty"`
	Category     string        `json:"category,omitempty"`
	Choices      []interface{} `json:"choices,omitempty"`
	Completer    string        `json:"completer,omitempty"`
	Default      interface{}   `json:"default,omitempty"`
	Description  string        `json:"description,omitempty"`
	IsGlobal     bool          `json:"is_global,omitempty"`
	IsGroup      bool          `json:"is_group,omitempty"`
	IsHidden     bool          `json:"is_hidden,omitempty"`
	IsMutex      bool          `json:"is_mutex,omitempty"`
	IsPositional bool          `json:"is_positional,omitempty"`
	IsRequired   bool          `json:"is_required,omitempty"`
	Name         string        `json:"name,omitempty"`
	Nargs        string        `json:"nargs,omitempty"`
	Type         string        `json:"type,omitempty"`
	Value        string        `json:"value,omitempty"`
}

type FlagAttr struct {
	InvertedSynopsis bool              `json:"inverted_synopsis,omitempty"`
	Property         *FlagAttrProperty `json:"property,omitempty"`
}

type FlagAttrProperty struct {
	Name string `json:"name,omitempty"`
}

type Command struct {
	Capsule     string              `json:"capsule,omitempty"`
	Commands    map[string]*Command `json:"commands,omitempty"`
	Constraints *Constraints        `json:"constraints,omitempty"`
	Flags       *Flag               `json:"flags,omitempty"`
	IsGlobal    bool                `json:"is_global,omitempty"`
	IsGroup     bool                `json:"is_group,omitempty"`
	IsHidden    bool                `json:"is_hidden,omitempty"`
	Name        string              `json:"name,omitempty"`
	Path        []string            `json:"path,omitempty"`
	Positionals []interface{}       `json:"positionals,omitempty"`
	Release     string              `json:"release,omitempty"`
	Sections    *CommandSections    `json:"sections,omitempty"`
}

type CommandSections struct {
	Description string `json:"DESCRIPTION,omitempty"`
	Notes       string `json:"NOTES,omitempty"`
}

type Constraints struct {
	Arguments    []interface{}          `json:"arguments,omitempty"`
	Attr         map[string]interface{} `json:"attr,omitempty"`
	Description  string                 `json:"description,omitempty"`
	IsGroup      bool                   `json:"is_group,omitempty"`
	IsHidden     bool                   `json:"is_hidden,omitempty"`
	IsMutex      bool                   `json:"is_mutex,omitempty"`
	IsPositional bool                   `json:"is_positional,omitempty"`
	IsRequired   bool                   `json:"is_required,omitempty"`
}

type Sections struct {
	Description string `json:"DESCRIPTION,omitempty"`
}
