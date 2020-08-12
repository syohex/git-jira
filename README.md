# git-jira

A utility for open JIRA URL by default web browse

## Setup

Set both `jira.baseURL` and `jira.projects` in your git config

```ini
[jira]
baseURL=https://jira.example.jp/browse/
projects=SOMEPROJ1,SOMEPROJ # project names are separated by comma
```

## Install

Build this command and put it in PATH

## How to use

Contain JIRA ticket ID in git branch name and run this command

```bash
% git switch -c syohex/SOMEPROJ-839
# some working
% git jira
# open JIRA ticket page on Web browser
```