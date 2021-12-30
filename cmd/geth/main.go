package main

import (
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

type StoreType struct {
	Path string
	T    func() interface{}
}

func oneProposer() interface{} {
	return Proposer{}
}

func onePolyInfo() interface{} {
	return PolyInfo{}
}

func oneClientIncentive() interface {
}

func main() {
	println("Hello golang")
	eth.Sync()

	lds := []StoreType{
		// {"/town/.wanchain/2.2-dev-bk/gwan/eplocaldb", oneProposer},
		// {"/town/.wanchain/2.2-dev-bk/gwan/rblocaldb", onePolyInfo},
		{"/town/.wanchain/2.2-dev-bk/gwan/incentive", oneProposer},
		// {"/town/.wanchain/2.2-dev-bk/gwan/pos", oneProposer},
		// {"/town/.wanchain/2.2-dev-bk/gwan/avgretdb", oneProposer},
	}

	for i := 0; i < len(lds); i++ {
		ld := lds[i]

		db, err := leveldb.OpenFile(ld.Path, nil)

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
			} else if len(value) < 10 {
				fmt.Printf("key: %s , value: %s\n", key, value)
			} else {
				proposer := ld.T()
				err := rlp.DecodeBytes(value, &proposer)
				if err != nil {
					println("err", err)
					continue
				}
				fmt.Printf("key: %s , value: %x\n", key, proposer)
			}
		}
	}

}
