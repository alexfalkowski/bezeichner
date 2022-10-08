# frozen_string_literal: true

module Bezeichner
  module Generator
    class PG
      def initialize
        uri = URI.parse('postgres://test:test@localhost:5432/test?sslmode=disable')
        @conn = ::PG.connect(uri.hostname, uri.port, nil, nil, uri.path[1..], uri.user, uri.password)
      end

      def create
        @conn.exec('CREATE SEQUENCE pg')
      end

      def destroy
        @conn.exec('DROP SEQUENCE pg')
      end
    end
  end
end
