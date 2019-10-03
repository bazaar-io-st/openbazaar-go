package core

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	util "gx/ipfs/QmNohiVssaPw3KVLZik59DBVGTSm2dGvYT9eoXt5DQ36Yz/go-ipfs-util"

	"github.com/OpenBazaar/openbazaar-go/pb"
	"github.com/OpenBazaar/openbazaar-go/repo"
	"github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// FormatRFC3339PB returns the given `google_protobuf.Timestamp` as a RFC3339
// formatted string
func FormatRFC3339PB(ts google_protobuf.Timestamp) string {
	return util.FormatRFC3339(time.Unix(ts.Seconds, int64(ts.Nanos)).UTC())
}

// BuildTransactionRecords - Used by the GET order API to build transaction records suitable to be included in the order response
func (n *OpenBazaarNode) BuildTransactionRecords(contract *pb.RicardianContract, records []*wallet.TransactionRecord, state pb.OrderState) ([]*pb.TransactionRecord, *pb.TransactionRecord, error) {
	paymentRecords := []*pb.TransactionRecord{}
	payments := make(map[string]*pb.TransactionRecord)
	order, err := repo.ToV5Order(contract.BuyerOrder, n.LookupCurrency)
	if err != nil {
		return nil, nil, err
	}
	wal, err := n.Multiwallet.WalletForCurrencyCode(order.Payment.AmountCurrency.Code)
	if err != nil {
		return paymentRecords, nil, err
	}

	// Consolidate any transactions with multiple outputs into a single record
	for _, r := range records {
		record, ok := payments[r.Txid]
		if ok {
			n, _ := new(big.Int).SetString(record.BigValue, 10)
			sum := new(big.Int).Add(n, &r.Value)
			record.BigValue = sum.String()
			payments[r.Txid] = record
		} else {
			tx := new(pb.TransactionRecord)
			tx.Txid = r.Txid
			tx.BigValue = r.Value.String()
			tx.Currency = order.Payment.AmountCurrency

			ts, err := ptypes.TimestampProto(r.Timestamp)
			if err != nil {
				return paymentRecords, nil, err
			}
			tx.Timestamp = ts
			ch, err := chainhash.NewHashFromStr(strings.TrimPrefix(tx.Txid, "0x"))
			if err != nil {
				return paymentRecords, nil, err
			}
			confirmations, height, err := wal.GetConfirmations(*ch)
			if err != nil {
				return paymentRecords, nil, err
			}
			tx.Height = height
			tx.Confirmations = confirmations
			payments[r.Txid] = tx
		}
	}
	for _, rec := range payments {
		paymentRecords = append(paymentRecords, rec)
	}
	var refundRecord *pb.TransactionRecord
	if contract != nil && (state == pb.OrderState_REFUNDED || state == pb.OrderState_DECLINED || state == pb.OrderState_CANCELED) && order != nil && order.Payment != nil {
		// For multisig we can use the outgoing from the payment address
		if order.Payment.Method == pb.Order_Payment_MODERATED || state == pb.OrderState_DECLINED || state == pb.OrderState_CANCELED {
			for _, rec := range payments {
				val, _ := new(big.Int).SetString(rec.BigValue, 10)
				if val.Cmp(big.NewInt(0)) < 0 {
					refundRecord = new(pb.TransactionRecord)
					refundRecord.Txid = rec.Txid
					refundRecord.BigValue = "-" + rec.BigValue
					refundRecord.Currency = rec.Currency
					refundRecord.Confirmations = rec.Confirmations
					refundRecord.Height = rec.Height
					refundRecord.Timestamp = rec.Timestamp
					break
				}
			}
		} else if contract.Refund != nil && contract.Refund.RefundTransaction != nil && contract.Refund.Timestamp != nil {
			refundRecord = new(pb.TransactionRecord)
			// Direct we need to use the transaction info in the contract's refund object
			ch, err := chainhash.NewHashFromStr(strings.TrimPrefix(contract.Refund.RefundTransaction.Txid, "0x"))
			if err != nil {
				return paymentRecords, refundRecord, err
			}
			confirmations, height, err := wal.GetConfirmations(*ch)
			if err != nil {
				return paymentRecords, refundRecord, nil
			}
			refundRecord.Txid = contract.Refund.RefundTransaction.Txid
			refundRecord.BigValue = contract.Refund.RefundTransaction.BigValue
			refundRecord.Currency = contract.Refund.RefundTransaction.ValueCurrency
			refundRecord.Timestamp = contract.Refund.Timestamp
			refundRecord.Confirmations = confirmations
			refundRecord.Height = height
		}
	}
	return paymentRecords, refundRecord, nil
}

// LookupCurrency looks up the CurrencyDefinition, first by crypto for the current network
// (mainnet or testnet) and then by fiat code
func (n *OpenBazaarNode) LookupCurrency(currencyCode string) (repo.CurrencyDefinition, error) {
	if n.TestnetEnable || n.RegressionTestEnable {
		if def, err := repo.TestnetCurrencies().Lookup(currencyCode); err == nil {
			return def, nil
		}
	} else {
		if def, err := repo.MainnetCurrencies().Lookup(currencyCode); err == nil {
			return def, nil
		}
	}
	return repo.FiatCurrencies().Lookup(currencyCode)
}

// exchangeRateCode strips the T off the currency code if we are on testnet or regtest.
func (n *OpenBazaarNode) exchangeRateCode(currencyCode string) string {
	if n.TestnetEnable || n.RegressionTestEnable {
		return strings.TrimPrefix(currencyCode, "T")
	}
	return currencyCode
}

func (n *OpenBazaarNode) ValidateMultiwalletHasPreferredCurrencies(data repo.SettingsData) error {
	if data.PreferredCurrencies == nil {
		return nil
	}
	for _, cc := range *data.PreferredCurrencies {
		_, err := n.Multiwallet.WalletForCurrencyCode(cc)
		if err != nil {
			return fmt.Errorf("preferred coin %s not found in multiwallet", cc)
		}
	}
	return nil
}
