import Vapor
import FluentMySQL
import Foundation

public class IndexCreator
{
    var articleLayout: HtmlTemplate

    init(articleLayout: HtmlTemplate)
    {
        self.articleLayout = articleLayout
    }

    func createIndex(content: [Contents]) throws -> String
    {
        let article = try self.replacePlaceholder(content: content)

        return article
    }

    private func replacePlaceholder(content: [Contents]) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()

        return articleLayoutString.replacingOccurrences(of: "%content%", with: try self.createContent(content: content))
    }

    private func createContent(content: [Contents])  throws -> String
    {
        var contentString = ""
        for contentItem in content {
            contentString += "<li><a href='/article/"
            contentString += String(contentItem.contentId!) + "'>"
            contentString += contentItem.title + "</a></li>"
        }

        return contentString
    }
}
