# frozen_string_literal: true

When('I request to identifiers with HTTP:') do |table|
  rows = table.rows_hash
  headers = { request_id: SecureRandom.uuid, user_agent: Bezeichner.server_config['transport']['grpc']['user_agent'] }

  @response = Bezeichner::V1.server_http.ids(rows['application'], rows['count'], headers)
end

Then('I should receive a successfuly identifiers from HTTP:') do |table|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  ids = resp['ids']
  rows = table.rows_hash

  expect(ids.length).to eq(rows['count'].to_i)
  expect(ids.first.length).to be > 0
end

Then('I should receive a not found identifiers from HTTP') do
  expect(@response.code).to eq(404)
end
