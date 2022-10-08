@startup
Feature: Server

  Server allows users to get different types of identifiers.

  Scenario Outline: Identifiers for existing applications
    When I request identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive identifiers from gRPC:
      | application | <application> |
      | count       | <count>       |

    Examples:
      | application | count |
      | uuid        | 1     |
      | uuid        | 2     |
      | ksuid       | 1     |
      | ksuid       | 2     |
      | ulid        | 1     |
      | ulid        | 2     |
      | snowflake   | 1     |
      | snowflake   | 2     |
      | pg          | 1     |
      | pg          | 2     |
      | redis       | 1     |
      | redis       | 2     |

  Scenario Outline: Identifiers for missing applications
    When I request identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found error from gRPC

    Examples:
      | application  | count |
      | not_found    | 1     |
      | invalid_kind | 1     |

  Scenario Outline: Identifiers for erroneous applications
    Given the system is having issues for the application:
      | application | <application> |
    When I request identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive an internal error from gRPC
    And the system should return to a healthy state for the following appliation:
      | application | <application> |

    Examples:
      | application | count |
      | pg          | 1     |
      | redis       | 1     |
