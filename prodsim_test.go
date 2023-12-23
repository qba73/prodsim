package prodsim_test

import (
	"testing"

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

}

func TestListStagesInNewEmptyProdLine(t *testing.T) {
	t.Parallel()

}

func TestListStagesInProductionLine(t *testing.T) {
	t.Parallel()

}
