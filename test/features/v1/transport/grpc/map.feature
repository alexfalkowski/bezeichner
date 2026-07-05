Feature: gRPC Map API
  These endpoints allows users to map identifiers.

  Scenario Outline: Map existing identifiers
    When I request to map identifiers with gRPC:
      | application | <application> |
      | request     | <request>     |
    Then I should receive mapped identifiers from gRPC:
      | results | <results> |

    Examples:
      | application | request   | results               |
      | uuid        | req1      | req1:resp1            |
      | uuid        | req1,req2 | req1:resp1,req2:resp2 |
      | uuid        | req2,req1 | req2:resp2,req1:resp1 |
      | uuid        | req1,req1 | req1:resp1,req1:resp1 |
      | ulid        | req1      | req1:ulid_resp1       |

  Scenario: Map maximum identifiers
    When I request to map 2 identifiers with gRPC:
      | application | uuid |
    Then I should receive 2 unmapped identifiers from gRPC

  Scenario: Map too many identifiers
    When I request to map 3 identifiers with gRPC:
      | application | uuid |
    Then I should receive an invalid argument error from gRPC

  Scenario Outline: Map non existing identifiers
    When I request to map identifiers with gRPC:
      | application | uuid      |
      | request | <request> |
    Then I should receive mapped identifiers from gRPC:
      | results | <results> |

    Examples:
      | request   | results          |
      | req3      | req3:            |
      | req1,req3 | req1:resp1,req3: |

  Scenario: Map identifiers for a missing application
    When I request to map identifiers with gRPC:
      | application | not_found |
      | request     | req1      |
    Then I should receive a not found error from gRPC
