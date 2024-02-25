package request_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_Invite(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))

	req := types.NewInviteRequest(os.Getenv("PROJECT_ID"), os.Getenv("EMAIL"))
	resp, err := client.Invite(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
