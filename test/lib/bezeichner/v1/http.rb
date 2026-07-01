# frozen_string_literal: true

module Bezeichner
  module V1
    # HTTP is a small test client for the Bezeichner v1 API over the HTTP RPC gateway.
    #
    # This client is used by the Cucumber feature tests in `test/features/**`. It is not a
    # general-purpose Ruby SDK.
    #
    # ## Transport model
    #
    # The Bezeichner service exposes both gRPC and HTTP. HTTP is implemented as an RPC
    # gateway that routes requests by **gRPC full method name**. Concretely, this client
    # POSTs JSON to paths like:
    #
    # - `/bezeichner.v1.Service/GenerateIdentifiers`
    # - `/bezeichner.v1.Service/ListApplications`
    # - `/bezeichner.v1.Service/MapIdentifiers`
    #
    # ## Request/response schema
    #
    # The JSON request bodies match the protobuf request messages:
    #
    # - {Bezeichner::V1::GenerateIdentifiersRequest}
    # - {Bezeichner::V1::ListApplicationsRequest}
    # - {Bezeichner::V1::MapIdentifiersRequest}
    #
    # The response bodies match the protobuf response messages and may include `meta`
    # (`requestId` and `userAgent` observability metadata) depending on server behavior.
    #
    # ## Error responses
    #
    # The HTTP gateway uses the same domain error categories as gRPC and renders
    # them as HTTP status responses. `InvalidArgument` is returned as HTTP 400,
    # `NotFound` as HTTP 404, and safe error messages are returned as
    # `text/error` bodies.
    #
    # ## Options
    #
    # Methods accept an optional `opts` hash that is passed through to the underlying
    # {Nonnative::HTTPClient#post} call (for example to set headers or timeouts in tests).
    class HTTP < Nonnative::HTTPClient
      # Calls GenerateIdentifiers over HTTP.
      #
      # @param application [String] configured generator application name
      # @param count [Integer] number of identifiers to generate; zero returns no IDs
      # @param opts [Hash] options forwarded to {Nonnative::HTTPClient#post}
      # @return [Object] HTTP response as returned by {Nonnative::HTTPClient#post}
      #
      # @example Generate three UUID identifiers for application "uuid"
      #   Bezeichner::V1.http.generate('uuid', 3)
      def generate(application, count, opts = {})
        post('/bezeichner.v1.Service/GenerateIdentifiers', { application:, count: }.to_json, opts)
      end

      # Calls ListApplications over HTTP.
      #
      # @param opts [Hash] options forwarded to {Nonnative::HTTPClient#post}
      # @return [Object] HTTP response as returned by {Nonnative::HTTPClient#post}
      def list(opts = {})
        post('/bezeichner.v1.Service/ListApplications', {}.to_json, opts)
      end

      # Calls MapIdentifiers over HTTP.
      #
      # @param application [String] configured mapper application name
      # @param ids [Array<String>] identifiers to classify as mapped or unmapped
      # @param opts [Hash] options forwarded to {Nonnative::HTTPClient#post}
      # @return [Object] HTTP response as returned by {Nonnative::HTTPClient#post}
      #
      # @example Classify identifiers using the configured application mapping
      #   Bezeichner::V1.http.map('uuid', %w[req1 req3])
      def map(application, ids, opts = {})
        post('/bezeichner.v1.Service/MapIdentifiers', { application:, ids: }.to_json, opts)
      end
    end
  end
end
