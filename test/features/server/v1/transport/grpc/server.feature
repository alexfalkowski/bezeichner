@startup
Feature: Server

  Server allows users to get different types of identifiers.

  Scenario Outline: Identifiers for existing applications
    When I request to identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a successfuly identifiers from gRPC:
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

  Scenario Outline: Identifiers for missing applications
    When I request to identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found identifiers from gRPC

    Examples:
      | application  | count |
      | not_found    | 1     |
      | invalid_kind | 1     |
