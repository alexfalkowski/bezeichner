# frozen_string_literal: true

module Bezeichner
  module V1
    class HTTP < Nonnative::HTTPClient
      def ids(application, count, headers = {})
        headers.merge!(content_type: :json, accept: :json)

        get("v1/ids/#{application}/#{count}", headers, 10)
      end
    end
  end
end
