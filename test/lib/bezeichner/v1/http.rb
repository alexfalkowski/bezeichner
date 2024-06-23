# frozen_string_literal: true

module Bezeichner
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate(application, count, opts = {})
        post('v1/generate', { application:, count: }.to_json, opts)
      end

      def map(ids, opts = {})
        post('v1/map', { ids: }.to_json, opts)
      end
    end
  end
end
