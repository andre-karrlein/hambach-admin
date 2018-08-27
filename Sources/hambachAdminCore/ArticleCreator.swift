import Vapor
import FluentMySQL
import Foundation

public class ArticleCreator
{
    var articleLayout: HtmlTemplate

    init(articleLayout: HtmlTemplate)
    {
        self.articleLayout = articleLayout
    }

    func createArticleWithouContent(userId: Int) throws -> String
    {
        let article = try self.replacePlaceholderWithoutContent(userId: userId)

        return article
    }

    func createArticle(content: Contents, userId: Int) throws -> String
    {
        let article = try self.replacePlaceholder(content: content, userId: userId)

        return article
    }

    private func replacePlaceholderWithoutContent(userId: Int) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()

        var currentUser = "André Karrlein"

        if (userId == 2) {
            currentUser = "Philipp Niedermeyer"
        }
        if (userId == 3) {
            currentUser = "Patrick Geißler"
        }

        var article = articleLayoutString.replacingOccurrences(of: "%title%", with: "")
        article = article.replacingOccurrences(of: "%user%", with: currentUser)
        article = article.replacingOccurrences(of: "%editor%", with: currentUser)
        article = article.replacingOccurrences(of: "%userId%", with: String(userId))
        article = article.replacingOccurrences(of: "%category%", with: "")
        article = article.replacingOccurrences(of: "%type%", with: "")
        article = article.replacingOccurrences(of: "%imageLink%", with: "")
        article = article.replacingOccurrences(of: "%date%", with: "")
        article = article.replacingOccurrences(of: "%article%", with: "")


        return article
    }

    private func replacePlaceholder(content: Contents, userId: Int) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()

        var name = "André Karrlein"
        var currentUser = "André Karrlein"

        if (content.creator == 2) {
            name = "Philipp Niedermeyer"
        }
        if (userId == 2) {
            currentUser = "Philipp Niedermeyer"
        }
        if (content.creator == 3) {
            name = "Patrick Geißler"
        }
        if (userId == 3) {
            currentUser = "Patrick Geißler"
        }

        var article = articleLayoutString.replacingOccurrences(of: "%title%", with: content.title)
        article = article.replacingOccurrences(of: "%user%", with: name)
        article = article.replacingOccurrences(of: "%editor%", with: currentUser)
        article = article.replacingOccurrences(of: "%userId%", with: String(userId))
        article = article.replacingOccurrences(of: "%category%", with: content.category)
        article = article.replacingOccurrences(of: "%type%", with: content.type)
        article = article.replacingOccurrences(of: "%imageLink%", with: content.titleImage)
        article = article.replacingOccurrences(of: "%date%", with: content.date)
        article = article.replacingOccurrences(of: "%article%", with: content.article)
        article = article.replacingOccurrences(of: "%articleId%", with: String(content.contentId!))


        return article
    }
}
