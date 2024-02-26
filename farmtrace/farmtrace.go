package farmtrace

import "time"

type FarmPractice string

const (
	Standard FarmPractice = "standard"
	Organic  FarmPractice = "organic"
)

type DocFormat string

const (
	PDF  DocFormat = "pdf"
	TXT  DocFormat = "txt"
	MD   DocFormat = "md"
	PNG  DocFormat = "png"
	JPG  DocFormat = "jpg"
	SVG  DocFormat = "svg"
	JSON DocFormat = "json"
	XML  DocFormat = "adoc"
)

type EntityType string

const (
	ETfarm    = "Farm"
	ETharvest = "Harvest"
)

type Farm struct {
	Pluscode string            `json:"pluscode"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Practice FarmPractice      `json:"practice"`
}

type URI string
type Hash []byte

type DocTrace struct {
	DocHash  Hash              `json:"dochash"`
	Schema   URI               `json:"uri"`
	W3CDID   URI               `json:"w3cdid"`
	Metadata map[string]string `json:"metadata,omitempty"`
	IsPublic bool              `json:"ispublic"`
	Format   DocFormat         `json:"format"`
}

type Harvest struct {
	FarmRef     Hash      `json:"farmref"`
	HarvestDate time.Time `json:"tmharvest"`
	Produce     URI       `json:"produce"`
	WeightGm    int64     `json:"gmwt"`
}

type DocLink struct {
	DocHash    Hash              `json:"dochash"`
	EntityHash Hash              `json:"entity"`
	Entity     EntityType        `json:"type"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}
