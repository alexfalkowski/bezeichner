# frozen_string_literal: true

When('I request to generate identifiers with HTTP:') do |table|
  rows = table.rows_hash
  headers = { request_id: SecureRandom.uuid, user_agent: Bezeichner.server_config['transport']['grpc']['user_agent'] }

  @response = Bezeichner::V1.server_http.generate(rows['application'], rows['count'], headers)
end

When('I request to map identifiers with HTTP:') do |table|
  rows = table.rows_hash
  headers = { request_id: SecureRandom.uuid, user_agent: Bezeichner.server_config['transport']['grpc']['user_agent'] }

  @response = Bezeichner::V1.server_http.map(rows['request'], headers)
end

Then('I should receive generated identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  ids = resp['ids']
  rows = table.rows_hash

  expect(ids.length).to eq(rows['count'].to_i)
  expect(ids.first.length).to be > 0
end

Then('I should receive mapped identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  ids = resp['ids']
  rows = table.rows_hash

  expect(ids).to eq(rows['response'].split(','))
end

Then('I should receive a not found error from HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive an internal error from HTTP') do
  expect(@response.code).to eq(500)
end
