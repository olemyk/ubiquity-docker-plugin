// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.ibm.com/almaden-containers/ubiquity/local/spectrumscale/connectors"
	"github.ibm.com/almaden-containers/ubiquity/resources"
)

type FakeSpectrumScaleConnector struct {
	GetClusterIdStub        func() (string, error)
	getClusterIdMutex       sync.RWMutex
	getClusterIdArgsForCall []struct{}
	getClusterIdReturns     struct {
		result1 string
		result2 error
	}
	IsFilesystemMountedStub        func(filesystemName string) (bool, error)
	isFilesystemMountedMutex       sync.RWMutex
	isFilesystemMountedArgsForCall []struct {
		filesystemName string
	}
	isFilesystemMountedReturns struct {
		result1 bool
		result2 error
	}
	MountFileSystemStub        func(filesystemName string) error
	mountFileSystemMutex       sync.RWMutex
	mountFileSystemArgsForCall []struct {
		filesystemName string
	}
	mountFileSystemReturns struct {
		result1 error
	}
	ListFilesystemsStub        func() ([]string, error)
	listFilesystemsMutex       sync.RWMutex
	listFilesystemsArgsForCall []struct{}
	listFilesystemsReturns     struct {
		result1 []string
		result2 error
	}
	GetFilesystemMountpointStub        func(filesystemName string) (string, error)
	getFilesystemMountpointMutex       sync.RWMutex
	getFilesystemMountpointArgsForCall []struct {
		filesystemName string
	}
	getFilesystemMountpointReturns struct {
		result1 string
		result2 error
	}
	CreateFilesetStub        func(filesystemName string, filesetName string, opts map[string]interface{}) error
	createFilesetMutex       sync.RWMutex
	createFilesetArgsForCall []struct {
		filesystemName string
		filesetName    string
		opts           map[string]interface{}
	}
	createFilesetReturns struct {
		result1 error
	}
	DeleteFilesetStub        func(filesystemName string, filesetName string) error
	deleteFilesetMutex       sync.RWMutex
	deleteFilesetArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	deleteFilesetReturns struct {
		result1 error
	}
	LinkFilesetStub        func(filesystemName string, filesetName string) error
	linkFilesetMutex       sync.RWMutex
	linkFilesetArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	linkFilesetReturns struct {
		result1 error
	}
	UnlinkFilesetStub        func(filesystemName string, filesetName string) error
	unlinkFilesetMutex       sync.RWMutex
	unlinkFilesetArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	unlinkFilesetReturns struct {
		result1 error
	}
	ListFilesetsStub        func(filesystemName string) ([]resources.VolumeMetadata, error)
	listFilesetsMutex       sync.RWMutex
	listFilesetsArgsForCall []struct {
		filesystemName string
	}
	listFilesetsReturns struct {
		result1 []resources.VolumeMetadata
		result2 error
	}
	ListFilesetStub        func(filesystemName string, filesetName string) (resources.VolumeMetadata, error)
	listFilesetMutex       sync.RWMutex
	listFilesetArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	listFilesetReturns struct {
		result1 resources.VolumeMetadata
		result2 error
	}
	IsFilesetLinkedStub        func(filesystemName string, filesetName string) (bool, error)
	isFilesetLinkedMutex       sync.RWMutex
	isFilesetLinkedArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	isFilesetLinkedReturns struct {
		result1 bool
		result2 error
	}
	ListFilesetQuotaStub        func(filesystemName string, filesetName string) (string, error)
	listFilesetQuotaMutex       sync.RWMutex
	listFilesetQuotaArgsForCall []struct {
		filesystemName string
		filesetName    string
	}
	listFilesetQuotaReturns struct {
		result1 string
		result2 error
	}
	SetFilesetQuotaStub        func(filesystemName string, filesetName string, quota string) error
	setFilesetQuotaMutex       sync.RWMutex
	setFilesetQuotaArgsForCall []struct {
		filesystemName string
		filesetName    string
		quota          string
	}
	setFilesetQuotaReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpectrumScaleConnector) GetClusterId() (string, error) {
	fake.getClusterIdMutex.Lock()
	fake.getClusterIdArgsForCall = append(fake.getClusterIdArgsForCall, struct{}{})
	fake.recordInvocation("GetClusterId", []interface{}{})
	fake.getClusterIdMutex.Unlock()
	if fake.GetClusterIdStub != nil {
		return fake.GetClusterIdStub()
	}
	return fake.getClusterIdReturns.result1, fake.getClusterIdReturns.result2
}

