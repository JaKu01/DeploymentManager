package tools

import (
	"github.com/JaKu01/DeploymentManager/model"
)

func FullDeployment(project *model.Project) error {
	if !DoesProjectExist(project) {
		err := CloneProject(project)
		if err != nil {
			return err
		}
	}

	err := CheckoutBranch(project)
	if err != nil {
		return err
	}

	err = PullProject(project)
	if err != nil {
		return err
	}

	err = DeployProject(project)
	if err != nil {
		return err
	}

	return nil
}
