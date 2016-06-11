package dropbox

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dropbox/dropbox-sdk-go-unofficial"
	"github.com/dropbox/dropbox-sdk-go-unofficial/sharing"
	"github.com/dropbox/dropbox-sdk-go-unofficial/files"
	ma "gx/ipfs/QmYzDkkgAEmrcNzFCiYo6L1dTX4EAG1gZkbtdbd9trL4vd/go-multiaddr"
	peer "gx/ipfs/QmbyvM8zRFDkbFdYyt1MnevUMJ62SiSGbfDFZ3Z8nkrzr4/go-libp2p-peer"
	mh "gx/ipfs/QmYf7ng2hG5XBtJA3tN34DQ2GUN5HNksEw1rLDkmr6vGku/go-multihash"
)


type DropBoxStorage struct {
	apiToken   string
}

func NewDropBoxStorage(apiToken string) *DropBoxStorage {
	return &DropBoxStorage{
		apiToken: apiToken,
	}
}

func (s *DropBoxStorage) Store(peerID peer.ID, ciphertext []byte) (ma.Multiaddr, error) {
	api := dropbox.Client(s.apiToken, dropbox.Options{Verbose: true})
	hash := sha256.Sum256(ciphertext)
	hex := hex.EncodeToString(hash)

	// Upload ciphertext
	uploadArg := files.NewCommitInfo("/" + hex)
	r := bytes.NewReader(ciphertext)
	_, err := api.Upload(uploadArg, r)
	if err != nil {
		return nil, err
	}

	// Set public sharing
	sharingArg := sharing.NewCreateSharedLinkArg("/" + hex)
	res, err := api.CreateSharedLink(sharingArg)
	if err != nil {
		return nil, err
	}

	// Create encoded multiaddr
	b, err := mh.Encode([]byte(res.Url), mh.SHA1)
	if err != nil {
		return nil, err
	}
	m, err := mh.Cast(b)
	if err != nil {
		return nil, err
	}

	addr, err := ma.NewMultiaddr("/ipfs/" + m.B58String() + "/https/")
	if err != nil {
		return nil, err
	}
	return addr, nil
}