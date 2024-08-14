package database

import (
	"testing"
)

func TestClient(t *testing.T) {
	type testResult struct {
		Id        int    `json:"id"`
		FirstName string `json:"first_name"`
	}
	sb := GetClient()
	var res []testResult
	// err := sb.DB.From("classes").Select("name", "subject").Eq("subject", "Math").Execute(&res)
	err := sb.DB.From("teachers").Select("*").Execute(&res)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Log(res)
}
