# frozen_string_literal: true

When('I request to map identifiers with HTTP:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Bezeichner.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Bezeichner::V1.http.map(rows['application'], rows['request'].split(','), opts)
end

When('I request to map {int} identifiers with HTTP:') do |count, table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  opts = Bezeichner.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Bezeichner::V1.http.map(rows['application'], count.times.map { SecureRandom.hex }, opts)
end

Then('I should receive mapped identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  rows = table.rows_hash

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(mapped_identifier_results(resp.fetch('ids', []))).to eq(mapped_identifiers(rows['results']))
end

Then('I should receive {int} unmapped identifiers from HTTP') do |count|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(mapped_identifier_results(resp.fetch('ids', []))).to all(satisfy { |result| !result.key?('mapped') })
  expect(resp.fetch('ids', []).length).to eq(count)
end
