@benchmark
Feature: Benchmark gRPC API
  Make sure these endpoints perform at their best.

  Scenario: Generate identifiers in a good time frame and memory.
    When I request to generate identifiers with gRPC which performs in 15 ms
    And the process 'server' should consume less than '70mb' of memory
