# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'pg'
require 'grpc/health/v1/health_services_pb'

require 'bezeichner/v1/http'
require 'bezeichner/v1/service_services_pb'
require 'bezeichner/generator/pg'

# Namespace for Ruby helpers used by Bezeichner's end-to-end (Cucumber) feature tests.
#
# The production service is implemented in Go. This Ruby code is **not** a general-purpose
# client library; it is a thin convenience layer used by tests under `test/features/**`.
#
# It provides:
# - Access to the loaded test configuration (`.config/server.yml`) via {Bezeichner.config}.
# - Preconfigured gRPC stubs for:
#   - The Bezeichner v1 API ({Bezeichner::V1.grpc})
#   - gRPC Health API ({Bezeichner.health_grpc})
# - A helper for provisioning a Postgres sequence used by the `pg` generator kind
#   ({Bezeichner.pg}).
#
# ## Addresses / ports
#
# These helpers assume the test server listens on:
# - HTTP: `http://localhost:11000`
# - gRPC: `localhost:12000`
#
# This matches the configuration used by the nonnative harness in tests.
#
# ## gRPC metadata / user agent
#
# gRPC stubs are constructed with channel args that include a deterministic user-agent,
# allowing tests to assert observability/metadata behavior.
module Bezeichner
  class << self
    # Loads and memoizes the service configuration used by the feature test harness.
    #
    # @return [Hash] configuration loaded by nonnative
    #
    # @example Load generator applications
    #   Bezeichner.config.dig('generator', 'applications')
    def config
      @config ||= Nonnative.configurations('.config/server.yml')
    end

    # Returns a gRPC Health stub for the running Bezeichner service.
    #
    # The server is expected to expose the standard gRPC Health Checking Protocol.
    #
    # @return [Grpc::Health::V1::Health::Stub]
    #
    # @example Check service health
    #   Bezeichner.health_grpc.check(Grpc::Health::V1::HealthCheckRequest.new(service: 'bezeichner.v1.Service'))
    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Bezeichner.user_agent)
    end

    # Returns a helper used by tests to provision the Postgres sequence required by the
    # `pg` generator kind.
    #
    # The Go service does not create sequences; tests create and drop them around scenarios.
    #
    # @return [Bezeichner::Generator::PG]
    #
    # @example Create sequence for pg generator
    #   Bezeichner.pg.create
    # @example Drop sequence
    #   Bezeichner.pg.destroy
    def pg
      @pg ||= Bezeichner::Generator::PG.new
    end

    # Builds and memoizes the gRPC channel arguments used by stubs created in this helper.
    #
    # @return [Hash] gRPC channel args
    #
    # @example Use when constructing a stub
    #   Bezeichner::V1::Service::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Bezeichner.user_agent)
    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Bezeichner-ruby-client/1.0 gRPC/1.0')
    end
  end

  # Namespace for the Bezeichner v1 API helpers.
  module V1
    class << self
      # Returns an HTTP RPC gateway client for the Bezeichner v1 API.
      #
      # HTTP is implemented as an RPC gateway that routes requests by gRPC full method name.
      # See {Bezeichner::V1::HTTP} for the concrete endpoints used.
      #
      # @return [Bezeichner::V1::HTTP]
      #
      # @example Generate identifiers over HTTP
      #   Bezeichner::V1.http.generate('public-uuid', 3)
      def http
        @http ||= Bezeichner::V1::HTTP.new('http://localhost:11000')
      end

      # Returns a gRPC stub for the Bezeichner v1 API.
      #
      # @return [Bezeichner::V1::Service::Stub]
      #
      # @example Generate identifiers over gRPC
      #   req = Bezeichner::V1::GenerateIdentifiersRequest.new(application: 'public-uuid', count: 3)
      #   Bezeichner::V1.grpc.generate_identifiers(req)
      def grpc
        @grpc ||= Bezeichner::V1::Service::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Bezeichner.user_agent)
      end
    end
  end
end
