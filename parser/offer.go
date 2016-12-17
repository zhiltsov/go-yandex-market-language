package parser

import xmlpath "gopkg.in/xmlpath.v2"

const (
	OffertTypeDefault     OfferType = ""
	OffertTypeVendorModel OfferType = "vendor.model"
	OffertTypeMedicine    OfferType = "medicine"
	OffertTypeBooks       OfferType = "books"
	OffertTypeAudioBooks  OfferType = "audiobooks"
	OffertTypeArtistTitle OfferType = "artist.title"
	OffertTypeEventTicket OfferType = "event-ticket"
	OffertTypeTour        OfferType = "tour"
)

type offerInterator func() *Offer
type OfferType string

type Offer struct {
	ID        string
	Type      OfferType
	Available bool
	Vendor    string
	Model     string
	URL       string
	Name      string
}

func (o OfferType) String() string {
	return string(o)
}

func (o *Offer) LoadFromNode(node *xmlpath.Node) {
	if val, ok := xmlpath.MustCompile("@id").String(node); ok {
		o.ID = val
	}

	if val, ok := xmlpath.MustCompile("@type").String(node); ok {
		o.Type = OfferType(val)
	}

	if val, ok := xmlpath.MustCompile("@available").String(node); ok {
		o.Available = (val == "true")
	} else {
		o.Available = true
	}

	if val, ok := xmlpath.MustCompile("vendor").String(node); ok {
		o.Vendor = val
	}

	if val, ok := xmlpath.MustCompile("model").String(node); ok {
		o.Model = val
	}

	if val, ok := xmlpath.MustCompile("url").String(node); ok {
		o.URL = val
	}

	if val, ok := xmlpath.MustCompile("name").String(node); ok {
		o.Name = val
	}
}