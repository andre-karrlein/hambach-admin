import Vapor

let factory = Factory()

public func routes(_ router: Router) throws
{
    router.get("/") { request -> HTTPResponse in
        let responseCreator = factory.createResponseCreator(template: "Login")
        return try responseCreator.createLoginResponse()
    }

    let sessions = router.grouped("/").grouped(SessionsMiddleware.self)

    sessions.get("home") { request -> Future<HTTPResponse> in
        let userId = try request.session()["user"] ?? "0"
        if (userId == "0") {
            throw Abort(.unauthorized, reason: "Not logged in!")
        }
        return request.withPooledConnection(to: .mysql) { db -> Future<HTTPResponse> in
            let responseCreator = factory.createResponseCreator(template: "Index")
            return db.query(Contents.self).all().map(to: HTTPResponse.self) { content in
                return try responseCreator.createResponse(content: content, page: "Index", type: "Allgemein", userId: Int(userId)!)
            }
        }
    }

    sessions.get("article", Int.parameter) { request -> Future<HTTPResponse> in
        let id = try request.parameters.next(Int.self)
        let userId = try request.session()["user"] ?? "0"
        if (userId == "0") {
            throw Abort(.unauthorized, reason: "Not logged in!")
        }
        return request.withPooledConnection(to: .mysql) { db -> Future<HTTPResponse> in
            return try Contents.find(id, on: db).map(to: HTTPResponse.self) { content in
                guard let content = content else {
                    throw Abort(.notFound, reason: "Could not find content.")
                }
                let responseCreator = factory.createResponseCreator(template: "Article")
                return try responseCreator.createResponse(content: [content], page: "Article", type: content.category.capitalized, userId: Int(userId)!)
            }
        }
    }

    sessions.get("create") { request -> HTTPResponse in
        let userId = try request.session()["user"] ?? "0"
        if (userId == "0") {
            throw Abort(.unauthorized, reason: "Not logged in!")
        }
        let responseCreator = factory.createResponseCreator(template: "Article")
        return try responseCreator.createResponse(content: [], page: "Article", type: "Allgemein", userId: Int(userId)!)
    }

    sessions.post(LoginRequest.self, at: "login") { request, data -> Response in
        //#TODO login handling
        try request.session()["user"] = data.user

        return request.redirect(to: "/home")
    }
    sessions.post(ContentData.self, at: "save") { request, data in
        return request.withPooledConnection(to: .mysql) { db -> Future<Response> in
            var titleImage = ""

            if (data.titleImage != nil) {
                titleImage = data.titleImage!
            }

            let content = Contents(
                title: data.title,
                creator: Int(data.userId)!,
                date: data.date,
                article: data.article,
                category: data.category,
                type: data.type,
                titleImage: titleImage
            )
            return content.save(on: db).map(to: Response.self) { _ in
                return request.redirect(to: "/home")
            }
        }
    }

    sessions.get("logout"){ request -> Response in
        try request.session()["user"] = "0"

        return request.redirect(to: "/")
    }
}
