package github

import (
	"testing"

	"github.com/google/go-github/github"
	"github.com/steadyequipment/goutility"

	"math/rand"
	"time"
)

const (
	accessToken = "REPLACE_WITH_TOKEN"

	existingOrgName  = "REPLACE_WITH_EXISTING_ORG"
	existingRepoName = "REPLACE_WITH_EXISTING_REPO"
	existingFileName = "REPLACE_WITH_EXISTING_FILE_WITH_NO_LEADING_SLASH"

	authorName  = "REPLACE_WITH_AUTHOR_NAME"
	authorEmail = "REPLACE_WITH_AUTHOR_EMAIL"
)

var (
	author = &github.CommitAuthor{
		Name:  github.String(authorName),
		Email: github.String(authorEmail),
	}
)

func getNonExistantRepoName() string {
	rand.Seed(time.Now().UnixNano())
	return goutility.RandomString(20)
}

func getNonExistantFileName(context string) string {
	rand.Seed(time.Now().UnixNano())
	return context + "/" + goutility.RandomString(20) + ".test.txt"
}

func makeClient(t *testing.T) *github.Client {
	result := MakeClientWithOauth2StaticToken(accessToken)
	if result == nil {
		t.Errorf("No Client returned")
	}
	return result
}

func TestMakeClientWithOauth2StaticToken(t *testing.T) {

	makeClient(t)
}

func findRepoByOrg(t *testing.T, client *github.Client, orgName string, repoName string) *github.Repository {
	result, error := FindRepoByOrg(client, orgName, repoName)
	if result == nil {
		t.Errorf("No Repository found")
	}
	if error != nil {
		t.Errorf("%q", error)
	}

	return result
}

func TestFindRepoByOrg(t *testing.T) {

	client := makeClient(t)
	findRepoByOrg(t, client, existingOrgName, existingRepoName)
}

func TestFindRepoByOrg_invalid(t *testing.T) {

	client := makeClient(t)
	repo, error := FindRepoByOrg(client, existingOrgName, getNonExistantRepoName())
	if repo != nil {
		t.Errorf("Repository %s returned for invalid repo", repo)
	}
	if error == nil {
		t.Errorf("Expected an error for an invalid repo")
	}
}

func getFileContents(t *testing.T, client *github.Client, repository *github.Repository, path string) *github.RepositoryContent {

	result, error := GetFileContents(client, repository, path)
	if result == nil {
		t.Errorf("No RepositoryContent returned")
	}
	if error != nil {
		t.Errorf("%q", error)
	}

	return result
}

func TestGetContents(t *testing.T) {

	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)
	getFileContents(t, client, repo, existingFileName)
}

func TestGetContents_invalid(t *testing.T) {

	invalidPath := getNonExistantFileName("GetContents_invalid")
	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)
	contents, error := GetFileContents(client, repo, invalidPath)
	if contents != nil {
		t.Errorf("Contents %s returned for invalid path %s", contents, invalidPath)
	}
	if error == nil {
		t.Errorf("Expected an error for an invalid path")
	}
}

func createFile(t *testing.T, client *github.Client, repository *github.Repository, path string, message string, contents string) *github.RepositoryContentResponse {
	options := &github.RepositoryContentFileOptions{
		Message:   github.String(message),
		Branch:    github.String("master"),
		Content:   []byte(contents),
		Author:    author,
		Committer: author,
	}

	result, error := CreateFile(client, repository, path, options)

	if result == nil {
		t.Errorf("No contents returned")
	}
	if error != nil {
		t.Errorf("%q", error)
	}
	return result
}

func TestCreateFile(t *testing.T) {
	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)
	createFile(t, client, repo, getNonExistantFileName("CreateFile"), "Test Create File", "This is a test of the...nevermind")
}

func TestCreateFile_existing(t *testing.T) {
	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("Test Create Existing File"),
		Branch:    github.String("master"),
		Content:   []byte("This is a test of the...nevermind"),
		Author:    author,
		Committer: author,
	}
	contents, error := CreateFile(client, repo, existingFileName, options)

	if contents != nil {
		t.Errorf("Contents %s returned for an existing file", contents)
	}
	if error == nil {
		t.Errorf("Expected an error for existing file")
	}
}

func TestUpdateFile(t *testing.T) {

	fileName := getNonExistantFileName("UpdateFile")

	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)

	createContentsResult := createFile(t, client, repo, fileName, "Test Update File : Create", "First!\n")

	if createContentsResult == nil {
		t.Errorf("Create Contents is nil")
		return
	}

	getContentsResult := getFileContents(t, client, repo, fileName)
	if getContentsResult == nil {
		t.Errorf("GetContents Contents is nil")
		return
	}

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("Test Update File : Update"),
		Branch:    github.String("master"),
		Content:   []byte("Firs oh, second."),
		Author:    author,
		Committer: author,
		SHA:       getContentsResult.SHA,
	}

	contents, error := UpdateFile(client, repo, fileName, options)

	if contents == nil {
		t.Errorf("No contents returned")
	}
	if error != nil {
		t.Errorf("%q", error)
	}
}

// Unnecessary, a file existing or not is already determined by the required call to GetContents
// func TestUpdateFile_nonexistant(t *testing.T)

func TestCreateOrUpdateFile_create(t *testing.T) {

	fileName := getNonExistantFileName("CreateOrUpdateFile_create")

	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("Test Create Or Update File : Create"),
		Branch:    github.String("master"),
		Content:   []byte("Firs oh, second."),
		Author:    author,
		Committer: author,
	}

	contents, error := CreateOrUpdateFile(client, repo, fileName, options)

	if contents == nil {
		t.Errorf("No contents returned")
	}
	if error != nil {
		t.Errorf("%q", error)
	}
}

func TestCreateOrUpdateFile_update(t *testing.T) {

	fileName := getNonExistantFileName("CreateOrUpdateFile_update")

	client := makeClient(t)
	repo := findRepoByOrg(t, client, existingOrgName, existingRepoName)

	createContentsResult := createFile(t, client, repo, fileName, "Test Create Or Update File : Update : Create", "First!\n")

	if createContentsResult == nil {
		t.Errorf("Create Contents is nil")
		return
	}

	options := &github.RepositoryContentFileOptions{
		Message:   github.String("Test Create Or Update File : Update : Update"),
		Branch:    github.String("master"),
		Content:   []byte("Firs oh, second."),
		Author:    author,
		Committer: author,
	}

	contents, error := CreateOrUpdateFile(client, repo, fileName, options)

	if contents == nil {
		t.Errorf("No contents returned")
	}
	if error != nil {
		t.Errorf("%q", error)
	}
}
