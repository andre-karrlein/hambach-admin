import Vapor

struct ContentData: Content {
    var articleId: String? = ""
    var title: String
    var date: String
    var article: String
    var userId: String
    var titleImage: String? = ""
    var category: String
    var type: String
}
