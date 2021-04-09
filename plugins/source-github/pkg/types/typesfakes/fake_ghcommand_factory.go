// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/maximilien/kn-source-github/pkg/types"
	typesa "github.com/maximilien/kn-source-pkg/pkg/types"
	"github.com/spf13/cobra"
)

type FakeGHCommandFactory struct {
	CreateCommandStub        func() *cobra.Command
	createCommandMutex       sync.RWMutex
	createCommandArgsForCall []struct {
	}
	createCommandReturns struct {
		result1 *cobra.Command
	}
	createCommandReturnsOnCall map[int]struct {
		result1 *cobra.Command
	}
	DeleteCommandStub        func() *cobra.Command
	deleteCommandMutex       sync.RWMutex
	deleteCommandArgsForCall []struct {
	}
	deleteCommandReturns struct {
		result1 *cobra.Command
	}
	deleteCommandReturnsOnCall map[int]struct {
		result1 *cobra.Command
	}
	DescribeCommandStub        func() *cobra.Command
	describeCommandMutex       sync.RWMutex
	describeCommandArgsForCall []struct {
	}
	describeCommandReturns struct {
		result1 *cobra.Command
	}
	describeCommandReturnsOnCall map[int]struct {
		result1 *cobra.Command
	}
	GHSourceFactoryStub        func() types.GHSourceFactory
	gHSourceFactoryMutex       sync.RWMutex
	gHSourceFactoryArgsForCall []struct {
	}
	gHSourceFactoryReturns struct {
		result1 types.GHSourceFactory
	}
	gHSourceFactoryReturnsOnCall map[int]struct {
		result1 types.GHSourceFactory
	}
	KnSourceFactoryStub        func() typesa.KnSourceFactory
	knSourceFactoryMutex       sync.RWMutex
	knSourceFactoryArgsForCall []struct {
	}
	knSourceFactoryReturns struct {
		result1 typesa.KnSourceFactory
	}
	knSourceFactoryReturnsOnCall map[int]struct {
		result1 typesa.KnSourceFactory
	}
	SourceCommandStub        func() *cobra.Command
	sourceCommandMutex       sync.RWMutex
	sourceCommandArgsForCall []struct {
	}
	sourceCommandReturns struct {
		result1 *cobra.Command
	}
	sourceCommandReturnsOnCall map[int]struct {
		result1 *cobra.Command
	}
	UpdateCommandStub        func() *cobra.Command
	updateCommandMutex       sync.RWMutex
	updateCommandArgsForCall []struct {
	}
	updateCommandReturns struct {
		result1 *cobra.Command
	}
	updateCommandReturnsOnCall map[int]struct {
		result1 *cobra.Command
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGHCommandFactory) CreateCommand() *cobra.Command {
	fake.createCommandMutex.Lock()
	ret, specificReturn := fake.createCommandReturnsOnCall[len(fake.createCommandArgsForCall)]
	fake.createCommandArgsForCall = append(fake.createCommandArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateCommand", []interface{}{})
	fake.createCommandMutex.Unlock()
	if fake.CreateCommandStub != nil {
		return fake.CreateCommandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createCommandReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) CreateCommandCallCount() int {
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	return len(fake.createCommandArgsForCall)
}

func (fake *FakeGHCommandFactory) CreateCommandCalls(stub func() *cobra.Command) {
	fake.createCommandMutex.Lock()
	defer fake.createCommandMutex.Unlock()
	fake.CreateCommandStub = stub
}

func (fake *FakeGHCommandFactory) CreateCommandReturns(result1 *cobra.Command) {
	fake.createCommandMutex.Lock()
	defer fake.createCommandMutex.Unlock()
	fake.CreateCommandStub = nil
	fake.createCommandReturns = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) CreateCommandReturnsOnCall(i int, result1 *cobra.Command) {
	fake.createCommandMutex.Lock()
	defer fake.createCommandMutex.Unlock()
	fake.CreateCommandStub = nil
	if fake.createCommandReturnsOnCall == nil {
		fake.createCommandReturnsOnCall = make(map[int]struct {
			result1 *cobra.Command
		})
	}
	fake.createCommandReturnsOnCall[i] = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) DeleteCommand() *cobra.Command {
	fake.deleteCommandMutex.Lock()
	ret, specificReturn := fake.deleteCommandReturnsOnCall[len(fake.deleteCommandArgsForCall)]
	fake.deleteCommandArgsForCall = append(fake.deleteCommandArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteCommand", []interface{}{})
	fake.deleteCommandMutex.Unlock()
	if fake.DeleteCommandStub != nil {
		return fake.DeleteCommandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCommandReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) DeleteCommandCallCount() int {
	fake.deleteCommandMutex.RLock()
	defer fake.deleteCommandMutex.RUnlock()
	return len(fake.deleteCommandArgsForCall)
}

func (fake *FakeGHCommandFactory) DeleteCommandCalls(stub func() *cobra.Command) {
	fake.deleteCommandMutex.Lock()
	defer fake.deleteCommandMutex.Unlock()
	fake.DeleteCommandStub = stub
}

func (fake *FakeGHCommandFactory) DeleteCommandReturns(result1 *cobra.Command) {
	fake.deleteCommandMutex.Lock()
	defer fake.deleteCommandMutex.Unlock()
	fake.DeleteCommandStub = nil
	fake.deleteCommandReturns = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) DeleteCommandReturnsOnCall(i int, result1 *cobra.Command) {
	fake.deleteCommandMutex.Lock()
	defer fake.deleteCommandMutex.Unlock()
	fake.DeleteCommandStub = nil
	if fake.deleteCommandReturnsOnCall == nil {
		fake.deleteCommandReturnsOnCall = make(map[int]struct {
			result1 *cobra.Command
		})
	}
	fake.deleteCommandReturnsOnCall[i] = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) DescribeCommand() *cobra.Command {
	fake.describeCommandMutex.Lock()
	ret, specificReturn := fake.describeCommandReturnsOnCall[len(fake.describeCommandArgsForCall)]
	fake.describeCommandArgsForCall = append(fake.describeCommandArgsForCall, struct {
	}{})
	fake.recordInvocation("DescribeCommand", []interface{}{})
	fake.describeCommandMutex.Unlock()
	if fake.DescribeCommandStub != nil {
		return fake.DescribeCommandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.describeCommandReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) DescribeCommandCallCount() int {
	fake.describeCommandMutex.RLock()
	defer fake.describeCommandMutex.RUnlock()
	return len(fake.describeCommandArgsForCall)
}

func (fake *FakeGHCommandFactory) DescribeCommandCalls(stub func() *cobra.Command) {
	fake.describeCommandMutex.Lock()
	defer fake.describeCommandMutex.Unlock()
	fake.DescribeCommandStub = stub
}

func (fake *FakeGHCommandFactory) DescribeCommandReturns(result1 *cobra.Command) {
	fake.describeCommandMutex.Lock()
	defer fake.describeCommandMutex.Unlock()
	fake.DescribeCommandStub = nil
	fake.describeCommandReturns = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) DescribeCommandReturnsOnCall(i int, result1 *cobra.Command) {
	fake.describeCommandMutex.Lock()
	defer fake.describeCommandMutex.Unlock()
	fake.DescribeCommandStub = nil
	if fake.describeCommandReturnsOnCall == nil {
		fake.describeCommandReturnsOnCall = make(map[int]struct {
			result1 *cobra.Command
		})
	}
	fake.describeCommandReturnsOnCall[i] = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) GHSourceFactory() types.GHSourceFactory {
	fake.gHSourceFactoryMutex.Lock()
	ret, specificReturn := fake.gHSourceFactoryReturnsOnCall[len(fake.gHSourceFactoryArgsForCall)]
	fake.gHSourceFactoryArgsForCall = append(fake.gHSourceFactoryArgsForCall, struct {
	}{})
	fake.recordInvocation("GHSourceFactory", []interface{}{})
	fake.gHSourceFactoryMutex.Unlock()
	if fake.GHSourceFactoryStub != nil {
		return fake.GHSourceFactoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gHSourceFactoryReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) GHSourceFactoryCallCount() int {
	fake.gHSourceFactoryMutex.RLock()
	defer fake.gHSourceFactoryMutex.RUnlock()
	return len(fake.gHSourceFactoryArgsForCall)
}

func (fake *FakeGHCommandFactory) GHSourceFactoryCalls(stub func() types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = stub
}

func (fake *FakeGHCommandFactory) GHSourceFactoryReturns(result1 types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = nil
	fake.gHSourceFactoryReturns = struct {
		result1 types.GHSourceFactory
	}{result1}
}

func (fake *FakeGHCommandFactory) GHSourceFactoryReturnsOnCall(i int, result1 types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = nil
	if fake.gHSourceFactoryReturnsOnCall == nil {
		fake.gHSourceFactoryReturnsOnCall = make(map[int]struct {
			result1 types.GHSourceFactory
		})
	}
	fake.gHSourceFactoryReturnsOnCall[i] = struct {
		result1 types.GHSourceFactory
	}{result1}
}

func (fake *FakeGHCommandFactory) KnSourceFactory() typesa.KnSourceFactory {
	fake.knSourceFactoryMutex.Lock()
	ret, specificReturn := fake.knSourceFactoryReturnsOnCall[len(fake.knSourceFactoryArgsForCall)]
	fake.knSourceFactoryArgsForCall = append(fake.knSourceFactoryArgsForCall, struct {
	}{})
	fake.recordInvocation("KnSourceFactory", []interface{}{})
	fake.knSourceFactoryMutex.Unlock()
	if fake.KnSourceFactoryStub != nil {
		return fake.KnSourceFactoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.knSourceFactoryReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) KnSourceFactoryCallCount() int {
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	return len(fake.knSourceFactoryArgsForCall)
}

func (fake *FakeGHCommandFactory) KnSourceFactoryCalls(stub func() typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = stub
}

func (fake *FakeGHCommandFactory) KnSourceFactoryReturns(result1 typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	fake.knSourceFactoryReturns = struct {
		result1 typesa.KnSourceFactory
	}{result1}
}

func (fake *FakeGHCommandFactory) KnSourceFactoryReturnsOnCall(i int, result1 typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	if fake.knSourceFactoryReturnsOnCall == nil {
		fake.knSourceFactoryReturnsOnCall = make(map[int]struct {
			result1 typesa.KnSourceFactory
		})
	}
	fake.knSourceFactoryReturnsOnCall[i] = struct {
		result1 typesa.KnSourceFactory
	}{result1}
}

func (fake *FakeGHCommandFactory) SourceCommand() *cobra.Command {
	fake.sourceCommandMutex.Lock()
	ret, specificReturn := fake.sourceCommandReturnsOnCall[len(fake.sourceCommandArgsForCall)]
	fake.sourceCommandArgsForCall = append(fake.sourceCommandArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceCommand", []interface{}{})
	fake.sourceCommandMutex.Unlock()
	if fake.SourceCommandStub != nil {
		return fake.SourceCommandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceCommandReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) SourceCommandCallCount() int {
	fake.sourceCommandMutex.RLock()
	defer fake.sourceCommandMutex.RUnlock()
	return len(fake.sourceCommandArgsForCall)
}

func (fake *FakeGHCommandFactory) SourceCommandCalls(stub func() *cobra.Command) {
	fake.sourceCommandMutex.Lock()
	defer fake.sourceCommandMutex.Unlock()
	fake.SourceCommandStub = stub
}

func (fake *FakeGHCommandFactory) SourceCommandReturns(result1 *cobra.Command) {
	fake.sourceCommandMutex.Lock()
	defer fake.sourceCommandMutex.Unlock()
	fake.SourceCommandStub = nil
	fake.sourceCommandReturns = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) SourceCommandReturnsOnCall(i int, result1 *cobra.Command) {
	fake.sourceCommandMutex.Lock()
	defer fake.sourceCommandMutex.Unlock()
	fake.SourceCommandStub = nil
	if fake.sourceCommandReturnsOnCall == nil {
		fake.sourceCommandReturnsOnCall = make(map[int]struct {
			result1 *cobra.Command
		})
	}
	fake.sourceCommandReturnsOnCall[i] = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) UpdateCommand() *cobra.Command {
	fake.updateCommandMutex.Lock()
	ret, specificReturn := fake.updateCommandReturnsOnCall[len(fake.updateCommandArgsForCall)]
	fake.updateCommandArgsForCall = append(fake.updateCommandArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateCommand", []interface{}{})
	fake.updateCommandMutex.Unlock()
	if fake.UpdateCommandStub != nil {
		return fake.UpdateCommandStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateCommandReturns
	return fakeReturns.result1
}

func (fake *FakeGHCommandFactory) UpdateCommandCallCount() int {
	fake.updateCommandMutex.RLock()
	defer fake.updateCommandMutex.RUnlock()
	return len(fake.updateCommandArgsForCall)
}

func (fake *FakeGHCommandFactory) UpdateCommandCalls(stub func() *cobra.Command) {
	fake.updateCommandMutex.Lock()
	defer fake.updateCommandMutex.Unlock()
	fake.UpdateCommandStub = stub
}

func (fake *FakeGHCommandFactory) UpdateCommandReturns(result1 *cobra.Command) {
	fake.updateCommandMutex.Lock()
	defer fake.updateCommandMutex.Unlock()
	fake.UpdateCommandStub = nil
	fake.updateCommandReturns = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) UpdateCommandReturnsOnCall(i int, result1 *cobra.Command) {
	fake.updateCommandMutex.Lock()
	defer fake.updateCommandMutex.Unlock()
	fake.UpdateCommandStub = nil
	if fake.updateCommandReturnsOnCall == nil {
		fake.updateCommandReturnsOnCall = make(map[int]struct {
			result1 *cobra.Command
		})
	}
	fake.updateCommandReturnsOnCall[i] = struct {
		result1 *cobra.Command
	}{result1}
}

func (fake *FakeGHCommandFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCommandMutex.RLock()
	defer fake.createCommandMutex.RUnlock()
	fake.deleteCommandMutex.RLock()
	defer fake.deleteCommandMutex.RUnlock()
	fake.describeCommandMutex.RLock()
	defer fake.describeCommandMutex.RUnlock()
	fake.gHSourceFactoryMutex.RLock()
	defer fake.gHSourceFactoryMutex.RUnlock()
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	fake.sourceCommandMutex.RLock()
	defer fake.sourceCommandMutex.RUnlock()
	fake.updateCommandMutex.RLock()
	defer fake.updateCommandMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGHCommandFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.GHCommandFactory = new(FakeGHCommandFactory)