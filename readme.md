
# nutanix 1.0

Nutanix API access

---
- #### Categories: infrastructure
- #### Image: gcr.io/direktiv/functions/nutanix 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/nutanix/issues
- #### URL: https://github.com/direktiv-apps/nutanix
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About nutanix

This function gives access to Nutanix APIs. PLease use the Nutanix's API Explorer to see the available API commands. More informtaion at [https://www.nutanix.dev](https://www.nutanix.dev).

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: nutanix
  image: gcr.io/direktiv/functions/nutanix:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: nutanix
  type: action
  action:
    function: nutanix
    secrets: ["nutanixpwd"]
    input: 
      auth:
        host: https://myserver:9440
        password: jq(.secrets.nutanixpwd)
        username: myuser@nutanix
        skipVerify: true
      api:
        path: "/api/nutanix/v3/vms/list"
        method: POST
        body:
          kind: vm
  catch:
  - error: "*"
```

   ### Secrets


- **nutanixpwd**: Password for the Nutanix user






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Nutanix API response.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "api_version": "3.1",
  "entities": [
    {
      "metadata": {
        "categories": {},
        "categories_mapping": {},
        "creation_time": "2021-12-02T12:37:54Z",
        "entity_version": "22",
        "kind": "vm",
        "last_update_time": "2022-08-05T10:40:01Z",
        "owner_reference": {
          "kind": "user",
          "name": "admin",
          "uuid": "00000000-0000-0000-0000-000000000000"
        }
      }
    }
  ]
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| nutanix | [interface{}](#interface)| `interface{}` |  | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| api | [PostParamsBodyAPI](#post-params-body-api)| `PostParamsBodyAPI` | ✓ | |  |  |
| auth | [PostParamsBodyAuth](#post-params-body-auth)| `PostParamsBodyAuth` | ✓ | |  |  |


#### <span id="post-params-body-api"></span> postParamsBodyApi

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| body | [interface{}](#interface)| `interface{}` |  | |  |  |
| method | string| `string` |  | `"GET"`| HTTP method to use | `POST` |
| path | string| `string` | ✓ | | API path to access | `/vms/list` |


#### <span id="post-params-body-auth"></span> postParamsBodyAuth

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` | ✓ | |  | `https://myserver:9440` |
| password | string| `string` |  | | Nutanix password |  |
| skipVerify | boolean| `bool` |  | | Skip SSL certificate verification |  |
| token | string| `string` |  | | Token authentication for e.g. Move |  |
| username | string| `string` |  | | Nutanix username |  |

 
