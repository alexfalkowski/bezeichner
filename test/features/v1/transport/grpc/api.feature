Feature: gRPC API
  These endpoints allows users to get different types of identifiers.

  Scenario Outline: Generate identifiers for existing applications
    When I request to generate identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive generated identifiers from gRPC:
      | application | <application> |
      | count       | <count>       |

    Examples:
      | application | count |
      | uuid        |     1 |
      | uuid        |     2 |
      | ksuid       |     1 |
      | ksuid       |     2 |
      | ulid        |     1 |
      | ulid        |     2 |
      | xid         |     1 |
      | xid         |     2 |
      | snowflake   |     1 |
      | snowflake   |     2 |
      | nanoid      |     1 |
      | nanoid      |     2 |
      | typeid      |     1 |
      | typeid      |     2 |
      | pg          |     1 |
      | pg          |     2 |

  Scenario Outline: Generate identifiers for missing applications
    When I request to generate identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found error from gRPC

    Examples:
      | application  | count |
      | not_found    |     1 |
      | invalid_kind |     1 |

  Scenario Outline: Generate identifiers for erroneous applications
    Given I set the proxy for service "<application>" to "close_all"
    And I should see "<application>" as unhealthy
    When I request to generate identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive an internal error from gRPC
    And I should reset the proxy for service "<application>"
    And I should see "<application>" as healthy

    Examples:
      | application | count |
      | pg          |     1 |

  Scenario Outline: Map existing identifiers
    When I request to map identifiers with gRPC:
      | request | <request> |
    Then I should receive mapped identifiers from gRPC:
      | response | <response> |

    Examples:
      | request   | response    |
      | req1      | resp1       |
      | req1,req2 | resp1,resp2 |

  Scenario Outline: Map non existing identifiers
    When I request to map identifiers with gRPC:
      | request | <request> |
    Then I should receive a not found error from gRPC

    Examples:
      | request   |
      | req3      |
      | req1,req3 |
