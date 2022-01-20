//go:generate protoc -I=. -I=$GOPATH/src/github.com/ElrondNetwork/elrond-go-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/ElrondNetwork/protobuf/protobuf --gogoslick_out=. multipleHeaderSigningProof.proto
package slash

import (
	"encoding/hex"
	"fmt"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/core/sliceUtil"
	"github.com/ElrondNetwork/elrond-go-core/data"
)

// GetPubKeys - returns all validator's public keys which have signed multiple headers
func (m *MultipleHeaderSigningProof) GetPubKeys() [][]byte {
	if m == nil {
		return nil
	}

	ret := make([][]byte, 0, len(m.SignersSlashData))
	for pubKey := range m.SignersSlashData {
		ret = append(ret, []byte(pubKey))
	}

	return ret
}

// GetLevel returns the ThreatLevel of a possible malicious validator
func (m *MultipleHeaderSigningProof) GetLevel(pubKey []byte) ThreatLevel {
	if m == nil {
		return Zero
	}

	slashData, exists := m.SignersSlashData[string(pubKey)]
	if !exists {
		return Zero
	}

	return slashData.ThreatLevel
}

// GetAllHeaders returns all header handlers stored in the proof
func (m *MultipleHeaderSigningProof) GetAllHeaders() []data.HeaderHandler {
	if m == nil {
		return nil
	}

	return m.HeadersV2.GetHeaderHandlers()
}

// GetHeaders returns all headers that have been signed by a possible malicious validator
func (m *MultipleHeaderSigningProof) GetHeaders(pubKey []byte) []data.HeaderHandler {
	if m == nil {
		return nil
	}

	slashData, exists := m.SignersSlashData[string(pubKey)]
	if !exists {
		return nil
	}

	bitmap := slashData.GetSignedHeadersBitMap()
	headers := m.HeadersV2.GetHeaderHandlers()

	ret := make([]data.HeaderHandler, 0)
	for idx, header := range headers {
		if sliceUtil.IsIndexSetInBitmap(uint32(idx), bitmap) {
			ret = append(ret, header)
		}
	}

	return ret
}

// GetProofTxData returns the necessary ProofTxData to issue a commitment slash tx
func (m *MultipleHeaderSigningProof) GetProofTxData() (*ProofTxData, error) {
	if m == nil {
		return nil, data.ErrNilPointerReceiver
	}

	pubKeys := m.GetPubKeys()
	if len(pubKeys) == 0 {
		return nil, data.ErrNotEnoughPublicKeysProvided
	}
	pubKey := pubKeys[0]
	headers := m.GetHeaders(pubKey)
	if len(headers) == 0 {
		return nil, data.ErrNotEnoughHeadersProvided
	}
	if check.IfNil(headers[0]) {
		return nil, data.ErrNilHeaderHandler
	}

	return &ProofTxData{
		Round:   headers[0].GetRound(),
		ShardID: headers[0].GetShardID(),
		ProofID: MultipleSigningProofID,
	}, nil
}

// NewMultipleSigningProof returns a MultipleSigningProofHandler from a slashing result
func NewMultipleSigningProof(slashResult map[string]SlashingResult) (MultipleSigningProofHandler, error) {
	if slashResult == nil {
		return nil, data.ErrNilSlashResult
	}

	headersInfo, err := getAllUniqueHeaders(slashResult)
	if err != nil {
		return nil, err
	}
	sortedHeaders, err := sortAndGetHeadersV2(headersInfo)
	if err != nil {
		return nil, err
	}

	hashIndexMap := calcHashIndexMap(headersInfo)
	signersSlashData := computeSignersSlashData(hashIndexMap, slashResult)

	return &MultipleHeaderSigningProof{
		HeadersV2:        *sortedHeaders,
		SignersSlashData: signersSlashData,
	}, nil
}

func getAllUniqueHeaders(slashResult map[string]SlashingResult) ([]data.HeaderInfoHandler, error) {
	headersInfo := make([]data.HeaderInfoHandler, 0, len(slashResult))
	hashes := make(map[string]struct{})

	for pubKey, res := range slashResult {
		hashesPerPubKey := make(map[string]struct{})
		for _, headerInfo := range res.Headers {
			hash, err := getAndCheckUniqueHeaderHash(headerInfo, hashesPerPubKey)
			if err != nil {
				return nil, fmt.Errorf("%w in slash result for public key: %s", err, hex.EncodeToString([]byte(pubKey)))
			}

			hashesPerPubKey[hash] = struct{}{}
			_, exists := hashes[hash]
			if exists {
				continue
			}

			hashes[hash] = struct{}{}
			headersInfo = append(headersInfo, headerInfo)
		}
	}

	return headersInfo, nil
}

func calcHashIndexMap(headersInfo []data.HeaderInfoHandler) map[string]uint32 {
	hashIndexMap := make(map[string]uint32)
	for idx, headerInfo := range headersInfo {
		hashIndexMap[string(headerInfo.GetHash())] = uint32(idx)
	}

	return hashIndexMap
}

func computeSignersSlashData(hashIndexMap map[string]uint32, slashResult map[string]SlashingResult) map[string]SignerSlashingData {
	signersSlashData := make(map[string]SignerSlashingData)
	bitMapSize := CalcBitmapSize(len(hashIndexMap))
	for pubKey, res := range slashResult {
		bitmap := make([]byte, bitMapSize)
		for _, header := range res.Headers {
			index, exists := hashIndexMap[string(header.GetHash())]
			if exists {
				sliceUtil.SetIndexInBitmap(index, bitmap)
			}
		}
		signersSlashData[pubKey] = SignerSlashingData{
			SignedHeadersBitMap: bitmap,
			ThreatLevel:         res.SlashingLevel,
		}
	}

	return signersSlashData
}
