package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/OpenBazaar/openbazaar-go/pb"
	"github.com/OpenBazaar/openbazaar-go/repo"
)

// MessagesDB represents the messages table
type MessagesDB struct {
	modelStore
}

// NewMessageStore return new MessagesDB
func NewMessageStore(db *sql.DB, lock *sync.Mutex) repo.MessageStore {
	return &MessagesDB{modelStore{db, lock}}
}

// Put will insert a record into the messages
func (o *MessagesDB) Put(messageID, orderID string, mType pb.Message_MessageType, peerID string, msg repo.Message, rErr string, receivedAt int64, pubkey []byte) error {
	o.lock.Lock()
	defer o.lock.Unlock()

	tx, err := o.db.Begin()
	if err != nil {
		return err
	}
	stm := `insert or replace into messages(messageID, orderID, message_type, message, peerID, err, received_at, pubkey, created_at) values(?,?,?,?,?,?,?,?,?)`
	stmt, err := tx.Prepare(stm)
	if err != nil {
		return err
	}

	msg0, err := msg.MarshalJSON()
	if err != nil {
		log.Errorf("err marshalling json: %v", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		messageID,
		orderID,
		int(mType),
		msg0,
		peerID,
		rErr,
		receivedAt,
		pubkey,
		time.Now().Unix(),
	)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			return fmt.Errorf("message put fail: %s and rollback failed: %s", err.Error(), rErr.Error())
		}
		return err
	}

	return tx.Commit()
}

// GetByOrderIDType returns the message for the specified order and message type
func (o *MessagesDB) GetByOrderIDType(orderID string, mType pb.Message_MessageType) (*repo.Message, string, string, error) {
	o.lock.Lock()
	defer o.lock.Unlock()
	var (
		msg0   []byte
		peerID string
		recErr string
	)

	stmt, err := o.db.Prepare("select message, peerID, err from messages where orderID=? and message_type=?")
	if err != nil {
		return nil, "", "", err
	}
	err = stmt.QueryRow(orderID, mType).Scan(&msg0, &peerID, &recErr)
	if err != nil {
		return nil, "", "", err
	}

	msg := new(repo.Message)

	if len(msg0) > 0 {
		err = msg.UnmarshalJSON(msg0)
		if err != nil {
			return nil, "", "", err
		}
	}

	return msg, recErr, peerID, nil
}

func (o *MessagesDB) GetAllErrored() ([]repo.OrderMessage, error) {
	o.lock.Lock()
	defer o.lock.Unlock()

	/*
		q := query{
			table:   "messages",
			columns: []string{"messageID", "orderID", "message_type", "message", "peerID", "err", "pubkey"},
			id:      "messageID",
		}
		stm, args := filterQuery(q)
	*/

	stmt := "select messageID, orderID, message_type, message, peerID, err, pubkey from messages where err!=? "
	var ret []repo.OrderMessage
	rows, err := o.db.Query(stmt, "")
	if err != nil {
		return ret, err
	}
	defer rows.Close()

	for rows.Next() {
		var messageID, orderID, peerID, rErr string
		var msg0, pkey []byte
		var mType int32
		err = rows.Scan(&messageID, &orderID, &mType, &msg0, &peerID, &rErr, &pkey)
		if err != nil {
			log.Error(err)
		}
		ret = append(ret, repo.OrderMessage{
			PeerID:      peerID,
			MessageID:   messageID,
			OrderID:     orderID,
			MessageType: mType,
			Message:     msg0,
			MsgErr:      rErr,
			PeerPubkey:  pkey,
		})
	}

	/*
		var ret []repo.OrderMessage
		for rows.Next() {
			var messageID, orderID, peerID, rErr string
			var mType pb.Message_MessageType
			var message, pubkey []byte
			if err := rows.Scan(&messageID, &orderID, &mType, &message, &peerID, &rErr, &pubkey); err != nil {
				return ret, err
			}
			msg := repo.OrderMessage{
				MessageID:   messageID,
				OrderID:     orderID,
				MessageType: int32(mType),
				MsgErr:      rErr,
				PeerID:      peerID,
				Message:     message,
				PeerPubkey:  pubkey,
			}

			ret = append(ret, msg)
		}
	*/
	return ret, nil
}
