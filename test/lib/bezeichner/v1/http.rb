# frozen_string_literal: true

module Bezeichner
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate(application, count, headers = {})
        headers.merge!(content_type: :json, accept: :json)

        get("v1/generate/#{application}/#{count}", headers, 10)
      end

      def map(ids, headers = {})
        headers.merge!(content_type: :json, accept: :json)

        get("v1/map/#{ids}", headers, 10)
      end
    end
  end
end
