import FluentMySQL
import Vapor

final class Contents: Model
{
    typealias Database = MySQLDatabase
    typealias ID = Int
    static let idKey: IDKey = \Contents.contentId
    static let entity = "content"

    var contentId: Int?
    var title: String
    var creator: Int
    var date: String
    var article: String
    var category: String
    var type: String
    var titleImage: String

    init(contentId: Int? = nil, title: String, creator: Int, date: String, article: String, category: String, type: String, titleImage: String)
    {
        self.contentId = contentId
        self.title = title
        self.creator = creator
        self.date = date
        self.article = article
        self.category = category
        self.type = type
        self.titleImage = titleImage
    }
}
