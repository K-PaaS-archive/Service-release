// This file was generated by counterfeiter
package fakes

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry/mariadb_ctrl/mariadb_helper"
	_ "github.com/go-sql-driver/mysql"
)

type FakeDBHelper struct {
	StartMysqldInModeStub        func(command string) error
	startMysqldInModeMutex       sync.RWMutex
	startMysqldInModeArgsForCall []struct {
		command string
	}
	startMysqldInModeReturns struct {
		result1 error
	}
	StartMysqlInJoinStub        func() (*exec.Cmd, error)
	startMysqlInJoinMutex       sync.RWMutex
	startMysqlInJoinArgsForCall []struct{}
	startMysqlInJoinReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	StartMysqlInBootstrapStub        func() (*exec.Cmd, error)
	startMysqlInBootstrapMutex       sync.RWMutex
	startMysqlInBootstrapArgsForCall []struct{}
	startMysqlInBootstrapReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	StopMysqlStub        func() error
	stopMysqlMutex       sync.RWMutex
	stopMysqlArgsForCall []struct{}
	stopMysqlReturns     struct {
		result1 error
	}
	StopStandaloneMysqlStub        func() error
	stopStandaloneMysqlMutex       sync.RWMutex
	stopStandaloneMysqlArgsForCall []struct{}
	stopStandaloneMysqlReturns     struct {
		result1 error
	}
	UpgradeStub        func() (output string, err error)
	upgradeMutex       sync.RWMutex
	upgradeArgsForCall []struct{}
	upgradeReturns     struct {
		result1 string
		result2 error
	}
	IsDatabaseReachableStub        func() bool
	isDatabaseReachableMutex       sync.RWMutex
	isDatabaseReachableArgsForCall []struct{}
	isDatabaseReachableReturns     struct {
		result1 bool
	}
	SeedStub        func() error
	seedMutex       sync.RWMutex
	seedArgsForCall []struct{}
	seedReturns     struct {
		result1 error
	}
}

func (fake *FakeDBHelper) StartMysqldInMode(command string) error {
	fake.startMysqldInModeMutex.Lock()
	fake.startMysqldInModeArgsForCall = append(fake.startMysqldInModeArgsForCall, struct {
		command string
	}{command})
	fake.startMysqldInModeMutex.Unlock()
	if fake.StartMysqldInModeStub != nil {
		return fake.StartMysqldInModeStub(command)
	} else {
		return fake.startMysqldInModeReturns.result1
	}
}

func (fake *FakeDBHelper) StartMysqldInModeCallCount() int {
	fake.startMysqldInModeMutex.RLock()
	defer fake.startMysqldInModeMutex.RUnlock()
	return len(fake.startMysqldInModeArgsForCall)
}

func (fake *FakeDBHelper) StartMysqldInModeArgsForCall(i int) string {
	fake.startMysqldInModeMutex.RLock()
	defer fake.startMysqldInModeMutex.RUnlock()
	return fake.startMysqldInModeArgsForCall[i].command
}

