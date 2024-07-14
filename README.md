# HackerNews Random
[![Go](https://github.com/stym06/hackernews-bot/actions/workflows/go.yml/badge.svg)](https://github.com/stym06/hackernews-bot/actions/workflows/go.yml)
![License](https://img.shields.io/github/license/stym06/hackernews-bot?style=flat-square)
![GitHub last commit](https://img.shields.io/github/last-commit/stym06/hackernews-bot?style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/stym06/hackernews-bot?style=flat-square)
![GitHub forks](https://img.shields.io/github/forks/stym06/hackernews-bot?style=flat-square)
![GitHub stars](https://img.shields.io/github/stars/stym06/hackernews-bot?style=flat-square)
![Language](https://img.shields.io/github/languages/top/stym06/hackernews-bot?style=flat-square)

<img src="https://images.pexels.com/photos/97050/pexels-photo-97050.jpeg" width="550" height="550">

This Go script fetches the top stories from Hacker News, selects a random story, and opens Chrome with the url

## Prerequisites

- Go 1.16 or later

## Quickstart
```
git clone github.com/stym06/hackernews-bot.git && cd hackernews-bot
make build
source ~/.zshrc
```

## Usage
When you run `hn` command now, it will open up Chrome with a random story from HackerNews

```
> hn show
  Tripping on Xenon Gas (2023)
  https://tripsitter.substack.com/p/xenon
```

