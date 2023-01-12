# Go Create Videos

- This project is a simple RESTful API for managing the video with the following details:
    - **Metadata:** A set of data that describes and gives information about other data.
    - **Annotations:** An annotation allows us to capture time related information about the video. For example, we may want to create an annotation that references to a part of the video between 00:04:00 to 00:05:00 that contains an advertisement. 
- This project allows you to store the data and access it via various API routes mentioned in the below section
- This project comes with a persistant lightweighted database `SQLite3`.

## Project Checklists

This project checks the following features:
- [x] Basic `API Key` authentication in place for all the APIs.
- [x] Have a peristant database `SQLite3`.
- [x] Creating the video with the relevant `metadata` and `annotations`.
- [x] List all `annotations` related to the Video.
- [x] Add `additional notes` by specifying the `annotation type` for the relevant video.
- [x] Update the `annotation details` for the relevant video.
- [x] Deleting the `annotation` for the relevant video.
- [x] Deleting the `video` from the system.
- [x] Sending you an error when the user tries to:
    - [x] Add duplicate video identified by the `video_url`
    - [x] Add `annotation` with **out of bound** timings of the provided video duration.
    - [x] Get all the `annotations` without providing a proper `video_url`.
    - [x] Update `additional details` without providing the necessary keys: `video_url`, `type`, `notes`.
    - [x] Update an `annotation details` without providing the necessary params: `video_url`, `type`.
    - [x] Deleting an `annotation` without providing the necessary keys: `video_url`, `type`.
    - [x] Deleting a `video` without providing a valid `video_url`.

## Getting Started

### Pre-Requisites

- [x] In case you'd like to run the project, please have `Go` in your system. Please verify it by typing the command `go version` in your terminal.
- [x] It is better to have `Docker` as well for seamless working of the project on a server.

### Running the Project

To run the project there are few ways, but I have chosen to use a very simple one:
- Run the below command after going into the project's directory
```
cd cmd && go run .
```

- In case you fail to get the response, use the below command after going into the project's directory:
```
go run ./..
```

**Important:** You need to be inside the same directory of `main.go`, so ensure the path before running the project.

### Development Server

The project run on the localhost port `8080`. Please access the APIs via `localhost:8080/api/v1/...`.

### Routes

For best maintainence of details, the routes section has a separate README which can be found here: [Route](https://github.com/aloklearning/go-create-video/tree/main/pkg/handlers)

## Assumptions

- The creation of the video and the annotations has been done in a single task, assuming it to be the part of the video itself. It can be added by submitting the full payload with the required data as mentioned in the [Routes README](https://github.com/aloklearning/go-create-video/tree/main/pkg/handlers)
- For the ease of computations and processing time related data, such as, `total_duration`, `start_time`, `end_time` are accepted as `INTEGER`. Further more, they have to be submitted in seconds data. For example, if you have a video having 5 minutes as a total duration, so you will submit the details `300` which in seconds equals **5 minutes**:
```json
{
    "video_url": "Some URL",
    "metadata": {
        "total_duration": 300
    }
}
```
    - Same goes for `start_time` and `end_time`.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/aloklearning/go-create-video/blob/main/LICENSE.md) file for details

## Links

- Install go via [Official Website of Go](https://go.dev/)
- Learn more about the famous go sqlite package [go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)




