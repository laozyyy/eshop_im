namespace go eshop.user_info

service UserService {
    GetOneUserResponse GetOneUser(1: GetOneUserRequest request)
    GetOneUserResponse GetOneUserByName(1: string name)
    InsertOneUserResponse InsertOneUser(1: User user)
}
struct InsertOneUserResponse {
    1: i64 code
    2: optional string errStr
}

struct User {
    1: string uid
    2: string name
    3: optional string phone
    4: optional string email
    5: string password
    6: i32 role
}

struct GetOneUserRequest {
    1: required string uid
}

struct GetOneUserResponse {
    1: optional User user
    2: i64 code
    3: optional string errStr
}
