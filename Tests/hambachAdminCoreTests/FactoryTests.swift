import Vapor
import XCTest
@testable import hambachCore

class FactoryTests: XCTestCase
{
    func testCanCreateApplication() throws
    {
        let factory = Factory()
        let actual = try factory.createApplication(.development)

        XCTAssertNotNil(actual)
    }
}

extension FactoryTests
{
    static var allTests =
        ("testCanCreateApplication", testCanCreateApplication),
    ]
}
