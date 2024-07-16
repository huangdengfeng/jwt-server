package errs

var Success = New(0, "success")
var Unknown = New(1000, "system error [%s]")
var BasArgs = New(1001, "bad args [%s]")
var RpcError = New(1002, "call remote error [%s]")
var AttrKeyLimit = New(1003, "Attribute Key[%s] not allowed")
var JwtError = New(1004, "token invalid [%s]")
var JwtTokenExpired = New(1005, "token Expired")
