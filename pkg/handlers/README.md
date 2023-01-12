# Routes

This section has been separately added, as it is believed that the information it has will take a lot of space in the main [README](https://github.com/aloklearning/go-create-video/blob/main/README.md). Which will look a bit cluttered and bad.

## Introduction

### Authentication 

- The API uses a very basic authetication method, which is passing a param key `api_key` in the **HEADER** of the client tool like `POSTMAN`.
- The correct api key could be found [here](https://github.com/aloklearning/go-create-video/blob/main/pkg/handlers/authentication.go)
- Passing the right api key will let you easily access the routes which we have in the project.

### Routes Data

- To add the video with the following `metadata` and `annotations`, the following are the pre-requisites in respect of payload and URL:
    - URL: `localhost:8080/api/v1/createVideo`
    - Payload should be submitted in the `raw` body like this. Assuming `Content-Type` has been added as `application/json` already:
        ```json
            {
                "video_url": "Random Video URL",
                "meta_data": {
                    "author": "SOME AUTHOR NAME",
                    "video_name": "New Video Name",
                    "created_at": "0001-01-01T00:00:00Z", // optional
                    "modified_at": "0001-01-01T00:00:00Z", // optional
                    "total_duration": 300
                },
                "annotations": [
                    {
                        "start_time": 0,
                        "end_time": 250,
                        "type": "None",
                        "annotation": "Some New Annotation",
                        "additional_notes": ["New Additional Notes"]
                    }
                    // You can add more to it
                ]
            }
        ```
    - If submitted correctly, you will get a `STATUS_CREATED 201`, and List of all Video along with your added video
    