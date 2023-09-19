// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package core

import (
	"pkg.berachain.dev/polaris/eth/core/types"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	InsertBlock(block *types.Block, receipts types.Receipts, logs []*types.Log) error
}

// InsertBlock inserts a block into the canonical chain and updates the state of the blockchain.
func (bc *blockchain) InsertBlock(block *types.Block, receipts types.Receipts, logs []*types.Log) error {
	var err error

	// TODO: prepare historical plugin here?
	// TBH still think we should deprecate it and run in another routine as indexer.

	// ***************************************** //
	// TODO: add safety check for canonicallness //
	// ***************************************** //

	// *********************************************** //
	// TODO: restructure this function / flow it sucks //
	// *********************************************** //
	blockHash, blockNum := block.Hash(), block.Number().Uint64()
	bc.logger.Info("finalizing evm block", "block_hash", blockHash.Hex(), "num_txs", len(receipts))

	// store the block header on the host chain
	err = bc.bp.StoreHeader(block.Header())
	if err != nil {
		bc.logger.Error("failed to store block header", "err", err)
		return err
	}

	// store the block, receipts, and txs on the host chain if historical plugin is supported
	if bc.hp != nil {
		if err = bc.hp.StoreBlock(block); err != nil {
			bc.logger.Error("failed to store block", "err", err)
			return err
		}
		if err = bc.hp.StoreReceipts(blockHash, receipts); err != nil {
			bc.logger.Error("failed to store receipts", "err", err)
			return err
		}
		if err = bc.hp.StoreTransactions(blockNum, blockHash, block.Transactions()); err != nil {
			bc.logger.Error("failed to store transactions", "err", err)
			return err
		}
	}

	// mark the current block, receipts, and logs
	if block != nil {
		bc.currentBlock.Store(block)
		bc.finalizedBlock.Store(block)

		bc.blockNumCache.Add(blockNum, block)
		bc.blockHashCache.Add(blockHash, block)

		for txIndex, tx := range block.Transactions() {
			bc.txLookupCache.Add(
				tx.Hash(),
				&types.TxLookupEntry{
					Tx:        tx,
					TxIndex:   uint64(txIndex),
					BlockNum:  blockNum,
					BlockHash: blockHash,
				},
			)
		}
	}
	if receipts != nil {
		bc.currentReceipts.Store(receipts)
		bc.receiptsCache.Add(blockHash, receipts)
	}
	if logs != nil {
		bc.pendingLogsFeed.Send(logs)
		bc.currentLogs.Store(logs)
		if len(logs) > 0 {
			bc.logsFeed.Send(logs)
		}
	}

	// Send chain events.
	bc.chainFeed.Send(ChainEvent{Block: block, Hash: blockHash, Logs: logs})
	bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})

	return nil
}
