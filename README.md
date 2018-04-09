## Blockchain Identity Manager

Written using GO SDK of Hyperledger Fabric, this library serves as a Identity Manager for any Blockchain.

### How to run
1. Clone This project

2. Install Dep
    go get -u github.com/golang/dep/cmd/dep

3. Install pakages via dep
    dep ensure

4. Compile the project
    go build

5. Make a folder called keys and put public and private keys in that folder

6. Run the project


### API's

* Register and enroll new users in Organization - **Org1**:

`curl -s -X POST http://localhost:8000/registerandenroll -H "content-type: application/x-www-form-urlencoded" -d 'username=Jim&password=secret'`

**OUTPUT:**

```
{
  "success": true,
  "message":"Succesfully registered and enrolled",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTQ4NjU1OTEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE0OTQ4NjE5OTF9.yWaJhFDuTvMQRaZIqg20Is5t-JJ_1BP58yrNLOKxtNI"
}
```
