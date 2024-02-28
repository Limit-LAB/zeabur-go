package request_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_UpdateRootDirectory(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))

	req := types.NewUpdateRootDirectoryRequest(os.Getenv("SERVICE_ID"), os.Getenv("ROOT_DIRECTORY"))
	resp, err := client.UpdateRootDirectory(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*resp)
}
