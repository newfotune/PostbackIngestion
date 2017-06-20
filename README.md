 A service to function as a small scale simulation of how to distribute data to third parties in real time.
 
 Redis is running on localhost port 6379 with tcp connection.
 
 To test, send a post request to *ingest.php* file.
 The post data should look something like this...
 
 #sample request
 
    (POST) http://{server_ip}/ingest.php
    (RAW POST DATA) 
    {  
        "endpoint":{  
            "method":"GET",
            "url":"http://sample_domain_endpoint.com/data?title={mascot}&image={location}&foo={bar}"
        },
        "data":[  
            {  
                "mascot":"Gopher",
                "location":"https://blog.golang.org/gopher/gopher.png"
            }
         ]
    }
    
Response will be writted to the requests.log file

#Sample response

    GET http://sample_domain_endpoint.com/data?title=Gopher&image=https%3A%2F%2Fblog.golang.org%2Fgopher%2Fgopher.png&foo=
