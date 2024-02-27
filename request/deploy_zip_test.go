package request_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_DeployZipProject(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))
	code, err := os.Open(os.Getenv("CODE_ZIP"))
	if err != nil {
		t.Fatal(err)
	}

	req := types.NewDeployRequest(
		os.Getenv("PROJECT_ID"), os.Getenv("SERVICE_ID"), os.Getenv("ENVIRONMENT_ID"), code,
	)
	resp, err := client.DeployZip(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
