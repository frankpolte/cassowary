package client

import (
	"net/http"

	"github.com/schollz/progressbar"
)

// Cassowary is the main struct with bootstraps the load test
type Cassowary struct {
	FileMode              bool
	IsTLS                 bool
	InputFile             string
	BaseURL               string
	ConcurrencyLevel      int
	Requests              int
	ExportMetrics         bool
	ExportMetricsFile     string
	PromExport            bool
	PromURL               string
	RequestHeader         []string
	DisableTerminalOutput bool
	DisableKeepAlive      bool
	Client                *http.Client
	Bar                   *progressbar.ProgressBar
}

// ResultMetrics are the aggregated metrics after the load test
type ResultMetrics struct {
	BaseURL           string                `json:"base_url"`
	TotalRequests     int                   `json:"total_requests"`
	FailedRequests    int                   `json:"failed_requests"`
	RequestsPerSecond float64               `json:"requests_per_second"`
	DNSMedian         float64               `json:"dns_median"`
	TCPStats          tcpStats              `json:"tcp_connect"`
	ProcessingStats   serverProcessingStats `json:"server_processing"`
	ContentStats      contentTransfer       `json:"content_transfer"`
}

type tcpStats struct {
	TCPMean   float64 `json:"mean"`
	TCPMedian float64 `json:"median"`
	TCP95p    float64 `json:"95th_percentile"`
}

type serverProcessingStats struct {
	ServerProcessingMean   float64 `json:"mean"`
	ServerProcessingMedian float64 `json:"median"`
	ServerProcessing95p    float64 `json:"95th_percentile"`
}

type contentTransfer struct {
	ContentTransferMean   float64 `json:"mean"`
	ContentTransferMedian float64 `json:"median"`
	ContentTransfer95p    float64 `json:"95th_percentile"`
}
