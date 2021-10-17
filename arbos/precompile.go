//
// Copyright 2021, Offchain Labs, Inc. All rights reserved.
//

package arbos

import (
	"log"
	"math/big"
	"reflect"
	"strings"
	"unicode"

	pre "github.com/offchainlabs/arbstate/arbos/precompiles"
	templates "github.com/offchainlabs/arbstate/precompiles/go"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
)

type ArbosPrecompile interface {
	GasToCharge(input []byte) uint64

	// Important fields: evm.StateDB and evm.Config.Tracer
	// NOTE: if precompileAddress != actingAsAddress, watch out! This is a delegatecall or callcode, so caller might be wrong. In that case, unless this precompile is pure, it should probably revert.
	Call(
		input []byte,
		precompileAddress common.Address,
		actingAsAddress common.Address,
		caller common.Address,
		value *big.Int,
		readOnly bool,
		evm *vm.EVM,
	) (output []byte, err error)
}

type Precompile struct {
	methods map[[4]byte]PrecompileMethod
}

type PrecompileMethod struct {
	name    string
	handler reflect.Method
	gascost reflect.Method
}

// Make a precompile for the given hardhat-to-geth bindings, ensuring that the implementer
// supports each method.
func makePrecompile(metadata *bind.MetaData, implementer interface{}) ArbosPrecompile {
	source, err := abi.JSON(strings.NewReader(metadata.ABI))
	if err != nil {
		log.Fatal("Bad ABI")
	}

	contract := reflect.TypeOf(implementer).Name()
	methods := make(map[[4]byte]PrecompileMethod)

	for _, method := range source.Methods {

		name := method.RawName
		capitalize := string(unicode.ToUpper(rune(name[0])))
		name = capitalize + name[1:]
		context := "Precompile " + contract + "'s " + name + "'s implementer "

		if len(method.ID) != 4 {
			log.Fatal("Method ID isn't 4 bytes")
		}
		id := *(*[4]byte)(method.ID)

		// check that the implementer has a supporting implementation for this method

		handler, ok := reflect.TypeOf(implementer).MethodByName(name)
		if !ok {
			log.Fatal("Precompile ", contract, " must implement ", name)
		}

		var needs = []reflect.Type{
			reflect.TypeOf(implementer),      // the contract itself
			reflect.TypeOf(common.Address{}), // the method's caller
		}

		switch method.StateMutability {
		case "pure":
		case "view":
			needs = append(needs, reflect.TypeOf(&state.StateDB{}))
		case "nonpayable":
			needs = append(needs, reflect.TypeOf(&state.StateDB{}))
		case "payable":
			needs = append(needs, reflect.TypeOf(&state.StateDB{}))
			needs = append(needs, reflect.TypeOf(&big.Int{}))
		default:
			log.Fatal("Unknown state mutability ", method.StateMutability)
		}

		for _, arg := range method.Inputs {
			needs = append(needs, arg.Type.GetType())
		}

		signature := handler.Type

		if signature.NumIn() != len(needs) {
			log.Fatal(context, "doesn't have the args\n\t", needs)
		}
		for i, arg := range needs {
			if signature.In(i) != arg {
				log.Fatal(
					context, "doesn't have the args\n\t", needs, "\n",
					"\tArg ", i, " is ", signature.In(i), " instead of ", arg,
				)
			}
		}

		var outputs = []reflect.Type{}
		for _, out := range method.Outputs {
			outputs = append(outputs, out.Type.GetType())
		}
		outputs = append(outputs, reflect.TypeOf((*error)(nil)).Elem())

		if signature.NumOut() != len(outputs) {
			log.Fatal("Precompile ", contract, "'s ", name, " implementer doesn't return ", outputs)
		}
		for i, out := range outputs {
			if signature.Out(i) != out {
				log.Fatal(
					context, "doesn't have the outputs\n\t", outputs, "\n",
					"\tReturn value ", i+1, " is ", signature.Out(i), " instead of ", out,
				)
			}
		}

		// ensure we have a matching gascost func

		gascost, ok := reflect.TypeOf(implementer).MethodByName(name + "GasCost")
		if !ok {
			log.Fatal("Precompile ", contract, " must implement ", name+"GasCost")
		}

		needs = []reflect.Type{
			reflect.TypeOf(implementer), // the contract itself
		}
		for _, arg := range method.Inputs {
			needs = append(needs, arg.Type.GetType())
		}

		signature = gascost.Type
		context = "Precompile " + contract + "'s " + name + "GasCost's implementer "

		if signature.NumIn() != len(needs) {
			log.Fatal(context, "doesn't have the args\n\t", needs)
		}
		for i, arg := range needs {
			if signature.In(i) != arg {
				log.Fatal(
					context, "doesn't have the args\n\t", needs, "\n",
					"\tArg ", i, " is ", signature.In(i), " instead of ", arg,
				)
			}
		}
		if signature.NumOut() != 1 || signature.Out(0) != reflect.TypeOf(&big.Int{}) {
			log.Fatal(context, "must return a *big.Int")
		}

		methods[id] = PrecompileMethod{
			name,
			handler,
			gascost,
		}
	}

	return Precompile{
		methods,
	}
}

func Precompiles() map[common.Address]ArbosPrecompile {
	return map[common.Address]ArbosPrecompile{
		addr("0x065"): makePrecompile(templates.ArbInfoMetaData, pre.ArbInfo{}),
		addr("0x100"): makePrecompile(templates.ArbSysMetaData, pre.ArbSys{}),
		addr("0x102"): makePrecompile(templates.ArbAddressTableMetaData, pre.ArbAddressTable{}),
		addr("0x103"): makePrecompile(templates.ArbBLSMetaData, pre.ArbBLS{}),
		addr("0x104"): makePrecompile(templates.ArbFunctionTableMetaData, pre.ArbFunctionTable{}),
		addr("0x105"): makePrecompile(templates.ArbosTestMetaData, pre.ArbosTest{}),
		addr("0x107"): makePrecompile(templates.ArbOwnerMetaData, pre.ArbOwner{}),
		addr("0x108"): makePrecompile(templates.ArbGasInfoMetaData, pre.ArbGasInfo{}),
		addr("0x109"): makePrecompile(templates.ArbAggregatorMetaData, pre.ArbAggregator{}),
		addr("0x110"): makePrecompile(templates.ArbRetryableTxMetaData, pre.ArbRetryableTx{}),
		addr("0x111"): makePrecompile(templates.ArbStatisticsMetaData, pre.ArbStatistics{}),
	}
}

func addr(s string) common.Address {
	return common.HexToAddress(s)
}

func (p Precompile) GasToCharge(input []byte) uint64 {
	return 0
}

func (p Precompile) Call(
	input []byte,
	precompileAddress common.Address,
	actingAsAddress common.Address,
	caller common.Address,
	value *big.Int,
	readOnly bool,
	evm *vm.EVM,
) (output []byte, err error) {
	return nil, nil
}
