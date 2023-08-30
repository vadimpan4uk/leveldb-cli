// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/liderman/leveldb-cli/cliutil"
)

/**
public class Block {

  private String id;
  private long number;
  private List<String> transactionIds;
}
 */
type Block struct {
	id string
	number int64
	transactionIds []byte
}

// The command get a value.
// It gets the value for the selected key.
//
// Returns a string containing information about the result of the operation.
func Get(key, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	if key == "" {
		return AppError(ErrKeyIsEmpty)
	}

	k, err := hex.DecodeString(key)
	if err != nil {
		return AppError(ErrKeyNotFound)
	}
	value, err := dbh.Get(k, nil)
	if err != nil {
		return AppError(ErrKeyNotFound)
	}

	buf := bytes.NewBuffer(value)
	block := &Block{}
	//if err := binary.Read(buf, binary.BigEndian, block); err != nil {
	//	fmt.Printf("err: %v", err)
	//	return AppError(ErrFailedDecodeBlock)
	//}
	if err := binary.Read(buf, binary.BigEndian, block); err != nil {
		fmt.Printf("err1: %v", err)
		//return AppError(ErrFailedDecodeBlock)
	}
	if err := binary.Read(buf, binary.LittleEndian, block); err != nil {
		fmt.Printf("err2: %v", err)
		return AppError(ErrFailedDecodeBlock)
	}
	fmt.Printf("Block: %v", block)

	return cliutil.ToString(format, value)
}
