# frozen_string_literal: true

module Bezeichner
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate(application, count, opts = {})
        post('/bezeichner.v1.Service/GenerateIdentifiers', { application:, count: }.to_json, opts)
      end

      def map(ids, opts = {})
        post('/bezeichner.v1.Service/MapIdentifiers', { ids: }.to_json, opts)
      end
    end
  end
end
