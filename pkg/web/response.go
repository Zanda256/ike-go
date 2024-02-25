package web

type Response struct {
	StatusCode int
	Headers    []byte
	Body       []byte
}

// =============================================================================
