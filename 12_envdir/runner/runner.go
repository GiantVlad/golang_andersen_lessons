package runner

import (
	"log"
	"os"
	"os/exec"
)

func ReadDir(dir string) (map[string]string, error) {
	envs := make(map[string]string)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, de := range files {
		if de.IsDir() {
			continue
		}
		if de.Name() == "" {
			continue
		}
		bites, _ := os.ReadFile(dir + "/" + de.Name())
		envs[de.Name()] = string(bites)
	}
	return envs, err
}

func RunCmd(cmd []string, env map[string]string) int {
	for key, val := range env {
		err := os.Setenv(key, val)
		if err != nil {
			continue
		}
	}
	command := exec.Command(cmd[0], cmd[1:]...)
	err := command.Start()
	if err != nil {
		log.Fatal(err) // ошибка запуска
	}
	log.Printf("Waiting for command to finish...")
	err = command.Wait() // ошибка выполнения
	log.Printf("Command finished with error: %v", err)
	return 0
}
