package application

import "cchalop1.com/deploy/internal/api/service"

func StartApplication(deployService *service.DeployService, deployId string) error {
	deploy, err := deployService.DatabaseAdapter.GetDeployById(deployId)
	if err != nil {
		return err
	}
	// TODO: connect to docker client
	deployService.DockerAdapter.Start(deploy.Name)
	deploy.Status = "Runing"
	deployService.DatabaseAdapter.UpdateDeploy(deploy)
	return nil
}
