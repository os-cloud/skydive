/*
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package flow

import (
	"encoding/binary"
	"hash"
	"strings"
)

func (fl *ICMPLayer) Hash(hasher hash.Hash) {
	if fl == nil {
		return
	}

	value32 := make([]byte, 4)
	binary.BigEndian.PutUint32(value32, uint32(fl.Type)<<24|uint32(fl.Code<<16|uint32(fl.ID)))
	hasher.Write(value32)
}

// Hash calculates a unique symetric flow layer hash
func (fl *FlowLayer) Hash(hasher hash.Hash) {
	if fl == nil {
		return
	}

	if strings.Compare(fl.A, fl.B) > 0 {
		hasher.Write([]byte(fl.A))
		hasher.Write([]byte(fl.B))
	} else {
		hasher.Write([]byte(fl.B))
		hasher.Write([]byte(fl.A))
	}

	value64 := make([]byte, 8)
	binary.BigEndian.PutUint64(value64, uint64(fl.ID))
	hasher.Write(value64)
}
