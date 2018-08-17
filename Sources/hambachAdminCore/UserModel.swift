import FluentMySQL
import Vapor

final class User: Model
{
    typealias Database = MySQLDatabase
    typealias ID = Int
    static let idKey: IDKey = \User.userId
    static let entity = "users"

    var userId: Int?
    var login: String
    var password: String
    var firstname: String
    var lastname: String

    init(userId: Int, login: String, password: String, firstname: String, lastname: String)
    {
        self.userId = userId
        self.login = login
        self.password = password
        self.firstname = firstname
        self.lastname = lastname
    }
}
