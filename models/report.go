package models

type Report struct {
	UserID        string `json:"user_id"`
	CallsSellers  int    `json:"calls_sellers"`
	CallsBuyers   int    `json:"calls_buyers"`
	IncomingCalls int    `json:"incoming_calls"`
	CrmEntries    int    `json:"crm_entries"`
	Statuses      int    `json:"statuses"`
	Banners       int    `json:"banners"`
	Stickers      int    `json:"stickers"`
	Date          string `json:"date"` // формат: YYYY-MM-DD
}
