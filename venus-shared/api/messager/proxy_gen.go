// Code generated by github.com/filecoin-project/venus/venus-devtool/api-gen. DO NOT EDIT.
package messager

import (
	"context"
	"time"

	address "github.com/filecoin-project/go-address"
	cid "github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/venus/venus-shared/types"
	mtypes "github.com/filecoin-project/venus/venus-shared/types/messager"
)

type IMessagerStruct struct {
	Internal struct {
		ActiveAddress            func(ctx context.Context, addr address.Address) error                                                                                      `perm:"admin"`
		ClearUnFillMessage       func(ctx context.Context, addr address.Address) (int, error)                                                                               `perm:"admin"`
		DeleteAddress            func(ctx context.Context, addr address.Address) error                                                                                      `perm:"admin"`
		DeleteNode               func(ctx context.Context, name string) error                                                                                               `perm:"admin"`
		ForbiddenAddress         func(ctx context.Context, addr address.Address) error                                                                                      `perm:"admin"`
		ForcePushMessage         func(ctx context.Context, account string, msg *types.Message, meta *mtypes.SendSpec) (string, error)                                       `perm:"admin"`
		ForcePushMessageWithId   func(ctx context.Context, id string, account string, msg *types.Message, meta *mtypes.SendSpec) (string, error)                            `perm:"write"`
		GetAddress               func(ctx context.Context, addr address.Address) (*mtypes.Address, error)                                                                   `perm:"admin"`
		GetMessageByFromAndNonce func(ctx context.Context, from address.Address, nonce uint64) (*mtypes.Message, error)                                                     `perm:"read"`
		GetMessageBySignedCid    func(ctx context.Context, cid cid.Cid) (*mtypes.Message, error)                                                                            `perm:"read"`
		GetMessageByUid          func(ctx context.Context, id string) (*mtypes.Message, error)                                                                              `perm:"read"`
		GetMessageByUnsignedCid  func(ctx context.Context, cid cid.Cid) (*mtypes.Message, error)                                                                            `perm:"read"`
		GetNode                  func(ctx context.Context, name string) (*mtypes.Node, error)                                                                               `perm:"admin"`
		GetSharedParams          func(ctx context.Context) (*mtypes.SharedSpec, error)                                                                                      `perm:"admin"`
		HasAddress               func(ctx context.Context, addr address.Address) (bool, error)                                                                              `perm:"read"`
		HasMessageByUid          func(ctx context.Context, id string) (bool, error)                                                                                         `perm:"read"`
		HasNode                  func(ctx context.Context, name string) (bool, error)                                                                                       `perm:"admin"`
		ListAddress              func(ctx context.Context) ([]*mtypes.Address, error)                                                                                       `perm:"admin"`
		ListBlockedMessage       func(ctx context.Context, addr address.Address, d time.Duration) ([]*mtypes.Message, error)                                                `perm:"admin"`
		ListFailedMessage        func(ctx context.Context) ([]*mtypes.Message, error)                                                                                       `perm:"admin"`
		ListMessage              func(ctx context.Context) ([]*mtypes.Message, error)                                                                                       `perm:"admin"`
		ListMessageByAddress     func(ctx context.Context, addr address.Address) ([]*mtypes.Message, error)                                                                 `perm:"admin"`
		ListMessageByFromState   func(ctx context.Context, from address.Address, state mtypes.MessageState, isAsc bool, pageIndex, pageSize int) ([]*mtypes.Message, error) `perm:"admin"`
		ListNode                 func(ctx context.Context) ([]*mtypes.Node, error)                                                                                          `perm:"admin"`
		MarkBadMessage           func(ctx context.Context, id string) error                                                                                                 `perm:"admin"`
		NetAddrsListen           func(ctx context.Context) (peer.AddrInfo, error)                                                                                           `perm:"read"`
		NetConnect               func(ctx context.Context, pi peer.AddrInfo) error                                                                                          `perm:"admin"`
		NetFindPeer              func(ctx context.Context, p peer.ID) (peer.AddrInfo, error)                                                                                `perm:"read"`
		NetPeers                 func(ctx context.Context) ([]peer.AddrInfo, error)                                                                                         `perm:"read"`
		PushMessage              func(ctx context.Context, msg *types.Message, meta *mtypes.SendSpec) (string, error)                                                       `perm:"write"`
		PushMessageWithId        func(ctx context.Context, id string, msg *types.Message, meta *mtypes.SendSpec) (string, error)                                            `perm:"write"`
		RecoverFailedMsg         func(ctx context.Context, addr address.Address) ([]string, error)                                                                          `perm:"admin"`
		ReplaceMessage           func(ctx context.Context, params *mtypes.ReplacMessageParams) (cid.Cid, error)                                                             `perm:"admin"`
		RepublishMessage         func(ctx context.Context, id string) error                                                                                                 `perm:"admin"`
		SaveNode                 func(ctx context.Context, node *mtypes.Node) error                                                                                         `perm:"admin"`
		Send                     func(ctx context.Context, params mtypes.QuickSendParams) (string, error)                                                                   `perm:"admin"`
		SetFeeParams             func(ctx context.Context, params *mtypes.AddressSpec) error                                                                                `perm:"admin"`
		SetLogLevel              func(ctx context.Context, level string) error                                                                                              `perm:"admin"`
		SetSelectMsgNum          func(ctx context.Context, addr address.Address, num uint64) error                                                                          `perm:"admin"`
		SetSharedParams          func(ctx context.Context, params *mtypes.SharedSpec) error                                                                                 `perm:"admin"`
		UpdateAllFilledMessage   func(ctx context.Context) (int, error)                                                                                                     `perm:"admin"`
		UpdateFilledMessageByID  func(ctx context.Context, id string) (string, error)                                                                                       `perm:"admin"`
		UpdateMessageStateByID   func(ctx context.Context, id string, state mtypes.MessageState) error                                                                      `perm:"admin"`
		UpdateNonce              func(ctx context.Context, addr address.Address, nonce uint64) error                                                                        `perm:"admin"`
		Version                  func(ctx context.Context) (types.Version, error)                                                                                           `perm:"read"`
		WaitMessage              func(ctx context.Context, id string, confidence uint64) (*mtypes.Message, error)                                                           `perm:"read"`
		WalletHas                func(ctx context.Context, addr address.Address) (bool, error)                                                                              `perm:"read"`
	}
}

