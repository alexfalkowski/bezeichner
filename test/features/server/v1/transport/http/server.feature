Feature: Server

  Server allows users to get different types of identifiers.

  Scenario Outline: Generate identifiers for existing applications
    When I request to generate identifiers with HTTP:
      | application | <application> |
      | count       | <count>       |
    Then I should receive generated identifiers from HTTP:
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
      | xid         | 1     |
      | xid         | 2     |
      | snowflake   | 1     |
      | snowflake   | 2     |
      | nanoid      | 1     |
      | nanoid      | 2     |
      | typeid      | 1     |
      | typeid      | 2     |
      | pg          | 1     |
      | pg          | 2     |
      | redis       | 1     |
      | redis       | 2     |

  Scenario Outline: Generate identifiers for missing applications
    When I request to generate identifiers with HTTP:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found error from HTTP

    Examples:
      | application  | count |
      | not_found    | 1     |
      | invalid_kind | 1     |

  Scenario Outline: Generate identifiers for erroneous applications
    Given the system is having issues for the application:
      | application | <application> |
    When I request to generate identifiers with HTTP:
      | application | <application> |
      | count       | <count>       |
    Then I should receive an internal error from HTTP
    And the system should return to a healthy state for the following application:
      | application | <application> |

    Examples:
      | application | count |
      | pg          | 1     |
      | redis       | 1     |

  Scenario Outline: Map existing identifiers
    When I request to map identifiers with HTTP:
      | request | <request> |
    Then I should receive mapped identifiers from HTTP:
      | response | <response> |

    Examples:
      | request   | response    |
      | req1      | resp1       |
      | req1,req2 | resp1,resp2 |

  Scenario Outline: Map non existing identifiers
    When I request to map identifiers with HTTP:
      | request | <request> |
    Then I should receive a not found error from HTTP

    Examples:
      | request   |
      | req3      |
      | req1,req3 |
