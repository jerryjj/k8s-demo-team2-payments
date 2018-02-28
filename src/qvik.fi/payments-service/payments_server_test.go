package main

import (
	"context"
	"testing"

	payments "qvik.fi/payments"
)

func TestPaymets(t *testing.T) {
	s := paymentsServer{}

	tests := []struct {
		wantStatus  payments.Status
		wantMessage string
	}{
		{
			payments.Status_OK,
			"",
		},
	}

	for _, tt := range tests {
		req := &payments.GetPSPStatusRequest{}
		resp, err := s.GetPSPStatus(context.Background(), req)
		if err != nil {
			t.Errorf("GetPSPStatus go unexpected error: %v", err)
		}
		if resp.Status != tt.wantStatus {
			t.Errorf("Expected different Status. Wanted=%s , Got=%s", tt.wantStatus, resp.Status)
		}
		if resp.StatusMessage != tt.wantMessage {
			t.Errorf("Expected different StatusMessage. Wanted=%s , Got=%s", tt.wantMessage, resp.StatusMessage)
		}
	}
}
