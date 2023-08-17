# GoDiscordBot

A Discord bot built using Golang and the Discord API, designed to provide useful commands and information.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Bot Commands](#bot-commands)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Introduction

GoDiscordBot is a Discord bot created with Golang and the Discord API. The bot enhances your Discord server by providing various commands to entertain and inform users. It can offer Golang tips, jokes, random facts, word meanings, IP information, and more.

## Features

- Golang tips for programming enthusiasts.
- Jokes to keep the server lively.
- Random facts to share interesting information.
- Word definitions for expanding vocabulary.
- IP information lookup for network-related queries.
- Help command to list available commands.

## Getting Started

### Prerequisites

Before you start, ensure you have the following:

- Go (1.16 or later): https://golang.org/doc/install
- Discord Bot Token: Create a bot on the [Discord Developer Portal](https://discord.com/developers/applications) and obtain the bot token.

### Installation

1. Clone this repository:

   ```sh
   git clone git@github.com:techno-stupid/Discord-Bot.git

   
2. Navigate to the project directory:

   ```sh
   cd Discord-Bot

3. Install dependencies:


   ```sh
   go mod tidy

4. Build the application:

   ```sh
   go build

## Usage

1. Invite the bot to your Discord server using the OAuth2 URL generated in the [Discord Developer Portal](https://discord.com/developers/applications)
2. Run the bot:
   ```sh
   go run main.go
3. The bot will come online and listen for commands in the server.

## Bot Commands

```!go tip: Get a random Golang programming tip.```
```!go joke: Receive a joke to brighten the mood.```
```!go fact: Learn a random interesting fact.```
```!go define {word}: Get the meaning of a word.```
```!go iplookup {ip}: Get information about an IP address.```
```!go help: List all available commands.```

# Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.
1. Fork the repository.
2. Create a new branch: git checkout -b feature/your-feature-name.
3. Make your changes and commit them: git commit -m 'Add some feature'.
4. Push to the branch: git push origin feature/your-feature-name.
5. Create a pull request.