func (fake *FakeDBHelper) StartMysqldInModeReturns(result1 error) {
	fake.StartMysqldInModeStub = nil
	fake.startMysqldInModeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StartMysqlInJoin() (*exec.Cmd, error) {
	fake.startMysqlInJoinMutex.Lock()
	fake.startMysqlInJoinArgsForCall = append(fake.startMysqlInJoinArgsForCall, struct{}{})
	fake.startMysqlInJoinMutex.Unlock()
	if fake.StartMysqlInJoinStub != nil {
		return fake.StartMysqlInJoinStub()
	} else {
		return fake.startMysqlInJoinReturns.result1, fake.startMysqlInJoinReturns.result2
	}
}

func (fake *FakeDBHelper) StartMysqlInJoinCallCount() int {
	fake.startMysqlInJoinMutex.RLock()
	defer fake.startMysqlInJoinMutex.RUnlock()
	return len(fake.startMysqlInJoinArgsForCall)
}

func (fake *FakeDBHelper) StartMysqlInJoinReturns(result1 *exec.Cmd, result2 error) {
	fake.StartMysqlInJoinStub = nil
	fake.startMysqlInJoinReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StartMysqlInBootstrap() (*exec.Cmd, error) {
	fake.startMysqlInBootstrapMutex.Lock()
	fake.startMysqlInBootstrapArgsForCall = append(fake.startMysqlInBootstrapArgsForCall, struct{}{})
	fake.startMysqlInBootstrapMutex.Unlock()
	if fake.StartMysqlInBootstrapStub != nil {
		return fake.StartMysqlInBootstrapStub()
	} else {
		return fake.startMysqlInBootstrapReturns.result1, fake.startMysqlInBootstrapReturns.result2
	}
}

func (fake *FakeDBHelper) StartMysqlInBootstrapCallCount() int {
	fake.startMysqlInBootstrapMutex.RLock()
	defer fake.startMysqlInBootstrapMutex.RUnlock()
	return len(fake.startMysqlInBootstrapArgsForCall)
}

func (fake *FakeDBHelper) StartMysqlInBootstrapReturns(result1 *exec.Cmd, result2 error) {
	fake.StartMysqlInBootstrapStub = nil
	fake.startMysqlInBootstrapReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StopMysql() error {
	fake.stopMysqlMutex.Lock()
	fake.stopMysqlArgsForCall = append(fake.stopMysqlArgsForCall, struct{}{})
	fake.stopMysqlMutex.Unlock()
	if fake.StopMysqlStub != nil {
		return fake.StopMysqlStub()
	} else {
		return fake.stopMysqlReturns.result1
	}
}

func (fake *FakeDBHelper) StopMysqlCallCount() int {
	fake.stopMysqlMutex.RLock()
	defer fake.stopMysqlMutex.RUnlock()
	return len(fake.stopMysqlArgsForCall)
}

func (fake *FakeDBHelper) StopMysqlReturns(result1 error) {
	fake.StopMysqlStub = nil
	fake.stopMysqlReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StopStandaloneMysql() error {
	fake.stopStandaloneMysqlMutex.Lock()
	fake.stopStandaloneMysqlArgsForCall = append(fake.stopStandaloneMysqlArgsForCall, struct{}{})
	fake.stopStandaloneMysqlMutex.Unlock()
	if fake.StopStandaloneMysqlStub != nil {
		return fake.StopStandaloneMysqlStub()
	} else {
		return fake.stopStandaloneMysqlReturns.result1
	}
}

func (fake *FakeDBHelper) StopStandaloneMysqlCallCount() int {
	fake.stopStandaloneMysqlMutex.RLock()
	defer fake.stopStandaloneMysqlMutex.RUnlock()
	return len(fake.stopStandaloneMysqlArgsForCall)
}

func (fake *FakeDBHelper) StopStandaloneMysqlReturns(result1 error) {
	fake.StopStandaloneMysqlStub = nil
	fake.stopStandaloneMysqlReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) Upgrade() (output string, err error) {
	fake.upgradeMutex.Lock()
	fake.upgradeArgsForCall = append(fake.upgradeArgsForCall, struct{}{})
	fake.upgradeMutex.Unlock()
	if fake.UpgradeStub != nil {
		return fake.UpgradeStub()
	} else {
		return fake.upgradeReturns.result1, fake.upgradeReturns.result2
	}
}

func (fake *FakeDBHelper) UpgradeCallCount() int {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return len(fake.upgradeArgsForCall)
}

func (fake *FakeDBHelper) UpgradeReturns(result1 string, result2 error) {
	fake.UpgradeStub = nil
	fake.upgradeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) IsDatabaseReachable() bool {
	fake.isDatabaseReachableMutex.Lock()
	fake.isDatabaseReachableArgsForCall = append(fake.isDatabaseReachableArgsForCall, struct{}{})
	fake.isDatabaseReachableMutex.Unlock()
	if fake.IsDatabaseReachableStub != nil {
		return fake.IsDatabaseReachableStub()
	} else {
		return fake.isDatabaseReachableReturns.result1
	}
}

func (fake *FakeDBHelper) IsDatabaseReachableCallCount() int {
	fake.isDatabaseReachableMutex.RLock()
	defer fake.isDatabaseReachableMutex.RUnlock()
	return len(fake.isDatabaseReachableArgsForCall)
}

func (fake *FakeDBHelper) IsDatabaseReachableReturns(result1 bool) {
	fake.IsDatabaseReachableStub = nil
	fake.isDatabaseReachableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDBHelper) Seed() error {
	fake.seedMutex.Lock()
	fake.seedArgsForCall = append(fake.seedArgsForCall, struct{}{})
	fake.seedMutex.Unlock()
	if fake.SeedStub != nil {
		return fake.SeedStub()
	} else {
		return fake.seedReturns.result1
	}
}

func (fake *FakeDBHelper) SeedCallCount() int {
	fake.seedMutex.RLock()
	defer fake.seedMutex.RUnlock()
	return len(fake.seedArgsForCall)
}

func (fake *FakeDBHelper) SeedReturns(result1 error) {
	fake.SeedStub = nil
	fake.seedReturns = struct {
		result1 error
	}{result1}
}

var _ mariadb_helper.DBHelper = new(FakeDBHelper)