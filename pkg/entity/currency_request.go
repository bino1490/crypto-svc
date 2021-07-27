package entity

type CurrencyRequest struct {
	ID       string `json:"id,omitempty"`
	FullName string `json:"fullname,omitempty"`
}

type SymbolRequest struct {
	ID           string `json:"id,omitempty"`
	BaseCurrency string `json:"baseCurrency,omitempty"`
	FeeCurrency  string `json:"feeCurrency,omitempty"`
}

type FinalRequest struct {
	Symbol      string `json:"symbol,omitempty"`
	ID          string `json:"id,omitempty"`
	FullName    string `json:"fullname,omitempty"`
	Ask         string `json:"ask,omitempty"`
	Bid         string `json:"bid,omitempty"`
	Last        string `json:"last,omitempty"`
	Open        string `json:"open,omitempty"`
	Low         string `json:"low,omitempty"`
	High        string `json:"high,omitempty"`
	FeeCurrency string `json:"feeCurrency,omitempty"`
}

type ResponseData struct {
	Currencies []FinalRequest `json:"currencies,omitempty"`
}
