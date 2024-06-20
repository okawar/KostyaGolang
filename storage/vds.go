package storage

import (
	"app/entity"
	"errors"
	"sync"
)

type VdsStorage struct {
	vdsMap map[uint64]entity.Vds
	next   uint64
	mutx   sync.RWMutex
}

var s *VdsStorage

func init() {
	s = new(VdsStorage)
}

func CreateVds(vds entity.Vds) entity.Vds {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	vds.Vid = s.next
	s.next++
	s.vdsMap[vds.Vid] = vds
	return vds
}

func UpdateVds(vds entity.Vds) error {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if _, exists := s.vdsMap[vds.Vid]; !exists {
		return errors.New("Vds не найден")
	}
	s.vdsMap[vds.Vid] = vds
	return nil
}

func DeleteVds(vid uint64) error {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	if _, exists := s.vdsMap[vid]; !exists {
		return errors.New("Vds не найден")
	}
	delete(s.vdsMap, vid)
	return nil
}

func GetAllVds() ([]entity.Vds, error) {
	s.mutx.RLock()
	s.mutx.RUnlock()
	vdsList := make([]entity.Vds, 0, len(s.vdsMap))
	for _, vds := range s.vdsMap {
		vdsList = append(vdsList, vds)
	}
	return vdsList, nil
}

func GetVdsByID(vid uint64) (*entity.Vds, error) {
	s.mutx.RLock()
	s.mutx.RUnlock()
	vds, exists := s.vdsMap[vid]
	if !exists {
		return nil, errors.New("Vds не найден")
	}
	return &vds, nil
}
