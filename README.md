# DataDog series
Undocumented DataDog features

## Disclaimer, or, Use at your own risk

This is **NOT** an official API from DataDog. You will not find any documentation regarding this code and usage.
The author assumes no responsibility or liability for any errors or problems regarding this API.
The information received with this API is provided on an "as is" basis with no guarantees of completeness, accuracy, usefulness or timeliness...

### Usage

```go
// set DATADOG_HOST, DATADOG_APP_KEY and DATADOG_API_KEY
client := ddseries.NewClient(http.DefaultClient)

// or

client := ddseries.Client{
    HttpClient: httpClient,
    
    Host:      datadogHost,
    Subdomain: "app",
    AppKey:    DatadogAppKey,
    ApiKey:    DatadogApiKey,
}

// Get RPM (Requests Per Minute) (Interval set to 60 seconds)
query := ddseries.Query{
    Q:        "sum:trace.http.request.hits{service:my-lovely-service}.as_count()",
    Interval: 60,
    From:     from,
    To:       to,
}

// Use timeouts. DD may hung for a long time
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

// Get parsed response
response, err := client.DoWithParsing(ctx, query)
if err != nil {
    log.Println(err)
}

// or 
// get raw response
response, err := client.Do(ctx, query)
```

#### Execute few queries in one request

```go
queryHits := ddseries.Query{
    Q:        "sum:trace.http.request.hits{service:my-lovely-service}.as_count()",
    Interval: 60,
    From:     from,
    To:       to,
}

queryErrors := ddseries.Query{
    Q:        "sum:trace.http.request.errors{service:my-lovely-service}.as_count()",
    Interval: 60,
    From:     from,
    To:       to,
}

queryApdex := ddseries.Query{
    Q:        "avg:trace.http.request.apdex.by.service{service:my-lovely-service}",
    Interval: 60,
    From:     from,
    To:       to,
}

response, err := client.Do(ctx, queryHits, queryErrors, queryApdex)

// or

queries := []ddseries.Query{queryHits, queryErrors, queryApdex}
response, err := client.Do(ctx, queries...)
```


### Example response

```json
{
  "responses":[
    {
      "status":"ok",
      "resp_version":2,
      "series":[
        {
          "unit":[
            {
              "family":"cache",
              "scale_factor":1.0,
              "name":"hit",
              "short_name":null,
              "plural":"hits",
              "id":39
            },
            null
          ],
          "query_index":0,
          "aggr":"sum",
          "scope":"service:my-lovely-service",
          "metric":"trace.http.request.hits",
          "expression":"sum:trace.http.request.hits{service:my-lovely-service}.as_count()",
          "tag_set":[

          ]
        }
      ],
      "to_date":1619958723000,
      "timing":"0.0662128925323",
      "query":"sum:trace.http.request.hits{service:my-lovely-service}.as_count()",
      "message":"",
      "res_type":"time_series",
      "interval":60,
      "times":[
        1619954820000.0,
        
        ...,
        
        1619958660000.0
      ],
      "from_date":1619955123000,
      "group_by":[

      ],
      "values":[
        [
          12345.0,
          
          ...,
          
          54321.0
        ]
      ]
    }
  ]
}

```
