# frozen_string_literal: true

When('I request to map identifiers with gRPC:') do |table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::MapIdentifiersRequest.new(application: rows['application'], ids: rows['request'].split(','))
  @response = Bezeichner::V1.grpc.map_identifiers(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

When('I request to map {int} identifiers with gRPC:') do |count, table|
  rows = table.rows_hash
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Bezeichner::V1::MapIdentifiersRequest.new(application: rows['application'], ids: count.times.map { SecureRandom.hex })
  @response = Bezeichner::V1.grpc.map_identifiers(request, Bezeichner.grpc_options(metadata:))
rescue StandardError => e
  @response = e
end

Then('I should receive mapped identifiers from gRPC:') do |table|
  rows = table.rows_hash

  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(mapped_identifier_results(@response.ids)).to eq(mapped_identifiers(rows['results']))
end

Then('I should receive {int} unmapped identifiers from gRPC') do |count|
  expect(@response.meta['requestId']).to eq(@request_id)
  expect(@response.meta['userAgent']).to include('Bezeichner-ruby-client/1.0 gRPC/1.0')
  expect(mapped_identifier_results(@response.ids)).to all(satisfy { |result| !result.key?('mapped') })
  expect(@response.ids.length).to eq(count)
end
