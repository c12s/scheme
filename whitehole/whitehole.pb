syntax = "proto3";

option go_package = "github.com/c12s/scheme/whitehole";

package whitehole;

//Service
service WhiteholeService{
        rpc Put(PutReq) returns(PutResp){}
        rpc List(ListReq) returns(Resp){}
        rpc Get(GetReq) returns(Resp){}
}

message PutReq {
        string task = 1;
        string detail 2;
        string level = 3;
}

message PutResp {
        string error = 1;
}
