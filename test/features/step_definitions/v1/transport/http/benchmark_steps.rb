# frozen_string_literal: true

When('I request to generate identifiers with HTTP which performs in {int} ms') do |time|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  expect { Bezeichner::V1.http.generate('ulid', 2, opts) }.to perform_under(time).ms
end
