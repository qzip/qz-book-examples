package farmtrace

import "time"

type FarmPractice string

const (
	Standard FarmPractice = "chemical"
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
	ETfarm    = "farm"
	ETharvest = "harvest"
	ETParty   = "party"
)

type Farm struct {
	Pluscode string            `json:"pluscode"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Practice FarmPractice      `json:"practice"`
}

type URI string

// W3Cdid did:qz:<namespace>#<Hash base32>
// did:qz:example.com#AFB123456ZX
type W3Cdid string
type Hash []byte

type DocTrace struct {
	ID       W3Cdid            `json:"id"`
	DocHash  Hash              `json:"dochash"`
	Schema   URI               `json:"uri"`
	Metadata map[string]string `json:"metadata,omitempty"`
	IsPublic bool              `json:"ispublic"`
	Format   DocFormat         `json:"format"`
}

type Harvest struct {
	ID          W3Cdid    `json:"id"`
	FarmRef     Hash      `json:"farmref"`
	HarvestDate time.Time `json:"tmharvest"`
	Produce     URI       `json:"produce"`
	WeightGm    int64     `json:"gmwt"`
	HarvestedBy W3Cdid    `json:"harvestedby"`
}

type DocLink struct {
	DocID    W3Cdid            `json:"docId"`
	EntityID W3Cdid            `json:"entityId"`
	Entity   EntityType        `json:"type"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Transfer struct {
	Seller     W3Cdid            `json:"seller"`
	Buyer      W3Cdid            `json:"buyer"`
	Harvest    W3Cdid            `json:"harvest"`
	Invoice    W3Cdid            `json:"invoice"`
	WeightGm   int64             `json:"gmwt"`
	TmTransfer time.Time         `json:"tmTransfer"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}
