Feature: HTTP List API
  These endpoints allows users to list configured identifier applications.

  Scenario: List configured applications
    When I request to list applications with HTTP
    Then I should receive configured applications from HTTP:
      | generator_applications | uuid:uuid,uuid_alias:uuid,ksuid:ksuid,ulid:ulid,snowflake:snowflake,xid:xid,nanoid:nanoid,typeid:typeid,invalid_kind:invalid_kind |
      | mapper_applications    | uuid,ulid                                                                                                                        |
      | generator_kinds        | ksuid,nanoid,snowflake,typeid,ulid,uuid,xid                                                                                      |
      | generate_count         | 2                                                                                                                                |
      | map_ids                | 2                                                                                                                                |
