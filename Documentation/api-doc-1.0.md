## Machine-L -- REST API Documentation
## ----- Doc version 1.0 *beta* [*under-construction*]

### Machine-L -- what is ?
Machine-L is a project to visualize statistics data. Basically for inquiry application, it tend to be applied on large data view matter. It aims to be simple, flexible and powerful.

### Machine-L Rest API -- what is ?
The Rest API (Application Programming Interface) as a purpose for being a Restful Service resolve the way to communicate and save stats data into database. To keep it simple, it uses a powerful No-SQL queries through MongoDB -- which means a system based on Json and Bson (mongo specific) format as exchange like any other document oriented data system like do, and keeps a lightweight approach with a high security of tokens and others sensitive data.

It is quite simple to deal with, and integrate with any application process like Machine-L do!
___
___

**API base-url**: *will be added soon*
<br>
## Signing To API
### base-url/api/o-sign/
method: post

#### Send Json:
```
{
  api_username:    "Joe-Biden",                      // 1
  api_pass:        "dghjfhkdhkkskkskksnsbyyyuzqcq",  // 2
  api_dev_email:   "joebidden@hotmail.com",          // 3
  api_owner_name:  "Organization Zen",               // 4
  api_owner_state: "Atlanta",                        // 5
}

```

##### params details:
1. string : the unique username 
2. string : password (recommend encrypted format)
3. string : the developer email
4. string : Organization who need the API
5. string : the state or country of the applicant


#### API response when successful(example):
```
{
  status: 1,
  rsp_code: 2472727,	// the 7 digits generated confirmation code 
  rsp_desc:  "User Signed Successful To Rest API, with status waiting for confirmation" 
}
```

#### API response when unsuccessful (example):
```
{
  status: 0,
  rsp_code: 201       // code always < 1000 when code errors
  rsp_desc: "Failed to register the Rest API: username already taken !"
}
```


#### status recap
1. code 0 : user signing unsuccessful -- check credentials
2. code 1 : waiting for (email) confirmation
3. code 2 : user account activated




___
## Confirming Signing To API
### base-url/api/o-sign/conf/
method: post
#### Send Json:
```
{
  api_username:    "Joe-Biden",                      // 1
  api_pass:        "dghjfhkdhkkskkskksnsbyyyuzqcq",  // 2
  conf_code:       2472727                           // 3
}

```

##### params details:
1. string : the unique username 
2. string : password (recommend encrypted format) <br> Anyway it will be encrypted anew  -- which is a double level pass protection.
3. int: the confirmation code as generated first by API


#### API response when successful(example):
```
{
  status: 2,
  rsp_desc:  "User registered Successful To Rest API, with status definitive" 
}
```

#### API response when unsuccessful (example):
```
{
  status: 6,
  rsp_desc: "The confirmation code sent is wrong!"
}
```

#### status recap
1. status 2 : user registered successfully to API
2. status 6 : confirmation code wrong
3. status 7 : credentials don't match signed records

when unsuccessful code > 5






___
## Connection To API 
### base-url/api/cmd/o-connect/
method: put, post

#### Send Json:
```
{
  api_username:    "Joe-Biden",                      // 1
  api_pass:        "dghjfhkdhkkskkskksnsbyyyuzqcq",  // 2
}

```

#### params details:
1. string : the unique username 
2. string : password (encrypted format if encrypted at API registration)

#### API response when successful (example):
```
{
  token:    "mtk-lozjfidfmjyàehvsjguehyfudjdpzgfz",		// the main token as a response
  rsp_code: 200,
  desc:     "User Signed Successful To Rest API",
}
```

#### API response when unsuccessful (example):
```
{
  rsp_code: 210,
  desc:     "User signed unsuccessful To Rest API – check username",
}
```

when unsuccessful rsp_code != 200 


#### code recap
1. code 200 : user registered successfully to API
2. code 210 : wrong username
3. code 211 : wrong password



___
## Query data
### base-url/api/cmd/query?user="username";...
method: put, post

___
## Submit data
### base-url/api/cmd/submit?user="username";...
method: post


