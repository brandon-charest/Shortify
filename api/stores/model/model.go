type Storage interface {
	GetEntryByID(string) (*Entry, error)
	GetVisitors(string) ([]Visitor, error)
	DeleteEntry(string) error
	IncreaseVisitCounter(string) error
	CreateEntry(Entry, string, string) error
	GetUserEntries(string) (map[string]Entry, error)
	RegisterVisitor(string, string, Visitor) error
	Close() error
}

type Entry struct {
	CreatedOn             time.Time
	LastVisit, Expiration *time.Time `json:",omitempty"`
	VisitCount            int
	URL                   string
}

type Visitor struct {
	IP, Referer, UserAgent string
	Timestamp              time.Time
}
