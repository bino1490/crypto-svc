# GOlang API - crypto-svc

## Collaborators ##
Developers: Bino Patric Prakash G

## UseCase
Microservice with crypto currency

***
## Fetch Dtails From In Memory DB API Specification

### HTTP Request

#### Allowed HTTP Method: GET, POST
 ```
URL: 
Fetch All Records:
GET http(s)://server/currency/all

Fetch SPECIFIC Records:
GET http://localhost:5000/currency/{symbole}
    http://localhost:5000/currency/ETHBTC

 ```

### Response Data Model

```
{
  "id": "ETH",
  "fullname": "Ethereum",
  "ask": "0.059798",
  "bid": "0.059784",
  "last": "0.059803",
  "open": "0.061361",
  "low": "0.058636",
  "high": "0.061358",
  "feeCurrency": "BTC"
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



