package application

import (
	"cchalop1.com/deploy/internal/api/dto"
	"cchalop1.com/deploy/internal/api/service"
)

func GetApplicationLogs(deployService *service.DeployService, deployId string) ([]dto.Logs, error) {
	deploy, err := deployService.DatabaseAdapter.GetDeployById(deployId)
	logs := []dto.Logs{}

	if err != nil {
		return logs, err
	}

	server, err := deployService.DatabaseAdapter.GetServerById(deploy.ServerId)

	if err != nil {
		return logs, err
	}

	// TODO: move all the connect client to a midleware
	deployService.DockerAdapter.ConnectClient(server)

	dockerLogs := deployService.DockerAdapter.GetLogsOfContainer(deploy.GetDockerName())

	for _, log := range dockerLogs {

		if len(log) < 30 {
			continue
		}

		datePart := log[0:30]
		messagePart := log[30:]
		logs = append(logs, dto.Logs{
			Date:    datePart,
			Message: messagePart,
		})
	}

	return logs, nil
}
