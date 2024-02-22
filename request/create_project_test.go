package request_test

import (
	"context"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_CreateProject(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))

	req := types.NewCreateProjectRequest(types.RegionSfo1)
	resp, err := client.CreateProject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
