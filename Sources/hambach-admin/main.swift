import hambachAdminCore

let factory = Factory()

do {
    let server = try factory.createApplication(env: .detect())
    try server.run()
} catch {
    print("Whoops! An error occured: \(error)")
}
