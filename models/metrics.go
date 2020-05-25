package models

type PerformanceMetrics struct {
	FinishDocumentTime float32 `json:"finishDocumentTime"`
	TTFB               float32 `json:"ttfb"`
	FCP                float32 `json:"fcp"`
}

type Metric struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
	Id    string  `json:"id"`
}
