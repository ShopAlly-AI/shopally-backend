package repository

import (
	"testing"
	"github.com/shopally-ai/pkg/domain"
)

func TestMockAlertRepository(t *testing.T) {
	repo := NewMockAlertRepository()

	sampleAlert := &domain.Alert{
		UserID:      "user-123",
		ProductID:   "product-abc",
		TargetPrice: 500.00,
		IsActive:    true,
	}

	var createdAlertID string 

	t.Run("CreateAlert_Success", func(t *testing.T) {
		err := repo.CreateAlert(sampleAlert)
		if err != nil {
			t.Fatalf("CreateAlert failed with error: %v", err)
		}
		
		if sampleAlert.ID == "" {
			t.Fatal("CreateAlert did not assign an ID to the alert")
		}
		createdAlertID = sampleAlert.ID
	})
	
	t.Run("GetAlert_Success", func(t *testing.T) {
		retrievedAlert, err := repo.GetAlert(createdAlertID)
		if err != nil {
			t.Fatalf("GetAlert failed with error: %v", err)
		}
		
		if retrievedAlert.ID != createdAlertID {
			t.Errorf("Retrieved alert ID mismatch: got %s, want %s", retrievedAlert.ID, createdAlertID)
		}
		
		if retrievedAlert.UserID != sampleAlert.UserID {
			t.Errorf("Retrieved alert UserID mismatch: got %s, want %s", retrievedAlert.UserID, sampleAlert.UserID)
		}
	})

	t.Run("GetAlert_NotFound", func(t *testing.T) {
		_, err := repo.GetAlert("non-existent-id")
		if err == nil {
			t.Fatal("GetAlert for a non-existent ID did not return an error")
		}
	})

	t.Run("DeleteAlert_Success", func(t *testing.T) {
		err := repo.DeleteAlert(createdAlertID)
		if err != nil {
			t.Fatalf("DeleteAlert failed with error: %v", err)
		}

		_, err = repo.GetAlert(createdAlertID)
		if err == nil {
			t.Fatal("Alert was not deleted as expected")
		}
	})

	t.Run("DeleteAlert_NotFound", func(t *testing.T) {
		err := repo.DeleteAlert("non-existent-id")
		if err == nil {
			t.Fatal("DeleteAlert for a non-existent ID did not return an error")
		}
	})
}
