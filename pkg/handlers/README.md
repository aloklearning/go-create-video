# Routes

This section has been separately added, as it is believed that the information it has will take a lot of space in the main [README](https://github.com/aloklearning/go-create-video/blob/main/README.md). Which will look a bit cluttered and bad.

## Introduction

### Authentication 

- The API uses a very basic authetication method, which is passing a param key `api_key` in the **HEADER** of the client tool like `POSTMAN`.
- The correct api key could be found [here](https://github.com/aloklearning/go-create-video/blob/main/pkg/handlers/authentication.go)
- Passing the right api key will let you easily access the routes which we have in the project.

### Routes Data

The data given here is to give you a generic idea and how to interact with the APIs. As the data is quite overwhelming and we are having a time constraint, I have kept the data to bare minimum possible. You will get to see more details, when you will run the project in your machine.

1. To add the video with the following `metadata` and `annotations`, the following are the pre-requisites in respect of payload and URL:
    - URL: `localhost:8080/api/v1/createVideo`
    - RQUEST TYPE: `POST`
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
    - In case you add an annotation having more `end_time` more than `total_duration` of the video, you will receieve `STATUS_BAD_REQUEST 400` with an error response like this:
    ```json
    {
        "error": "Your annotations end time {end_time} is out of bounds of duration of the video {total_duration}"
    }
    ```

2. To get all the Videos in the system, here are the following details:
    - URL: `localhost:8080/api/v1/videos`
    - REQUEST TYPE: `GET`
    - You need to pass the `video_url`, `type`, and `notes` in the param key inside `form-data` of `Body`.
    - Here is the sucess response with a `STATUS_OK 200`:
    ```json
    [
        {
            "video_id": "0620a10d-680e-4e61-8b8f-46f44c82b42f",
            "video_url": "Random Video URL",
            "meta_data": {
                "author": "AUTHOR NAME",
                "video_name": "New Video Name",
                "created_at": "2023-01-12T16:32:39.643285+05:30",
                "modified_at": "2023-01-12T16:32:39.643285+05:30",
                "total_duration": 300
            },
            "annotations": [
                {
                    "annotation_id": "1d420b1f-ac1c-40bc-9571-17cf00f09c6f",
                    "start_time": 0,
                    "end_time": 250,
                    "type": "None",
                    "annotation": "Some New Annotation",
                    "additional_notes": [
                        "New Additional Notes"
                    ]
                },
                ...
            ]
        }
    ]
    ```
    - On having no data, you will get `STATUS_OK 200` with a data `null`.

3. To list all annotations related to the video, here are the following details:
    - URL: `localhost:8080/api/v1/annotations`
    - REQUEST TYPE: `GET`
    - You need to pass the `video_url` in the param key inside `form-data` of `Body`.
    - On success you will receive all the annotations like this with `STATUS_OK 200`:
    ```json
    [
        {
            "annotation_id": "c844d5ed-490a-40de-b84c-6c9b22ffbcd6",
            "start_time": 0,
            "end_time": 250,
            "type": "None",
            "annotation": "Some New Annotation",
            "additional_notes": [
                "New Additional Notes"
            ]
        },
        {
            "annotation_id": "fcfb83b8-cf60-48f4-b8fb-ceb0dadce327",
            "start_time": 0,
            "end_time": 200,
            "type": "New None",
            "annotation": "Another Annotation",
            "additional_notes": [
                "Another Additional Notes"
            ]
        }
    ]
    ```
    - On error, you will receive error message with `STATUS_NOT_FOUND 404`.

4. To update Additional Notes with the Annotation type related to the video, here are the following details:
    - URL: `localhost:8080/api/v1/updateAdditionalNotes`
    - REQUEST TYPE: `PUT`
    - You need to pass the `video_url`, `type`, and `notes` in the param key inside `form-data` of `Body`. 
    - On success you will receive all the annotations with your updated annotation like below with `STATUS_CREATED 201`:
    ```json
    {
        "video_id": "9fe0a603-6b5e-4a3e-91ff-b4eedd1538e8",
        "video_url": "Random Video URL",
        "meta_data": {
            "author": "AUTHOR NAME",
            "video_name": "New Video Name",
            "created_at": "2023-01-12T14:14:32.505628+05:30",
            "modified_at": "2023-01-12T14:14:40.727834+05:30",
            "total_duration": 300
        },
        "annotations": [
            {
                "annotation_id": "c844d5ed-490a-40de-b84c-6c9b22ffbcd6",
                "start_time": 0,
                "end_time": 250,
                "type": "None",
                "annotation": "Some New Annotation",
                "additional_notes": [
                    "New Additional Notes",
                    "Checking with timings modification" // This one was added when you passed notes
                ]
            },
            ...
        ]
    }
    ```
    - On error, you will receive error message with `STATUS_BAD_REQUEST 400`

5. To update Annotation related to the video, please see the following details:
    - URL: `localhost:8080/api/v1/updateAnnotation/{video_url}/{type}`. 
        - Video URL and Annotation type has to be passed as a route param. Failing which you will get a `STATUS_BAD_REQUEST 400`.
    - REQUEST TYPE: `PUT`
    - Payload has to be a full object of annotation which has to be updated inside the video system:
    ```json
    {
        "start_time": 0,
        "end_time": 150,
        "type": "None",
        "annotation": "Updated New Annotation",
        "additional_notes": [
            "Updated Details"
        ]
    }
    ```
    - On successful submission of the data, you will recieve an updated video annotations along with the relevent video data with `STATUS_OK 200`:
    ```json
    {
        "video_id": "d62a7dc2-03e3-402f-b59f-6e88e8279e60",
        "video_url": "Random Video URL",
        "meta_data": {
            "author": "RANDOM AUTHOR",
            "video_name": "New Video Name",
            "created_at": "2023-01-12T17:00:27.128479+05:30",
            "modified_at": "2023-01-12T17:00:37.420777+05:30",
            "total_duration": 300
        },
        "annotations": [
            {
                "annotation_id": "82a10be3-d863-4734-bab9-e840cf754490",
                "start_time": 0,
                "end_time": 150,
                "type": "None",
                "annotation": "Updated New Annotation",
                "additional_notes": [
                    "Updated Details"
                ]
            },
            ...
        ]
    }
    ```
    - On having an out of bound `end_time`, you will receive `STATUS_BAD_REQUEST 400` with an error message.

6. To delete an annotation from the relevant video, here are the following details to perform the task:
    - URL: `localhost:8080/api/v1/deleteAnnotation`
    - REQUEST TYPE: `DELETE`
    - You need to pass the `video_url`, and `type` in the param key inside `form-data` of `Body`.
    - On successful operation, you will receieve **Changed Annotation data** with the relevant video information with `STATUS_OK 200`:
    ```json
    {
        "video_id": "d62a7dc2-03e3-402f-b59f-6e88e8279e60",
        "video_url": "Random Video URL",
        "meta_data": {
            "author": "Rohita",
            "video_name": "New Video Name",
            "created_at": "2023-01-12T17:00:27.128479+05:30",
            "modified_at": "2023-01-12T17:03:25.661425+05:30",
            "total_duration": 300
        },
        "annotations": [
            // Removed annotation doesn't exists in the system for the relevant video
            {
                "annotation_id": "c11919d9-1948-4daf-95ea-c21ea12aacf6",
                "start_time": 0,
                "end_time": 200,
                "type": "New None",
                "annotation": "Another Annotation",
                "additional_notes": [
                    "Another Additional Notes"
                ]
            }
        ]
    }
    ```
    - On error, you will receive `STATUS_BAD_REQUEST 400` with the error message of the possible cause.

7. To delete the video from the system, follow the following:
    - URL: `localhost:8080/api/v1/deleteVideo`
    - REQUEST TYPE: `DELETE`
    - You need to pass the `video_url`, and `type` in the param key inside `form-data` of `Body`.
    - On success, you will receieve `STATUS_OK 200` with this response:
    ```json
        "{Success: Video deleted successfully}"
    ```
    - On error, you will receive `STATUS_BAD_REQUEST 400` with the error message of the possible cause.
