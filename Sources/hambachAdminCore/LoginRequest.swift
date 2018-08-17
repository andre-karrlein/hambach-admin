import Vapor

struct LoginRequest: Content {
    var user: String
}
