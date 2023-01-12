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

All the assumptions has been made around the agenda of *achieving the work with more efficiency*:

- The creation of the video and the annotations has been done in a single task, assuming it to be the part of the video itself. It can be added by submitting the full payload with the required data as mentioned in the [Routes README](https://github.com/aloklearning/go-create-video/tree/main/pkg/handlers)
- It has been assumed that for the ease of computations and processing time related data, such as, `total_duration`, `start_time`, `end_time` can be accepted as `INTEGER`. Furthermore, they have to be submitted in **seconds**. For example, if you have a video having 5 minutes as a total duration, so you will submit the details `300` which in seconds equals **5 minutes** (*same goes with the `start_time` and `end_time`*):
```json
{
    "video_url": "Some URL",
    "metadata": {
        "total_duration": 300
    }
}
```
- There has been assumption that a very basic API Key authentication implemented in this project will work do the job, as we want to see the API peforming different set of jobs. It is nothing but a stored key based check with the user passing the key in the `Header` with the key `api_key`. More details will be found in the [Routes README](https://github.com/aloklearning/go-create-video/tree/main/pkg/handlers).
- `Additional Notes` has been assumed as a part of `Annotations` related to a particular video. Further more it has been assumed that it is a `slice/list/array` of the notes, which could be altered by adding a new item not updating the existing item in the list.
- Assumptions has been made that there is no hard neccessity of adding test cases, which save a bit of time while completing the project.
- Assumptions has been made around not having a **complete** `Relational Database`. The tables were created having the idea to be able to store the data and to be a part of the video. Although the data model will explain how it has been linked. I have talked more about this thing in the [Improvement Section](https://github.com/aloklearning/go-create-video#improvements) below.
- It has been assumed that the by `Annotation Details` addition we mean the whole `Annotation Details` added to the **list/slice/array** of the videos items and **not** the specific items inside the Annotation Details.

## Improvements

- It is a universal truth that the project can always be improved. And while working with the this project, I do have some improvements pointers, which I believe could make the project more `robuts`, `clean`, and `efficient`. 
- Due to the time contraint, it was a bit difficult to achieve all of them, but if provided time, I could talk or add those suggestions in the future:
    - **Error responses/messages** in a some places could be added and in some places could be improved with a better `status code handling`.
    - It is very much evident that the project doesn't have any **test cases** running around the project. It can be added to things like:
        - Handlers
        - DB Connection Check
        - Methods
    - Database tables like `annotations` could have been added with a `FOREIGN KEY` relationship, making the `DELETE` and `UPDATE` function handled more from the SQL query side, rather a bit workaround.
    - More convenient **vairable names** and a bit of using **functional programming** could have save some repitition in the project. Although I have tried using the best and convenient name possible and tried not to use a lot of duplicate functions or logics.


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/aloklearning/go-create-video/blob/main/LICENSE.md) file for details

## Links

- Install go via [Official Website of Go](https://go.dev/)
- Learn more about the famous go sqlite package [go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)




