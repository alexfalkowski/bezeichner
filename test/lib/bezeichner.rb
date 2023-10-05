# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'pg'
require 'grpc/health/v1/health_services_pb'

require 'bezeichner/v1/http'
require 'bezeichner/v1/service_services_pb'
require 'bezeichner/generator/pg'

module Bezeichner
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def server_config
      @server_config ||= YAML.load_file('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:9090', :this_channel_is_insecure)
    end

    def pg
      @pg ||= Bezeichner::Generator::PG.new
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Bezeichner::V1::HTTP.new('http://localhost:8080')
      end

      def server_grpc
        @server_grpc ||= Bezeichner::V1::Service::Stub.new('localhost:9090', :this_channel_is_insecure)
      end
    end
  end
end
