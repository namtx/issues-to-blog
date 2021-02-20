package main

import (
	"context"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/namtx/issues-to-blog/helper"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := "33871a9f95079a27b2f395e3e57010cb948b826d"
	owner := "namtx"
	sourceRepository := "til"
	targetBranch := "main"
	targetDirectory := "issues"
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	issues, _, err := client.Issues.ListByRepo(ctx, owner, sourceRepository, nil)
	if (err != nil) {
		panic(err.Error())
	}

	refs, _, err := client.Git.GetRef(ctx, owner, "issues-to-blog", "refs/heads/"+targetBranch)
	if (err != nil) {
		panic(err.Error())
	}

	entries := []*github.TreeEntry{}

	for _, issue := range issues {
		entries = append(
			entries, 
			&github.TreeEntry{
				Path: github.String(targetDirectory + helper.ToSnakeCase(helper.TrimSpecial(*issue.Title))), 
				Type: github.String("blob"), Content: github.String(*issue.Body), 
				Mode: github.String("100644"),
			},
		)
	}

	tree, _, err := client.Git.CreateTree(ctx, owner, "issues-to-blog", *refs.Object.SHA, entries)
	if (err != nil) {
		panic(err.Error())
	}
	parent, _, err := client.Repositories.GetCommit(ctx, owner, "issues-to-blog", *refs.Object.SHA)
	if (err != nil) {
		panic(err.Error())
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{Date: &date, Name: github.String("namtx"), Email: github.String("namtx.93@gmail.com")}
	commit := &github.Commit{Author: author, Message: github.String("test commit"), Tree: tree, Parents: []*github.Commit{parent.Commit}}

	newCommit, _, err := client.Git.CreateCommit(ctx, owner, "issues-to-blog", commit)

	refs.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, owner, "issues-to-blog", refs, false)
	if (err != nil) {
		panic(err.Error())
	}
}
