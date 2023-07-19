# frozen_string_literal: true

When('I request to generate identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Bezeichner.server_config['transport']['grpc']['user_agent'] }

  request = Bezeichner::V1::GenerateIdentifiersRequest.new(application: rows['application'], count: rows['count'].to_i)
  @response = Bezeichner::V1.server_grpc.generate_identifiers(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to map identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Bezeichner.server_config['transport']['grpc']['user_agent'] }

  request = Bezeichner::V1::MapIdentifiersRequest.new(ids: rows['request'].split(','))
  @response = Bezeichner::V1.server_grpc.map_identifiers(request, { metadata: })
rescue StandardError => e
  @response = e
end

Then('I should receive generated identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.ids.length).to eq(rows['count'].to_i)
  expect(@response.ids.first.length).to be > 0
end

Then('I should receive mapped identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.ids).to eq(rows['response'].split(','))
end

Then('I should receive a not found error from gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive an internal error from gRPC') do
  expect(@response).to be_a(GRPC::Internal)
end