func (s *IMessagerStruct) ActiveAddress(p0 context.Context, p1 address.Address) error {
	return s.Internal.ActiveAddress(p0, p1)
}
func (s *IMessagerStruct) ClearUnFillMessage(p0 context.Context, p1 address.Address) (int, error) {
	return s.Internal.ClearUnFillMessage(p0, p1)
}
func (s *IMessagerStruct) DeleteAddress(p0 context.Context, p1 address.Address) error {
	return s.Internal.DeleteAddress(p0, p1)
}
func (s *IMessagerStruct) DeleteNode(p0 context.Context, p1 string) error {
	return s.Internal.DeleteNode(p0, p1)
}
func (s *IMessagerStruct) ForbiddenAddress(p0 context.Context, p1 address.Address) error {
	return s.Internal.ForbiddenAddress(p0, p1)
}
func (s *IMessagerStruct) ForcePushMessage(p0 context.Context, p1 string, p2 *types.Message, p3 *mtypes.SendSpec) (string, error) {
	return s.Internal.ForcePushMessage(p0, p1, p2, p3)
}
func (s *IMessagerStruct) ForcePushMessageWithId(p0 context.Context, p1 string, p2 string, p3 *types.Message, p4 *mtypes.SendSpec) (string, error) {
	return s.Internal.ForcePushMessageWithId(p0, p1, p2, p3, p4)
}
func (s *IMessagerStruct) GetAddress(p0 context.Context, p1 address.Address) (*mtypes.Address, error) {
	return s.Internal.GetAddress(p0, p1)
}
func (s *IMessagerStruct) GetMessageByFromAndNonce(p0 context.Context, p1 address.Address, p2 uint64) (*mtypes.Message, error) {
	return s.Internal.GetMessageByFromAndNonce(p0, p1, p2)
}
func (s *IMessagerStruct) GetMessageBySignedCid(p0 context.Context, p1 cid.Cid) (*mtypes.Message, error) {
	return s.Internal.GetMessageBySignedCid(p0, p1)
}
func (s *IMessagerStruct) GetMessageByUid(p0 context.Context, p1 string) (*mtypes.Message, error) {
	return s.Internal.GetMessageByUid(p0, p1)
}
func (s *IMessagerStruct) GetMessageByUnsignedCid(p0 context.Context, p1 cid.Cid) (*mtypes.Message, error) {
	return s.Internal.GetMessageByUnsignedCid(p0, p1)
}
func (s *IMessagerStruct) GetNode(p0 context.Context, p1 string) (*mtypes.Node, error) {
	return s.Internal.GetNode(p0, p1)
}
func (s *IMessagerStruct) GetSharedParams(p0 context.Context) (*mtypes.SharedSpec, error) {
	return s.Internal.GetSharedParams(p0)
}
func (s *IMessagerStruct) HasAddress(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.HasAddress(p0, p1)
}
func (s *IMessagerStruct) HasMessageByUid(p0 context.Context, p1 string) (bool, error) {
	return s.Internal.HasMessageByUid(p0, p1)
}
func (s *IMessagerStruct) HasNode(p0 context.Context, p1 string) (bool, error) {
	return s.Internal.HasNode(p0, p1)
}
func (s *IMessagerStruct) ListAddress(p0 context.Context) ([]*mtypes.Address, error) {
	return s.Internal.ListAddress(p0)
}
func (s *IMessagerStruct) ListBlockedMessage(p0 context.Context, p1 address.Address, p2 time.Duration) ([]*mtypes.Message, error) {
	return s.Internal.ListBlockedMessage(p0, p1, p2)
}
func (s *IMessagerStruct) ListFailedMessage(p0 context.Context) ([]*mtypes.Message, error) {
	return s.Internal.ListFailedMessage(p0)
}
func (s *IMessagerStruct) ListMessage(p0 context.Context) ([]*mtypes.Message, error) {
	return s.Internal.ListMessage(p0)
}
func (s *IMessagerStruct) ListMessageByAddress(p0 context.Context, p1 address.Address) ([]*mtypes.Message, error) {
	return s.Internal.ListMessageByAddress(p0, p1)
}
func (s *IMessagerStruct) ListMessageByFromState(p0 context.Context, p1 address.Address, p2 mtypes.MessageState, p3 bool, p4, p5 int) ([]*mtypes.Message, error) {
	return s.Internal.ListMessageByFromState(p0, p1, p2, p3, p4, p5)
}
func (s *IMessagerStruct) ListNode(p0 context.Context) ([]*mtypes.Node, error) {
	return s.Internal.ListNode(p0)
}
func (s *IMessagerStruct) MarkBadMessage(p0 context.Context, p1 string) error {
	return s.Internal.MarkBadMessage(p0, p1)
}
func (s *IMessagerStruct) NetAddrsListen(p0 context.Context) (peer.AddrInfo, error) {
	return s.Internal.NetAddrsListen(p0)
}
func (s *IMessagerStruct) NetConnect(p0 context.Context, p1 peer.AddrInfo) error {
	return s.Internal.NetConnect(p0, p1)
}
func (s *IMessagerStruct) NetFindPeer(p0 context.Context, p1 peer.ID) (peer.AddrInfo, error) {
	return s.Internal.NetFindPeer(p0, p1)
}
func (s *IMessagerStruct) NetPeers(p0 context.Context) ([]peer.AddrInfo, error) {
	return s.Internal.NetPeers(p0)
}
func (s *IMessagerStruct) PushMessage(p0 context.Context, p1 *types.Message, p2 *mtypes.SendSpec) (string, error) {
	return s.Internal.PushMessage(p0, p1, p2)
}
func (s *IMessagerStruct) PushMessageWithId(p0 context.Context, p1 string, p2 *types.Message, p3 *mtypes.SendSpec) (string, error) {
	return s.Internal.PushMessageWithId(p0, p1, p2, p3)
}
func (s *IMessagerStruct) RecoverFailedMsg(p0 context.Context, p1 address.Address) ([]string, error) {
	return s.Internal.RecoverFailedMsg(p0, p1)
}
func (s *IMessagerStruct) ReplaceMessage(p0 context.Context, p1 *mtypes.ReplacMessageParams) (cid.Cid, error) {
	return s.Internal.ReplaceMessage(p0, p1)
}
func (s *IMessagerStruct) RepublishMessage(p0 context.Context, p1 string) error {
	return s.Internal.RepublishMessage(p0, p1)
}
func (s *IMessagerStruct) SaveNode(p0 context.Context, p1 *mtypes.Node) error {
	return s.Internal.SaveNode(p0, p1)
}
func (s *IMessagerStruct) Send(p0 context.Context, p1 mtypes.QuickSendParams) (string, error) {
	return s.Internal.Send(p0, p1)
}
func (s *IMessagerStruct) SetFeeParams(p0 context.Context, p1 *mtypes.AddressSpec) error {
	return s.Internal.SetFeeParams(p0, p1)
}
func (s *IMessagerStruct) SetLogLevel(p0 context.Context, p1 string) error {
	return s.Internal.SetLogLevel(p0, p1)
}
func (s *IMessagerStruct) SetSelectMsgNum(p0 context.Context, p1 address.Address, p2 uint64) error {
	return s.Internal.SetSelectMsgNum(p0, p1, p2)
}
func (s *IMessagerStruct) SetSharedParams(p0 context.Context, p1 *mtypes.SharedSpec) error {
	return s.Internal.SetSharedParams(p0, p1)
}
func (s *IMessagerStruct) UpdateAllFilledMessage(p0 context.Context) (int, error) {
	return s.Internal.UpdateAllFilledMessage(p0)
}
func (s *IMessagerStruct) UpdateFilledMessageByID(p0 context.Context, p1 string) (string, error) {
	return s.Internal.UpdateFilledMessageByID(p0, p1)
}
func (s *IMessagerStruct) UpdateMessageStateByID(p0 context.Context, p1 string, p2 mtypes.MessageState) error {
	return s.Internal.UpdateMessageStateByID(p0, p1, p2)
}
func (s *IMessagerStruct) UpdateNonce(p0 context.Context, p1 address.Address, p2 uint64) error {
	return s.Internal.UpdateNonce(p0, p1, p2)
}
func (s *IMessagerStruct) Version(p0 context.Context) (types.Version, error) {
	return s.Internal.Version(p0)
}
func (s *IMessagerStruct) WaitMessage(p0 context.Context, p1 string, p2 uint64) (*mtypes.Message, error) {
	return s.Internal.WaitMessage(p0, p1, p2)
}
func (s *IMessagerStruct) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.WalletHas(p0, p1)
}
