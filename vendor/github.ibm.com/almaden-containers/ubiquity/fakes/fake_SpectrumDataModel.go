// This file was generated by counterfeiter
package fakes

import (
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.ibm.com/almaden-containers/ubiquity/local/spectrumscale"
)

type FakeSpectrumDataModel struct {
	CreateVolumeTableStub        func() error
	createVolumeTableMutex       sync.RWMutex
	createVolumeTableArgsForCall []struct{}
	createVolumeTableReturns     struct {
		result1 error
	}
	SetClusterIdStub        func(string)
	setClusterIdMutex       sync.RWMutex
	setClusterIdArgsForCall []struct {
		arg1 string
	}
	GetClusterIdStub        func() string
	getClusterIdMutex       sync.RWMutex
	getClusterIdArgsForCall []struct{}
	getClusterIdReturns     struct {
		result1 string
	}
	DeleteVolumeStub        func(name string) error
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		name string
	}
	deleteVolumeReturns struct {
		result1 error
	}
	InsertFilesetVolumeStub        func(fileset, volumeName string, filesystem string, opts map[string]interface{}) error
	insertFilesetVolumeMutex       sync.RWMutex
	insertFilesetVolumeArgsForCall []struct {
		fileset    string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}
	insertFilesetVolumeReturns struct {
		result1 error
	}
	InsertLightweightVolumeStub        func(fileset, directory, volumeName string, filesystem string, opts map[string]interface{}) error
	insertLightweightVolumeMutex       sync.RWMutex
	insertLightweightVolumeArgsForCall []struct {
		fileset    string
		directory  string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}
	insertLightweightVolumeReturns struct {
		result1 error
	}
	InsertFilesetQuotaVolumeStub        func(fileset, quota, volumeName string, filesystem string, opts map[string]interface{}) error
	insertFilesetQuotaVolumeMutex       sync.RWMutex
	insertFilesetQuotaVolumeArgsForCall []struct {
		fileset    string
		quota      string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}
	insertFilesetQuotaVolumeReturns struct {
		result1 error
	}
	GetVolumeStub        func(name string) (spectrumscale.SpectrumScaleVolume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		name string
	}
	getVolumeReturns struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}
	ListVolumesStub        func() ([]spectrumscale.SpectrumScaleVolume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpectrumDataModel) CreateVolumeTable() error {
	fake.createVolumeTableMutex.Lock()
	fake.createVolumeTableArgsForCall = append(fake.createVolumeTableArgsForCall, struct{}{})
	fake.recordInvocation("CreateVolumeTable", []interface{}{})
	fake.createVolumeTableMutex.Unlock()
	if fake.CreateVolumeTableStub != nil {
		return fake.CreateVolumeTableStub()
	}
	return fake.createVolumeTableReturns.result1
}

func (fake *FakeSpectrumDataModel) CreateVolumeTableCallCount() int {
	fake.createVolumeTableMutex.RLock()
	defer fake.createVolumeTableMutex.RUnlock()
	return len(fake.createVolumeTableArgsForCall)
}

func (fake *FakeSpectrumDataModel) CreateVolumeTableReturns(result1 error) {
	fake.CreateVolumeTableStub = nil
	fake.createVolumeTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) SetClusterId(arg1 string) {
	fake.setClusterIdMutex.Lock()
	fake.setClusterIdArgsForCall = append(fake.setClusterIdArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetClusterId", []interface{}{arg1})
	fake.setClusterIdMutex.Unlock()
	if fake.SetClusterIdStub != nil {
		fake.SetClusterIdStub(arg1)
	}
}

func (fake *FakeSpectrumDataModel) SetClusterIdCallCount() int {
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	return len(fake.setClusterIdArgsForCall)
}

func (fake *FakeSpectrumDataModel) SetClusterIdArgsForCall(i int) string {
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	return fake.setClusterIdArgsForCall[i].arg1
}

func (fake *FakeSpectrumDataModel) GetClusterId() string {
	fake.getClusterIdMutex.Lock()
	fake.getClusterIdArgsForCall = append(fake.getClusterIdArgsForCall, struct{}{})
	fake.recordInvocation("GetClusterId", []interface{}{})
	fake.getClusterIdMutex.Unlock()
	if fake.GetClusterIdStub != nil {
		return fake.GetClusterIdStub()
	}
	return fake.getClusterIdReturns.result1
}

func (fake *FakeSpectrumDataModel) GetClusterIdCallCount() int {
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	return len(fake.getClusterIdArgsForCall)
}

func (fake *FakeSpectrumDataModel) GetClusterIdReturns(result1 string) {
	fake.GetClusterIdStub = nil
	fake.getClusterIdReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpectrumDataModel) DeleteVolume(name string) error {
	fake.deleteVolumeMutex.Lock()
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteVolume", []interface{}{name})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(name)
	}
	return fake.deleteVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) DeleteVolumeArgsForCall(i int) string {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return fake.deleteVolumeArgsForCall[i].name
}

