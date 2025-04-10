package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils"
	"errors"
	"fmt"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type GitService struct{}

var fileUtil *utils.FileUtil

func getAccessTocken() string {
	return os.Getenv("REPO_ACCESS_TOKEN")
}

func getRepoLocalPath() string {
	return os.Getenv("REPO_LOCAL_PATH")
}

func (g *GitService) cloneRepo() (*git.Repository, error) {
	var repoURL = os.Getenv("REPO_REMOTE_PATH") + ".git"
	var accessToken = getAccessTocken()
	var username = os.Getenv("GIT_USERNAME")
	var repoLocalPath = getRepoLocalPath()

	if _, err := os.Stat(repoLocalPath); os.IsNotExist(err) {
		return git.PlainClone(repoLocalPath, false, &git.CloneOptions{
			URL: repoURL,
			Auth: &http.BasicAuth{
				Username: username, // any non-empty string works
				Password: accessToken,
			},
		})
	}
	return git.PlainOpen(repoLocalPath)
}

func (g *GitService) writeFileToRepo(repo *git.Repository, filePathInRepo string, fileBody string) error {
	// 📄 Write your file
	filePath := getRepoLocalPath() + "/" + filePathInRepo
	err := fileUtil.CreateFileIfNotExists(filePath, fileBody)

	if err != nil {
		fmt.Println(err.Error())
		return errors.New("Error creating file")
	}

	// ✅ Add, commit, push
	w, err := repo.Worktree()
	if err != nil {
		return errors.New("Work Tree Error")
	}

	fmt.Println(filePathInRepo)
	_, err = w.Add(filePathInRepo)
	if err != nil {
		return errors.New("Error Adding file to git")
	}

	_, err = w.Commit("Automated commit from Go", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Go Bot",
			Email: "bot@example.com",
			When:  time.Now(),
		},
	})

	if err != nil {
		return errors.New("Error commiting to git")
	}

	err = repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: "your-username",
			Password: getAccessTocken(),
		},
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

	return nil
}

func (g *GitService) PushToGitHub(path string, content string) error {
	repo, err := g.cloneRepo()
	if err != nil {
		return err
	}
	return g.writeFileToRepo(repo, path, content)
}
