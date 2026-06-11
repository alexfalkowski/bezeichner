Feature: Configuration
  Invalid configuration should stop the server from starting.

  Scenario Outline: Reject invalid generator application lists
    When I try to start the server with config "<config>"
    Then the server should fail to start

    Examples:
      | config                            |
      | duplicate_generator_names.yml     |
      | empty_generator_application.yml   |
      | empty_generator_kind.yml          |
      | nil_generator_application.yml     |
