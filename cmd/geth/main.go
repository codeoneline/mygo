package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/codeoneline/mygo/common"
	"github.com/codeoneline/mygo/eth"
	"github.com/codeoneline/mygo/rlp"
	"github.com/syndtr/goleveldb/leveldb"
)

// epoch leader
type Proposer struct {
	PubSec256     []byte
	PubBn256      []byte
	Probabilities *big.Int
}

func (p *Proposer) toString() string {
	return hex.EncodeToString(p.PubBn256) + " " + hex.EncodeToString(p.PubBn256) + " " + p.Probabilities.String()
}

type Polynomial []big.Int
type PolyInfo struct {
	poly Polynomial
	s    *big.Int
}

type ClientIncentive struct {
	ValidatorAddr common.Address
	WalletAddr    common.Address
	Incentive     *big.Int
}

type RefundInfo struct {
	Addr   common.Address
	Amount *big.Int
}

type StoreType struct {
	Path string
	T    func(key string, value []byte) interface{}
}

func oneProposer(key string, value []byte) interface{} {
	proposer := Proposer{}
	err := rlp.DecodeBytes(value, &proposer)
	if err != nil {
		println("err", err)
		return nil
	}
	return proposer.toString()
}
func oneRB(key string, value []byte) interface{} {
	proposer := Proposer{}
	err := rlp.DecodeBytes(value, &proposer)
	if err != nil {
		println("err", err)
		return nil
	}
	return proposer.toString()
}

func onePolyInfo(key string, value []byte) interface{} {
	polyInfo := PolyInfo{}
	err := rlp.DecodeBytes(value, &polyInfo)
	if err != nil {
		println("err", err)
		return nil
	}
	return polyInfo
}

var gC = false

func oneIncentive(key string, value []byte) interface{} {
	if strings.Contains(key, "epoch_pay_detail") {
		if !gC {
			gC = true
			clientIncentive := [][]ClientIncentive{}
			err := rlp.DecodeBytes(value, &clientIncentive)
			if err != nil {
				println("err", err)
				return nil
			}
			println("len =", len(clientIncentive))
			k := 0
			for i := 0; i < len(clientIncentive); i++ {
				k += len(clientIncentive[i])
				println("len2 =", len(clientIncentive[i]), "total = ", k)
			}
			return clientIncentive
		}
		return ""
	} else if strings.Contains(key, "epoch_remain") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	} else if strings.Contains(key, "epoch_total") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	} else if strings.Contains(key, "epoch_block_number") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	} else if strings.Contains(key, "all_total") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	} else if strings.Contains(key, "total_remain") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	} else if strings.Contains(key, "run_times") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	}
	return nil
}

func oneSt(key string, value []byte) interface{} {
	return nil
}

func onePos(key string, value []byte) interface{} {
	if strings.Contains(key, "StakeOutEpochKey") {
		stakeInfo := []RefundInfo{}
		err := rlp.DecodeBytes(value, &stakeInfo)
		if err != nil {
			println("err", err)
			return nil
		}
		return stakeInfo
	} else if strings.Contains(key, "SLOT_LEADER_SC_CALL_TIMES") {
		o := big.Int{}
		o.SetBytes(value)
		return o.String()
	}
	return nil
}

func oneAvg(key string, value []byte) interface{} {
	return nil
}

func oneFork(key string, value []byte) interface{} {
	return nil
}

func main() {
	println("Hello golang")
	eth.Sync()

	lds := []StoreType{
		{"/town/.wanchain/2.2-dev/gwan/eplocaldb", oneProposer},
		{"/town/.wanchain/2.2-dev/gwan/rblocaldb", oneRB},
		{"/town/.wanchain/2.2-dev/gwan/pos", onePos},
		{"/town/.wanchain/2.2-dev/gwan/incentive", oneIncentive},
		// {"/town/.wanchain/2.2-dev/gwan/stlocaldb", oneSt},
		// {"/town/.wanchain/2.2-dev/gwan/avgretdb", oneAvg},
		// {"/town/.wanchain/2.2-dev/gwan/forkdb", oneFork},
	}

	for i := 0; i < len(lds); i++ {
		ld := lds[i]

		db, err := leveldb.OpenFile(ld.Path, nil)

		fmt.Println("****", ld.Path)

		if err != nil {
			println("err", err)
			return
		}

		defer db.Close()

		iter := db.NewIterator(nil, nil)
		for iter.Next() {
			// Remember that the contents of the returned slice should not be modified, and
			// only valid until the next call to Next.
			key := iter.Key()
			value := iter.Value()
			// println("key = ", key, "value = ", value)

			if strings.Contains(string(key), "keyCount") {
				fmt.Printf("key: %s , value: %d\n", key, value[0])
			} else if strings.Contains(string(key), "_key_") {
				fmt.Printf("key: %s , value: %s\n", key, value)
			} else {
				proposer := ld.T(string(key), value)
				if proposer != nil {
					fmt.Printf("key: %s , value: %s \n", key, proposer)
				} else {
					fmt.Printf("key: %s , value: failed to parse \n", key)
				}
			}
		}
	}

}
