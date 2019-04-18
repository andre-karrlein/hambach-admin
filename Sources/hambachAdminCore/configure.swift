import Vapor
import FluentMySQL

public func configure(_ config: inout Config, _ env: inout Environment, _ services: inout Services) throws {
    let host = NIOServerConfig.default(hostname: "0.0.0.0")
    services.register(host)

    var middleware = MiddlewareConfig.default()
    middleware.use(FileMiddleware.self)
    services.register(middleware)

    try services.register(FluentMySQLProvider())

    let router = EngineRouter.default()
    try routes(router)
    services.register(router, as: Router.self)

    let migrations = MigrationConfig()
    services.register(migrations)

    let mysqlConfig = MySQLDatabaseConfig(hostname: "db.karrlein.com", port: 3306, username: "hambach", password: "spvggHambach1933", database: "hambach")
    services.register(mysqlConfig)

    config.prefer(MemoryKeyedCache.self, for: KeyedCache.self)
}
