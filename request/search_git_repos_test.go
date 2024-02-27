package request_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/request"
	"github.com/limit-lab/zeabur-go/types"
)

func TestClient_SearchGitRepos(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))

	req := types.NewSearchGitReposRequest(10, "")
	resp, err := client.SearchGitRepos(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
