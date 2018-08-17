import Foundation

public class HtmlTemplate
{
    private let path: String

    public init(path: String)
    {
        self.path = path
    }

    public func getTemplate() throws -> String
    {
        return try readTemplate(path: getCurrentDir() + path)
    }

    func getCurrentDir() -> String
    {
        let fileManager = FileManager.default
        return fileManager.currentDirectoryPath
    }

    func readTemplate(path: String) throws -> String
    {
        let template = try String(contentsOfFile: path, encoding: String.Encoding.utf8)
        return template
    }
}
