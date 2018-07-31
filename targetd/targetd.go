package targetd

import (
	"errors"
	"log"

	"github.com/powerman/rpc-codec/jsonrpc2"
)

//Endpoint represents our targetd backend/endpoint
type Endpoint struct {
	URL string
}

//createVolumeRequest provides all of the arguments we need for a std create request
type createVolumeRequest struct {
	Pool string `json:"pool,omitempty"`
	Name string `json:"name,omitempty"`
	Size int64  `json:"size,omitempty"`
}

type deleteVolumeRequest struct {
	Name string `json:"name,omitempty"`
	Pool string `json:"pool,omitempty"`
}

type cloneVolumeRequest struct {
}

type export struct {
	InitiatorWWN string
	Lun          int32
	VolName      string
	VolSize      int32
	VolUUID      string
	Pool         string
}

type exports []export

type createExportRequest struct {
	Pool         string `json:"pool,omitempty"`
	Vol          string `json:"vol,omitempty"`
	InitiatorWWN string `json:"initiator_wwn,omitempty"`
	Lun          int32  `json:"lun,omitempty"`
}

type deleteExportRequest struct {
	Pool         string `json:"pool,omitempty"`
	Vol          string `json:"vol,omitempty"`
	InitiatorWWN string `json:"initiator_wwn,omitempty"`
}

func (e *Endpoint) getConnection() (*jsonrpc2.Client, error) {
	client := jsonrpc2.NewHTTPClient(e.URL)
	if client == nil {
		log.Printf("error creating connection to targetd endpoint: %s", e.URL)
		return nil, errors.New("failed to connect to endpoint")
	}
	return client, nil
}

//NewTargetd creates a new targetd instance to issue commands against
func (e *Endpoint) NewTargetd(url string) Endpoint {
	return Endpoint{
		URL: url,
	}
}

//CreateVolume issues a create request to the targetd endpoint
func (e *Endpoint) CreateVolume(name, pool string, size int64) error {
	req := createVolumeRequest{
		Pool: pool,
		Name: name,
		Size: size,
	}
	client, err := e.getConnection()
	if err != nil {
		return err
	}
	err = client.Call("vol_create", req, nil)
	return err
}

// DeleteVolume issues a Delete request of the specified volume to the targetd endpoint
func (e *Endpoint) DeleteVolume(n, p string) error {
	req := deleteVolumeRequest{
		Name: n,
		Pool: p,
	}
	client, err := e.getConnection()
	if err != nil {
		return err
	}
	err = client.Call("vol_destroy", req, nil)
	return err
}

//CloneVolume issues a Clone request of the specified volume to the targetd endpoint
func (e *Endpoint) CloneVolume() error {
	return nil
}

//ListExports issues a request for the targetd endpoint to report all of it's current exports
func (e *Endpoint) ListExports() error {
	return nil

}

//ListInitiators issues a request for the targetd endpoint to report all currently configured initiators
func (e *Endpoint) ListInitiators() error {
	return nil

}

//CreateExport issues a create iscsi export request to the targetd endpoint
func (e *Endpoint) CreateExport() error {
	return nil

}

//RemoveExport issues  a delete iscsi exprot request to the targetd endpoint
func (e *Endpoint) RemoveExport() error {
	return nil

}

//ListVolumes issues a request to fetch all volumes currently provisioned on the targetd endpoint
func (e *Endpoint) ListVolumes() error {
	return nil

}

//ListAccessGroups issues a request to fetch all currently configured access groups on the targetd endpoint
func (e *Endpoint) ListAccessGroups() error {
	return nil

}

//CreateAccessGroup issues a request to create a new access group on the targetd endpoint
func (e *Endpoint) CreateAccessGroup() error {
	return nil

}

//DeleteAccessGroup issues a request to delete an access group on the targetd endpoint
func (e *Endpoint) DeleteAccessGroup() error {
	return nil

}
