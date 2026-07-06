package calculator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/shared"
	"testing"
)

func setup() {
	shared.InitValidator()
}

func TestHandlerCalculateSuccess(t *testing.T) {
	setup()
	h := newHandler(&service{})

	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{"addition", "2+3", 5},
		{"precedence", "2+3x4", 14},
		{"parentheses", "(2+3)x4", 20},
		{"power", "2^3", 8},
		{"sqrt", "s25+5", 10},
		{"percentage", "50%+1", 1.5},
		{"complex", "(-2+5)^2x(10-3)", 63},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(calculationRequest{
				Expression: tt.expression,
			})

			req := httptest.NewRequest(
				http.MethodPost,
				"/calculate",
				bytes.NewBuffer(body),
			)

			rec := httptest.NewRecorder()

			h.calculate(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected status %d got %d", http.StatusOK, rec.Code)
			}

			var resp calculationResponse
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Fatalf("decode response: %v", err)
			}

			if !shared.AlmostEqual(resp.Result, tt.expected) {
				t.Fatalf(
					"expr=%q expected=%v got=%v",
					tt.expression,
					tt.expected,
					resp.Result,
				)
			}
		})
	}
}

func TestHandlerCalculateInvalidExpression(t *testing.T) {
	setup()
	h := newHandler(&service{})

	tests := []string{
		"2+",
		"(",
		"2++3",
		"abc",
	}

	for _, expr := range tests {
		t.Run(expr, func(t *testing.T) {
			body, _ := json.Marshal(calculationRequest{
				Expression: expr,
			})

			req := httptest.NewRequest(
				http.MethodPost,
				"/calculate",
				bytes.NewBuffer(body),
			)

			rec := httptest.NewRecorder()

			h.calculate(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf(
					"expected status %d got %d",
					http.StatusBadRequest,
					rec.Code,
				)
			}
		})
	}
}

func TestHandlerCalculateValidationError(t *testing.T) {
	setup()
	h := newHandler(&service{})

	body := []byte(`{"expression":""}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	h.calculate(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func TestHandlerCalculateInvalidJSON(t *testing.T) {
	setup()
	h := newHandler(&service{})

	req := httptest.NewRequest(
		http.MethodPost,
		"/calculate",
		bytes.NewBufferString(`{"expression":`),
	)

	rec := httptest.NewRecorder()

	h.calculate(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
