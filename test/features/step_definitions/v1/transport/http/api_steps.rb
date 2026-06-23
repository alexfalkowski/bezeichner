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

When('I request to map {int} existing identifiers with HTTP for application {string}') do |count, application|
  @request_id = SecureRandom.uuid
  opts = Bezeichner.http_options(
    headers: {
      request_id: @request_id, user_agent: 'Bezeichner-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    }
  )

  @response = Bezeichner::V1.http.map(application, Array.new(count, 'req1'), opts)
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

Then('I should receive mapped identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  ids = resp['ids']
  rows = table.rows_hash

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(ids).to eq(rows['response'].split(','))
end

Then('I should receive {int} mapped identifiers from HTTP') do |count|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta']['requestId']).to eq(@request_id)
  expect(resp['meta']['userAgent']).to eq('Bezeichner-ruby-client/1.0 HTTP/1.0')
  expect(resp['ids']).to eq(Array.new(count, 'resp1'))
end

Then('I should receive a not found error from HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive an internal error from HTTP') do
  expect(@response.code).to eq(500)
end

Then('I should receive an invalid argument error from HTTP') do
  expect(@response.code).to eq(400)
end
