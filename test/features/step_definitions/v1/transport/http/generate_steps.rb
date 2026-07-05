# frozen_string_literal: true

When('I request to generate identifiers with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Bezeichner.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Bezeichner::V1.http.generate(rows['application'], rows['count'].to_i, opts)
end

Then('I should receive generated identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  ids = resp['ids']
  rows = table.rows_hash

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(ids.length).to eq(rows['count'].to_i)
  expect(ids).to all(satisfy { |id| id.start_with?("#{rows['application']}_") })
end
