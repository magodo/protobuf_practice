/*************************************************************************
 Author: Zhaoting Weng
 Created Time: Wed 02 Jan 2019 05:18:10 PM CST
 Description:
 ************************************************************************/

var fs = require("fs")
var p = require("node-protobuf") // note there is no .Protobuf part anymore
// WARNING: next call will throw if desc file is invalid
var pb = new p(fs.readFileSync("foo.desc")) // obviously you can use async methods, it's for simplicity reasons
var obj = {
    "number": "123",
    "type": undefined
}

var buf = pb.serialize(obj, "tutorial.Person.PhoneNumber") // you get Buffer here, send it via socket.write, etc.
var newObj = pb.parse(buf, "tutorial.Person.PhoneNumber") // you get plain object here, it should be exactly the same as obj

console.log(newObj)
