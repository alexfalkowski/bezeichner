Feature: gRPC API
  These endpoints allows users to get different types of identifiers.

  Scenario: List configured applications
    When I request to list applications with gRPC
    Then I should receive configured applications from gRPC:
      | generator_applications | uuid:uuid,uuid_alias:uuid,ksuid:ksuid,ulid:ulid,snowflake:snowflake,xid:xid,nanoid:nanoid,typeid:typeid,invalid_kind:invalid_kind |
      | mapper_applications    | uuid,ulid                                                                                                                        |
      | generator_kinds        | ksuid,nanoid,snowflake,typeid,ulid,uuid,xid                                                                                      |
      | generate_count         | 2                                                                                                                                |
      | map_ids                | 2                                                                                                                                |

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
      | uuid_alias  |     1 |
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

  Scenario: Generate maximum identifiers for existing applications
    When I request to generate identifiers with gRPC:
      | application | uuid |
      | count       | 2 |
    Then I should receive generated identifiers from gRPC:
      | application | uuid |
      | count       | 2 |

  Scenario: Generate too many identifiers for existing applications
    When I request to generate identifiers with gRPC:
      | application | uuid |
      | count       | 3 |
    Then I should receive an invalid argument error from gRPC

  Scenario Outline: Generate identifiers for missing applications
    When I request to generate identifiers with gRPC:
      | application | <application> |
      | count       | <count>       |
    Then I should receive a not found error from gRPC

    Examples:
      | application  | count |
      | not_found    |     1 |
      | invalid_kind |     1 |

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
