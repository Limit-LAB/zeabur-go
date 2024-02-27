package helper_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/limit-lab/zeabur-go/helper"
	"github.com/limit-lab/zeabur-go/request"
)

func TestHelper_DeployProject(t *testing.T) {
	client := request.NewClient(os.Getenv("API_KEY"))
	helper := helper.NewHelper(client)
	code, err := os.Open(os.Getenv("CODE_ZIP"))
	if err != nil {
		t.Fatal(err)
	}

	domain, err := helper.DeployZipProject(context.Background(),
		code, os.Getenv("SERVICE_NAME"), os.Getenv("SERVICE_NAME"),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("url: ", domain)
}
