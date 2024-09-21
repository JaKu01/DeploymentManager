package tools

import (
	"github.com/JaKu01/DeploymentManager/model"
	"log"
	"os"
	"os/exec"
)

func CloneProject(project *model.Project) error {
	err := CreateProjectFolder(project)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", project.Url, ".")
	cmd.Dir = "../" + project.Name
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	log.Println("CloneProject: " + string(output))
	return err
}

func CheckoutBranch(project *model.Project) error {
	cmd := exec.Command("git", "checkout", project.Branch)
	cmd.Dir = "../" + project.Name
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	log.Println("Checkout Branch: " + string(output))
	return err
}

func PullProject(project *model.Project) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = "../" + project.Name
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	log.Println("Pull Project: " + string(output))
	return err
}

func DeployProject(project *model.Project) error {
	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Dir = "../" + project.Name + "/" + project.DockerComposePath
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	log.Println("Deploy Project: " + string(output))
	return err
}

func CreateProjectFolder(project *model.Project) error {
	if DoesProjectExist(project) {
		return nil
	}
	err := os.Mkdir("../"+project.Name, 0755)
	return err
}

func DoesProjectExist(project *model.Project) bool {
	_, err := os.Stat("../" + project.Name)
	return !os.IsNotExist(err)
}
