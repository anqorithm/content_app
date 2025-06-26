# content_app
REST APIs that can be used to upload and view video/audio contents, along with a simple UI page for local tests.

## How to run the app?
- Install docker
- Create an AWS account and an S3 bucket then update the following in the README file:
    1. AWS_ACCESS_KEY_ID
    2. AWS_SECRET_ACCESS_KEY
    3. AWS_REGION
    4. AWS_S3_BUCKET
- Run `docker compose up`
- The UI can be accessed from http://localhost:3001/ . After uploading the items, you can click on one of them using the list on the top of the page in order to see its content.
- Make sure to add `.env` file to `.gitignore` to not commit your creds. I kept it for now for easier build and test.

# Project documentation
For more info, refer to [this page](https://www.notion.so/Content-Management-and-Discovery-Service-21b9bf2fee4a8017b6c5fe8fbfe46bfd?source=copy_link)