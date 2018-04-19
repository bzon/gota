package gota

import (
	"fmt"
	"testing"
)

var nexus = Nexus{
	HostURL:        "http://localhost:8081",
	Username:       "admin",
	Password:       "admin123",
	SiteRepository: "site",
}

var component = NexusComponent{
	SourceFile: "./assets/index.html",
	Filename:   "index.html",
	Directory:  "go_upload_test",
}

func TestNexusUpload(t *testing.T) {
	resp, err := nexus.NexusUpload(component)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.Request.URL.RequestURI())
	if resp.StatusCode != 201 {
		t.Fatalf("wanted 201 got %d", resp.StatusCode)
	}
}
