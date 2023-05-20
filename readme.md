<h1 align="center" style="display: block; font-size: 2.5em; font-weight: bold; margin-block-start: 1em; margin-block-end: 1em;">
    Go-Whatsapp
</h1>

## Introduction

Welcome to the WhatsApp Bot project developed in Golang using the Whatsmeow library! This README.md file provides an introduction and guidance on how to use and set up this bot.

## Overview

This project aims to create a versatile and interactive WhatsApp bot using the Golang programming language. The bot leverages the powerful features of the Whatsmeow library, which acts as an interface to interact with the WhatsApp Web API. Additionally, this bot integrates with OpenAI's external resources, allowing it to leverage natural language processing capabilities to provide intelligent responses.

## Features

-   **WhatsApp Bot**: This bot allows you to automate various tasks and interact with WhatsApp directly from your Golang application.
-   **Two-way Communication**: The bot can both send messages to WhatsApp contacts and groups, as well as receive incoming messages and respond accordingly.
-   **OpenAI Integration**: By leveraging OpenAI's external resources, the bot can understand and respond to natural language queries and provide intelligent and contextually relevant answers.

## Prerequisites

Before you begin, make sure you have the following set up:

-   Go installed on your machine (version 1.20 or later).
-   A valid WhatsApp account with an active phone number.
-   API credentials from OpenAI, allowing access to the external resources for natural language processing.

## Getting Started

To get started with the WhatsApp bot, follow these steps:

1. Clone or download this repository to your local machine.
2. Install the necessary dependencies using the package manager of your choice.
3. Set up your API credentials from OpenAI by following their documentation.
4. Configure the bot with your WhatsApp account details and OpenAI API credentials.

## Setting up the Environment

To configure the environment for the WhatsApp bot, follow these steps:

-   Create a .env file in the project's root directory.
-   Open the .env file and add the following variables:

    ### OpenAI API credentials

    `OPENAI_API_KEY=your_openai_api_key`

    Make sure to replace the placeholder values with your own WhatsApp phone number, device name, and OpenAI API key.

    ### How to get your OpenAI API key

    To obtain your OpenAI API key, please follow the instructions below:

    i. Visit the OpenAI website at https://openai.com and sign in to your account (or create a new one if you don't have an account).
    ii. Navigate to your dashboard or account settings to access your API key.
    iii. Generate a new API key if you haven't done so already.
    iv. Copy the generated API key to your clipboard.

5. Run the bot using the provided command or by building and executing the Golang code.
6. Enjoy interacting with the WhatsApp bot and explore its capabilities!

## Contributing

Contributions to this project are welcome! If you encounter any issues, have suggestions for improvements, or would like to add new features, please submit a pull request or open an issue in the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

-   The Whatsmeow library for providing a convenient interface to interact with the WhatsApp Web API.
-   OpenAI for their powerful natural language processing capabilities and external resources.

Please refer to the detailed documentation and code comments for more information on how to use and customize the WhatsApp bot. We hope you find this project useful and have fun experimenting with it!
