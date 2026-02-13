# frozen_string_literal: true

module Bezeichner
  module Generator
    # PG is a small helper used by feature tests to provision the Postgres sequence
    # required by the service's `pg` generator kind.
    #
    # The Go service uses a Postgres sequence to generate identifiers when a generator
    # application is configured with `kind: pg`. It queries:
    #
    #   SELECT nextval($1::regclass)
    #
    # where the sequence name is the configured generator application name.
    #
    # In the test harness, the generator application is named `pg`, so this helper
    # creates and drops a sequence named `pg` around scenarios.
    #
    # This class is intended for test setup/teardown only; it is not part of a
    # production Ruby client library.
    #
    # @example Create the sequence before tests
    #   Bezeichner.pg.create
    #
    # @example Drop the sequence after tests
    #   Bezeichner.pg.destroy
    class PG
      # Creates a new helper connected to the local test Postgres instance.
      #
      # Connection details are currently hard-coded to match the development/test
      # environment:
      #
      # - host: localhost
      # - port: 5432
      # - database: test
      # - user: test
      # - password: test
      #
      # @raise [URI::InvalidURIError] if the URI is malformed
      # @raise [PG::Error] if the connection cannot be established
      def initialize
        uri = URI.parse('postgres://test:test@localhost:5432/test?sslmode=disable')
        @conn = ::PG.connect(uri.hostname, uri.port, nil, nil, uri.path[1..], uri.user, uri.password)
      end

      # Creates the `pg` sequence.
      #
      # @return [PG::Result] result from Postgres
      # @raise [PG::Error] if the sequence already exists or the command fails
      def create
        @conn.exec('CREATE SEQUENCE pg')
      end

      # Drops the `pg` sequence.
      #
      # @return [PG::Result] result from Postgres
      # @raise [PG::Error] if the sequence does not exist or the command fails
      def destroy
        @conn.exec('DROP SEQUENCE pg')
      end
    end
  end
end
