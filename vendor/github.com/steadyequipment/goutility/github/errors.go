package github

import (
	"fmt"
	"github.com/steadyequipment/goutility"
)

const (
	CodeContext = "goutility/github"

	NoClientProvidedCode goutility.ErrorCode = 1 << iota
	ListingReposInOrgCode

	NoRepoProvidedCode

	PathIsNotAFileIsADirectoryCode

	NoOptionsProvidedCode

	UnableToDetermineStatusOfFileCode
)

// region No Client Provided
type NoClientProvidedError struct {
	goutility.ErrorType
}

func MakeNoClientProvidedError() NoClientProvidedError {
	return NoClientProvidedError{
		ErrorType: goutility.MakeErrorWithCode("No Client provided", NoClientProvidedCode, CodeContext),
	}
}

// endregion

// region Listing Repos In Org
type ListingReposInOrgError struct {
	goutility.ErrorType

	orgName         string
	underlyingError error
}

func MakeListingReposInOrgError(orgName string, underlyingError error) ListingReposInOrgError {
	message := fmt.Sprintf("Error while listing repos in org '%s' in Github: %q", orgName, underlyingError)
	return ListingReposInOrgError{
		ErrorType:       goutility.MakeErrorWithCode(message, ListingReposInOrgCode, CodeContext),
		orgName:         orgName,
		underlyingError: underlyingError,
	}
}

func (this ListingReposInOrgError) OrgName() string {
	return this.orgName
}

func (this ListingReposInOrgError) UnderlyingError() error {
	return this.underlyingError
}

// endregion

// region Unable To Find Repo
type UnableToFindRepoError struct {
	goutility.ErrorType

	orgName  string
	repoName string
}

func MakeUnableToFindRepoError(orgName string, repoName string) UnableToFindRepoError {
	message := fmt.Sprintf("Unable to find repo '%s/%s'", orgName, repoName)
	return UnableToFindRepoError{
		ErrorType: goutility.MakeErrorWithCode(message, ListingReposInOrgCode, CodeContext),
		orgName:   orgName,
		repoName:  repoName,
	}
}

func (this UnableToFindRepoError) OrgName() string {
	return this.orgName
}

func (this UnableToFindRepoError) RepoName() string {
	return this.repoName
}

// endregion

// region No Repo Provided
type NoRepoProvidedError struct {
	goutility.ErrorType
}

func MakeNoRepoProvidedError() NoRepoProvidedError {
	return NoRepoProvidedError{
		ErrorType: goutility.MakeErrorWithCode("No Repository provided", NoRepoProvidedCode, CodeContext),
	}
}

// endregion

// region Path Is Not A File Is A Directory
type PathIsNotAFileIsADirectoryError struct {
	goutility.ErrorType

	path string
}

func MakeNotAFileIsADirectoryError(path string) PathIsNotAFileIsADirectoryError {
	message := fmt.Sprintf("'%s' is not a file but a directory", path)
	return PathIsNotAFileIsADirectoryError{
		ErrorType: goutility.MakeErrorWithCode(message, PathIsNotAFileIsADirectoryCode, CodeContext),
		path:      path,
	}
}

func (this PathIsNotAFileIsADirectoryError) Path() string {
	return this.path
}

// endregion

// region No Options Provided
type NoOptionsProvidedError struct {
	goutility.ErrorType
}

func MakeNoOptionsProvidedError() NoOptionsProvidedError {
	return NoOptionsProvidedError{
		ErrorType: goutility.MakeErrorWithCode("No Options provided", NoOptionsProvidedCode, CodeContext),
	}
}

// endregion

// region Unable To Determine Status of File
type UnableToDetermineStatusOfFileError struct {
	goutility.ErrorType

	path string
}

func MakeUnableToDetermineStatusOfFileError(path string) UnableToDetermineStatusOfFileError {
	message := fmt.Sprintf("Unable to determine status of '%s' to determine whether to create or update", path)
	return UnableToDetermineStatusOfFileError{
		ErrorType: goutility.MakeErrorWithCode(message, UnableToDetermineStatusOfFileCode, CodeContext),
		path:      path,
	}
}

func (this UnableToDetermineStatusOfFileError) Path() string {
	return this.path
}

// endregion
