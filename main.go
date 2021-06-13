package main

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"text/template"
	"time"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type Label struct {
	Name string
	Color string
}

type Issue struct {
	Title string
	Body string
	Labels []Label
}

func main() {
	accessToken := os.Getenv("INPUT_ACCESSTOKEN")
	sourceRepository := os.Getenv("INPUT_SOURCEREPOSITORY")
	targetRepository := os.Getenv("INPUT_TARGETREPOSITORY")
	targetBranch := os.Getenv("INPUT_TARGETBRANCH")
	targetDirectory := os.Getenv("INPUT_TARGETDIRECTORY")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, "")
	if (err != nil) {
		panic(err.Error())
	}
	owner := *user.Login
	ownerEmailAddress := *user.Email

	refs, _, err := client.Git.GetRef(ctx, owner, targetRepository, "refs/heads/"+targetBranch)
	if (err != nil) {
		panic(err.Error())
	}
	entries := []*github.TreeEntry{}

	options := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions {
			Page: 1,
			PerPage: 50,
		},
	}
	issues, response, err := client.Issues.ListByRepo(ctx, owner, sourceRepository, options)
	if (err != nil) {
		panic(err.Error())
	}
	for (response.NextPage != 0) {
		options = &github.IssueListByRepoOptions {
			ListOptions: github.ListOptions {
				Page: response.NextPage,
				PerPage: 50,
			},
		}
		nextIssues, nextResponse, err := client.Issues.ListByRepo(ctx, owner, sourceRepository, options)
		if (err != nil) {
			panic(err.Error())
		}
		response = nextResponse
		issues = append(issues, nextIssues...)
	}
	for _, issue := range issues {
		issueDTO := Issue {
			Body: *issue.Body,
			Title: *issue.Title,
		}
		labels := []Label{}
		for _, label := range issue.Labels {
			labels = append(labels, Label{Color: *label.Color, Name: *label.Name})
		}
		issueDTO.Labels = labels

		entries = append(
			entries,
			&github.TreeEntry{
				Path: github.String(targetDirectory+"/"+ buildFileName(*issue.CreatedAt, *issue.Title)),
				Type: github.String("blob"),
				Content: github.String(buildFileContent(issueDTO)),
				Mode: github.String("100644"),
			},
		)
	}

	tree, _, err := client.Git.CreateTree(ctx, owner, targetRepository, *refs.Object.SHA, entries)
	if (err != nil) {
		panic(err.Error())
	}
	parent, _, err := client.Repositories.GetCommit(ctx, owner, targetRepository, *refs.Object.SHA)
	if (err != nil) {
		panic(err.Error())
	}
	parent.Commit.SHA = parent.SHA

	date := time.Now()
	author := &github.CommitAuthor{
		Date: &date,
		Name: github.String(owner),
		Email: github.String(ownerEmailAddress),
	}
	commit := &github.Commit{
		Author: author,
		Message: github.String("issues-to-blog-actions#"+os.Getenv("GITHUB_RUN_NUMBER")),
		Tree: tree,
		Parents: []*github.Commit{parent.Commit},
	}

	newCommit, _, err := client.Git.CreateCommit(ctx, owner, targetRepository, commit)

	refs.Object.SHA = newCommit.SHA
	_, _, err = client.Git.UpdateRef(ctx, owner, targetRepository, refs, true)
	if (err != nil) {
		panic(err.Error())
	}
}

func buildFileContent(issue Issue) string {
	templateStr :=
`---
layout: post
label: til
title: "{{.Title}}"
---

<p>
  {{range .Labels}}
  <span class="issue-label" style="background-color: #{{.Color}}">{{.Name}}</span>
  {{end}}
</p>
{{.Body}}
`

	tmpl, err := template.New("template").Parse(templateStr)
	if (err != nil) {
		panic(err.Error())
	}
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, issue)

	return fmt.Sprintln(buf)
}

func buildFileName(createdAt time.Time, title string) string {
	return fmt.Sprintf("%d-%d-%d-%s.md", createdAt.Year(), createdAt.Month(), createdAt.Day(), url.QueryEscape(title))
}
