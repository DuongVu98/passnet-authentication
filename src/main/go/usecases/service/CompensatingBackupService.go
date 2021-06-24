package service

import "github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"

type CompensatingBackupService struct {
	compensatingBackup map[string]command.BaseCompensating
}

func (service CompensatingBackupService) Store(compensating command.BaseCompensating, eventId string) {
	service.compensatingBackup[eventId] = compensating
}

var compensatingBackupService = CompensatingBackupService{compensatingBackup: make(map[string]command.BaseCompensating)}
func GetCompensatingBackupService() CompensatingBackupService {
	return compensatingBackupService
}
