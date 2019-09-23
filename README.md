# ddctl
Golang cli client for https://dynamicdisplay.recurse.com

To compile and run this
```
$ make && ./ddctl hello world
Response Status: 201 CREATED
Response Body: {
  "author": "golang", 
  "created": "Mon, 23 Sep 2019 19:11:29 GMT", 
  "data": 101, 
  "id": 101
}
```

Main page for the dynamic display board on floor 5:
- https://dynamicdisplay.recurse.com

View past messages here:
- https://dynamicdisplay.recurse.com/matrix/api/message

