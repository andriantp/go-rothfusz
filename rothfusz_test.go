package rothfusz

import (
	"encoding/json"
	"testing"

	"github.com/andriantp/go-rothfusz/rothfusz"
)

func TestRothfusz(t *testing.T) {
	temp := float64(32)
	rh := float64(70)

	repo := rothfusz.NewRothfusz(
		26.7, // min valid temp
		85,   // humid RH threshold
	)

	result := repo.CalculateHeatIndex(temp, rh)
	js, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		t.Fatalf("Failed to construct result:%v", err)
	}
	t.Logf("result:\n%s", string(js))
}
