package db

import (
	"testing"
)

func TestSaveCleanedFinalizedSlot(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)

	slot := uint64(100)
	if err := db.SaveCleanedFinalizedSlot(slot); err != nil {
		t.Errorf("failed to save cleaned finalized slot %v", err)
	}
}

func TestGetCleanedFinalizedSlot_NotFound(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)

	var slot uint64
	slot, err := db.GetCleanedFinalizedSlot()
	if err != nil {
		t.Error("got DB error when reading cleaned finalized slot")
	}
	if slot != 0 {
		t.Error("expect 0 if DB doesn't have last cleaned finalized slot")
	}
}

func TestGetCleanedFinalizedSlot(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)

	slot := uint64(100)
	if err := db.SaveCleanedFinalizedSlot(slot); err != nil {
		t.Fatalf("failed to save cleaned finalized slot %v", err)
	}

	readSlot, err := db.GetCleanedFinalizedSlot()
	if err != nil {
		t.Fatalf("failed to read cleaned finalized slot from DB %v", err)
	}
	if readSlot != slot {
		t.Error("got wrong result when reading cleaned finalized slot from DB")
	}
}