func (fake *FakeSpectrumDataModel) DeleteVolumeReturns(result1 error) {
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolume(fileset string, volumeName string, filesystem string, opts map[string]interface{}) error {
	fake.insertFilesetVolumeMutex.Lock()
	fake.insertFilesetVolumeArgsForCall = append(fake.insertFilesetVolumeArgsForCall, struct {
		fileset    string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}{fileset, volumeName, filesystem, opts})
	fake.recordInvocation("InsertFilesetVolume", []interface{}{fileset, volumeName, filesystem, opts})
	fake.insertFilesetVolumeMutex.Unlock()
	if fake.InsertFilesetVolumeStub != nil {
		return fake.InsertFilesetVolumeStub(fileset, volumeName, filesystem, opts)
	}
	return fake.insertFilesetVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeCallCount() int {
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	return len(fake.insertFilesetVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeArgsForCall(i int) (string, string, string, map[string]interface{}) {
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	return fake.insertFilesetVolumeArgsForCall[i].fileset, fake.insertFilesetVolumeArgsForCall[i].volumeName, fake.insertFilesetVolumeArgsForCall[i].filesystem, fake.insertFilesetVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeReturns(result1 error) {
	fake.InsertFilesetVolumeStub = nil
	fake.insertFilesetVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolume(fileset string, directory string, volumeName string, filesystem string, opts map[string]interface{}) error {
	fake.insertLightweightVolumeMutex.Lock()
	fake.insertLightweightVolumeArgsForCall = append(fake.insertLightweightVolumeArgsForCall, struct {
		fileset    string
		directory  string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}{fileset, directory, volumeName, filesystem, opts})
	fake.recordInvocation("InsertLightweightVolume", []interface{}{fileset, directory, volumeName, filesystem, opts})
	fake.insertLightweightVolumeMutex.Unlock()
	if fake.InsertLightweightVolumeStub != nil {
		return fake.InsertLightweightVolumeStub(fileset, directory, volumeName, filesystem, opts)
	}
	return fake.insertLightweightVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeCallCount() int {
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	return len(fake.insertLightweightVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeArgsForCall(i int) (string, string, string, string, map[string]interface{}) {
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	return fake.insertLightweightVolumeArgsForCall[i].fileset, fake.insertLightweightVolumeArgsForCall[i].directory, fake.insertLightweightVolumeArgsForCall[i].volumeName, fake.insertLightweightVolumeArgsForCall[i].filesystem, fake.insertLightweightVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeReturns(result1 error) {
	fake.InsertLightweightVolumeStub = nil
	fake.insertLightweightVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolume(fileset string, quota string, volumeName string, filesystem string, opts map[string]interface{}) error {
	fake.insertFilesetQuotaVolumeMutex.Lock()
	fake.insertFilesetQuotaVolumeArgsForCall = append(fake.insertFilesetQuotaVolumeArgsForCall, struct {
		fileset    string
		quota      string
		volumeName string
		filesystem string
		opts       map[string]interface{}
	}{fileset, quota, volumeName, filesystem, opts})
	fake.recordInvocation("InsertFilesetQuotaVolume", []interface{}{fileset, quota, volumeName, filesystem, opts})
	fake.insertFilesetQuotaVolumeMutex.Unlock()
	if fake.InsertFilesetQuotaVolumeStub != nil {
		return fake.InsertFilesetQuotaVolumeStub(fileset, quota, volumeName, filesystem, opts)
	}
	return fake.insertFilesetQuotaVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeCallCount() int {
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	return len(fake.insertFilesetQuotaVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeArgsForCall(i int) (string, string, string, string, map[string]interface{}) {
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	return fake.insertFilesetQuotaVolumeArgsForCall[i].fileset, fake.insertFilesetQuotaVolumeArgsForCall[i].quota, fake.insertFilesetQuotaVolumeArgsForCall[i].volumeName, fake.insertFilesetQuotaVolumeArgsForCall[i].filesystem, fake.insertFilesetQuotaVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeReturns(result1 error) {
	fake.InsertFilesetQuotaVolumeStub = nil
	fake.insertFilesetQuotaVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) GetVolume(name string) (spectrumscale.SpectrumScaleVolume, bool, error) {
	fake.getVolumeMutex.Lock()
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVolume", []interface{}{name})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(name)
	}
	return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
}

func (fake *FakeSpectrumDataModel) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].name
}

func (fake *FakeSpectrumDataModel) GetVolumeReturns(result1 spectrumscale.SpectrumScaleVolume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpectrumDataModel) ListVolumes() ([]spectrumscale.SpectrumScaleVolume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
}

func (fake *FakeSpectrumDataModel) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeSpectrumDataModel) ListVolumesReturns(result1 []spectrumscale.SpectrumScaleVolume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumDataModel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeTableMutex.RLock()
	defer fake.createVolumeTableMutex.RUnlock()
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpectrumDataModel) recordInvocation(key string, args []interface{}) {
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

var _ spectrumscale.SpectrumDataModel = new(FakeSpectrumDataModel)
