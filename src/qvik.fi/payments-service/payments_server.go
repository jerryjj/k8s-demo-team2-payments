package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	payments "qvik.fi/payments"
)

// Payments Service server type
type paymentsServer struct{}

func (s *paymentsServer) GetPSPStatus(ctx context.Context,
	r *payments.GetPSPStatusRequest) (*payments.GetPSPStatusResponse, error) {
	log.Debugf("GetPSPStatus()")

	wantedStatus := 200
	url := fmt.Sprintf("http://httpstat.us/%d", wantedStatus)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Failed to create request: %v", err)
		return &payments.GetPSPStatusResponse{
			Status:        payments.Status_ERROR,
			StatusMessage: fmt.Sprintf("Failed to create request: %v", err),
		}, nil
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Failed to execute query: %v", err)
		return &payments.GetPSPStatusResponse{
			Status:        payments.Status_ERROR,
			StatusMessage: fmt.Sprintf("Failed to execute request: %v", err),
		}, nil
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	if resp.StatusCode == wantedStatus {
		log.Debugf("GetPSPStatus successful")
		return &payments.GetPSPStatusResponse{
			Status:        payments.Status_OK,
			StatusMessage: resp.Status,
		}, nil
	}

	return &payments.GetPSPStatusResponse{
		Status: payments.Status_ERROR,
		StatusMessage: fmt.Sprintf("received invalid status. wanted=%v , received=%v",
			resp.StatusCode, wantedStatus),
	}, nil
}
