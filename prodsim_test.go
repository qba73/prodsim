package prodsim_test

import (
	"testing"
	"time"

	"github.com/qba73/prodsim"
)

func TestCreateDefaultProductionLine(t *testing.T) {
	t.Parallel()

	pl := prodsim.NewProductionLine()
	if pl.Verbose {
		t.Errorf("want false, got true")
	}
}

func TestAddStageToProductionLine(t *testing.T) {
	t.Parallel()

	pl := prodsim.NewProductionLine()
	pl.AddStage("stage1", prodsim.NewDummyStage(5*time.Second, 500*time.Millisecond))
	got := pl.ListStages()

	if len(got) != 1 {
		t.Fatalf("want 1 stage, got %d", len(got))
	}

	want := "stage1"
	gotName := got[0].Name
	if want != got[0].Name {
		t.Errorf("want %s, got %s", want, gotName)
	}
}

func TestListStagesInNewEmptyProductionLine(t *testing.T) {
	t.Parallel()

	pl := prodsim.NewProductionLine()
	pl.AddStage("stage1", prodsim.NewDummyStage(5*time.Second, 500*time.Millisecond))
	got := pl.ListStages()

	if len(got) != 1 {
		t.Fatalf("want 1 stage, got %d", len(got))
	}

	want := "stage1"
	gotStageName := got[0].Name
	if want != gotStageName {
		t.Errorf("want %s, got %s", want, gotStageName)
	}
}

func TestAddMultipleStagesToProductionLine(t *testing.T) {
	t.Parallel()

	pl := prodsim.NewProductionLine()

	pl.AddStage("stage1", prodsim.NewDummyStage(5*time.Second, 500*time.Millisecond))
	pl.AddStage("stage2", prodsim.NewDummyStage(5*time.Second, 500*time.Millisecond))
	pl.AddStage("stage3", prodsim.NewDummyStage(5*time.Second, 500*time.Millisecond))

	got := pl.ListStages()

	if len(got) != 3 {
		t.Fatalf("want 3 stages, got %d", len(got))
	}

	// Check that stages remain in order
	want := "stage1"
	gotStageName := got[0].Name
	if want != gotStageName {
		t.Errorf("want %s, got %s", want, gotStageName)
	}

	want = "stage2"
	gotStageName = got[1].Name
	if want != gotStageName {
		t.Errorf("want %s, got %s", want, gotStageName)
	}

	want = "stage3"
	gotStageName = got[2].Name
	if want != gotStageName {
		t.Errorf("want %s, got %s", want, gotStageName)
	}
}
