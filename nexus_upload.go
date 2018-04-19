package gota

import (
	"net/http"
	"os"
)

type Nexus struct {
	SiteRepository, HostURL, Username, Password string
}

type NexusComponent struct {
	SourceFile, Filename, Directory string
}

func (n *Nexus) NexusUpload(c NexusComponent) (*http.Response, error) {
	file, err := os.Open(c.SourceFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	uri := n.HostURL + "/repository/" + n.SiteRepository + "/" + c.Directory + "/" + c.Filename
	req, err := http.NewRequest("PUT", uri, file)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(n.Username, n.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
