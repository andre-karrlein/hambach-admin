import Vapor
import Foundation

public class Factory
{
    public init() {}

    public func createApplication(env: Environment) throws -> Application
    {
        var config = Config.default()
        var services = Services.default()
        var env = env

        try configure(&config, &env, &services)

        let app = try Application(config: config, environment: env, services: services)
        return app
    }

    public func createResponseCreator(template: String) -> ResponseCreator
    {
        let layout = self.createTemplate(page: "Layout")
        let navbar = self.createNavBar()
        let login = self.createTemplate(page: "Login")
        let articleCreator = self.createArticleCreator(template: template)
        let indexCreator = self.createIndexCreator(template: template)

        return ResponseCreator(layout: layout, navbar: navbar, login: login, articleCreator: articleCreator, indexCreator: indexCreator)
    }

    private func createTemplate(page: String) -> HtmlTemplate
    {
        return HtmlTemplate(path: "/html/" + page + ".html")
    }

    private func createNavBar() -> HtmlTemplate
    {
        return HtmlTemplate(path: "/html/MainNavbar.html")
    }

    private func createArticleCreator(template: String) -> ArticleCreator
    {
        return ArticleCreator(articleLayout: HtmlTemplate(path: "/html/" + template + ".html"))
    }

    private func createIndexCreator(template: String) -> IndexCreator
    {
        return IndexCreator(articleLayout: HtmlTemplate(path: "/html/" + template + ".html"))
    }
}
