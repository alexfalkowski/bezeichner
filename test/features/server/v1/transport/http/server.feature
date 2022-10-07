@startup
Feature: Server

  Server allows users to get different types of identifiers.

  Scenario Outline: Identifiers for existing applications
    When I request to identifiers with HTTP:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a successfuly identifiers from HTTP:
      | application | <application> |
      | count       | <count>       |

    Examples:
      | application | count |
      | uuid        | 1     |
      | uuid        | 2     |

  Scenario Outline: Identifiers for missing applications
    When I request to identifiers with HTTP:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found identifiers from HTTP

    Examples:
      | application  | count |
      | not_found    | 1     |
      | invalid_kind | 1     |
