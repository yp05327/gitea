// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package migrations

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	base "code.gitea.io/gitea/modules/migration"
	"github.com/stretchr/testify/assert"
)

func TestAzureDevOpsDownloadRepo(t *testing.T) {
	// Skip tests if Gitlab token is not found
	azureDevOpsPersonalAccessToken := os.Getenv("Azure_DEVOPS_READ_TOKEN")
	if azureDevOpsPersonalAccessToken == "" {
		t.Skip("skipped test because Azure_DEVOPS_READ_TOKEN was not in the environment")
	}

	resp, err := http.Get("https://dev.azure.com/go-gitea/test_repo")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Skipf("Can't access test repo, skipping %s", t.Name())
	}

	downloader, err := NewAzureDevOpsDownloaderV7(context.Background(), "https://dev.azure.com", "", "", azureDevOpsPersonalAccessToken, "go-gitea", "test_repo")
	if err != nil {
		t.Fatalf("NewAzureDevOpsDownloaderV7 is nil: %v", err)
	}

	repo, err := downloader.GetRepoInfo()
	assert.NoError(t, err)
	assertRepositoryEqual(t, &base.Repository{
		Name:          "test_repo",
		Owner:         "go-gitea",
		Description:   "Test repository for testing migration from azure devops to gitea",
		CloneURL:      "https://dev.azure.com/go-gitea/test_repo.git",
		OriginalURL:   "https://dev.azure.com/go-gitea/test_repo",
		DefaultBranch: "main",
	}, repo)

	topics, err := downloader.GetTopics()
	assert.NoError(t, err)
	assert.Len(t, topics, 2)
	assert.EqualValues(t, []string{"migration", "test"}, topics)

	milestones, err := downloader.GetMilestones()
	assert.NoError(t, err)
	assertMilestonesEqual(t, []*base.Milestone{
		{
			Title:   "1.1.0",
			Created: time.Date(2019, 11, 28, 8, 42, 44, 575000000, time.UTC),
			Updated: timePtr(time.Date(2019, 11, 28, 8, 42, 44, 575000000, time.UTC)),
			State:   "active",
		},
		{
			Title:   "1.0.0",
			Created: time.Date(2019, 11, 28, 8, 42, 30, 301000000, time.UTC),
			Updated: timePtr(time.Date(2019, 11, 28, 15, 57, 52, 401000000, time.UTC)),
			Closed:  timePtr(time.Date(2019, 11, 28, 15, 57, 52, 401000000, time.UTC)),
			State:   "closed",
		},
	}, milestones)

	labels, err := downloader.GetLabels()
	assert.NoError(t, err)
	assertLabelsEqual(t, []*base.Label{
		{
			Name:  "bug",
			Color: "d9534f",
		},
		{
			Name:  "confirmed",
			Color: "d9534f",
		},
		{
			Name:  "critical",
			Color: "d9534f",
		},
		{
			Name:  "discussion",
			Color: "428bca",
		},
		{
			Name:  "documentation",
			Color: "f0ad4e",
		},
		{
			Name:  "duplicate",
			Color: "7f8c8d",
		},
		{
			Name:  "enhancement",
			Color: "5cb85c",
		},
		{
			Name:  "suggestion",
			Color: "428bca",
		},
		{
			Name:  "support",
			Color: "f0ad4e",
		},
	}, labels)

	releases, err := downloader.GetReleases()
	assert.NoError(t, err)
	assertReleasesEqual(t, []*base.Release{
		{
			TagName:         "v0.9.99",
			TargetCommitish: "master",
			Name:            "First Release",
			Body:            "A test release",
			Created:         time.Date(2019, 11, 9, 16, 49, 21, 0, time.UTC),
			Published:       time.Date(2019, 11, 12, 20, 12, 10, 0, time.UTC),
			PublisherID:     1669571,
			PublisherName:   "mrsdizzie",
		},
	}, releases)

	// downloader.GetIssues()
	issues, isEnd, err := downloader.GetIssues(1, 2)
	assert.NoError(t, err)
	assert.False(t, isEnd)
	assertIssuesEqual(t, []*base.Issue{
		{
			Number:     1,
			Title:      "Please add an animated gif icon to the merge button",
			Content:    "I just want the merge button to hurt my eyes a little. \xF0\x9F\x98\x9D ",
			Milestone:  "1.0.0",
			PosterID:   18600385,
			PosterName: "guillep2k",
			State:      "closed",
			Created:    time.Date(2019, 11, 9, 17, 0, 29, 0, time.UTC),
			Updated:    time.Date(2019, 11, 12, 20, 29, 53, 0, time.UTC),
			Labels: []*base.Label{
				{
					Name:        "bug",
					Color:       "d73a4a",
					Description: "Something isn't working",
				},
				{
					Name:        "good first issue",
					Color:       "7057ff",
					Description: "Good for newcomers",
				},
			},
			Reactions: []*base.Reaction{
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "+1",
				},
			},
			Closed: timePtr(time.Date(2019, 11, 12, 20, 22, 22, 0, time.UTC)),
		},
		{
			Number:     2,
			Title:      "Test issue",
			Content:    "This is test issue 2, do not touch!",
			Milestone:  "1.1.0",
			PosterID:   1669571,
			PosterName: "mrsdizzie",
			State:      "closed",
			Created:    time.Date(2019, 11, 12, 21, 0, 6, 0, time.UTC),
			Updated:    time.Date(2019, 11, 12, 22, 7, 14, 0, time.UTC),
			Labels: []*base.Label{
				{
					Name:        "duplicate",
					Color:       "cfd3d7",
					Description: "This issue or pull request already exists",
				},
			},
			Reactions: []*base.Reaction{
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "heart",
				},
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "laugh",
				},
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "-1",
				},
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "confused",
				},
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "hooray",
				},
				{
					UserID:   1669571,
					UserName: "mrsdizzie",
					Content:  "+1",
				},
			},
			Closed: timePtr(time.Date(2019, 11, 12, 21, 1, 31, 0, time.UTC)),
		},
	}, issues)
}
