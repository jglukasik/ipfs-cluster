package clientInterface

import (
	"context"

	"github.com/ipfs/ipfs-cluster/api"

	cid "github.com/ipfs/go-cid"
	shell "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-cmdkit/files"
	peer "github.com/libp2p/go-libp2p-peer"
)

// ClientIface defines the interface to be used by API clients to interact with
// the ipfs-cluster-service
type ClientIface interface {
	IPFS() *shell.Shell
	ID() (api.ID, error)

	Peers() ([]api.ID, error)
	PeerAdd(pid peer.ID) (api.ID, error)
	PeerRm(id peer.ID) error

	Add(paths []string, params *api.AddParams, out chan<- *api.AddedOutput) error
	Pin(ci *cid.Cid, replicationFactorMin, replicationFactorMax int, name string) error
	Unpin(ci *cid.Cid) error

	Allocations(filter api.PinType) ([]api.Pin, error)
	Allocation(ci *cid.Cid) (api.Pin, error)

	Status(ci *cid.Cid, local bool) (api.GlobalPinInfo, error)
	StatusAll(local bool) ([]api.GlobalPinInfo, error)

	Sync(ci *cid.Cid, local bool) (api.GlobalPinInfo, error)
	SyncAll(local bool) ([]api.GlobalPinInfo, error)

	Recover(ci *cid.Cid, local bool) (api.GlobalPinInfo, error)
	RecoverAll(local bool) ([]api.GlobalPinInfo, error)

	Version() (api.Version, error)

	GetConnectGraph() (api.ConnectGraphSerial, error)
	WaitFor(ctx context.Context, fp api.StatusFilterParams) (api.GlobalPinInfo, error)
	AddMultiFile(multiFileR *files.MultiFileReader, params *api.AddParams, out chan<- *api.AddedOutput) error
}
