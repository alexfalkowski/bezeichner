# frozen_string_literal: true

module Bezeichner
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate(application, count, opts = {})
        get("v1/generate/#{application}/#{count}", opts)
      end

      def map(ids, opts = {})
        get("v1/map/#{ids}", opts)
      end
    end
  end
end
