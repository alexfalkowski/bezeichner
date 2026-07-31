Feature: Configuration
  Invalid configuration should stop the server from starting.

  Scenario Outline: Reject invalid generator application lists
    When I try to start the server with config "<config>"
    Then the server should fail to start

    Examples:
      | config                            |
      | duplicate_generator_names.yaml     |
      | empty_generator_application.yaml   |
      | empty_generator_kind.yaml          |
      | nil_generator_application.yaml     |
