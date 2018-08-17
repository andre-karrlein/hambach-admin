import Vapor
import Foundation

public class ResponseCreator
{
    var layout: HtmlTemplate
    var navbar: HtmlTemplate
    var login: HtmlTemplate
    var articleCreator: ArticleCreator
    var indexCreator: IndexCreator

    public init(layout: HtmlTemplate, navbar: HtmlTemplate, login: HtmlTemplate, articleCreator: ArticleCreator, indexCreator: IndexCreator)
    {
        self.layout = layout
        self.navbar = navbar
        self.login = login
        self.articleCreator = articleCreator
        self.indexCreator = indexCreator
    }

    func createResponse(content: [Contents], page: String, type: String, userId: Int) throws -> HTTPResponse
    {
        if (page == "Index") {
            return HTTPResponse(status: .ok, body: try self.createIndexResponseBody(content: content, type: type))
        }
        return HTTPResponse(status: .ok, body: try self.createArticleResponseBody(content: content, userId: userId))
    }

    func createLoginResponse() throws -> HTTPResponse
    {
        let layoutString = try self.layout.getTemplate()
        let loginString = try self.login.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: "")
        let body = layout.replacingOccurrences(of: "%content%", with: loginString)
        return HTTPResponse(status: .ok, body: body)
    }

    private func createArticleResponseBody(content: [Contents], userId: Int) throws -> String
    {
        let layoutString = try self.layout.getTemplate()
        let navbarString = try self.navbar.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: navbarString)
        var article = ""
        if (content.isEmpty) {
            article = try self.articleCreator.createArticleWithouContent(userId: userId)
        } else {
            article = try self.articleCreator.createArticle(content: content[0], userId: userId)
        }

        return layout.replacingOccurrences(of: "%content%", with: article)
    }

    private func createIndexResponseBody(content: [Contents], type: String) throws -> String
    {
        let layoutString = try self.layout.getTemplate()
        let navbarString = try self.navbar.getTemplate()
        let layout = layoutString.replacingOccurrences(of: "%navbar%", with: navbarString)
        let article = try self.indexCreator.createIndex(content: content)

        return layout.replacingOccurrences(of: "%content%", with: article)
    }
}
