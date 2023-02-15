package handler

import (
	"testing"
	"context"
	"net/url"
)

func TestGetReadQryCond(t *testing.T) {

	ctx := context.TODO()
	
	curUrl := "http://localhost/read?laptop=HP"
	
	u, err := url.Parse(curUrl)
	if err != nil {
		t.Fatal(err)
	}

	result := GetReadQryCond(ctx, u)
	if result != `laptop='HP'` {
		t.Fatalf("Error. [%v]", result)
	}

	t.Logf("Success")

}
