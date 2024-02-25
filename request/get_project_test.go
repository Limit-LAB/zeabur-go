package request_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_GetProject(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))

	req := types.NewGetProjectRequest(os.Getenv("PROJECT_ID"))
	resp, err := client.GetProject(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
