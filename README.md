# url-shortner

A URL Shortner service written in Go. It accepts a URL as a POST Parameter and returns a shortened URL.
<br/><br/>

<!-- 
### Endpoints

/shorten : used for shortening supplied URL -->


## Steps to run the service
**N.B**
docker and git needs to be present on the system.
<br/>

### 1. Getting the Docker Image ready
You may either pull the image from the Docker Hub using the following command.
```shell
$ docker pull neogopher/url-shortner
```
Or you may build it yourself by cloning this repo and then building the image from the supplied Dockerfile.
```shell
$ git clone https://github.com/NeoGopher/url-shortner.git
$ cd url-shortner/
$ docker build -t neogopher/url-shortner:latest
```

### 2. Running a container from the image.
You can run the container using the below command. 
```shell
$ docker run -d -p 8080:8080 neogopher/url-shortner:latest
```

### 3. Hitting the API using CURL or similar tools.
```shell
$ curl -X POST -d '{"url": "https://docs.docker.com/engine/reference/builder/"}' http://localhost:8080/shorten
```
<br/>

## API Usage

The endpoint accepts valid JSON and responds with valid JSON.
<br/> <br/>

### POST: /shorten

---

Please provide your URL as parameter "url" in Request Body.

#### Request Body:

```json
{
  "url": "https://docs.docker.com/engine/reference/builder/"
}
```

#### Fields:

- **url (String)** : URL to be shortened


#### Response Body:
Status Code : 200

```json
{
  "data": {
    "shortUrl": "http://localhost:8080/07bf7aef"
  },
  "message": "shortened URL generated",
  "status":200
}
```

<br/>
