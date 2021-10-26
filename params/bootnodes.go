// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://a96d4305f24a1f9c98abd315b87076f46c26f6b61a7816fcbfe023899d788ea7a93f4709a3fcd8d373d1260214884a71643afe3026c0aeb21e264995bf776332@18.178.150.45:30303",
	"enode://3b6471f7e03e2df7708587d9e6c7f4d066e07d763a22b19e4875176a7f4ca915ca63d8b1ed39fe74e39b0a32d1c5ac14d80a81d8005c6b635d0917d7e35aef4f@54.95.233.91:30303",
	"enode://b57061828bc1bd4839a9e70bc99be4d642d9c95537338aa9c794d36fa2977f0e6c39fdccbcd5672581deaddd20ac4e228b5b71e535cb51c5ad53b5f0a401ab08@54.199.251.38:30303",
	"enode://9ee8b1bb71fadcc53410a3810b8176ff8ab517f1939d1afebed29dabf53bf39951efeb5a4265f7e3f1ee89c0686f61d240bf50027fdb505addc02186885269a1@13.231.144.5:30303",
	"enode://b57061828bc1bd4839a9e70bc99be4d642d9c95537338aa9c794d36fa2977f0e6c39fdccbcd5672581deaddd20ac4e228b5b71e535cb51c5ad53b5f0a401ab08@54.199.251.38:30303",
}
var TestnetBootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