func (fake *FakeSpectrumScaleConnector) GetClusterIdCallCount() int {
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	return len(fake.getClusterIdArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) GetClusterIdReturns(result1 string, result2 error) {
	fake.GetClusterIdStub = nil
	fake.getClusterIdReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) IsFilesystemMounted(filesystemName string) (bool, error) {
	fake.isFilesystemMountedMutex.Lock()
	fake.isFilesystemMountedArgsForCall = append(fake.isFilesystemMountedArgsForCall, struct {
		filesystemName string
	}{filesystemName})
	fake.recordInvocation("IsFilesystemMounted", []interface{}{filesystemName})
	fake.isFilesystemMountedMutex.Unlock()
	if fake.IsFilesystemMountedStub != nil {
		return fake.IsFilesystemMountedStub(filesystemName)
	}
	return fake.isFilesystemMountedReturns.result1, fake.isFilesystemMountedReturns.result2
}

func (fake *FakeSpectrumScaleConnector) IsFilesystemMountedCallCount() int {
	fake.isFilesystemMountedMutex.RLock()
	defer fake.isFilesystemMountedMutex.RUnlock()
	return len(fake.isFilesystemMountedArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) IsFilesystemMountedArgsForCall(i int) string {
	fake.isFilesystemMountedMutex.RLock()
	defer fake.isFilesystemMountedMutex.RUnlock()
	return fake.isFilesystemMountedArgsForCall[i].filesystemName
}

func (fake *FakeSpectrumScaleConnector) IsFilesystemMountedReturns(result1 bool, result2 error) {
	fake.IsFilesystemMountedStub = nil
	fake.isFilesystemMountedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) MountFileSystem(filesystemName string) error {
	fake.mountFileSystemMutex.Lock()
	fake.mountFileSystemArgsForCall = append(fake.mountFileSystemArgsForCall, struct {
		filesystemName string
	}{filesystemName})
	fake.recordInvocation("MountFileSystem", []interface{}{filesystemName})
	fake.mountFileSystemMutex.Unlock()
	if fake.MountFileSystemStub != nil {
		return fake.MountFileSystemStub(filesystemName)
	}
	return fake.mountFileSystemReturns.result1
}

func (fake *FakeSpectrumScaleConnector) MountFileSystemCallCount() int {
	fake.mountFileSystemMutex.RLock()
	defer fake.mountFileSystemMutex.RUnlock()
	return len(fake.mountFileSystemArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) MountFileSystemArgsForCall(i int) string {
	fake.mountFileSystemMutex.RLock()
	defer fake.mountFileSystemMutex.RUnlock()
	return fake.mountFileSystemArgsForCall[i].filesystemName
}

func (fake *FakeSpectrumScaleConnector) MountFileSystemReturns(result1 error) {
	fake.MountFileSystemStub = nil
	fake.mountFileSystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) ListFilesystems() ([]string, error) {
	fake.listFilesystemsMutex.Lock()
	fake.listFilesystemsArgsForCall = append(fake.listFilesystemsArgsForCall, struct{}{})
	fake.recordInvocation("ListFilesystems", []interface{}{})
	fake.listFilesystemsMutex.Unlock()
	if fake.ListFilesystemsStub != nil {
		return fake.ListFilesystemsStub()
	}
	return fake.listFilesystemsReturns.result1, fake.listFilesystemsReturns.result2
}

func (fake *FakeSpectrumScaleConnector) ListFilesystemsCallCount() int {
	fake.listFilesystemsMutex.RLock()
	defer fake.listFilesystemsMutex.RUnlock()
	return len(fake.listFilesystemsArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) ListFilesystemsReturns(result1 []string, result2 error) {
	fake.ListFilesystemsStub = nil
	fake.listFilesystemsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) GetFilesystemMountpoint(filesystemName string) (string, error) {
	fake.getFilesystemMountpointMutex.Lock()
	fake.getFilesystemMountpointArgsForCall = append(fake.getFilesystemMountpointArgsForCall, struct {
		filesystemName string
	}{filesystemName})
	fake.recordInvocation("GetFilesystemMountpoint", []interface{}{filesystemName})
	fake.getFilesystemMountpointMutex.Unlock()
	if fake.GetFilesystemMountpointStub != nil {
		return fake.GetFilesystemMountpointStub(filesystemName)
	}
	return fake.getFilesystemMountpointReturns.result1, fake.getFilesystemMountpointReturns.result2
}

func (fake *FakeSpectrumScaleConnector) GetFilesystemMountpointCallCount() int {
	fake.getFilesystemMountpointMutex.RLock()
	defer fake.getFilesystemMountpointMutex.RUnlock()
	return len(fake.getFilesystemMountpointArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) GetFilesystemMountpointArgsForCall(i int) string {
	fake.getFilesystemMountpointMutex.RLock()
	defer fake.getFilesystemMountpointMutex.RUnlock()
	return fake.getFilesystemMountpointArgsForCall[i].filesystemName
}

func (fake *FakeSpectrumScaleConnector) GetFilesystemMountpointReturns(result1 string, result2 error) {
	fake.GetFilesystemMountpointStub = nil
	fake.getFilesystemMountpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) CreateFileset(filesystemName string, filesetName string, opts map[string]interface{}) error {
	fake.createFilesetMutex.Lock()
	fake.createFilesetArgsForCall = append(fake.createFilesetArgsForCall, struct {
		filesystemName string
		filesetName    string
		opts           map[string]interface{}
	}{filesystemName, filesetName, opts})
	fake.recordInvocation("CreateFileset", []interface{}{filesystemName, filesetName, opts})
	fake.createFilesetMutex.Unlock()
	if fake.CreateFilesetStub != nil {
		return fake.CreateFilesetStub(filesystemName, filesetName, opts)
	}
	return fake.createFilesetReturns.result1
}

func (fake *FakeSpectrumScaleConnector) CreateFilesetCallCount() int {
	fake.createFilesetMutex.RLock()
	defer fake.createFilesetMutex.RUnlock()
	return len(fake.createFilesetArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) CreateFilesetArgsForCall(i int) (string, string, map[string]interface{}) {
	fake.createFilesetMutex.RLock()
	defer fake.createFilesetMutex.RUnlock()
	return fake.createFilesetArgsForCall[i].filesystemName, fake.createFilesetArgsForCall[i].filesetName, fake.createFilesetArgsForCall[i].opts
}

func (fake *FakeSpectrumScaleConnector) CreateFilesetReturns(result1 error) {
	fake.CreateFilesetStub = nil
	fake.createFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) DeleteFileset(filesystemName string, filesetName string) error {
	fake.deleteFilesetMutex.Lock()
	fake.deleteFilesetArgsForCall = append(fake.deleteFilesetArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("DeleteFileset", []interface{}{filesystemName, filesetName})
	fake.deleteFilesetMutex.Unlock()
	if fake.DeleteFilesetStub != nil {
		return fake.DeleteFilesetStub(filesystemName, filesetName)
	}
	return fake.deleteFilesetReturns.result1
}

func (fake *FakeSpectrumScaleConnector) DeleteFilesetCallCount() int {
	fake.deleteFilesetMutex.RLock()
	defer fake.deleteFilesetMutex.RUnlock()
	return len(fake.deleteFilesetArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) DeleteFilesetArgsForCall(i int) (string, string) {
	fake.deleteFilesetMutex.RLock()
	defer fake.deleteFilesetMutex.RUnlock()
	return fake.deleteFilesetArgsForCall[i].filesystemName, fake.deleteFilesetArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) DeleteFilesetReturns(result1 error) {
	fake.DeleteFilesetStub = nil
	fake.deleteFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) LinkFileset(filesystemName string, filesetName string) error {
	fake.linkFilesetMutex.Lock()
	fake.linkFilesetArgsForCall = append(fake.linkFilesetArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("LinkFileset", []interface{}{filesystemName, filesetName})
	fake.linkFilesetMutex.Unlock()
	if fake.LinkFilesetStub != nil {
		return fake.LinkFilesetStub(filesystemName, filesetName)
	}
	return fake.linkFilesetReturns.result1
}

func (fake *FakeSpectrumScaleConnector) LinkFilesetCallCount() int {
	fake.linkFilesetMutex.RLock()
	defer fake.linkFilesetMutex.RUnlock()
	return len(fake.linkFilesetArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) LinkFilesetArgsForCall(i int) (string, string) {
	fake.linkFilesetMutex.RLock()
	defer fake.linkFilesetMutex.RUnlock()
	return fake.linkFilesetArgsForCall[i].filesystemName, fake.linkFilesetArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) LinkFilesetReturns(result1 error) {
	fake.LinkFilesetStub = nil
	fake.linkFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) UnlinkFileset(filesystemName string, filesetName string) error {
	fake.unlinkFilesetMutex.Lock()
	fake.unlinkFilesetArgsForCall = append(fake.unlinkFilesetArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("UnlinkFileset", []interface{}{filesystemName, filesetName})
	fake.unlinkFilesetMutex.Unlock()
	if fake.UnlinkFilesetStub != nil {
		return fake.UnlinkFilesetStub(filesystemName, filesetName)
	}
	return fake.unlinkFilesetReturns.result1
}

func (fake *FakeSpectrumScaleConnector) UnlinkFilesetCallCount() int {
	fake.unlinkFilesetMutex.RLock()
	defer fake.unlinkFilesetMutex.RUnlock()
	return len(fake.unlinkFilesetArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) UnlinkFilesetArgsForCall(i int) (string, string) {
	fake.unlinkFilesetMutex.RLock()
	defer fake.unlinkFilesetMutex.RUnlock()
	return fake.unlinkFilesetArgsForCall[i].filesystemName, fake.unlinkFilesetArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) UnlinkFilesetReturns(result1 error) {
	fake.UnlinkFilesetStub = nil
	fake.unlinkFilesetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) ListFilesets(filesystemName string) ([]resources.VolumeMetadata, error) {
	fake.listFilesetsMutex.Lock()
	fake.listFilesetsArgsForCall = append(fake.listFilesetsArgsForCall, struct {
		filesystemName string
	}{filesystemName})
	fake.recordInvocation("ListFilesets", []interface{}{filesystemName})
	fake.listFilesetsMutex.Unlock()
	if fake.ListFilesetsStub != nil {
		return fake.ListFilesetsStub(filesystemName)
	}
	return fake.listFilesetsReturns.result1, fake.listFilesetsReturns.result2
}

func (fake *FakeSpectrumScaleConnector) ListFilesetsCallCount() int {
	fake.listFilesetsMutex.RLock()
	defer fake.listFilesetsMutex.RUnlock()
	return len(fake.listFilesetsArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) ListFilesetsArgsForCall(i int) string {
	fake.listFilesetsMutex.RLock()
	defer fake.listFilesetsMutex.RUnlock()
	return fake.listFilesetsArgsForCall[i].filesystemName
}

func (fake *FakeSpectrumScaleConnector) ListFilesetsReturns(result1 []resources.VolumeMetadata, result2 error) {
	fake.ListFilesetsStub = nil
	fake.listFilesetsReturns = struct {
		result1 []resources.VolumeMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) ListFileset(filesystemName string, filesetName string) (resources.VolumeMetadata, error) {
	fake.listFilesetMutex.Lock()
	fake.listFilesetArgsForCall = append(fake.listFilesetArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("ListFileset", []interface{}{filesystemName, filesetName})
	fake.listFilesetMutex.Unlock()
	if fake.ListFilesetStub != nil {
		return fake.ListFilesetStub(filesystemName, filesetName)
	}
	return fake.listFilesetReturns.result1, fake.listFilesetReturns.result2
}

func (fake *FakeSpectrumScaleConnector) ListFilesetCallCount() int {
	fake.listFilesetMutex.RLock()
	defer fake.listFilesetMutex.RUnlock()
	return len(fake.listFilesetArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) ListFilesetArgsForCall(i int) (string, string) {
	fake.listFilesetMutex.RLock()
	defer fake.listFilesetMutex.RUnlock()
	return fake.listFilesetArgsForCall[i].filesystemName, fake.listFilesetArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) ListFilesetReturns(result1 resources.VolumeMetadata, result2 error) {
	fake.ListFilesetStub = nil
	fake.listFilesetReturns = struct {
		result1 resources.VolumeMetadata
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) IsFilesetLinked(filesystemName string, filesetName string) (bool, error) {
	fake.isFilesetLinkedMutex.Lock()
	fake.isFilesetLinkedArgsForCall = append(fake.isFilesetLinkedArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("IsFilesetLinked", []interface{}{filesystemName, filesetName})
	fake.isFilesetLinkedMutex.Unlock()
	if fake.IsFilesetLinkedStub != nil {
		return fake.IsFilesetLinkedStub(filesystemName, filesetName)
	}
	return fake.isFilesetLinkedReturns.result1, fake.isFilesetLinkedReturns.result2
}

func (fake *FakeSpectrumScaleConnector) IsFilesetLinkedCallCount() int {
	fake.isFilesetLinkedMutex.RLock()
	defer fake.isFilesetLinkedMutex.RUnlock()
	return len(fake.isFilesetLinkedArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) IsFilesetLinkedArgsForCall(i int) (string, string) {
	fake.isFilesetLinkedMutex.RLock()
	defer fake.isFilesetLinkedMutex.RUnlock()
	return fake.isFilesetLinkedArgsForCall[i].filesystemName, fake.isFilesetLinkedArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) IsFilesetLinkedReturns(result1 bool, result2 error) {
	fake.IsFilesetLinkedStub = nil
	fake.isFilesetLinkedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) ListFilesetQuota(filesystemName string, filesetName string) (string, error) {
	fake.listFilesetQuotaMutex.Lock()
	fake.listFilesetQuotaArgsForCall = append(fake.listFilesetQuotaArgsForCall, struct {
		filesystemName string
		filesetName    string
	}{filesystemName, filesetName})
	fake.recordInvocation("ListFilesetQuota", []interface{}{filesystemName, filesetName})
	fake.listFilesetQuotaMutex.Unlock()
	if fake.ListFilesetQuotaStub != nil {
		return fake.ListFilesetQuotaStub(filesystemName, filesetName)
	}
	return fake.listFilesetQuotaReturns.result1, fake.listFilesetQuotaReturns.result2
}

func (fake *FakeSpectrumScaleConnector) ListFilesetQuotaCallCount() int {
	fake.listFilesetQuotaMutex.RLock()
	defer fake.listFilesetQuotaMutex.RUnlock()
	return len(fake.listFilesetQuotaArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) ListFilesetQuotaArgsForCall(i int) (string, string) {
	fake.listFilesetQuotaMutex.RLock()
	defer fake.listFilesetQuotaMutex.RUnlock()
	return fake.listFilesetQuotaArgsForCall[i].filesystemName, fake.listFilesetQuotaArgsForCall[i].filesetName
}

func (fake *FakeSpectrumScaleConnector) ListFilesetQuotaReturns(result1 string, result2 error) {
	fake.ListFilesetQuotaStub = nil
	fake.listFilesetQuotaReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumScaleConnector) SetFilesetQuota(filesystemName string, filesetName string, quota string) error {
	fake.setFilesetQuotaMutex.Lock()
	fake.setFilesetQuotaArgsForCall = append(fake.setFilesetQuotaArgsForCall, struct {
		filesystemName string
		filesetName    string
		quota          string
	}{filesystemName, filesetName, quota})
	fake.recordInvocation("SetFilesetQuota", []interface{}{filesystemName, filesetName, quota})
	fake.setFilesetQuotaMutex.Unlock()
	if fake.SetFilesetQuotaStub != nil {
		return fake.SetFilesetQuotaStub(filesystemName, filesetName, quota)
	}
	return fake.setFilesetQuotaReturns.result1
}

func (fake *FakeSpectrumScaleConnector) SetFilesetQuotaCallCount() int {
	fake.setFilesetQuotaMutex.RLock()
	defer fake.setFilesetQuotaMutex.RUnlock()
	return len(fake.setFilesetQuotaArgsForCall)
}

func (fake *FakeSpectrumScaleConnector) SetFilesetQuotaArgsForCall(i int) (string, string, string) {
	fake.setFilesetQuotaMutex.RLock()
	defer fake.setFilesetQuotaMutex.RUnlock()
	return fake.setFilesetQuotaArgsForCall[i].filesystemName, fake.setFilesetQuotaArgsForCall[i].filesetName, fake.setFilesetQuotaArgsForCall[i].quota
}

func (fake *FakeSpectrumScaleConnector) SetFilesetQuotaReturns(result1 error) {
	fake.SetFilesetQuotaStub = nil
	fake.setFilesetQuotaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumScaleConnector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	fake.isFilesystemMountedMutex.RLock()
	defer fake.isFilesystemMountedMutex.RUnlock()
	fake.mountFileSystemMutex.RLock()
	defer fake.mountFileSystemMutex.RUnlock()
	fake.listFilesystemsMutex.RLock()
	defer fake.listFilesystemsMutex.RUnlock()
	fake.getFilesystemMountpointMutex.RLock()
	defer fake.getFilesystemMountpointMutex.RUnlock()
	fake.createFilesetMutex.RLock()
	defer fake.createFilesetMutex.RUnlock()
	fake.deleteFilesetMutex.RLock()
	defer fake.deleteFilesetMutex.RUnlock()
	fake.linkFilesetMutex.RLock()
	defer fake.linkFilesetMutex.RUnlock()
	fake.unlinkFilesetMutex.RLock()
	defer fake.unlinkFilesetMutex.RUnlock()
	fake.listFilesetsMutex.RLock()
	defer fake.listFilesetsMutex.RUnlock()
	fake.listFilesetMutex.RLock()
	defer fake.listFilesetMutex.RUnlock()
	fake.isFilesetLinkedMutex.RLock()
	defer fake.isFilesetLinkedMutex.RUnlock()
	fake.listFilesetQuotaMutex.RLock()
	defer fake.listFilesetQuotaMutex.RUnlock()
	fake.setFilesetQuotaMutex.RLock()
	defer fake.setFilesetQuotaMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpectrumScaleConnector) recordInvocation(key string, args []interface{}) {
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

var _ connectors.SpectrumScaleConnector = new(FakeSpectrumScaleConnector)
