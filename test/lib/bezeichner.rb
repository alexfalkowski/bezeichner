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
    def config
      @config ||= Nonnative.configurations('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Bezeichner.user_agent)
    end

    def pg
      @pg ||= Bezeichner::Generator::PG.new
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Bezeichner-ruby-client/1.0 gRPC/1.0')
    end
  end

  module V1
    class << self
      def http
        @http ||= Bezeichner::V1::HTTP.new('http://localhost:11000')
      end

      def grpc
        @grpc ||= Bezeichner::V1::Service::Stub.new('localhost:12000', :this_channel_is_insecure, channel_args: Bezeichner.user_agent)
      end
    end
  end
end
