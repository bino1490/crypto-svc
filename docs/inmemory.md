---
id: inmemory
title: Fetch Dtails From In Memory DB 
sidebar_label: Fetch Dtails From DB
---
***
## Fetch Dtails From In Memory DB API Specification

### HTTP Request

#### Allowed HTTP Method: GET, POST
 ```
URL: 
Fetch All Records:
GET http(s)://server/in-memory

Fetch SPECIFIC Records:
GET http://localhost:8080/in-memory/active-tabs

POST A Record:
Fetch SPECIFIC Records:
POST http://localhost:8080/in-memory

 ```

### Response Data Model

```
{
    "key": "active-tabs",
    "value": "getir"
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

curl --location --request POST 'http://localhost:8080/in-memory' \
--header 'Content-Type: application/json' \
--data-raw ' {
        "key": "active-tabs",
        "value": "getir"
    }
'

### HTTP Response

```
[
    {
        "key": "active-tabs",
        "value": "getir"
    }
]
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

