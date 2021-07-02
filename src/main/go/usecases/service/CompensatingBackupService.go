package service

import "github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"

type CompensatingBackupService struct {
	compensatingBackup map[string]command.BaseCompensating
}

func (service CompensatingBackupService) Store(compensating command.BaseCompensating, eventId string) {
	service.compensatingBackup[eventId] = compensating
}

func (service CompensatingBackupService) Find(eventId string) command.BaseCompensating {
	return service.compensatingBackup[eventId]
}

func (service CompensatingBackupService) Remove(eventId string) {
	delete(service.compensatingBackup, eventId)
}

var compensatingBackupService = CompensatingBackupService{compensatingBackup: make(map[string]command.BaseCompensating)}
func GetCompensatingBackupService() CompensatingBackupService {
	return compensatingBackupService
}
