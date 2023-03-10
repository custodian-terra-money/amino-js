// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/cosmos/amino-js/go/src"
)

var done chan bool

func main() {
	done = make(chan bool)

	js.Global().Set("Amino", map[string]interface{}{
		// @formatter:off
		// Bech32
		"encodeBech32": js.FuncOf(EncodeBech32),
		"decodeBech32": js.FuncOf(DecodeBech32),

		// Basic encoding
		"encodeByte":      js.FuncOf(EncodeByte),
		"encodeByteSlice": js.FuncOf(EncodeByteSlice),
		"encodeInt8":      js.FuncOf(EncodeInt8),
		"encodeInt16":     js.FuncOf(EncodeInt16),
		"encodeInt32":     js.FuncOf(EncodeInt32),
		"encodeInt64":     js.FuncOf(EncodeInt64),
		"encodeVarint":    js.FuncOf(EncodeVarint),
		"encodeUint8":     js.FuncOf(EncodeUint8),
		"encodeUint16":    js.FuncOf(EncodeUint16),
		"encodeUint32":    js.FuncOf(EncodeUint32),
		"encodeUint64":    js.FuncOf(EncodeUint64),
		"encodeUvarint":   js.FuncOf(EncodeUvarint),
		"encodeFloat32":   js.FuncOf(EncodeFloat32),
		"encodeFloat64":   js.FuncOf(EncodeFloat64),
		"encodeBool":      js.FuncOf(EncodeBool),
		"encodeString":    js.FuncOf(EncodeString),
		"encodeTime":      js.FuncOf(EncodeTime),

		// Basic decoding
		"decodeByte":      js.FuncOf(DecodeByte),
		"decodeByteSlice": js.FuncOf(DecodeByteSlice),
		"decodeInt8":      js.FuncOf(DecodeInt8),
		"decodeInt16":     js.FuncOf(DecodeInt16),
		"decodeInt32":     js.FuncOf(DecodeInt32),
		"decodeInt64":     js.FuncOf(DecodeInt64),
		"decodeVarint":    js.FuncOf(DecodeVarint),
		"decodeUint8":     js.FuncOf(DecodeUint8),
		"decodeUint16":    js.FuncOf(DecodeUint16),
		"decodeUint32":    js.FuncOf(DecodeUint32),
		"decodeUint64":    js.FuncOf(DecodeUint64),
		"decodeUvarint":   js.FuncOf(DecodeUvarint),
		"decodeFloat32":   js.FuncOf(DecodeFloat32),
		"decodeFloat64":   js.FuncOf(DecodeFloat64),
		"decodeBool":      js.FuncOf(DecodeBool),
		"decodeString":    js.FuncOf(DecodeString),
		"decodeTime":      js.FuncOf(DecodeTime),

		// Meta
		"decodeDisambPrefixBytes": js.FuncOf(DecodeDisambPrefixBytes),
		"nameToDisfix":            js.FuncOf(NameToDisfix),
		"byteSliceSize":           js.FuncOf(ByteSliceSize),
		"uvarintSize":             js.FuncOf(UvarintSize),
		"varintSize":              js.FuncOf(VarintSize),

		// Typed encoding
		"encodeMultiStoreProofOp":              encodeDecodeType(src.EncodeMultiStoreProofOp),
		"encodeIAVLAbsenceOp":                  encodeDecodeType(src.EncodeIAVLAbsenceOp),
		"encodeIAVLValueOp":                    encodeDecodeType(src.EncodeIAVLValueOp),
		"encodePrivKeyLedgerSecp256k1":         encodeDecodeType(src.EncodePrivKeyLedgerSecp256k1),
		"encodeInfo":                           encodeDecodeType(src.EncodeInfo),
		"encodeBIP44Params":                    encodeDecodeType(src.EncodeBIP44Params),
		"encodeLocalInfo":                      encodeDecodeType(src.EncodeLocalInfo),
		"encodeLedgerInfo":                     encodeDecodeType(src.EncodeLedgerInfo),
		"encodeOfflineInfo":                    encodeDecodeType(src.EncodeOfflineInfo),
		"encodeMultiInfo":                      encodeDecodeType(src.EncodeMultiInfo),
		"encodeMsg":                            encodeDecodeType(src.EncodeMsg),
		"encodeTx":                             encodeDecodeType(src.EncodeTx),
		"encodeAccount":                        encodeDecodeType(src.EncodeAccount),
		"encodeVestingAccount":                 encodeDecodeType(src.EncodeVestingAccount),
		"encodeBaseAccount":                    encodeDecodeType(src.EncodeBaseAccount),
		"encodeBaseVestingAccount":             encodeDecodeType(src.EncodeBaseVestingAccount),
		"encodeContinuousVestingAccount":       encodeDecodeType(src.EncodeContinuousVestingAccount),
		"encodeDelayedVestingAccount":          encodeDecodeType(src.EncodeDelayedVestingAccount),
		"encodeStdTx":                          encodeDecodeType(src.EncodeStdTx),
		"encodeMsgSend":                        encodeDecodeType(src.EncodeMsgSend),
		"encodeMsgMultiSend":                   encodeDecodeType(src.EncodeMsgMultiSend),
		"encodeMsgVerifyInvariant":             encodeDecodeType(src.EncodeMsgVerifyInvariant),
		"encodeMsgWithdrawDelegatorReward":     encodeDecodeType(src.EncodeMsgWithdrawDelegatorReward),
		"encodeMsgWithdrawValidatorCommission": encodeDecodeType(src.EncodeMsgWithdrawValidatorCommission),
		"encodeMsgSetWithdrawAddress":          encodeDecodeType(src.EncodeMsgSetWithdrawAddress),
		"encodeContent":                        encodeDecodeType(src.EncodeContent),
		"encodeMsgSubmitProposal":              encodeDecodeType(src.EncodeMsgSubmitProposal),
		"encodeMsgDeposit":                     encodeDecodeType(src.EncodeMsgDeposit),
		"encodeMsgVote":                        encodeDecodeType(src.EncodeMsgVote),
		"encodeTextProposal":                   encodeDecodeType(src.EncodeTextProposal),
		"encodeSoftwareUpgradeProposal":        encodeDecodeType(src.EncodeSoftwareUpgradeProposal),
		"encodeCommunityPoolSpendProposal":     encodeDecodeType(src.EncodeCommunityPoolSpendProposal),
		"encodeMsgIBCTransfer":                 encodeDecodeType(src.EncodeMsgIBCTransfer),
		"encodeMsgIBCReceive":                  encodeDecodeType(src.EncodeMsgIBCReceive),
		"encodeParameterChangeProposal":        encodeDecodeType(src.EncodeParameterChangeProposal),
		"encodeMsgUnjail":                      encodeDecodeType(src.EncodeMsgUnjail),
		"encodeMsgCreateValidator":             encodeDecodeType(src.EncodeMsgCreateValidator),
		"encodeMsgEditValidator":               encodeDecodeType(src.EncodeMsgEditValidator),
		"encodeMsgDelegate":                    encodeDecodeType(src.EncodeMsgDelegate),
		"encodeMsgUndelegate":                  encodeDecodeType(src.EncodeMsgUndelegate),
		"encodeMsgBeginRedelegate":             encodeDecodeType(src.EncodeMsgBeginRedelegate),
		"encodeBlockchainMessage":              encodeDecodeType(src.EncodeBlockchainMessage),
		"encodeBcBlockRequestMessage":          encodeDecodeType(src.EncodeBcBlockRequestMessage),
		"encodeBcBlockResponseMessage":         encodeDecodeType(src.EncodeBcBlockResponseMessage),
		"encodeBcNoBlockResponseMessage":       encodeDecodeType(src.EncodeBcNoBlockResponseMessage),
		"encodeBcStatusResponseMessage":        encodeDecodeType(src.EncodeBcStatusResponseMessage),
		"encodeBcStatusRequestMessage":         encodeDecodeType(src.EncodeBcStatusRequestMessage),
		"encodeConsensusMessage":               encodeDecodeType(src.EncodeConsensusMessage),
		"encodeNewRoundStepMessage":            encodeDecodeType(src.EncodeNewRoundStepMessage),
		"encodeNewValidBlockMessage":           encodeDecodeType(src.EncodeNewValidBlockMessage),
		"encodeProposalMessage":                encodeDecodeType(src.EncodeProposalMessage),
		"encodeProposalPOLMessage":             encodeDecodeType(src.EncodeProposalPOLMessage),
		"encodeBlockPartMessage":               encodeDecodeType(src.EncodeBlockPartMessage),
		"encodeVoteMessage":                    encodeDecodeType(src.EncodeVoteMessage),
		"encodeHasVoteMessage":                 encodeDecodeType(src.EncodeHasVoteMessage),
		"encodeVoteSetMaj23Message":            encodeDecodeType(src.EncodeVoteSetMaj23Message),
		"encodeVoteSetBitsMessage":             encodeDecodeType(src.EncodeVoteSetBitsMessage),
		"encodeWALMessage":                     encodeDecodeType(src.EncodeWALMessage),
		"encodeMsgInfo":                        encodeDecodeType(src.EncodeMsgInfo),
		"encodeTimeoutInfo":                    encodeDecodeType(src.EncodeTimeoutInfo),
		"encodeEndHeightMessage":               encodeDecodeType(src.EncodeEndHeightMessage),
		"encodePubKey":                         encodeDecodeType(src.EncodePubKey),
		"encodePrivKey":                        encodeDecodeType(src.EncodePrivKey),
		"encodePubKeyEd25519":                  encodeDecodeType(src.EncodePubKeyEd25519),
		"encodePrivKeyEd25519":                 encodeDecodeType(src.EncodePrivKeyEd25519),
		"encodePubKeySecp256k1":                encodeDecodeType(src.EncodePubKeySecp256k1),
		"encodePrivKeySecp256k1":               encodeDecodeType(src.EncodePrivKeySecp256k1),
		"encodePubKeyMultisigThreshold":        encodeDecodeType(src.EncodePubKeyMultisigThreshold),
		"encodeEvidenceMessage":                encodeDecodeType(src.EncodeEvidenceMessage),
		"encodeEvidenceListMessage":            encodeDecodeType(src.EncodeEvidenceListMessage),
		"encodeMempoolMessage":                 encodeDecodeType(src.EncodeMempoolMessage),
		"encodeTxMessage":                      encodeDecodeType(src.EncodeTxMessage),
		"encodePacket":                         encodeDecodeType(src.EncodePacket),
		"encodePacketPing":                     encodeDecodeType(src.EncodePacketPing),
		"encodePacketPong":                     encodeDecodeType(src.EncodePacketPong),
		"encodePacketMsg":                      encodeDecodeType(src.EncodePacketMsg),
		"encodePexMessage":                     encodeDecodeType(src.EncodePexMessage),
		"encodePexRequestMessage":              encodeDecodeType(src.EncodePexRequestMessage),
		"encodePexAddrsMessage":                encodeDecodeType(src.EncodePexAddrsMessage),
		"encodeRemoteSignerMsg":                encodeDecodeType(src.EncodeRemoteSignerMsg),
		"encodePubKeyRequest":                  encodeDecodeType(src.EncodePubKeyRequest),
		"encodePubKeyResponse":                 encodeDecodeType(src.EncodePubKeyResponse),
		"encodeSignVoteRequest":                encodeDecodeType(src.EncodeSignVoteRequest),
		"encodeSignedVoteResponse":             encodeDecodeType(src.EncodeSignedVoteResponse),
		"encodeSignProposalRequest":            encodeDecodeType(src.EncodeSignProposalRequest),
		"encodeSignedProposalResponse":         encodeDecodeType(src.EncodeSignedProposalResponse),
		"encodePingRequest":                    encodeDecodeType(src.EncodePingRequest),
		"encodePingResponse":                   encodeDecodeType(src.EncodePingResponse),
		"encodeTMEventData":                    encodeDecodeType(src.EncodeTMEventData),
		"encodeEventDataNewBlock":              encodeDecodeType(src.EncodeEventDataNewBlock),
		"encodeEventDataNewBlockHeader":        encodeDecodeType(src.EncodeEventDataNewBlockHeader),
		"encodeEventDataTx":                    encodeDecodeType(src.EncodeEventDataTx),
		"encodeEventDataRoundState":            encodeDecodeType(src.EncodeEventDataRoundState),
		"encodeEventDataNewRound":              encodeDecodeType(src.EncodeEventDataNewRound),
		"encodeEventDataCompleteProposal":      encodeDecodeType(src.EncodeEventDataCompleteProposal),
		"encodeEventDataVote":                  encodeDecodeType(src.EncodeEventDataVote),
		"encodeEventDataValidatorSetUpdates":   encodeDecodeType(src.EncodeEventDataValidatorSetUpdates),
		"encodeEventDataString":                encodeDecodeType(src.EncodeEventDataString),
		"encodeEvidence":                       encodeDecodeType(src.EncodeEvidence),
		"encodeDuplicateVoteEvidence":          encodeDecodeType(src.EncodeDuplicateVoteEvidence),
		"encodeMockGoodEvidence":               encodeDecodeType(src.EncodeMockGoodEvidence),
		"encodeMockRandomGoodEvidence":         encodeDecodeType(src.EncodeMockRandomGoodEvidence),
		"encodeMockBadEvidence":                encodeDecodeType(src.EncodeMockBadEvidence),

		// Typed decoding
		"decodeMultiStoreProofOp":              encodeDecodeType(src.DecodeMultiStoreProofOp),
		"decodeIAVLAbsenceOp":                  encodeDecodeType(src.DecodeIAVLAbsenceOp),
		"decodeIAVLValueOp":                    encodeDecodeType(src.DecodeIAVLValueOp),
		"decodePrivKeyLedgerSecp256k1":         encodeDecodeType(src.DecodePrivKeyLedgerSecp256k1),
		"decodeInfo":                           encodeDecodeType(src.DecodeInfo),
		"decodeBIP44Params":                    encodeDecodeType(src.DecodeBIP44Params),
		"decodeLocalInfo":                      encodeDecodeType(src.DecodeLocalInfo),
		"decodeLedgerInfo":                     encodeDecodeType(src.DecodeLedgerInfo),
		"decodeOfflineInfo":                    encodeDecodeType(src.DecodeOfflineInfo),
		"decodeMultiInfo":                      encodeDecodeType(src.DecodeMultiInfo),
		"decodeMsg":                            encodeDecodeType(src.DecodeMsg),
		"decodeTx":                             encodeDecodeType(src.DecodeTx),
		"decodeAccount":                        encodeDecodeType(src.DecodeAccount),
		"decodeVestingAccount":                 encodeDecodeType(src.DecodeVestingAccount),
		"decodeBaseAccount":                    encodeDecodeType(src.DecodeBaseAccount),
		"decodeBaseVestingAccount":             encodeDecodeType(src.DecodeBaseVestingAccount),
		"decodeContinuousVestingAccount":       encodeDecodeType(src.DecodeContinuousVestingAccount),
		"decodeDelayedVestingAccount":          encodeDecodeType(src.DecodeDelayedVestingAccount),
		"decodeStdTx":                          encodeDecodeType(src.DecodeStdTx),
		"decodeMsgSend":                        encodeDecodeType(src.DecodeMsgSend),
		"decodeMsgMultiSend":                   encodeDecodeType(src.DecodeMsgMultiSend),
		"decodeMsgVerifyInvariant":             encodeDecodeType(src.DecodeMsgVerifyInvariant),
		"decodeMsgWithdrawDelegatorReward":     encodeDecodeType(src.DecodeMsgWithdrawDelegatorReward),
		"decodeMsgWithdrawValidatorCommission": encodeDecodeType(src.DecodeMsgWithdrawValidatorCommission),
		"decodeMsgSetWithdrawAddress":          encodeDecodeType(src.DecodeMsgSetWithdrawAddress),
		"decodeContent":                        encodeDecodeType(src.DecodeContent),
		"decodeMsgSubmitProposal":              encodeDecodeType(src.DecodeMsgSubmitProposal),
		"decodeMsgDeposit":                     encodeDecodeType(src.DecodeMsgDeposit),
		"decodeMsgVote":                        encodeDecodeType(src.DecodeMsgVote),
		"decodeTextProposal":                   encodeDecodeType(src.DecodeTextProposal),
		"decodeSoftwareUpgradeProposal":        encodeDecodeType(src.DecodeSoftwareUpgradeProposal),
		"decodeCommunityPoolSpendProposal":     encodeDecodeType(src.DecodeCommunityPoolSpendProposal),
		"decodeMsgIBCTransfer":                 encodeDecodeType(src.DecodeMsgIBCTransfer),
		"decodeMsgIBCReceive":                  encodeDecodeType(src.DecodeMsgIBCReceive),
		"decodeParameterChangeProposal":        encodeDecodeType(src.DecodeParameterChangeProposal),
		"decodeMsgUnjail":                      encodeDecodeType(src.DecodeMsgUnjail),
		"decodeMsgCreateValidator":             encodeDecodeType(src.DecodeMsgCreateValidator),
		"decodeMsgEditValidator":               encodeDecodeType(src.DecodeMsgEditValidator),
		"decodeMsgDelegate":                    encodeDecodeType(src.DecodeMsgDelegate),
		"decodeMsgUndelegate":                  encodeDecodeType(src.DecodeMsgUndelegate),
		"decodeMsgBeginRedelegate":             encodeDecodeType(src.DecodeMsgBeginRedelegate),
		"decodeBlockchainMessage":              encodeDecodeType(src.DecodeBlockchainMessage),
		"decodeBcBlockRequestMessage":          encodeDecodeType(src.DecodeBcBlockRequestMessage),
		"decodeBcBlockResponseMessage":         encodeDecodeType(src.DecodeBcBlockResponseMessage),
		"decodeBcNoBlockResponseMessage":       encodeDecodeType(src.DecodeBcNoBlockResponseMessage),
		"decodeBcStatusResponseMessage":        encodeDecodeType(src.DecodeBcStatusResponseMessage),
		"decodeBcStatusRequestMessage":         encodeDecodeType(src.DecodeBcStatusRequestMessage),
		"decodeConsensusMessage":               encodeDecodeType(src.DecodeConsensusMessage),
		"decodeNewRoundStepMessage":            encodeDecodeType(src.DecodeNewRoundStepMessage),
		"decodeNewValidBlockMessage":           encodeDecodeType(src.DecodeNewValidBlockMessage),
		"decodeProposalMessage":                encodeDecodeType(src.DecodeProposalMessage),
		"decodeProposalPOLMessage":             encodeDecodeType(src.DecodeProposalPOLMessage),
		"decodeBlockPartMessage":               encodeDecodeType(src.DecodeBlockPartMessage),
		"decodeVoteMessage":                    encodeDecodeType(src.DecodeVoteMessage),
		"decodeHasVoteMessage":                 encodeDecodeType(src.DecodeHasVoteMessage),
		"decodeVoteSetMaj23Message":            encodeDecodeType(src.DecodeVoteSetMaj23Message),
		"decodeVoteSetBitsMessage":             encodeDecodeType(src.DecodeVoteSetBitsMessage),
		"decodeWALMessage":                     encodeDecodeType(src.DecodeWALMessage),
		"decodeMsgInfo":                        encodeDecodeType(src.DecodeMsgInfo),
		"decodeTimeoutInfo":                    encodeDecodeType(src.DecodeTimeoutInfo),
		"decodeEndHeightMessage":               encodeDecodeType(src.DecodeEndHeightMessage),
		"decodePubKey":                         encodeDecodeType(src.DecodePubKey),
		"decodePrivKey":                        encodeDecodeType(src.DecodePrivKey),
		"decodePubKeyEd25519":                  encodeDecodeType(src.DecodePubKeyEd25519),
		"decodePrivKeyEd25519":                 encodeDecodeType(src.DecodePrivKeyEd25519),
		"decodePubKeySecp256k1":                encodeDecodeType(src.DecodePubKeySecp256k1),
		"decodePrivKeySecp256k1":               encodeDecodeType(src.DecodePrivKeySecp256k1),
		"decodePubKeyMultisigThreshold":        encodeDecodeType(src.DecodePubKeyMultisigThreshold),
		"decodeEvidenceMessage":                encodeDecodeType(src.DecodeEvidenceMessage),
		"decodeEvidenceListMessage":            encodeDecodeType(src.DecodeEvidenceListMessage),
		"decodeMempoolMessage":                 encodeDecodeType(src.DecodeMempoolMessage),
		"decodeTxMessage":                      encodeDecodeType(src.DecodeTxMessage),
		"decodePacket":                         encodeDecodeType(src.DecodePacket),
		"decodePacketPing":                     encodeDecodeType(src.DecodePacketPing),
		"decodePacketPong":                     encodeDecodeType(src.DecodePacketPong),
		"decodePacketMsg":                      encodeDecodeType(src.DecodePacketMsg),
		"decodePexMessage":                     encodeDecodeType(src.DecodePexMessage),
		"decodePexRequestMessage":              encodeDecodeType(src.DecodePexRequestMessage),
		"decodePexAddrsMessage":                encodeDecodeType(src.DecodePexAddrsMessage),
		"decodeRemoteSignerMsg":                encodeDecodeType(src.DecodeRemoteSignerMsg),
		"decodePubKeyRequest":                  encodeDecodeType(src.DecodePubKeyRequest),
		"decodePubKeyResponse":                 encodeDecodeType(src.DecodePubKeyResponse),
		"decodeSignVoteRequest":                encodeDecodeType(src.DecodeSignVoteRequest),
		"decodeSignedVoteResponse":             encodeDecodeType(src.DecodeSignedVoteResponse),
		"decodeSignProposalRequest":            encodeDecodeType(src.DecodeSignProposalRequest),
		"decodeSignedProposalResponse":         encodeDecodeType(src.DecodeSignedProposalResponse),
		"decodePingRequest":                    encodeDecodeType(src.DecodePingRequest),
		"decodePingResponse":                   encodeDecodeType(src.DecodePingResponse),
		"decodeTMEventData":                    encodeDecodeType(src.DecodeTMEventData),
		"decodeEventDataNewBlock":              encodeDecodeType(src.DecodeEventDataNewBlock),
		"decodeEventDataNewBlockHeader":        encodeDecodeType(src.DecodeEventDataNewBlockHeader),
		"decodeEventDataTx":                    encodeDecodeType(src.DecodeEventDataTx),
		"decodeEventDataRoundState":            encodeDecodeType(src.DecodeEventDataRoundState),
		"decodeEventDataNewRound":              encodeDecodeType(src.DecodeEventDataNewRound),
		"decodeEventDataCompleteProposal":      encodeDecodeType(src.DecodeEventDataCompleteProposal),
		"decodeEventDataVote":                  encodeDecodeType(src.DecodeEventDataVote),
		"decodeEventDataValidatorSetUpdates":   encodeDecodeType(src.DecodeEventDataValidatorSetUpdates),
		"decodeEventDataString":                encodeDecodeType(src.DecodeEventDataString),
		"decodeEvidence":                       encodeDecodeType(src.DecodeEvidence),
		"decodeDuplicateVoteEvidence":          encodeDecodeType(src.DecodeDuplicateVoteEvidence),
		"decodeMockGoodEvidence":               encodeDecodeType(src.DecodeMockGoodEvidence),
		"decodeMockRandomGoodEvidence":         encodeDecodeType(src.DecodeMockRandomGoodEvidence),
		"decodeMockBadEvidence":                encodeDecodeType(src.DecodeMockBadEvidence),
		// @formatter:on
	})

	<-done
}
