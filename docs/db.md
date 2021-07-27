---
id: db
title: Fetch Dtails From Mongo DB 
sidebar_label: Fetch Dtails From Mongo DB
---
***
## Fetch Dtails From Mongo DBa API Specification

### HTTP Request

#### HTTP Method: POST
 ```
URL: http(s)://server/records
 ```

### Data Model

```
 {
        "startDate": "2016-01-25",
        "endDate": "2016-01-30",
        "minCount": 1000,
        "maxCount": 3000
    }

```


#### HTTP Request Body

<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Required / Optional</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>Content Request Body</td>
    <td>String</td>
    <td>Required</td>
    <td>Body Content-type: application/json Encoded as a JSON Object for the fields to be updated. Request body format should match flat json format shown in data model section.</td>
  </tr>
</table>

#### Request Syntax Example

The following URLs are examples:

curl --location --request POST 'http://localhost:8080/records' \
--header 'Content-Type: application/json' \
--data-raw ' {
        "startDate": "2016-01-25",
        "endDate": "2016-01-30",
        "minCount": 1000,
        "maxCount": 3000
    }
'

### HTTP Response

```
{
    "code": "0",
    "msg": "Success",
    "data": [
        {
            "key": "ZpoHRnZT",
            "createdAt": "2016-01-29T13:18:38.649Z",
            "totalCount": 2337
        },
        {
            "key": "bxoQiSKL",
            "createdAt": "2016-01-29T01:59:53.494Z",
            "totalCount": 2991
        },
        {
            "key": "NOdGNUDn",
            "createdAt": "2016-01-28T07:10:33.558Z",
            "totalCount": 2813
        }
    ]
}
```

#### HTTP Response Headers

<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>Content-Type</td>
    <td>String</td>
    <td>application/json</td>
  </tr>
</table>

#### HTTP Response Body

<table>
  <tr>
    <th>Name</th>
    <th>Type</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>Success or Failure acknowledgement</td>
    <td>JSON</td>
    <td>JSON payload</td>
  </tr>
</table>

#### HTTP Response Codes

<table width="100%">
  <tr>
    <th>HTTP Status Code</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>200</td>
    <td>OK</td>
  </tr>
  <tr>
    <td>400</td>
    <td>Bad Request</td>
  </tr>
  <tr>
    <td>403</td>
    <td>Forbidden</td>
  </tr>
  <tr>
    <td>500</td>
    <td>Internal Server Error</td>
  </tr>
</table>

