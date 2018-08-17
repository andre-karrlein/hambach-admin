import XCTest
import Foundation
@testable import hambachCore

class HtmlTemplateTests: XCTestCase
{
    func testCanGetTemplate() throws
    {
        let template = HtmlTemplate(path: "/../html/Index.html")
        let fileManager = FileManager.default
        let path = fileManager.currentDirectoryPath

        let actual = try template.getTemplate()
        let expected = try String(contentsOfFile: path, encoding: String.Encoding.utf8)

        XCTAssertEqual(expected, actual)
    }
}
