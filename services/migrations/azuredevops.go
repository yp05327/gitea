// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package migrations

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"code.gitea.io/gitea/modules/log"
	base "code.gitea.io/gitea/modules/migration"
	"code.gitea.io/gitea/modules/structs"

	azuredevops "github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/git"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/tfvc"
)

var (
	_ base.Downloader        = &AzureDevOpsDownloaderV7{}
	_ base.DownloaderFactory = &AzureDevOpsDownloaderV7Factory{}
)

func init() {
	RegisterDownloaderFactory(&AzureDevOpsDownloaderV7Factory{})
}

// AzureDevOpsDownloaderV7Factory defines a azure devops downloader factory
type AzureDevOpsDownloaderV7Factory struct{}

// New returns a Downloader related to this factory according MigrateOptions
func (f *AzureDevOpsDownloaderV7Factory) New(ctx context.Context, opts base.MigrateOptions) (base.Downloader, error) {
	u, err := url.Parse(opts.CloneAddr)
	if err != nil {
		return nil, err
	}

	baseURL := u.Scheme + "://" + u.Host
	fields := strings.Split(u.Path, "/")
	oldOwner := fields[1]
	oldName := strings.TrimSuffix(fields[2], ".git")

	log.Trace("Create azure devops downloader BaseURL: %s %s/%s", baseURL, oldOwner, oldName)

	return NewAzureDevOpsDownloaderV7(ctx, baseURL, opts.AuthUsername, opts.AuthPassword, opts.AuthToken, oldOwner, oldName)
}

// GitServiceType returns the type of git service
func (f *AzureDevOpsDownloaderV7Factory) GitServiceType() structs.GitServiceType {
	return structs.AzureDevOpsService
}

// AzureDevOpsDownloaderV7 implements a Downloader interface to get repository information
// from azure devops via APIv7
// TODO: implement rate limit handling
type AzureDevOpsDownloaderV7 struct {
	base.NullDownloader
	ctx           context.Context
	coreClient    core.Client
	tfvcClient    tfvc.Client
	baseURL       string
	repoOwner     string
	repoName      string
	userName      string
	password      string
	maxPerPage    int
	SkipReactions bool
	SkipReviews   bool
}

// NewAzureDevOpsDownloaderV7 creates a azure devops Downloader via v7 API
func NewAzureDevOpsDownloaderV7(ctx context.Context, baseURL, userName, password, token, repoOwner, repoName string) (*AzureDevOpsDownloaderV7, error) {
	downloader := AzureDevOpsDownloaderV7{
		userName:   userName,
		baseURL:    baseURL,
		password:   password,
		ctx:        ctx,
		repoOwner:  repoOwner,
		repoName:   repoName,
		maxPerPage: 100,
	}

	connection := azuredevops.NewPatConnection(baseURL, token)
	if token == "" && userName != "" && password != "" {
		connection.AuthorizationString = azuredevops.CreateBasicAuthHeaderValue(userName, password)
	} else {
		return nil, fmt.Errorf("no token or username/password provided")
	}

	// init clients
	var err error
	downloader.coreClient, err = core.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}
	downloader.tfvcClient, err = tfvc.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}
	return &downloader, nil
}

// String implements Stringer
func (g *AzureDevOpsDownloaderV7) String() string {
	return fmt.Sprintf("migration from azure devops server %s %s/%s", g.baseURL, g.repoOwner, g.repoName)
}

func (g *AzureDevOpsDownloaderV7) LogString() string {
	if g == nil {
		return "<AzureDownloaderV7 nil>"
	}
	return fmt.Sprintf("<AzureDownloaderV7 %s %s/%s>", g.baseURL, g.repoOwner, g.repoName)
}

// SetContext set context
func (g *AzureDevOpsDownloaderV7) SetContext(ctx context.Context) {
	g.ctx = ctx
}

func (g *AzureDevOpsDownloaderV7) getProjectID() string {
	return fmt.Sprintf("%v/%v", g.repoOwner, g.repoName)
}

// GetRepoInfo returns a repository information
func (g *AzureDevOpsDownloaderV7) GetRepoInfo() (*base.Repository, error) {
	pjID := g.getProjectID()
	coreClient, err := core.NewClient(g.getConnection())
	if err != nil {
		return nil, err
	}
	pj, err := coreClient.GetProject(g.ctx, core.GetProjectArgs{ProjectId: &pjID})
	if err != nil {
		return nil, err
	}
	g.setRate(&resp.Rate)

	// convert azure devops project repo to stand Repo
	return &base.Repository{
		Owner:         g.repoOwner,
		Name:          *pj.Name,
		IsPrivate:     pj.Visibility == &core.ProjectVisibilityValues.Private,
		Description:   *pj.Description,
		OriginalURL:   *pj.Url,
		CloneURL:      *pj.Url,
		DefaultBranch: "", // TODO
	}, nil
}

// GetTopics return topics
// TODO

// GetMilestones returns milestones
// TODO

// GetLabels returns labels
func (g *AzureDevOpsDownloaderV7) GetLabels() ([]*base.Label, error) {
	g.tfvcClient.GetLabels(g.ctx, tfvc.GetLabelsArgs{
		RequestData: &git.TfvcLabelRequestData{
			Owner:      &g.repoOwner,
			Repository: &g.repoName,
		},
	})
	// TODO
}
