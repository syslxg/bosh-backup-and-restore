// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/executor"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
)

type FakeDeployment struct {
	IsBackupableStub        func() bool
	isBackupableMutex       sync.RWMutex
	isBackupableArgsForCall []struct{}
	isBackupableReturns     struct {
		result1 bool
	}
	isBackupableReturnsOnCall map[int]struct {
		result1 bool
	}
	BackupableInstancesStub        func() []orchestrator.Instance
	backupableInstancesMutex       sync.RWMutex
	backupableInstancesArgsForCall []struct{}
	backupableInstancesReturns     struct {
		result1 []orchestrator.Instance
	}
	backupableInstancesReturnsOnCall map[int]struct {
		result1 []orchestrator.Instance
	}
	HasUniqueCustomArtifactNamesStub        func() bool
	hasUniqueCustomArtifactNamesMutex       sync.RWMutex
	hasUniqueCustomArtifactNamesArgsForCall []struct{}
	hasUniqueCustomArtifactNamesReturns     struct {
		result1 bool
	}
	hasUniqueCustomArtifactNamesReturnsOnCall map[int]struct {
		result1 bool
	}
	CheckArtifactDirStub        func() error
	checkArtifactDirMutex       sync.RWMutex
	checkArtifactDirArgsForCall []struct{}
	checkArtifactDirReturns     struct {
		result1 error
	}
	checkArtifactDirReturnsOnCall map[int]struct {
		result1 error
	}
	IsRestorableStub        func() bool
	isRestorableMutex       sync.RWMutex
	isRestorableArgsForCall []struct{}
	isRestorableReturns     struct {
		result1 bool
	}
	isRestorableReturnsOnCall map[int]struct {
		result1 bool
	}
	RestorableInstancesStub        func() []orchestrator.Instance
	restorableInstancesMutex       sync.RWMutex
	restorableInstancesArgsForCall []struct{}
	restorableInstancesReturns     struct {
		result1 []orchestrator.Instance
	}
	restorableInstancesReturnsOnCall map[int]struct {
		result1 []orchestrator.Instance
	}
	PreBackupLockStub        func(orchestrator.LockOrderer, executor.Executor) error
	preBackupLockMutex       sync.RWMutex
	preBackupLockArgsForCall []struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}
	preBackupLockReturns struct {
		result1 error
	}
	preBackupLockReturnsOnCall map[int]struct {
		result1 error
	}
	BackupStub        func(executor.Executor) error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct {
		arg1 executor.Executor
	}
	backupReturns struct {
		result1 error
	}
	backupReturnsOnCall map[int]struct {
		result1 error
	}
	PostBackupUnlockStub        func(bool, orchestrator.LockOrderer, executor.Executor) error
	postBackupUnlockMutex       sync.RWMutex
	postBackupUnlockArgsForCall []struct {
		arg1 bool
		arg2 orchestrator.LockOrderer
		arg3 executor.Executor
	}
	postBackupUnlockReturns struct {
		result1 error
	}
	postBackupUnlockReturnsOnCall map[int]struct {
		result1 error
	}
	RestoreStub        func() error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct{}
	restoreReturns     struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	CleanupPreviousStub        func() error
	cleanupPreviousMutex       sync.RWMutex
	cleanupPreviousArgsForCall []struct{}
	cleanupPreviousReturns     struct {
		result1 error
	}
	cleanupPreviousReturnsOnCall map[int]struct {
		result1 error
	}
	InstancesStub        func() []orchestrator.Instance
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct{}
	instancesReturns     struct {
		result1 []orchestrator.Instance
	}
	instancesReturnsOnCall map[int]struct {
		result1 []orchestrator.Instance
	}
	CustomArtifactNamesMatchStub        func() error
	customArtifactNamesMatchMutex       sync.RWMutex
	customArtifactNamesMatchArgsForCall []struct{}
	customArtifactNamesMatchReturns     struct {
		result1 error
	}
	customArtifactNamesMatchReturnsOnCall map[int]struct {
		result1 error
	}
	PreRestoreLockStub        func(orchestrator.LockOrderer, executor.Executor) error
	preRestoreLockMutex       sync.RWMutex
	preRestoreLockArgsForCall []struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}
	preRestoreLockReturns struct {
		result1 error
	}
	preRestoreLockReturnsOnCall map[int]struct {
		result1 error
	}
	PostRestoreUnlockStub        func(orchestrator.LockOrderer, executor.Executor) error
	postRestoreUnlockMutex       sync.RWMutex
	postRestoreUnlockArgsForCall []struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}
	postRestoreUnlockReturns struct {
		result1 error
	}
	postRestoreUnlockReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateLockingDependenciesStub        func(orderer orchestrator.LockOrderer) error
	validateLockingDependenciesMutex       sync.RWMutex
	validateLockingDependenciesArgsForCall []struct {
		orderer orchestrator.LockOrderer
	}
	validateLockingDependenciesReturns struct {
		result1 error
	}
	validateLockingDependenciesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeployment) IsBackupable() bool {
	fake.isBackupableMutex.Lock()
	ret, specificReturn := fake.isBackupableReturnsOnCall[len(fake.isBackupableArgsForCall)]
	fake.isBackupableArgsForCall = append(fake.isBackupableArgsForCall, struct{}{})
	fake.recordInvocation("IsBackupable", []interface{}{})
	fake.isBackupableMutex.Unlock()
	if fake.IsBackupableStub != nil {
		return fake.IsBackupableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isBackupableReturns.result1
}

func (fake *FakeDeployment) IsBackupableCallCount() int {
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	return len(fake.isBackupableArgsForCall)
}

func (fake *FakeDeployment) IsBackupableReturns(result1 bool) {
	fake.IsBackupableStub = nil
	fake.isBackupableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) IsBackupableReturnsOnCall(i int, result1 bool) {
	fake.IsBackupableStub = nil
	if fake.isBackupableReturnsOnCall == nil {
		fake.isBackupableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isBackupableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) BackupableInstances() []orchestrator.Instance {
	fake.backupableInstancesMutex.Lock()
	ret, specificReturn := fake.backupableInstancesReturnsOnCall[len(fake.backupableInstancesArgsForCall)]
	fake.backupableInstancesArgsForCall = append(fake.backupableInstancesArgsForCall, struct{}{})
	fake.recordInvocation("BackupableInstances", []interface{}{})
	fake.backupableInstancesMutex.Unlock()
	if fake.BackupableInstancesStub != nil {
		return fake.BackupableInstancesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.backupableInstancesReturns.result1
}

func (fake *FakeDeployment) BackupableInstancesCallCount() int {
	fake.backupableInstancesMutex.RLock()
	defer fake.backupableInstancesMutex.RUnlock()
	return len(fake.backupableInstancesArgsForCall)
}

func (fake *FakeDeployment) BackupableInstancesReturns(result1 []orchestrator.Instance) {
	fake.BackupableInstancesStub = nil
	fake.backupableInstancesReturns = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) BackupableInstancesReturnsOnCall(i int, result1 []orchestrator.Instance) {
	fake.BackupableInstancesStub = nil
	if fake.backupableInstancesReturnsOnCall == nil {
		fake.backupableInstancesReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.Instance
		})
	}
	fake.backupableInstancesReturnsOnCall[i] = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) HasUniqueCustomArtifactNames() bool {
	fake.hasUniqueCustomArtifactNamesMutex.Lock()
	ret, specificReturn := fake.hasUniqueCustomArtifactNamesReturnsOnCall[len(fake.hasUniqueCustomArtifactNamesArgsForCall)]
	fake.hasUniqueCustomArtifactNamesArgsForCall = append(fake.hasUniqueCustomArtifactNamesArgsForCall, struct{}{})
	fake.recordInvocation("HasUniqueCustomArtifactNames", []interface{}{})
	fake.hasUniqueCustomArtifactNamesMutex.Unlock()
	if fake.HasUniqueCustomArtifactNamesStub != nil {
		return fake.HasUniqueCustomArtifactNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.hasUniqueCustomArtifactNamesReturns.result1
}

func (fake *FakeDeployment) HasUniqueCustomArtifactNamesCallCount() int {
	fake.hasUniqueCustomArtifactNamesMutex.RLock()
	defer fake.hasUniqueCustomArtifactNamesMutex.RUnlock()
	return len(fake.hasUniqueCustomArtifactNamesArgsForCall)
}

func (fake *FakeDeployment) HasUniqueCustomArtifactNamesReturns(result1 bool) {
	fake.HasUniqueCustomArtifactNamesStub = nil
	fake.hasUniqueCustomArtifactNamesReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) HasUniqueCustomArtifactNamesReturnsOnCall(i int, result1 bool) {
	fake.HasUniqueCustomArtifactNamesStub = nil
	if fake.hasUniqueCustomArtifactNamesReturnsOnCall == nil {
		fake.hasUniqueCustomArtifactNamesReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasUniqueCustomArtifactNamesReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) CheckArtifactDir() error {
	fake.checkArtifactDirMutex.Lock()
	ret, specificReturn := fake.checkArtifactDirReturnsOnCall[len(fake.checkArtifactDirArgsForCall)]
	fake.checkArtifactDirArgsForCall = append(fake.checkArtifactDirArgsForCall, struct{}{})
	fake.recordInvocation("CheckArtifactDir", []interface{}{})
	fake.checkArtifactDirMutex.Unlock()
	if fake.CheckArtifactDirStub != nil {
		return fake.CheckArtifactDirStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkArtifactDirReturns.result1
}

func (fake *FakeDeployment) CheckArtifactDirCallCount() int {
	fake.checkArtifactDirMutex.RLock()
	defer fake.checkArtifactDirMutex.RUnlock()
	return len(fake.checkArtifactDirArgsForCall)
}

func (fake *FakeDeployment) CheckArtifactDirReturns(result1 error) {
	fake.CheckArtifactDirStub = nil
	fake.checkArtifactDirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CheckArtifactDirReturnsOnCall(i int, result1 error) {
	fake.CheckArtifactDirStub = nil
	if fake.checkArtifactDirReturnsOnCall == nil {
		fake.checkArtifactDirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkArtifactDirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) IsRestorable() bool {
	fake.isRestorableMutex.Lock()
	ret, specificReturn := fake.isRestorableReturnsOnCall[len(fake.isRestorableArgsForCall)]
	fake.isRestorableArgsForCall = append(fake.isRestorableArgsForCall, struct{}{})
	fake.recordInvocation("IsRestorable", []interface{}{})
	fake.isRestorableMutex.Unlock()
	if fake.IsRestorableStub != nil {
		return fake.IsRestorableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isRestorableReturns.result1
}

func (fake *FakeDeployment) IsRestorableCallCount() int {
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	return len(fake.isRestorableArgsForCall)
}

func (fake *FakeDeployment) IsRestorableReturns(result1 bool) {
	fake.IsRestorableStub = nil
	fake.isRestorableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) IsRestorableReturnsOnCall(i int, result1 bool) {
	fake.IsRestorableStub = nil
	if fake.isRestorableReturnsOnCall == nil {
		fake.isRestorableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isRestorableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDeployment) RestorableInstances() []orchestrator.Instance {
	fake.restorableInstancesMutex.Lock()
	ret, specificReturn := fake.restorableInstancesReturnsOnCall[len(fake.restorableInstancesArgsForCall)]
	fake.restorableInstancesArgsForCall = append(fake.restorableInstancesArgsForCall, struct{}{})
	fake.recordInvocation("RestorableInstances", []interface{}{})
	fake.restorableInstancesMutex.Unlock()
	if fake.RestorableInstancesStub != nil {
		return fake.RestorableInstancesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restorableInstancesReturns.result1
}

func (fake *FakeDeployment) RestorableInstancesCallCount() int {
	fake.restorableInstancesMutex.RLock()
	defer fake.restorableInstancesMutex.RUnlock()
	return len(fake.restorableInstancesArgsForCall)
}

func (fake *FakeDeployment) RestorableInstancesReturns(result1 []orchestrator.Instance) {
	fake.RestorableInstancesStub = nil
	fake.restorableInstancesReturns = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) RestorableInstancesReturnsOnCall(i int, result1 []orchestrator.Instance) {
	fake.RestorableInstancesStub = nil
	if fake.restorableInstancesReturnsOnCall == nil {
		fake.restorableInstancesReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.Instance
		})
	}
	fake.restorableInstancesReturnsOnCall[i] = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) PreBackupLock(arg1 orchestrator.LockOrderer, arg2 executor.Executor) error {
	fake.preBackupLockMutex.Lock()
	ret, specificReturn := fake.preBackupLockReturnsOnCall[len(fake.preBackupLockArgsForCall)]
	fake.preBackupLockArgsForCall = append(fake.preBackupLockArgsForCall, struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}{arg1, arg2})
	fake.recordInvocation("PreBackupLock", []interface{}{arg1, arg2})
	fake.preBackupLockMutex.Unlock()
	if fake.PreBackupLockStub != nil {
		return fake.PreBackupLockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.preBackupLockReturns.result1
}

func (fake *FakeDeployment) PreBackupLockCallCount() int {
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	return len(fake.preBackupLockArgsForCall)
}

func (fake *FakeDeployment) PreBackupLockArgsForCall(i int) (orchestrator.LockOrderer, executor.Executor) {
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	return fake.preBackupLockArgsForCall[i].arg1, fake.preBackupLockArgsForCall[i].arg2
}

func (fake *FakeDeployment) PreBackupLockReturns(result1 error) {
	fake.PreBackupLockStub = nil
	fake.preBackupLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PreBackupLockReturnsOnCall(i int, result1 error) {
	fake.PreBackupLockStub = nil
	if fake.preBackupLockReturnsOnCall == nil {
		fake.preBackupLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.preBackupLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Backup(arg1 executor.Executor) error {
	fake.backupMutex.Lock()
	ret, specificReturn := fake.backupReturnsOnCall[len(fake.backupArgsForCall)]
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct {
		arg1 executor.Executor
	}{arg1})
	fake.recordInvocation("Backup", []interface{}{arg1})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.backupReturns.result1
}

func (fake *FakeDeployment) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeDeployment) BackupArgsForCall(i int) executor.Executor {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return fake.backupArgsForCall[i].arg1
}

func (fake *FakeDeployment) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) BackupReturnsOnCall(i int, result1 error) {
	fake.BackupStub = nil
	if fake.backupReturnsOnCall == nil {
		fake.backupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.backupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PostBackupUnlock(arg1 bool, arg2 orchestrator.LockOrderer, arg3 executor.Executor) error {
	fake.postBackupUnlockMutex.Lock()
	ret, specificReturn := fake.postBackupUnlockReturnsOnCall[len(fake.postBackupUnlockArgsForCall)]
	fake.postBackupUnlockArgsForCall = append(fake.postBackupUnlockArgsForCall, struct {
		arg1 bool
		arg2 orchestrator.LockOrderer
		arg3 executor.Executor
	}{arg1, arg2, arg3})
	fake.recordInvocation("PostBackupUnlock", []interface{}{arg1, arg2, arg3})
	fake.postBackupUnlockMutex.Unlock()
	if fake.PostBackupUnlockStub != nil {
		return fake.PostBackupUnlockStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.postBackupUnlockReturns.result1
}

func (fake *FakeDeployment) PostBackupUnlockCallCount() int {
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	return len(fake.postBackupUnlockArgsForCall)
}

func (fake *FakeDeployment) PostBackupUnlockArgsForCall(i int) (bool, orchestrator.LockOrderer, executor.Executor) {
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	return fake.postBackupUnlockArgsForCall[i].arg1, fake.postBackupUnlockArgsForCall[i].arg2, fake.postBackupUnlockArgsForCall[i].arg3
}

func (fake *FakeDeployment) PostBackupUnlockReturns(result1 error) {
	fake.PostBackupUnlockStub = nil
	fake.postBackupUnlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PostBackupUnlockReturnsOnCall(i int, result1 error) {
	fake.PostBackupUnlockStub = nil
	if fake.postBackupUnlockReturnsOnCall == nil {
		fake.postBackupUnlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.postBackupUnlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Restore() error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct{}{})
	fake.recordInvocation("Restore", []interface{}{})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restoreReturns.result1
}

func (fake *FakeDeployment) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeDeployment) RestoreReturns(result1 error) {
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) RestoreReturnsOnCall(i int, result1 error) {
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Cleanup() error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeDeployment) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeDeployment) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CleanupPrevious() error {
	fake.cleanupPreviousMutex.Lock()
	ret, specificReturn := fake.cleanupPreviousReturnsOnCall[len(fake.cleanupPreviousArgsForCall)]
	fake.cleanupPreviousArgsForCall = append(fake.cleanupPreviousArgsForCall, struct{}{})
	fake.recordInvocation("CleanupPrevious", []interface{}{})
	fake.cleanupPreviousMutex.Unlock()
	if fake.CleanupPreviousStub != nil {
		return fake.CleanupPreviousStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupPreviousReturns.result1
}

func (fake *FakeDeployment) CleanupPreviousCallCount() int {
	fake.cleanupPreviousMutex.RLock()
	defer fake.cleanupPreviousMutex.RUnlock()
	return len(fake.cleanupPreviousArgsForCall)
}

func (fake *FakeDeployment) CleanupPreviousReturns(result1 error) {
	fake.CleanupPreviousStub = nil
	fake.cleanupPreviousReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CleanupPreviousReturnsOnCall(i int, result1 error) {
	fake.CleanupPreviousStub = nil
	if fake.cleanupPreviousReturnsOnCall == nil {
		fake.cleanupPreviousReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupPreviousReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Instances() []orchestrator.Instance {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct{}{})
	fake.recordInvocation("Instances", []interface{}{})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.instancesReturns.result1
}

func (fake *FakeDeployment) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeDeployment) InstancesReturns(result1 []orchestrator.Instance) {
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) InstancesReturnsOnCall(i int, result1 []orchestrator.Instance) {
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.Instance
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []orchestrator.Instance
	}{result1}
}

func (fake *FakeDeployment) CustomArtifactNamesMatch() error {
	fake.customArtifactNamesMatchMutex.Lock()
	ret, specificReturn := fake.customArtifactNamesMatchReturnsOnCall[len(fake.customArtifactNamesMatchArgsForCall)]
	fake.customArtifactNamesMatchArgsForCall = append(fake.customArtifactNamesMatchArgsForCall, struct{}{})
	fake.recordInvocation("CustomArtifactNamesMatch", []interface{}{})
	fake.customArtifactNamesMatchMutex.Unlock()
	if fake.CustomArtifactNamesMatchStub != nil {
		return fake.CustomArtifactNamesMatchStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.customArtifactNamesMatchReturns.result1
}

func (fake *FakeDeployment) CustomArtifactNamesMatchCallCount() int {
	fake.customArtifactNamesMatchMutex.RLock()
	defer fake.customArtifactNamesMatchMutex.RUnlock()
	return len(fake.customArtifactNamesMatchArgsForCall)
}

func (fake *FakeDeployment) CustomArtifactNamesMatchReturns(result1 error) {
	fake.CustomArtifactNamesMatchStub = nil
	fake.customArtifactNamesMatchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) CustomArtifactNamesMatchReturnsOnCall(i int, result1 error) {
	fake.CustomArtifactNamesMatchStub = nil
	if fake.customArtifactNamesMatchReturnsOnCall == nil {
		fake.customArtifactNamesMatchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.customArtifactNamesMatchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PreRestoreLock(arg1 orchestrator.LockOrderer, arg2 executor.Executor) error {
	fake.preRestoreLockMutex.Lock()
	ret, specificReturn := fake.preRestoreLockReturnsOnCall[len(fake.preRestoreLockArgsForCall)]
	fake.preRestoreLockArgsForCall = append(fake.preRestoreLockArgsForCall, struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}{arg1, arg2})
	fake.recordInvocation("PreRestoreLock", []interface{}{arg1, arg2})
	fake.preRestoreLockMutex.Unlock()
	if fake.PreRestoreLockStub != nil {
		return fake.PreRestoreLockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.preRestoreLockReturns.result1
}

func (fake *FakeDeployment) PreRestoreLockCallCount() int {
	fake.preRestoreLockMutex.RLock()
	defer fake.preRestoreLockMutex.RUnlock()
	return len(fake.preRestoreLockArgsForCall)
}

func (fake *FakeDeployment) PreRestoreLockArgsForCall(i int) (orchestrator.LockOrderer, executor.Executor) {
	fake.preRestoreLockMutex.RLock()
	defer fake.preRestoreLockMutex.RUnlock()
	return fake.preRestoreLockArgsForCall[i].arg1, fake.preRestoreLockArgsForCall[i].arg2
}

func (fake *FakeDeployment) PreRestoreLockReturns(result1 error) {
	fake.PreRestoreLockStub = nil
	fake.preRestoreLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PreRestoreLockReturnsOnCall(i int, result1 error) {
	fake.PreRestoreLockStub = nil
	if fake.preRestoreLockReturnsOnCall == nil {
		fake.preRestoreLockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.preRestoreLockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PostRestoreUnlock(arg1 orchestrator.LockOrderer, arg2 executor.Executor) error {
	fake.postRestoreUnlockMutex.Lock()
	ret, specificReturn := fake.postRestoreUnlockReturnsOnCall[len(fake.postRestoreUnlockArgsForCall)]
	fake.postRestoreUnlockArgsForCall = append(fake.postRestoreUnlockArgsForCall, struct {
		arg1 orchestrator.LockOrderer
		arg2 executor.Executor
	}{arg1, arg2})
	fake.recordInvocation("PostRestoreUnlock", []interface{}{arg1, arg2})
	fake.postRestoreUnlockMutex.Unlock()
	if fake.PostRestoreUnlockStub != nil {
		return fake.PostRestoreUnlockStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.postRestoreUnlockReturns.result1
}

func (fake *FakeDeployment) PostRestoreUnlockCallCount() int {
	fake.postRestoreUnlockMutex.RLock()
	defer fake.postRestoreUnlockMutex.RUnlock()
	return len(fake.postRestoreUnlockArgsForCall)
}

func (fake *FakeDeployment) PostRestoreUnlockArgsForCall(i int) (orchestrator.LockOrderer, executor.Executor) {
	fake.postRestoreUnlockMutex.RLock()
	defer fake.postRestoreUnlockMutex.RUnlock()
	return fake.postRestoreUnlockArgsForCall[i].arg1, fake.postRestoreUnlockArgsForCall[i].arg2
}

func (fake *FakeDeployment) PostRestoreUnlockReturns(result1 error) {
	fake.PostRestoreUnlockStub = nil
	fake.postRestoreUnlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) PostRestoreUnlockReturnsOnCall(i int, result1 error) {
	fake.PostRestoreUnlockStub = nil
	if fake.postRestoreUnlockReturnsOnCall == nil {
		fake.postRestoreUnlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.postRestoreUnlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) ValidateLockingDependencies(orderer orchestrator.LockOrderer) error {
	fake.validateLockingDependenciesMutex.Lock()
	ret, specificReturn := fake.validateLockingDependenciesReturnsOnCall[len(fake.validateLockingDependenciesArgsForCall)]
	fake.validateLockingDependenciesArgsForCall = append(fake.validateLockingDependenciesArgsForCall, struct {
		orderer orchestrator.LockOrderer
	}{orderer})
	fake.recordInvocation("ValidateLockingDependencies", []interface{}{orderer})
	fake.validateLockingDependenciesMutex.Unlock()
	if fake.ValidateLockingDependenciesStub != nil {
		return fake.ValidateLockingDependenciesStub(orderer)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateLockingDependenciesReturns.result1
}

func (fake *FakeDeployment) ValidateLockingDependenciesCallCount() int {
	fake.validateLockingDependenciesMutex.RLock()
	defer fake.validateLockingDependenciesMutex.RUnlock()
	return len(fake.validateLockingDependenciesArgsForCall)
}

func (fake *FakeDeployment) ValidateLockingDependenciesArgsForCall(i int) orchestrator.LockOrderer {
	fake.validateLockingDependenciesMutex.RLock()
	defer fake.validateLockingDependenciesMutex.RUnlock()
	return fake.validateLockingDependenciesArgsForCall[i].orderer
}

func (fake *FakeDeployment) ValidateLockingDependenciesReturns(result1 error) {
	fake.ValidateLockingDependenciesStub = nil
	fake.validateLockingDependenciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) ValidateLockingDependenciesReturnsOnCall(i int, result1 error) {
	fake.ValidateLockingDependenciesStub = nil
	if fake.validateLockingDependenciesReturnsOnCall == nil {
		fake.validateLockingDependenciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateLockingDependenciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDeployment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isBackupableMutex.RLock()
	defer fake.isBackupableMutex.RUnlock()
	fake.backupableInstancesMutex.RLock()
	defer fake.backupableInstancesMutex.RUnlock()
	fake.hasUniqueCustomArtifactNamesMutex.RLock()
	defer fake.hasUniqueCustomArtifactNamesMutex.RUnlock()
	fake.checkArtifactDirMutex.RLock()
	defer fake.checkArtifactDirMutex.RUnlock()
	fake.isRestorableMutex.RLock()
	defer fake.isRestorableMutex.RUnlock()
	fake.restorableInstancesMutex.RLock()
	defer fake.restorableInstancesMutex.RUnlock()
	fake.preBackupLockMutex.RLock()
	defer fake.preBackupLockMutex.RUnlock()
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	fake.postBackupUnlockMutex.RLock()
	defer fake.postBackupUnlockMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.cleanupPreviousMutex.RLock()
	defer fake.cleanupPreviousMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	fake.customArtifactNamesMatchMutex.RLock()
	defer fake.customArtifactNamesMatchMutex.RUnlock()
	fake.preRestoreLockMutex.RLock()
	defer fake.preRestoreLockMutex.RUnlock()
	fake.postRestoreUnlockMutex.RLock()
	defer fake.postRestoreUnlockMutex.RUnlock()
	fake.validateLockingDependenciesMutex.RLock()
	defer fake.validateLockingDependenciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeployment) recordInvocation(key string, args []interface{}) {
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

var _ orchestrator.Deployment = new(FakeDeployment)
