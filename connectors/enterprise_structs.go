package connectors

import "time"

// EnterpriseAPIResponseModel0 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel0 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel1 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel1 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel2 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel2 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel3 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel3 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel4 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel4 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel5 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel5 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel6 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel6 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel7 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel7 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel8 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel8 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel9 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel9 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel10 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel10 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel11 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel11 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel12 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel12 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel13 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel13 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel14 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel14 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel15 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel15 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel16 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel16 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel17 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel17 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel18 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel18 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel19 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel19 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel20 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel20 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel21 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel21 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel22 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel22 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel23 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel23 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel24 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel24 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel25 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel25 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel26 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel26 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel27 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel27 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel28 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel28 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel29 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel29 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel30 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel30 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel31 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel31 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel32 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel32 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel33 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel33 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel34 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel34 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel35 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel35 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel36 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel36 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel37 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel37 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel38 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel38 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel39 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel39 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel40 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel40 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel41 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel41 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel42 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel42 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel43 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel43 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel44 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel44 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel45 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel45 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel46 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel46 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel47 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel47 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel48 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel48 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel49 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel49 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel50 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel50 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel51 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel51 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel52 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel52 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel53 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel53 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel54 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel54 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel55 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel55 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel56 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel56 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel57 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel57 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel58 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel58 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel59 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel59 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel60 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel60 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel61 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel61 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel62 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel62 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel63 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel63 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel64 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel64 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel65 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel65 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel66 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel66 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel67 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel67 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel68 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel68 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel69 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel69 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel70 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel70 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel71 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel71 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel72 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel72 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel73 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel73 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel74 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel74 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel75 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel75 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel76 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel76 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel77 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel77 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel78 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel78 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel79 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel79 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel80 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel80 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel81 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel81 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel82 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel82 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel83 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel83 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel84 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel84 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel85 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel85 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel86 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel86 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel87 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel87 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel88 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel88 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel89 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel89 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel90 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel90 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel91 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel91 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel92 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel92 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel93 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel93 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel94 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel94 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel95 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel95 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel96 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel96 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel97 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel97 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel98 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel98 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

// EnterpriseAPIResponseModel99 represents a complex multi-layered JSON payload
type EnterpriseAPIResponseModel99 struct {
	ID            string    `json:"id"`
	Object        string    `json:"object"`
	Created       int64     `json:"created"`
	Model         string    `json:"model"`
	Choices       []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     struct {
			Tokens        []string  `json:"tokens"`
			TokenLogprobs []float64 `json:"token_logprobs"`
			TopLogprobs   []map[string]float64 `json:"top_logprobs"`
			TextOffset    []int     `json:"text_offset"`
		} `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Metadata map[string]interface{} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	CorrelationID string `json:"correlation_id"`
}

