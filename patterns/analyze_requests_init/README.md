### **analyze_requests_init** 

pattern is available [here](./pattern) .  

Input file [example](https://github.com/michalswi/honeypot-results/blob/main/full599.log) .

Output base on the input file (from above):
```
[
    {
        "id": 1,
        "Request line": "GET /cgi-bin/authLogin.cgi HTTP/1.1",
        "Remote Address": "185.180.143.11:33208",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
        "review": "The request is a GET method to access the authLogin.cgi script. The User-Agent appears to be a legitimate browser."
    },
    ...
]
```