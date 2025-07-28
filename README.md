# S3 Bucket Downloader

A simple Go utility to download all files from an AWS S3 bucket to a local `storage` folder, preserving the directory structure.

## Features

- Downloads all objects from an S3 bucket (handles pagination)
- Saves files locally in the `storage` directory
- Preserves S3 folder structure

## Usage

1. Set up AWS credentials (via environment variables or config file).
2. Build and run:

```sh
go run . 2>&1 | tee output.txt
```

## Environment Configuration

Create a .env file based on the example provided:

1. Copy the example file to create a new .env:

    ```sh
    cp .env.example .env
    ```
2. Open the newly created `.env` file and fill in the necessary values.

These environment variables are required for authenticating with AWS and downloading files from your S3 bucket.

## Requirements

- Go 1.18+
- AWS credentials with S3 read access

## License

MIT